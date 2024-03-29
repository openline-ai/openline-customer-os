import { useForm } from 'react-inverted-form';
import React, { useRef, useState, useEffect } from 'react';

import { produce } from 'immer';
import { useQueryClient } from '@tanstack/react-query';
import { useDebounce, useDeepCompareEffect } from 'rooks';

import { Flex } from '@ui/layout/Flex';
import { useDisclosure } from '@ui/utils';
import { FormInput } from '@ui/form/Input';
import { Check } from '@ui/media/icons/Check';
import { File02 } from '@ui/media/icons/File02';
import { Edit03 } from '@ui/media/icons/Edit03';
import { FormSelect } from '@ui/form/SyncSelect';
import { IconButton } from '@ui/form/IconButton';
import { Heading } from '@ui/typography/Heading';
import { DateTimeUtils } from '@spaces/utils/date';
import { Collapse } from '@ui/transitions/Collapse';
import { toastError } from '@ui/presentation/Toast';
import { DatePicker } from '@ui/form/DatePicker/DatePicker';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { Card, CardBody, CardFooter, CardHeader } from '@ui/presentation/Card';
import { useUpdateContractMutation } from '@organization/src/graphql/updateContract.generated';
import {
  Contract,
  ContractStatus,
  ContractUpdateInput,
  ContractRenewalCycle,
} from '@graphql/types';
import {
  GetContractsQuery,
  useGetContractsQuery,
} from '@organization/src/graphql/getContracts.generated';
import { ContractSubtitle } from '@organization/src/components/Tabs/panels/AccountPanel/Contract/ContractSubtitle';
import { useUpdatePanelModalStateContext } from '@organization/src/components/Tabs/panels/AccountPanel/context/AccountModalsContext';

import { Services } from './Services/Services';
import { FormPeriodInput } from './PeriodInput';
import { RenewalARRCard } from './RenewalARR/RenewalARRCard';
import { ServiceLineItemsModal } from './ServiceLineItemsModal';
import { ContractDTO, TimeToRenewalForm } from './Contract.dto';
import { ContractBillingDetailsModal } from './ContractBillingDetailsModal';
import { ContractStatusSelect } from './contractStatuses/ContractStatusSelect';
import {
  paymentDueOptions,
  autorenewalOptions,
  billingFrequencyOptions,
  contractBillingCycleOptions,
} from '../utils';

interface ContractCardProps {
  data: Contract;
  organizationId: string;
  organizationName: string;
}

export const ContractCard = ({
  data,
  organizationName,
  organizationId,
}: ContractCardProps) => {
  const queryKey = useGetContractsQuery.getKey({ id: organizationId });
  const queryClient = useQueryClient();
  const timeoutRef = useRef<NodeJS.Timeout | null>(null);
  const [isExpanded, setIsExpanded] = useState(!data?.contractSigned);
  const formId = `contract-form-${data.metadata.id}`;
  const { setIsPanelModalOpen } = useUpdatePanelModalStateContext();
  const { onOpen, onClose, isOpen } = useDisclosure({
    id: 'billing-details-modal',
  });
  const {
    onOpen: onServiceLineItemsOpen,
    onClose: onServiceLineItemClose,
    isOpen: isServceItemsModalOpen,
  } = useDisclosure({
    id: 'service-line-items-modal',
  });

  const client = getGraphQLClient();

  // this is needed to block scroll on safari when modal is open, scrollbar overflow issue
  useEffect(() => {
    if (isOpen || isServceItemsModalOpen) {
      setIsPanelModalOpen(true);
    }
    if (!isOpen && !isServceItemsModalOpen) {
      setIsPanelModalOpen(false);
    }
  }, [isOpen, isServceItemsModalOpen]);

  const updateContract = useUpdateContractMutation(client, {
    onMutate: ({ input: { patch, contractId, ...input } }) => {
      queryClient.cancelQueries({ queryKey });
      queryClient.setQueryData<GetContractsQuery>(queryKey, (currentCache) => {
        return produce(currentCache, (draft) => {
          const previousContracts = draft?.['organization']?.['contracts'];
          const updatedContractIndex = previousContracts?.findIndex(
            (contract) => contract.metadata.id === data?.metadata?.id,
          );
          if (draft?.['organization']?.['contracts']) {
            draft['organization']['contracts']?.map((contractData, index) => {
              if (index !== updatedContractIndex) {
                return contractData;
              }
              const result = Object.entries(input).find(
                ([_, value]) => value === '0001-01-01T00:00:00.000000Z',
              );

              return {
                ...contractData,
                ...input,
                ...(result ? { [result[0]]: null } : {}),
              };
            });
          }
        });
      });
      const previousEntries =
        queryClient.getQueryData<GetContractsQuery>(queryKey);

      return { previousEntries };
    },
    onError: (error, { input }, context) => {
      queryClient.setQueryData<GetContractsQuery>(
        queryKey,
        context?.previousEntries,
      );

      const invalidDate =
        DateTimeUtils.isBefore(input.contractEnded, input.serviceStarted) ||
        DateTimeUtils.isBefore(input.contractEnded, input.contractSigned);

      toastError(
        `${
          invalidDate
            ? 'The contract must end after the service start or signing date'
            : 'Failed to update contract'
        }`,
        `update-contract-error-${error}`,
      );
    },
    onSettled: () => {
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }

      timeoutRef.current = setTimeout(() => {
        queryClient.invalidateQueries({ queryKey });

        queryClient.invalidateQueries({ queryKey: ['GetTimeline.infinite'] });
      }, 1000);
    },
  });

  const updateContractDebounced = useDebounce(
    (variables: { input: ContractUpdateInput }) => {
      updateContract.mutate({
        ...variables,
        input: {
          ...variables.input,
          patch: true,
        },
      });
    },
    500,
  );
  const defaultValues = ContractDTO.toForm(data ?? {});
  const { setDefaultValues, state } = useForm<TimeToRenewalForm>({
    formId,
    defaultValues,
    debug: true,
    stateReducer: (state, action, next) => {
      if (action.type === 'FIELD_CHANGE') {
        switch (action.payload.name) {
          case 'renewalPeriods':
            return next;
          case 'name': {
            updateContractDebounced(
              ContractDTO.toPayload({
                contractId: data.metadata.id,
                name: action.payload.value,
              }),
            );

            return next;
          }
          case 'contractRenewalCycle': {
            let renewalPeriods = '1';

            if (action.payload.value.value === 'MULTI_YEAR') {
              renewalPeriods = '2';
            }

            updateContract.mutate(
              ContractDTO.toPayload({
                contractId: data.metadata.id,
                contractRenewalCycle:
                  state.values.contractRenewalCycle?.value === 'MULTI_YEAR'
                    ? ContractRenewalCycle.AnnualRenewal
                    : state.values.contractRenewalCycle?.value,
                renewalPeriods,
              }),
            );

            return {
              ...next,
              values: {
                ...next.values,
                renewalPeriods,
              },
            };
          }
          case 'serviceStarted':
          case 'endedAt':
          case 'invoicingStartDate':
            updateContract.mutate(
              ContractDTO.toPayload({
                contractId: data.metadata.id,
                [action.payload.name]: action.payload.value
                  ? action.payload.value
                  : '0001-01-01T00:00:00.000000Z',
              }),
            );

            return {
              ...next,
              values: {
                ...next.values,
                [action.payload.name]: action.payload.value ?? null,
              },
            };
          case 'autoRenew':
          case 'billingCycle':
          case 'billingEnabled':
            updateContract.mutate(
              ContractDTO.toPayload({
                contractId: data.metadata.id,
                [action.payload.name]: action.payload.value?.value,
              }),
            );

            return next;
          case 'dueDays':
            updateContract.mutate(
              ContractDTO.toPayload({
                contractId: data.metadata.id,
                billingDetails: {
                  dueDays: action.payload.value?.value,
                },
              }),
            );

            return next;
          case 'contractUrl':
            updateContractDebounced(
              ContractDTO.toPayload({
                contractId: data.metadata.id,
                contractUrl: action.payload.value,
              }),
            );

            return next;
          default: {
            return next;
          }
        }
      }

      if (action.type === 'FIELD_BLUR') {
        if (action.payload.name === 'renewalPeriods') {
          updateContract.mutate(
            ContractDTO.toPayload({
              contractId: data.metadata.id,
              renewalPeriods:
                state.values?.contractRenewalCycle?.value === 'MULTI_YEAR'
                  ? parseInt(action.payload?.value || '2')
                  : action.payload?.value
                  ? parseInt(action.payload?.value)
                  : undefined,
            }),
          );

          return {
            ...next,
            values: {
              ...next.values,
              renewalPeriods: action.payload?.value || '2',
            },
          };
        }
      }

      return next;
    },
  });

  useDeepCompareEffect(() => {
    setDefaultValues(defaultValues);
  }, [defaultValues]);

  useEffect(() => {
    return () => {
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      updateContractDebounced.flush();
    };
  }, []);

  return (
    <Card
      as='section'
      px='4'
      py='3'
      w='full'
      size='lg'
      variant='outline'
      cursor='default'
      border='1px solid'
      borderColor='gray.200'
      bg='gray.50'
      transition='all 0.2s ease-out'
    >
      <CardHeader
        as={Flex}
        p='0'
        role='button'
        pb={isExpanded ? 2 : 0}
        w='full'
        flexDir='column'
        _hover={
          !isExpanded
            ? {
                '#edit-contract-icon': {
                  opacity: 1,
                  transition: 'opacity 0.2s linear',
                },
              }
            : {}
        }
        sx={
          !isExpanded
            ? {
                '#edit-contract-icon': {
                  opacity: 0,
                  transition: 'opacity 0.2s linear',
                },
              }
            : {}
        }
        onClick={() => (!isExpanded ? setIsExpanded(true) : null)}
      >
        <Flex justifyContent='space-between' w='full' flex={1}>
          <Heading
            size='sm'
            color='gray.700'
            noOfLines={1}
            lineHeight={1.4}
            display='inline'
            whiteSpace='nowrap'
          >
            {!isExpanded && state.values.name}

            {isExpanded && (
              <FormInput
                fontWeight='semibold'
                fontSize='inherit'
                height='fit-content'
                name='name'
                formId={formId}
                borderBottom='none'
                _hover={{
                  borderBottom: 'none',
                }}
              />
            )}
          </Heading>

          <Flex alignItems='center' gap={2} ml={2}>
            {!isExpanded && (
              <Edit03
                mr={1}
                color='gray.400'
                boxSize='4'
                id='edit-contract-icon'
              />
            )}

            <IconButton
              aria-label='Edit billing details'
              size='xs'
              variant='ghost'
              icon={<File02 color='gray.400' />}
              onClick={() => onOpen()}
            />

            <ContractStatusSelect
              status={data.contractStatus}
              contractId={data.metadata.id}
              renewsAt={data?.opportunities?.[0]?.renewedAt}
              onUpdateContract={updateContract}
              serviceStarted={data.serviceStarted}
              organizationName={
                data?.billingDetails?.organizationLegalName ||
                organizationName ||
                'Unnamed'
              }
              nextInvoiceDate={data?.billingDetails?.nextInvoicing}
            />

            {isExpanded && (
              <IconButton
                size='xs'
                variant='ghost'
                aria-label='Collapse'
                icon={<Check color='gray.400' />}
                onClick={() => setIsExpanded(false)}
              />
            )}
          </Flex>
        </Flex>

        {!isExpanded && (
          <Flex
            bg='transparent'
            _hover={{
              bg: 'transparent',
              svg: { opacity: 1, transition: 'opacity 0.2s linear' },
            }}
            sx={{ svg: { opacity: 0, transition: 'opacity 0.2s linear' } }}
            fontSize='sm'
            fontWeight='normal'
            color='gray.500'
            p={0}
            height='fit-content'
            alignItems='flex-start'
            justifyContent='flex-start'
          >
            <ContractSubtitle data={data} />
          </Flex>
        )}
      </CardHeader>
      {isExpanded && (
        <CardBody as={Flex} p='0' flexDir='column' w='full'>
          <Flex gap='4' mb={2} flexGrow={0}>
            <DatePicker
              label='Service starts'
              placeholder='Service starts date'
              formId={formId}
              name='serviceStarted'
              inset='120% auto auto 0px'
              calendarIconHidden
              value={state.values.serviceStarted}
            />
            <FormSelect
              label='Renewal term'
              placeholder='Contract renews'
              isLabelVisible
              name='autoRenew'
              formId={formId}
              options={autorenewalOptions}
              chakraStyles={{
                menuList: (props, state) => ({
                  ...props,
                  minW: '190px',
                }),
              }}
            />
          </Flex>
          <Flex gap='4' flexGrow={0} mb={2}>
            <FormSelect
              label='Contract term'
              placeholder='Contract term'
              isLabelVisible
              name='contractRenewalCycle'
              formId={formId}
              options={billingFrequencyOptions}
            />
            {state.values.contractRenewalCycle?.value === 'MULTI_YEAR' && (
              <FormPeriodInput
                formId={formId}
                label='Renews every'
                name='renewalPeriods'
                placeholder='Renews every'
              />
            )}
          </Flex>
          <Flex gap='4' flexGrow={0} mb={2}>
            <FormSelect
              label='Billing is'
              placeholder='Enable billing'
              isLabelVisible
              name='billingEnabled'
              formId={formId}
              options={[
                { label: 'Enabled', value: true },
                { label: 'Disabled', value: false },
              ]}
            />
            <FormSelect
              label='Billing period'
              placeholder='Billing period'
              isLabelVisible
              name='billingCycle'
              formId={formId}
              options={contractBillingCycleOptions}
            />
          </Flex>
          <Flex gap='4' flexGrow={0} mb={2}>
            <DatePicker
              label='Invoicing starts'
              placeholder='Invoicing starts'
              minDate={state.values.serviceStarted}
              formId={formId}
              name='invoicingStartDate'
              inset='120% auto auto 0px'
              calendarIconHidden
              value={state.values.invoicingStartDate}
            />
            <FormSelect
              label='Payment due'
              placeholder='Payment due'
              isLabelVisible
              name='dueDays'
              formId={formId}
              options={paymentDueOptions}
            />
          </Flex>
        </CardBody>
      )}
      <CardFooter p='0' mt={1} w='full' flexDir='column'>
        <Collapse
          delay={{ enter: 0.2 }}
          in={!!data?.opportunities && !!data.contractRenewalCycle}
          animateOpacity
          startingHeight={0}
        >
          {data?.opportunities && data.contractRenewalCycle && (
            <RenewalARRCard
              hasEnded={data.contractStatus === ContractStatus.Ended}
              startedAt={data.serviceStarted}
              renewCycle={data.contractRenewalCycle}
              currency={data.currency}
              opportunity={data.opportunities?.[0]}
            />
          )}
        </Collapse>
        <Services
          data={data?.contractLineItems}
          currency={data?.currency}
          onModalOpen={onServiceLineItemsOpen}
        />

        <ContractBillingDetailsModal
          isOpen={isOpen}
          contractId={data.metadata.id}
          onClose={onClose}
          organizationName={organizationName}
          notes={data?.billingDetails?.invoiceNote}
        />

        <ServiceLineItemsModal
          isOpen={isServceItemsModalOpen}
          contractId={data.metadata.id}
          onClose={onServiceLineItemClose}
          contractName={data.contractName}
          currency={data.currency}
          contractLineItems={data?.contractLineItems ?? []}
          organizationName={organizationName}
          notes={data?.billingDetails?.invoiceNote}
        />
      </CardFooter>
    </Card>
  );
};
