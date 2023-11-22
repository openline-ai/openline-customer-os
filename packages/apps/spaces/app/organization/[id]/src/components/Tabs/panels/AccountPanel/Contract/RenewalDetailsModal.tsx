'use client';
// TODO uncomment when forecast
import { useParams } from 'next/navigation';
import { useRef, useMemo, useState, useEffect } from 'react';

import { produce } from 'immer';
import { useQueryClient } from '@tanstack/react-query';

import { Dot } from '@ui/media/Dot';
import { Box } from '@ui/layout/Box';
import { Flex } from '@ui/layout/Flex';
import { Text } from '@ui/typography/Text';
import { FeaturedIcon } from '@ui/media/Icon';
import { Heading } from '@ui/typography/Heading';
import { FormSelect } from '@ui/form/SyncSelect';
import { toastError } from '@ui/presentation/Toast';
import { Button, ButtonGroup } from '@ui/form/Button';
import { AutoresizeTextarea } from '@ui/form/Textarea';
import { CurrencyInput } from '@ui/form/CurrencyInput';
import { CurrencyDollar } from '@ui/media/icons/CurrencyDollar';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { ClockFastForward } from '@ui/media/icons/ClockFastForward';
import { Opportunity, OpportunityRenewalLikelihood } from '@graphql/types';
import { useGetUsersQuery } from '@organizations/graphql/getUsers.generated';
import { useUpdateOpportunityMutation } from '@organization/src/graphql/updateOpportunity.generated';
import {
  GetContractsQuery,
  useGetContractsQuery,
} from '@organization/src/graphql/getContracts.generated';
import {
  Modal,
  ModalBody,
  ModalFooter,
  ModalHeader,
  ModalContent,
  ModalOverlay,
  ModalCloseButton,
} from '@ui/overlay/Modal';

interface RenewalDetailsProps {
  isOpen: boolean;
  data: Opportunity;
  onClose: () => void;
}

export const RenewalDetailsModal = ({
  data,
  isOpen,
  onClose,
}: RenewalDetailsProps) => {
  const orgId = useParams()?.id as string;

  const client = getGraphQLClient();
  const queryClient = useQueryClient();

  const formId = 'renewalDetailsForm';
  const timeoutRef = useRef<NodeJS.Timeout | null>(null);
  const [probability, setLikelihood] = useState<
    OpportunityRenewalLikelihood | undefined | null
  >((data?.renewalLikelihood as OpportunityRenewalLikelihood) ?? null);
  const [amount, setAmount] = useState<string>(data?.amount?.toString() || '');
  const [reason, setReason] = useState<string>(data?.generalNotes || '');
  const [owner, setOwner] = useState<null | { value: string; label: string }>(
    null,
  );
  const { data: usersData } = useGetUsersQuery(client, {
    pagination: {
      limit: 50,
      page: 1,
    },
  });

  const options = useMemo(() => {
    return usersData?.users?.content
      ?.filter((e) => Boolean(e.firstName) || Boolean(e.lastName))
      ?.map((o) => ({
        value: o.id,
        label: `${o.firstName} ${o.lastName}`.trim(),
      }));
  }, [usersData?.users?.content?.length]);

  const getContractsQueryKey = useGetContractsQuery.getKey({
    id: orgId,
  });

  const updateOpportunityMutation = useUpdateOpportunityMutation(client, {
    onMutate: ({ input }) => {
      queryClient.cancelQueries(getContractsQueryKey);

      queryClient.setQueryData<GetContractsQuery>(
        getContractsQueryKey,
        (currentCache) => {
          if (!currentCache || !currentCache?.organization) return;

          return produce(currentCache, (draft) => {
            if (draft?.['organization']?.['contracts']) {
              draft['organization']['contracts']?.map((contractData, index) => {
                return (contractData.opportunities ?? []).map((opportunity) => {
                  if ((opportunity as Opportunity).id === input.opportunityId) {
                    return {
                      ...opportunity,
                      amount: input.amount,
                      renewalUpdatedByUserAt: new Date().toISOString(),
                    };
                  }

                  return opportunity;
                });
              });
            }
          });
        },
      );
      const previousEntries =
        queryClient.getQueryData<GetContractsQuery>(getContractsQueryKey);

      return { previousEntries };
    },
    onError: (_, __, context) => {
      queryClient.setQueryData<GetContractsQuery>(
        getContractsQueryKey,
        context?.previousEntries,
      );
      toastError(
        'Failed to update renewal details',
        'update-renewal-details-error',
      );
    },
    onSettled: () => {
      onClose();

      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
      timeoutRef.current = setTimeout(() => {
        queryClient.invalidateQueries(getContractsQueryKey);
      }, 1000);
    },
  });

  const handleSet = () => {
    updateOpportunityMutation.mutate({
      input: {
        opportunityId: data.id,
        generalNotes: reason,
        amount: parseFloat(amount),
      },
    });
  };

  useEffect(() => {
    return () => {
      if (timeoutRef.current) {
        clearTimeout(timeoutRef.current);
      }
    };
  }, []);

  return (
    <Modal isOpen={isOpen} onClose={onClose}>
      <ModalOverlay />
      <ModalContent
        borderRadius='2xl'
        backgroundImage='/backgrounds/organization/circular-bg-pattern.png'
        backgroundRepeat='no-repeat'
        sx={{
          backgroundPositionX: '1px',
          backgroundPositionY: '-7px',
        }}
      >
        <ModalCloseButton />
        <ModalHeader>
          <FeaturedIcon size='lg' colorScheme='primary'>
            <ClockFastForward />
          </FeaturedIcon>
          <Heading fontSize='lg' mt='4'>
            Renewal details
          </Heading>
        </ModalHeader>
        <ModalBody as={Flex} flexDir='column' pb='0' gap={4}>
          <FormSelect
            formId={formId}
            name='arrForecast'
            isClearable
            isDisabled
            value={owner}
            isLoading={false}
            placeholder='Owner'
            backspaceRemovesValue
            onChange={setOwner}
            options={options}
            label='Owner'
            isLabelVisible
          />

          <div>
            <Text
              fontWeight='semibold'
              fontSize='sm'
              mb={2}
              id='likelihood-options-button'
            >
              Likelihood
            </Text>
            <ButtonGroup
              w='full'
              isAttached
              aria-describedby='likelihood-oprions-button'
            >
              <Button
                variant='outline'
                w='full'
                leftIcon={<Dot />}
                onClick={() =>
                  setLikelihood(OpportunityRenewalLikelihood.ZeroRenewal)
                }
                bg={
                  probability === OpportunityRenewalLikelihood.ZeroRenewal
                    ? 'gray.100'
                    : 'white'
                }
              >
                Zero
              </Button>
              <Button
                w='full'
                variant='outline'
                leftIcon={<Dot colorScheme='error' />}
                onClick={() =>
                  setLikelihood(OpportunityRenewalLikelihood.LowRenewal)
                }
                bg={
                  probability === OpportunityRenewalLikelihood.LowRenewal
                    ? 'gray.100'
                    : 'white'
                }
              >
                Low
              </Button>
              <Button
                w='full'
                variant='outline'
                leftIcon={<Dot colorScheme='warning' />}
                onClick={() =>
                  setLikelihood(OpportunityRenewalLikelihood.MediumRenewal)
                }
                bg={
                  probability === OpportunityRenewalLikelihood.MediumRenewal
                    ? 'gray.100'
                    : 'white'
                }
              >
                Medium
              </Button>
              <Button
                w='full'
                variant='outline'
                leftIcon={<Dot colorScheme='success' />}
                onClick={() =>
                  setLikelihood(OpportunityRenewalLikelihood.HighRenewal)
                }
                bg={
                  probability === OpportunityRenewalLikelihood.HighRenewal
                    ? 'gray.100'
                    : 'white'
                }
              >
                High
              </Button>
            </ButtonGroup>
            {data?.renewalUpdatedByUserId && (
              <Text color='gray.500' fontSize='xs' mt={2}>
                Last updated by{' '}
              </Text>
            )}
          </div>

          {data?.amount && (
            <CurrencyInput
              w='full'
              placeholder='Amount'
              label='ARR forecast'
              isLabelVisible
              value={amount}
              onChange={(value) => setAmount(value)}
              min={0}
              leftElement={
                <Box color='gray.500'>
                  <CurrencyDollar height='16px' />
                </Box>
              }
            />
          )}

          {!!probability && (
            <div>
              <Text as='label' htmlFor='reason' fontSize='sm'>
                <b>Reason for change</b> (optional)
              </Text>
              <AutoresizeTextarea
                pt='0'
                id='reason'
                value={reason}
                spellCheck='false'
                onChange={(e) => setReason(e.target.value)}
                placeholder={`What is the reason for updating these details`}
              />
            </div>
          )}
        </ModalBody>
        <ModalFooter p='6'>
          <Button variant='outline' w='full' onClick={onClose}>
            Cancel
          </Button>
          <Button
            ml='3'
            w='full'
            variant='outline'
            colorScheme='primary'
            onClick={handleSet}
          >
            Update
          </Button>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};
