'use client';

import { useForm } from 'react-inverted-form';
import React, { useMemo, useState, useEffect } from 'react';

import { produce } from 'immer';
import { ScaleFade } from '@chakra-ui/transition';
import { useDebounce, useWillUnmount } from 'rooks';
import { useQueryClient } from '@tanstack/react-query';
import { VatInput } from '@settings/components/Tabs/panels/BillingPanel/VatInput';
import { LogoUploader } from '@settings/components/LogoUploadComponent/LogoUploader';
import { useTenantSettingsQuery } from '@settings/graphql/getTenantSettings.generated';
import { useUpdateTenantSettingsMutation } from '@settings/graphql/updateTenantSettings.generated';
import { useTenantBillingProfilesQuery } from '@settings/graphql/getTenantBillingProfiles.generated';
import { useCreateBillingProfileMutation } from '@settings/graphql/createTenantBillingProfile.generated';
import { useTenantUpdateBillingProfileMutation } from '@settings/graphql/updateTenantBillingProfile.generated';
import { TenantBillingPanelDetailsForm } from '@settings/components/Tabs/panels/BillingPanel/TenantBillingDetailsForm';

import { Box } from '@ui/layout/Box';
import { Flex } from '@ui/layout/Flex';
import { Gb } from '@ui/media/logos/Gb';
import { Us } from '@ui/media/logos/Us';
import { Eu } from '@ui/media/logos/Eu';
import { Text } from '@ui/typography/Text';
import { Switch } from '@ui/form/Switch';
import { FormInput } from '@ui/form/Input';
import { FormSelect } from '@ui/form/SyncSelect';
import { Heading } from '@ui/typography/Heading';
import { Divider } from '@ui/presentation/Divider';
import { TenantBillingProfile } from '@graphql/types';
import { FormSwitch } from '@ui/form/Switch/FromSwitch';
import { Invoice } from '@shared/components/Invoice/Invoice';
import { Card, CardBody, CardHeader } from '@ui/layout/Card';
import { countryOptions } from '@shared/util/countryOptions';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';

import {
  TenantBillingDetails,
  TenantBillingDetailsDto,
} from './TenantBillingProfile.dto';
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

export const BillingPanel = () => {
  const client = getGraphQLClient();
  const queryClient = useQueryClient();

  const { data: tenantSettingsData } = useTenantSettingsQuery(client);
  const queryKey = useTenantSettingsQuery.getKey();

  const updateTenantSettingsMutation = useUpdateTenantSettingsMutation(client, {
    onMutate: ({ input: { patch, ...newSettings } }) => {
      queryClient.cancelQueries({ queryKey });
      console.log('🏷️ ----- : here ');
      const previousSettings = tenantSettingsData?.tenantSettings;
      queryClient.setQueryData(queryKey, {
        tenantSettings: {
          ...previousSettings,
          ...newSettings,
        },
      });

      return { previousSettings };
    },
    onError: (err, newSettings, context) => {
      queryClient.setQueryData(queryKey, context?.previousSettings);
    },
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey });
    },
  });

  const isInvoicingEnabled =
    tenantSettingsData?.tenantSettings?.invoicingEnabled;
  const [isInvoiceProviderFocused, setIsInvoiceProviderFocused] =
    useState<boolean>(false);
  const [isInvoiceProviderDetailsHovered, setIsInvoiceProviderDetailsHovered] =
    useState<boolean>(false);

  const tenantBillingProfileId = data?.tenantBillingProfiles?.[0]?.id ?? '';
  const queryKey = useTenantBillingProfilesQuery.getKey();

  const createBillingProfileMutation = useCreateBillingProfileMutation(client, {
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey });
    },
  });
  const updateBillingProfileMutation = useTenantUpdateBillingProfileMutation(
    client,
    {
      onMutate: ({ input }) => {
        queryClient.cancelQueries({ queryKey });

        useTenantBillingProfilesQuery.mutateCacheEntry(queryClient)(
          (cacheEntry) => {
            return produce(cacheEntry, (draft) => {
              const selectedProfile = draft?.tenantBillingProfiles?.findIndex(
                (profileId) =>
                  profileId.id === data?.tenantBillingProfiles?.[0]?.id,
              );

              if (
                selectedProfile &&
                draft?.tenantBillingProfiles?.[selectedProfile]
              ) {
                draft.tenantBillingProfiles[selectedProfile] = {
                  ...draft.tenantBillingProfiles[selectedProfile],
                  ...(input as TenantBillingProfile),
                };
              }
            });
          },
        );
      },
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey });
      },
    },
  );
  const formId = 'tenant-billing-profile-form';
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
      billedTo: {
        addressLine1: '29 Maple Lane',
        addressLine2: 'Springfield, Haven County',
        locality: 'San Francisco',
        zip: '89302',
        country: 'United States',
        email: 'invoices@acme.com',
        name: 'Acme Corp.',
      },
    }),
    [],
  );

  const newDefaults = new TenantBillingDetailsDto();

  const handleUpdateData = useDebounce((d: TenantBillingDetails) => {
    const payload = TenantBillingDetailsDto.toPayload(d);
    updateBillingProfileMutation.mutate({
      input: {
        id: tenantBillingProfileId,
        ...payload,
      },
    });
  }, 2500);
  const { state, setDefaultValues } = useForm({
    formId,
    defaultValues: newDefaults,
    stateReducer: (_, action, next) => {
      if (action.type === 'FIELD_CHANGE') {
        switch (action.payload.name) {
          case 'country':
          case 'canPayWithDirectDebitSEPA':
          case 'canPayWithDirectDebitACH':
          case 'canPayWithDirectDebitBacs':
          case 'canPayWithCard':
          case 'canPayWithPigeon': {
            const payload = TenantBillingDetailsDto.toPayload(next.values);
            updateBillingProfileMutation.mutate({
              input: {
                id: tenantBillingProfileId,
                ...payload,
              },
            });

            return next;
          }
          case 'vatNumber':
          case 'sendInvoicesFrom':
          case 'organizationLegalName':
          case 'addressLine1':
          case 'addressLine2':
          case 'addressLine3':
          case 'zip':
          case 'locality': {
            handleUpdateData({
              ...next.values,
            });

            return next;
          }
          default:
            return next;
        }
      }
      if (action.type === 'FIELD_BLUR') {
        switch (action.payload.name) {
          case 'vatNumber':
          case 'sendInvoicesFrom':
          case 'organizationLegalName':
          case 'addressLine1':
          case 'addressLine2':
          case 'addressLine3':
          case 'zip':
          case 'locality': {
            handleUpdateData.flush();

            return next;
          }
          default:
            return next;
        }
      }

      return next;
    },
  });

  useEffect(() => {
    return handleUpdateData.flush();
  }, []);

  useEffect(() => {
    if (isFetchedAfterMount && !data?.tenantBillingProfiles.length) {
      createBillingProfileMutation.mutate({
        input: {
          canPayWithDirectDebitACH: false,
          canPayWithDirectDebitSEPA: false,
          canPayWithDirectDebitBacs: false,
          canPayWithCard: false,
          canPayWithPigeon: false,
          sendInvoicesFrom: '',
          vatNumber: '',
        },
      });
    }
  }, [isFetchedAfterMount, data]);

  useEffect(() => {
    if (
      isFetchedAfterMount &&
      !!data?.tenantBillingProfiles.length &&
      data?.tenantBillingProfiles?.[0]
    ) {
      const newDefaults = new TenantBillingDetailsDto(
        data?.tenantBillingProfiles?.[0] as TenantBillingProfile,
      );
      setDefaultValues(newDefaults);
    }
  }, [isFetchedAfterMount, data]);

  return (
    <Flex>
      <Card
        flex='1'
        w='full'
        h='100vh'
        bg='#FCFCFC'
        flexDirection='column'
        boxShadow='none'
        background='gray.25'
        maxW={400}
        borderRight='1px solid'
        borderColor='gray.300'
        overflowY='scroll'
        borderRadius='none'
      >
        <CardHeader px='6' pb='0' pt='4'>
          <Heading as='h1' fontSize='lg' color='gray.700' pt={1}>
            <b>Billing</b>
          </Heading>
        </CardHeader>
        <CardBody as={Flex} flexDir='column' px='6' w='full' gap={4}>
          <FormInput
            autoComplete='off'
            label='Organization legal name'
            placeholder='Legal name'
            isLabelVisible
            labelProps={{
              fontSize: 'sm',
              mb: 0,
              fontWeight: 'semibold',
            }}
            name='legalName'
            formId={formId}
            onMouseEnter={() => setIsInvoiceProviderDetailsHovered(true)}
            onMouseLeave={() => setIsInvoiceProviderDetailsHovered(false)}
            onFocus={() => setIsInvoiceProviderFocused(true)}
            onBlur={() => setIsInvoiceProviderFocused(false)}
          />
          <Flex
            flexDir='column'
            onMouseEnter={() => setIsInvoiceProviderDetailsHovered(true)}
            onMouseLeave={() => setIsInvoiceProviderDetailsHovered(false)}
          >
            <FormInput
              autoComplete='off'
              label='Billing address'
              placeholder='Address line 1'
              isLabelVisible
              labelProps={{
                fontSize: 'sm',
                mb: 0,
                fontWeight: 'semibold',
              }}
              name='addressLine1'
              formId={formId}
              onFocus={() => setIsInvoiceProviderFocused(true)}
              onBlur={() => setIsInvoiceProviderFocused(false)}
            />
            <FormInput
              autoComplete='off'
              label='Billing address line 2'
              name='addressLine2'
              placeholder='Address line 2'
              formId={formId}
              onFocus={() => setIsInvoiceProviderFocused(true)}
              onBlur={() => setIsInvoiceProviderFocused(false)}
            />

            <Flex gap={2}>
              <FormInput
                autoComplete='off'
                label='Billing address locality'
                name='locality'
                placeholder='City'
                formId={formId}
                onFocus={() => setIsInvoiceProviderFocused(true)}
                onBlur={() => setIsInvoiceProviderFocused(false)}
              />
              <FormInput
                autoComplete='off'
                label='Billing address zip/Postal code'
                name='zip'
                placeholder='ZIP/Postal code'
                formId={formId}
                onFocus={() => setIsInvoiceProviderFocused(true)}
                onBlur={() => setIsInvoiceProviderFocused(false)}
              />
            </Flex>
            <FormSelect
              name='country'
              placeholder='Country'
              formId={formId}
              options={countryOptions}
            />
            <VatInput
              formId={formId}
              name='vatNumber'
              autoComplete='off'
              label='VAT number'
              isLabelVisible
              labelProps={{
                fontSize: 'sm',
                mb: 0,
                mt: 4,
                fontWeight: 'semibold',
              }}
              textOverflow='ellipsis'
              placeholder='VAT number'
              onFocus={() => setIsInvoiceProviderFocused(true)}
              onBlur={() => setIsInvoiceProviderFocused(false)}
            />

            <FormInput
              autoComplete='off'
              label='Send invoice from'
              isLabelVisible
              labelProps={{
                fontSize: 'sm',
                mb: 0,
                mt: 4,
                fontWeight: 'semibold',
              }}
              formId={formId}
              name='email'
              textOverflow='ellipsis'
              placeholder='Email'
              type='email'
              isInvalid={
                !!state.values.email?.length &&
                !emailRegex.test(state.values.email)
              }
              onFocus={() => setIsInvoiceProviderFocused(true)}
              onBlur={() => setIsInvoiceProviderFocused(false)}
            />
          </Flex>

          <Flex position='relative' alignItems='center'>
            <Text color='gray.500' fontSize='xs' whiteSpace='nowrap' mr={2}>
              Customer can pay using
            </Text>
            <Divider background='gray.200' />
          </Flex>

          <FormSwitch
            name='canPayWithCard'
            formId={formId}
            size='sm'
            label={
              <Text fontSize='sm' fontWeight='semibold' whiteSpace='nowrap'>
                Credit or Debit cards
              </Text>
            }
          />

          <Flex flexDir='column' gap={2}>
            <Text fontSize='sm' fontWeight='semibold' whiteSpace='nowrap'>
              Direct debit via
            </Text>
            <FormSwitch
              name='canPayWithDirectDebitSEPA'
              formId={formId}
              size='sm'
              label={
                <Text
                  fontSize='sm'
                  fontWeight='medium'
                  whiteSpace='nowrap'
                  as='label'
                >
                  <Eu mr={2} />
                  SEPA
                </Text>
              }
            />
            <FormSwitch
              name='canPayWithDirectDebitACH'
              formId={formId}
              size='sm'
              label={
                <Text
                  fontSize='sm'
                  fontWeight='medium'
                  whiteSpace='nowrap'
                  as='label'
                >
                  <Us mr={2} />
                  ACH
                </Text>
              }
            />

            <FormSwitch
              name='canPayWithDirectDebitBacs'
              formId={formId}
              size='sm'
              label={
                <Text
                  fontSize='sm'
                  fontWeight='medium'
                  whiteSpace='nowrap'
                  as='label'
                >
                  <Gb mr={2} />
                  Bacs
                </Text>
              }
            />
          </Flex>
          {/*<Flex justifyContent='space-between' alignItems='center'>*/}
          {/*  <Text fontSize='sm' fontWeight='semibold' whiteSpace='nowrap'>*/}
          {/*    Bank transfer*/}
          {/*  </Text>*/}
          {/*  <Switch size='sm' />*/}
          {/*</Flex>*/}
          <FormSwitch
            name='canPayWithPigeon'
            formId={formId}
            size='sm'
            label={
              <Text fontSize='sm' fontWeight='semibold' whiteSpace='nowrap'>
                Carrier pigeon
              </Text>
            }
          />
        </CardBody>

          <ScaleFade in={isInvoicingEnabled}>
              {isInvoicingEnabled && (
                  <TenantBillingPanelDetailsForm
                      setIsDomesticBankingDetailsSectionFocused={
                          setIsDomesticBankingDetailsSectionFocused
                      }
                      setIsInvoiceProviderDetailsHovered={
                          setIsInvoiceProviderDetailsHovered
                      }
                      setIsInvoiceProviderFocused={setIsInvoiceProviderFocused}
                      setIsInternationalBankingDetailsSectionFocused={
                          setIsInternationalBankingDetailsSectionFocused
                      }
                      setIsDomesticBankingDetailsSectionHovered={
                          setIsDomesticBankingDetailsSectionHovered
                      }
                      setIsInternationalBankingDetailsSectionHovered={
                          setIsInternationalBankingDetailsSectionHovered
                      }
                  />
              )}
          </ScaleFade>
      </Card>
      <Box borderRight='1px solid' borderColor='gray.300' maxH='100vh'>
        <Invoice
          isInvoiceProviderFocused={
            isInvoiceProviderFocused || isInvoiceProviderDetailsHovered
          }
          from={{
            addressLine1: state.values.addressLine1 ?? '',
            addressLine2: state.values.addressLine2 ?? '',
            locality: state.values.locality ?? '',
            zip: state.values.zip ?? '',
            country: state?.values?.country?.label ?? '',
            email: state?.values?.email ?? '',
            name: state.values?.legalName ?? '',
            vatNumber: state.values?.vatNumber ?? '',
          }}
          {...invoicePreviewStaticData}
        />
      </Box>
    </Flex>
  );
};
