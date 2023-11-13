'use client';

import React, { useMemo } from 'react';
// import { useMemo } from 'react';
import { useParams } from 'next/navigation';

import { Box } from '@ui/layout/Box';
import { useDisclosure } from '@ui/utils';
import { Select } from '@ui/form/SyncSelect';
import { ActivityHeart } from '@ui/media/icons/ActivityHeart';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
// import { EmptyContracts } from '@organization/src/components/Tabs/panels/AccountPanel/EmptyContracts';
import { ContractCard } from '@organization/src/components/Tabs/panels/AccountPanel/Contract/ContractCard';
import { ARRForecast } from '@organization/src/components/Tabs/panels/AccountPanel/ARRForecast/ARRForecast';
import { useOrganizationAccountDetailsQuery } from '@organization/src/graphql/getAccountPanelDetails.generated';

import { Notes } from './Notes';
import { RenewalForecastType } from './RenewalForecast';
import { AccountPanelSkeleton } from './AccountPanelSkeleton';
import { OrganizationPanel } from '../OrganizationPanel/OrganizationPanel';

export const AccountPanel = () => {
  const id = useParams()?.id as string;
  // Moved to upperscope due to error in safari https://linear.app/customer-os/issue/COS-619/scrollbar-overlaps-the-renewal-modals-in-safari

  // Todo modify and connect modals
  // const renewalLikelihoodUpdateModal = useDisclosure({
  //   id: 'renewal-likelihood-update-modal',
  // });
  const arrForecastInfoModal = useDisclosure({
    id: 'arr-forecast-info-modal',
  });
  const addServiceInfoModal = useDisclosure({
    id: 'add-service-info-modal',
  });

  // const renewalForecastUpdateModal = useDisclosure({
  //   id: 'renewal-renewal-update-modal',
  // });
  // const renewalForecastInfoModal = useDisclosure({
  //   id: 'renewal-renewal-info-modal',
  // });

  const client = getGraphQLClient();
  const { data, isInitialLoading } = useOrganizationAccountDetailsQuery(
    client,
    { id },
  );
  const isModalOpen = useMemo(() => {
    return arrForecastInfoModal.isOpen || addServiceInfoModal.isOpen;
  }, [arrForecastInfoModal.isOpen, addServiceInfoModal.isOpen]);

  if (isInitialLoading || !arrForecastInfoModal) {
    return <AccountPanelSkeleton />;
  }

  // TODO uncomment after integrating with BE
  // if (no contracts available) {
  //   return <EmptyContracts name={data?.organization?.name || ''} />;
  // }

  return (
    <OrganizationPanel
      title='Account'
      withFade
      actionItem={
        <Box>
          <Select
            isSearchable={false}
            isClearable={false}
            isMulti={false}
            value={{
              label: 'Customer',
              value: 'customer',
            }}
            onChange={(e) => {}}
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
              container: (props, state) => {
                const isCustomer = state.getValue()[0]?.value === 'customer';

                return {
                  ...props,
                  px: 2,
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
              input: (props) => {
                return {
                  ...props,
                  maxHeight: '22px',
                  minH: 'auto',
                  p: 0,
                };
              },
              inputContainer: (props) => {
                return {
                  ...props,
                  maxHeight: '22px',
                  minH: 'auto',
                  p: 0,
                };
              },

              control: (props) => {
                return {
                  ...props,
                  w: '100%',
                  border: 'none',
                };
              },

              menuList: (props) => {
                return {
                  ...props,
                  w: 'fit-content',
                  left: '-32px',
                  // top: '-10px',
                };
              },
            }}
            leftElement={<ActivityHeart color='success.500' mr='1' />}
          />
        </Box>
      }
      shouldBlockPanelScroll={isModalOpen}
    >
      <ARRForecast
        infoModal={arrForecastInfoModal}
        name={data?.organization?.name || ''}
        isInitialLoading={isInitialLoading}
        renewalProbability={
          data?.organization?.accountDetails?.renewalLikelihood?.probability
        }
        aRRForecast={
          data?.organization?.accountDetails
            ?.renewalForecast as RenewalForecastType
        }
      />

      <ContractCard
        name={data?.organization?.name || ''}
        data={null}
        serviceModal={addServiceInfoModal}
      />

      <Notes id={id} data={data?.organization} />
    </OrganizationPanel>
  );
};
