'use client';

import React from 'react';

import { produce } from 'immer';
import { useQueryClient } from '@tanstack/react-query';

import { Flex } from '@ui/layout/Flex';
import { Clock } from '@ui/media/icons/Clock';
import { InvoiceStatus } from '@graphql/types';
import { CheckCircle } from '@ui/media/icons/CheckCircle';
import { SlashCircle01 } from '@ui/media/icons/SlashCircle01';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { useInvoicesMeta } from '@shared/state/InvoicesMeta.atom';
import { renderStatusNode } from '@shared/components/Invoice/Cells';
import { Menu, MenuItem, MenuList, MenuButton } from '@ui/overlay/Menu';
import { useVoidInvoiceMutation } from '@shared/graphql/voidInvoice.generated';
import { useInfiniteGetInvoicesQuery } from '@shared/graphql/getInvoices.generated';
import { useUpdateInvoiceStatusMutation } from '@shared/graphql/updateInvoiceStatus.generated';
import {
  GetInvoiceQuery,
  useGetInvoiceQuery,
} from '@shared/graphql/getInvoice.generated';

export const StatusMenuButton = ({
  status,
  id,
}: {
  id: string;
  status?: InvoiceStatus | null;
}) => {
  const client = getGraphQLClient();
  const queryClient = useQueryClient();
  const [invoicesMeta] = useInvoicesMeta();
  const queryKey = useGetInvoiceQuery.getKey({ id });
  const invoicesList = useInfiniteGetInvoicesQuery.getKey({
    ...invoicesMeta.getInvoices,
  });
  const { mutate: updateInvoiceStatus } = useUpdateInvoiceStatusMutation(
    client,
    {
      onMutate: ({ input }) => {
        const prevData = queryClient.getQueryData<GetInvoiceQuery>(queryKey);
        const prevListData =
          queryClient.getQueryData<GetInvoiceQuery>(invoicesList);
        useGetInvoiceQuery.mutateCacheEntry(queryClient, { id })(
          (cacheEntry) => {
            return produce(cacheEntry, (draft) => {
              draft['invoice']['status'] = input.status;
            });
          },
        );
        useInfiniteGetInvoicesQuery.mutateCacheEntry(queryClient, {
          ...invoicesMeta.getInvoices,
        })((cacheEntry) => {
          return produce(cacheEntry, (draft) => {
            draft?.pages.map((page, index) => {
              const selectedProfile = page.invoices?.content?.findIndex(
                (invoice) => invoice.metadata.id === id,
              );
              if (
                selectedProfile >= 0 &&
                page?.invoices?.content?.[selectedProfile]
              ) {
                draft.pages[index].invoices.content[selectedProfile] = {
                  ...draft.pages[index].invoices.content[selectedProfile],
                  status: input.status,
                };
              }
            });
          });
        });

        return { prevData, prevListData };
      },
      onSuccess: (data, variables, context) => {
        queryClient.invalidateQueries({ queryKey });
        queryClient.invalidateQueries({ queryKey: invoicesList });
      },
      onError: (error, _, context) => {
        queryClient.setQueryData<GetInvoiceQuery>(queryKey, context?.prevData);
        queryClient.setQueryData(invoicesList, context?.prevListData);
      },
    },
  );
  // this is needed cause updating status to void with updateInvoiceStatus mutation does not work correctly
  const { mutate: voidInvoice } = useVoidInvoiceMutation(client, {
    onMutate: () => {
      const prevData = queryClient.getQueryData<GetInvoiceQuery>(queryKey);
      const prevListData =
        queryClient.getQueryData<GetInvoiceQuery>(invoicesList);
      useGetInvoiceQuery.mutateCacheEntry(queryClient, { id })((cacheEntry) => {
        return produce(cacheEntry, (draft) => {
          draft['invoice']['status'] = InvoiceStatus.Void;
        });
      });
      useInfiniteGetInvoicesQuery.mutateCacheEntry(queryClient, {
        ...invoicesMeta.getInvoices,
      })((cacheEntry) => {
        return produce(cacheEntry, (draft) => {
          draft?.pages.map((page, index) => {
            const selectedProfile = page.invoices?.content?.findIndex(
              (invoice) => invoice.metadata.id === id,
            );
            if (
              selectedProfile >= 0 &&
              page?.invoices?.content?.[selectedProfile]
            ) {
              draft.pages[index].invoices.content[selectedProfile] = {
                ...draft.pages[index].invoices.content[selectedProfile],
                status: InvoiceStatus.Void,
              };
            }
          });
        });
      });

      return { prevData, prevListData };
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey });
      queryClient.invalidateQueries({ queryKey: invoicesList });
    },
    onError: (error, _, context) => {
      queryClient.setQueryData<GetInvoiceQuery>(queryKey, context?.prevData);
      queryClient.setQueryData(invoicesList, context?.prevListData);
    },
  });

  const handleUpdateStatus = (newStatus: InvoiceStatus) => {
    updateInvoiceStatus({
      input: {
        id,
        status: newStatus,
        patch: true,
      },
    });
  };

  return (
    <Menu>
      <MenuButton as={Flex} align='center' aria-label='Status' border='none'>
        {renderStatusNode(status)}
      </MenuButton>
      <MenuList>
        {status !== InvoiceStatus.Paid && (
          <MenuItem
            alignItems='center'
            color='gray.700'
            onClick={() => handleUpdateStatus(InvoiceStatus.Paid)}
          >
            <CheckCircle mr={2} color='gray.500' />
            Paid
          </MenuItem>
        )}
        {status !== InvoiceStatus.Void && (
          <MenuItem
            alignItems='center'
            color='gray.700'
            onClick={() => voidInvoice({ id })}
          >
            <SlashCircle01 mr={2} color='gray.500' />
            Void
          </MenuItem>
        )}

        {status !== InvoiceStatus.Due && (
          <MenuItem
            alignItems='center'
            color='gray.700'
            onClick={() => handleUpdateStatus(InvoiceStatus.Due)}
          >
            <Clock mr={2} color='gray.500' />
            Due
          </MenuItem>
        )}
      </MenuList>
    </Menu>
  );
};
