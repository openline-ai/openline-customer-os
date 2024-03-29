'use client';

import React, { useMemo } from 'react';
import { useField } from 'react-inverted-form';

import { Flex } from '@ui/layout/Flex';
import { Text } from '@ui/typography/Text';
import { Select } from '@ui/form/SyncSelect';

import { currencyIcon, mapCurrencyToOptions } from './utils';

export const BankTransferCurrencySelect = ({
  formId,
  currency,
}: {
  formId: string;
  currency?: string | null;
}) => {
  const { getInputProps } = useField('currency', formId);
  const { id, onChange } = getInputProps();
  const currencyOptions = useMemo(() => mapCurrencyToOptions(), []);

  return (
    <Select
      id={id}
      placeholder='Account currency'
      name='contractRenewalCycle'
      blurInputOnSelect
      onChange={(e) => {
        onChange(e);
      }}
      options={[
        {
          options: [
            { label: 'USD', value: 'USD' },
            { label: 'GBP', value: 'GBP' },
            { label: 'EUR', value: 'EUR' },
          ],
        },

        ...currencyOptions,
      ]}
      formatOptionLabel={(option, { context }) => {
        return (
          <Flex alignItems='center'>
            <Flex
              w={context === 'value' ? 'auto' : 7}
              justifyContent={context === 'value' ? 'center' : 'flex-end'}
              alignItems='center'
              minW={context === 'value' ? '14px' : 'auto'}
              padding={context === 'value' ? '0 2px' : 'auto'}
            >
              {currencyIcon?.[option.value]}
            </Flex>
            <Text className='option-label' ml={3}>
              {option.value}
            </Text>
          </Flex>
        );
      }}
      defaultValue={{ label: currency, value: currency }}
      chakraStyles={{
        container: (props, state) => {
          if (
            !state?.selectProps?.menuIsOpen &&
            state.hasValue &&
            !state.isFocused
          ) {
            return {
              display: 'flex',
              alignItems: 'center',
              width: 'fit-content',
              maxW: 'fit-content',
              willChange: 'width',
              transition: 'width 0.2s',
            };
          }

          return {
            ...props,
            w: '100%',
            overflow: 'visible',
            willChange: 'width',
            transition: 'width 0.2s',
            _hover: { cursor: 'pointer' },
          };
        },
        control: (props, state) => {
          if (
            !state?.selectProps?.menuIsOpen &&
            state.hasValue &&
            !state.isFocused
          ) {
            return {
              height: '24px',
              maxH: '24px',
              width: 'max-content',
              minW: '24px',
              borderRadius: '30px',
              border: '1px solid',
              borderColor: 'gray.200',
              padding: '2px',

              display: 'flex',
              justifyContent: 'center',
              alignItems: 'center',
              fontSize: '12px',

              '& .option-label': {
                display: 'none',
              },
              '& svg': {
                marginLeft: '1px',
                height: '12px',
              },
            };
          }

          return {
            ...props,
            w: '100%',
            overflow: 'visible',
            _hover: { cursor: 'pointer' },
          };
        },
        groupHeading: (props) => ({
          display: 'none',
        }),
        group: (props) => ({
          borderBottom: '1px solid',
          borderColor: 'gray.200',
        }),
      }}
    />
  );
};
