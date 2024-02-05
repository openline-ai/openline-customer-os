'use client';
import { useForm } from 'react-inverted-form';
import React, { useRef, useMemo, useState } from 'react';

import { useQueryClient } from '@tanstack/react-query';
import { useTenantBillingProfilesQuery } from '@settings/graphql/getTenantBillingProfiles.generated';

import { Box } from '@ui/layout/Box';
import { Button } from '@ui/form/Button';
import { FeaturedIcon } from '@ui/media/Icon';
import { File02 } from '@ui/media/icons/File02';
import { Grid, GridItem } from '@ui/layout/Grid';
import { Heading } from '@ui/typography/Heading';
import { Invoice } from '@shared/components/Invoice/Invoice';
import { countryOptions } from '@shared/util/countryOptions';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { toastError, toastSuccess } from '@ui/presentation/Toast';
import { useUpdateContractMutation } from '@organization/src/graphql/updateContract.generated';
import {
  Modal,
  ModalFooter,
  ModalHeader,
  ModalContent,
  ModalOverlay,
} from '@ui/overlay/Modal';
import {
  GetContractQuery,
  useGetContractQuery,
} from '@organization/src/graphql/getContract.generated';

import { BillingDetailsDto } from './BillingDetails.dto';
import { ContractBillingDetailsForm } from './ContractBillingDetailsForm';

interface SubscriptionServiceModalProps {
  isOpen: boolean;
  contractId: string;
  onClose: () => void;
  organizationName: string;
  data?: GetContractQuery['contract'] | null;
}
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

export const ContractBillingDetailsModal = ({
  isOpen,
  onClose,
  contractId,
  organizationName,
  data,
}: SubscriptionServiceModalProps) => {
  const initialRef = useRef(null);
  const formId = `billing-details-form-${contractId}`;

  const [isBillingDetailsFocused, setIsBillingDetailsFocused] =
    useState<boolean>(false);

  const [isBillingDetailsHovered, setIsBillingDetailsHovered] =
    useState<boolean>(false);
  const queryKey = useGetContractQuery.getKey({ id: contractId });

  const queryClient = useQueryClient();
  const client = getGraphQLClient();
  const timeoutRef = useRef<NodeJS.Timeout | null>(null);
  const { data: tenantBillingProfile } = useTenantBillingProfilesQuery(client);

  const updateContract = useUpdateContractMutation(client, {
    onError: (error) => {
      toastError(
        'Failed to update billing details',
        `update-contract-error-${error}`,
      );
    },
    onSuccess: () => {
      toastSuccess(
        'Billing details updated',
        `update-contract-success-${contractId}`,
      );
      onClose();
    },
    onSettled: () => {
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      timeoutRef.current = setTimeout(() => {
        queryClient.invalidateQueries({ queryKey });
      }, 1000);
    },
  });
  const defaultValues = new BillingDetailsDto({
    ...(data ?? {}),
    organizationLegalName: data?.organizationLegalName || organizationName,
  });

  const { state } = useForm({
    debug: true,
    formId,
    defaultValues,
    stateReducer: (_, action, next) => {
      if (action.type === 'FIELD_CHANGE') {
        if (action.payload.name === 'invoiceEmail') {
          return {
            ...next,
            values: {
              ...next.values,
              invoiceEmail: action.payload.value.split(' ').join('').trim(),
            },
          };
        }
      }

      return next;
    },
  });

  const handleApplyChanges = () => {
    const payload = BillingDetailsDto.toPayload(state.values);

    updateContract.mutate({
      input: {
        contractId,
        ...payload,
      },
    });
  };
  const invoicePreviewStaticData = useMemo(
    () => ({
      status: 'Preview',
      invoiceNumber: 'INV-003',
      lines: [
        {
          amount: 100,
          createdAt: new Date().toISOString(),
          id: 'dummy-id',
          name: 'Professional tier',
          price: 50,
          quantity: 2,
          totalAmount: 100,
          vat: 0,
        },
      ],
      tax: 0,
      note: '',
      total: 100,
      dueDate: new Date().toISOString(),
      subtotal: 100,
      issueDate: new Date().toISOString(),
      from: tenantBillingProfile?.tenantBillingProfiles?.[0]
        ? {
            addressLine1:
              tenantBillingProfile?.tenantBillingProfiles?.[0]?.addressLine1 ??
              '',
            addressLine2:
              tenantBillingProfile?.tenantBillingProfiles?.[0].addressLine2,
            locality:
              tenantBillingProfile?.tenantBillingProfiles?.[0]?.locality ?? '',
            zip: tenantBillingProfile?.tenantBillingProfiles?.[0]?.zip ?? '',
            country: tenantBillingProfile?.tenantBillingProfiles?.[0].country
              ? countryOptions.find(
                  (country) =>
                    country.value ===
                    tenantBillingProfile?.tenantBillingProfiles?.[0]?.country,
                )?.label
              : '',
            email: '',
            name: tenantBillingProfile?.tenantBillingProfiles?.[0]?.legalName,
          }
        : {
            addressLine1: '29 Maple Lane',
            addressLine2: 'Springfield, Haven County',
            locality: 'San Francisco',
            zip: '89302',
            country: 'United States',
            email: 'invoices@acme.com',
            name: 'Acme Corp.',
          },
    }),
    [tenantBillingProfile?.tenantBillingProfiles?.[0]],
  );
  const isEmailValid = useMemo(() => {
    return (
      !!state.values.invoiceEmail?.length &&
      !emailRegex.test(state.values.invoiceEmail)
    );
  }, [state?.values?.invoiceEmail]);

  return (
    <Modal
      isOpen={isOpen}
      onClose={onClose}
      initialFocusRef={initialRef}
      size='4xl'
    >
      <ModalOverlay />
      <ModalContent borderRadius='2xl'>
        <Grid h='100%' templateColumns='356px 1fr'>
          <GridItem
            rowSpan={1}
            colSpan={1}
            h='100%'
            display='flex'
            flexDir='column'
            justifyContent='space-between'
            bg='gray.25'
            borderRight='1px solid'
            borderColor='gray.200'
            borderTopLeftRadius='2xl'
            borderBottomLeftRadius='2xl'
            backgroundImage='/backgrounds/organization/circular-bg-pattern.png'
            backgroundRepeat='no-repeat'
            sx={{
              backgroundPositionX: '1px',
              backgroundPositionY: '-7px',
            }}
          >
            <ModalHeader>
              <FeaturedIcon size='lg' colorScheme='primary'>
                <File02 color='primary.600' />
              </FeaturedIcon>
              <Heading fontSize='lg' mt='4'>
                {data?.organizationLegalName ||
                  organizationName ||
                  "Unnamed's "}{' '}
                contract details
              </Heading>
            </ModalHeader>
            <ContractBillingDetailsForm
              formId={formId}
              isEmailValid={isEmailValid}
              onSetIsBillingDetailsHovered={setIsBillingDetailsHovered}
              onSetIsBillingDetailsFocused={setIsBillingDetailsFocused}
            />
            <ModalFooter p='6'>
              <Button variant='outline' w='full' onClick={onClose}>
                Cancel
              </Button>
              <Button
                ml='3'
                w='full'
                variant='outline'
                colorScheme='primary'
                loadingText='Applying changes...'
                onClick={handleApplyChanges}
              >
                Done
              </Button>
            </ModalFooter>
          </GridItem>
          <GridItem>
            <Box width='100%'>
              <Invoice
                isBilledToFocused={
                  isBillingDetailsFocused || isBillingDetailsHovered
                }
                currency={state?.values?.currency?.value}
                billedTo={{
                  addressLine1: state.values.addressLine1 ?? '',
                  addressLine2: state.values.addressLine2 ?? '',
                  locality: state.values.locality ?? '',
                  zip: state.values.zip ?? '',
                  country: state?.values?.country?.label ?? '',
                  email: state.values.invoiceEmail ?? '',
                  name: state.values?.organizationLegalName ?? '',
                }}
                {...invoicePreviewStaticData}
              />
            </Box>
          </GridItem>
        </Grid>
      </ModalContent>
    </Modal>
  );
};