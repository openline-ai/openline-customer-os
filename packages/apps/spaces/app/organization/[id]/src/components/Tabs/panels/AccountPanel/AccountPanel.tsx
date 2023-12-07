'use client';

import { useParams } from 'next/navigation';
import React, { FC, PropsWithChildren } from 'react';

import { Box } from '@ui/layout/Box';
import { Flex } from '@ui/layout/Flex';
import { Contract } from '@graphql/types';
import { Select } from '@ui/form/SyncSelect';
import { ActivityHeart } from '@ui/media/icons/ActivityHeart';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { useGetContractsQuery } from '@organization/src/graphql/getContracts.generated';
import { contractButtonSelect } from '@organization/src/components/Tabs/shared/contractSelectStyles';
import { ARRForecast } from '@organization/src/components/Tabs/panels/AccountPanel/ARRForecast/ARRForecast';

import { Notes } from './Notes';
import { EmptyContracts } from './EmptyContracts';
import { ContractCard } from './Contract/ContractCard';
import { AccountPanelSkeleton } from './AccountPanelSkeleton';
import { OrganizationPanel } from '../OrganizationPanel/OrganizationPanel';
import {
  useAccountPanelStateContext,
  AccountModalsContextProvider,
} from './context/AccountModalsContext';

const AccountPanelComponent = () => {
  const client = getGraphQLClient();
  const id = useParams()?.id as string;

  const { isModalOpen } = useAccountPanelStateContext();
  const { data, isInitialLoading } = useGetContractsQuery(client, {
    id,
  });

  if (isInitialLoading) {
    return <AccountPanelSkeleton />;
  }

  if (!data?.organization?.contracts?.length) {
    return (
      <EmptyContracts name={data?.organization?.name || ''}>
        <Notes id={id} data={data?.organization} />
      </EmptyContracts>
    );
  }

  return (
    <OrganizationPanel
      title='Account'
      withFade
      actionItem={
        <Box display='none'>
          <Select
            isSearchable={false}
            isClearable={false}
            isMulti={false}
            value={{
              label: 'Customer',
              value: 'customer',
            }}
            options={[
              {
                label: 'Customer',
                value: 'customer',
              },
              {
                label: 'Prospect',
                value: 'prospect',
              },
            ]}
            chakraStyles={{
              ...contractButtonSelect,
              container: (props, state) => {
                const isCustomer = state.getValue()[0]?.value === 'customer';

                return {
                  ...props,
                  px: 2,
                  pointerEvents: 'none',
                  py: '1px',
                  border: '1px solid',
                  borderColor: isCustomer ? 'success.200' : 'gray.300',
                  backgroundColor: isCustomer ? 'success.50' : 'transparent',
                  color: isCustomer ? 'success.700' : 'gray.500',

                  borderRadius: '2xl',
                  fontSize: 'xs',
                  maxHeight: '22px',

                  '& > div': {
                    p: 0,
                    border: 'none',
                    fontSize: 'xs',
                    maxHeight: '22px',
                    minH: 'auto',
                  },
                };
              },
              valueContainer: (props, state) => {
                const isCustomer = state.getValue()[0]?.value === 'customer';

                return {
                  ...props,
                  p: 0,
                  border: 'none',
                  fontSize: 'xs',
                  maxHeight: '22px',
                  minH: 'auto',
                  color: isCustomer ? 'success.700' : 'gray.500',
                };
              },
              singleValue: (props) => {
                return {
                  ...props,
                  maxHeight: '22px',
                  p: 0,
                  minH: 'auto',
                  color: 'inherit',
                };
              },
              menuList: (props) => {
                return {
                  ...props,
                  w: 'fit-content',
                  left: '-32px',
                };
              },
            }}
            leftElement={<ActivityHeart color='success.500' mr='1' />}
          />
        </Box>
      }
      shouldBlockPanelScroll={isModalOpen}
    >
      {!!data?.organization?.contracts && (
        <>
          <ARRForecast
            renewalSunnary={data?.organization?.accountDetails?.renewalSummary}
            name={data?.organization?.name || ''}
            isInitialLoading={isInitialLoading}
          />
          {data?.organization?.contracts.map((contract) => (
            <Flex
              key={`contract-card-${contract.id}`}
              flexDir='column'
              gap={4}
              mb={4}
            >
              <ContractCard
                organizationId={id}
                organizationName={data?.organization?.name ?? ''}
                data={(contract as Contract) ?? undefined}
              />
            </Flex>
          ))}
        </>
      )}

      <Notes id={id} data={data?.organization} />
    </OrganizationPanel>
  );
};

export const AccountPanel: FC<PropsWithChildren> = () => (
  <AccountModalsContextProvider>
    <AccountPanelComponent />
  </AccountModalsContextProvider>
);
