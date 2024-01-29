'use client';

import React from 'react';

import { Box } from '@ui/layout/Box';
import { Flex } from '@ui/layout/Flex';
import { Tag } from '@ui/presentation/Tag';
import { Text } from '@ui/typography/Text';
import { InvoiceLine } from '@graphql/types';
import { Heading } from '@ui/typography/Heading';
import { Grid, GridItem } from '@ui/layout/Grid';
import { DateTimeUtils } from '@spaces/utils/date';
import { Divider } from '@ui/presentation/Divider';
import { formatCurrency } from '@spaces/utils/getFormattedCurrencyNumber';

import { ServicesTable } from './ServicesTable';
type Address = {
  zip: string;
  email: string;
  name?: string;
  country: string;
  locality: string;
  addressLine: string;
  addressLine2?: string;
};

type InvoiceProps = {
  tax: number;
  note: string;
  from: Address;
  total: number;
  dueDate: string;
  status?: string;
  subtotal: number;
  currency?: string;
  issueDate: string;
  billedTo: Address;
  amountDue?: number;
  invoiceNumber: string;
  lines: Array<InvoiceLine>;
  isBilledToFocused?: boolean;
  isInvoiceProviderFocused?: boolean;
  domesticBankingDetails?: string | null;
  internationalBankingDetails?: string | null;
};

export function Invoice({
  invoiceNumber,
  issueDate,
  dueDate,
  billedTo,
  from,
  lines,
  subtotal,
  tax,
  total,
  note,
  amountDue,
  status,
  isBilledToFocused,
  isInvoiceProviderFocused,
  currency = 'USD',
  domesticBankingDetails,
  internationalBankingDetails,
}: InvoiceProps) {
  const isOutOfFocus = isBilledToFocused || isInvoiceProviderFocused;

  return (
    <Flex px={4} flexDir='column' w='inherit' overflowY='auto'>
      <Flex flexDir='column' mt={2}>
        <Flex alignItems='center'>
          <Heading as='h1' fontSize='3xl' fontWeight='bold'>
            Invoice
          </Heading>
          {status && (
            <Box ml={4}>
              <Tag variant='outline' colorScheme='gray'>
                {status}
              </Tag>
            </Box>
          )}
        </Flex>

        <Heading as='h2' fontSize='sm' fontWeight='regular' color='gray.500'>
          N° {invoiceNumber}
        </Heading>

        <Flex mt={2} justifyContent='space-evenly'>
          <Flex
            flexDir='column'
            flex={1}
            w={170}
            py={2}
            px={2}
            borderRight={'1px solid'}
            filter={isOutOfFocus ? 'blur(2px)' : 'none'}
            transition='filter 0.25s ease-in-out'
            borderTop='1px solid'
            borderBottom='1px solid'
            borderColor='gray.300'
          >
            <Text fontWeight='semibold' mb={1} fontSize='sm'>
              Issued
            </Text>
            <Text fontSize='sm' mb={4}>
              {DateTimeUtils.format(
                issueDate,
                DateTimeUtils.dateWithAbreviatedMonth,
              )}
            </Text>
            <Text fontWeight='semibold' mb={1} fontSize='sm'>
              Due
            </Text>
            <Text fontSize='sm'>
              {DateTimeUtils.format(
                dueDate,
                DateTimeUtils.dateWithAbreviatedMonth,
              )}
            </Text>
          </Flex>
          <Flex
            flexDir='column'
            w={170}
            py={2}
            px={3}
            borderTop='1px solid'
            borderBottom='1px solid'
            borderRight={'1px solid'}
            borderColor={'gray.300'}
            filter={isInvoiceProviderFocused ? 'blur(2px)' : 'none'}
            transition='filter 0.25s ease-in-out'
            position='relative'
            sx={{
              '&:after': {
                content: '""',
                bg: 'transparent',
                border: '2px solid',
                position: 'absolute',
                top: 0,
                bottom: 0,
                left: 0,
                right: 0,
                opacity: isBilledToFocused ? 1 : 0,
                transition: 'opacity 0.25s ease-in-out',
              },
            }}
          >
            <Text fontWeight='semibold' mb={0.5} fontSize='sm'>
              Billed to
            </Text>
            <Text fontSize='sm' fontWeight='medium' mb={1} lineHeight={1.2}>
              {billedTo.name}
            </Text>

            <Text fontSize='sm' lineHeight={1.2}>
              {billedTo.addressLine}
              <Text as='span' display='block' lineHeight={1.2}>
                {billedTo.addressLine2}
              </Text>
            </Text>
            <Text fontSize='sm' lineHeight={1.2}>
              {billedTo.locality} {billedTo.locality && ', '} {billedTo.zip}
            </Text>
            <Text fontSize='sm' lineHeight={1.2}>
              {billedTo.country}
            </Text>
            <Text fontSize='sm' lineHeight={1.2}>
              {billedTo.email}
            </Text>
          </Flex>
          <Flex
            flexDir='column'
            flex={1}
            w={170}
            py={2}
            px={3}
            borderTop='1px solid'
            borderBottom='1px solid'
            borderColor={'gray.300'}
            filter={isBilledToFocused ? 'blur(2px)' : 'none'}
            transition='filter 0.25s ease-in-out'
            position='relative'
            sx={{
              '&:after': {
                content: '""',
                bg: 'transparent',
                border: '2px solid',
                position: 'absolute',
                top: 0,
                bottom: 0,
                left: 0,
                right: -4,
                opacity: isInvoiceProviderFocused ? 1 : 0,
                transition: 'opacity 0.25s ease-in-out',
              },
            }}
          >
            <Text fontWeight='semibold' mb={1} fontSize='sm'>
              From
            </Text>
            <Text fontSize='sm' fontWeight='medium' mb={1} lineHeight={1.2}>
              {from.name}
            </Text>

            <Text fontSize='sm' lineHeight={1.2}>
              {from.addressLine}
              <Text as='span' display='block' lineHeight={1.2}>
                {from.addressLine2}
              </Text>
            </Text>
            <Text fontSize='sm' lineHeight={1.2}>
              {from.locality} {from.locality && ', '} {from.zip}
            </Text>
            <Text fontSize='sm' lineHeight={1.2}>
              {from.country}
            </Text>
            <Text fontSize='sm' lineHeight={1.2}>
              {from.email}
            </Text>
          </Flex>
        </Flex>
      </Flex>

      <Flex
        mt={4}
        flexDir='column'
        filter={isOutOfFocus ? 'blur(2px)' : 'none'}
        transition='filter 0.25s ease-in-out'
      >
        <ServicesTable services={lines} currency={currency} />
        <Flex flexDir='column' alignSelf='flex-end' w='50%' maxW='300px' mt={4}>
          <Flex justifyContent='space-between'>
            <Text fontSize='sm' fontWeight='medium'>
              Subtotal
            </Text>
            <Text fontSize='sm' ml={2} color='gray.600'>
              {formatCurrency(subtotal, 2, currency)}
            </Text>
          </Flex>
          <Divider orientation='horizontal' my={1} borderColor='gray.300' />
          <Flex justifyContent='space-between'>
            <Text fontSize='sm'>Tax</Text>
            <Text fontSize='sm' ml={2} color='gray.600'>
              {formatCurrency(tax, 2, currency)}
            </Text>
          </Flex>
          <Divider orientation='horizontal' my={1} borderColor='gray.300' />
          <Flex justifyContent='space-between'>
            <Text fontSize='sm' fontWeight='medium'>
              Total
            </Text>
            <Text fontSize='sm' ml={2} color='gray.600'>
              {formatCurrency(total, 2, currency)}
            </Text>
          </Flex>
          <Divider orientation='horizontal' my={1} borderColor='gray.500' />
          <Flex justifyContent='space-between'>
            <Text fontSize='sm' fontWeight='semibold'>
              Amount due
            </Text>
            <Text fontSize='sm' fontWeight='semibold' ml={2}>
              {formatCurrency(amountDue || total, 2, currency)}
            </Text>
          </Flex>
          <Divider orientation='horizontal' my={1} borderColor='gray.500' />

          {note && (
            <Flex>
              <Text fontSize='sm' fontWeight='medium'>
                Note:
              </Text>
              <Text fontSize='sm' ml={2} color='gray.500'>
                {note}
              </Text>
            </Flex>
          )}
        </Flex>
      </Flex>

      {domesticBankingDetails && internationalBankingDetails && (
        <Grid
          templateColumns={'50% 50%'}
          marginTop={40}
          minH={100}
          borderTop='1px solid'
          borderBottom='1px solid'
          borderColor='gray.300'
          maxW={600}
          filter={isOutOfFocus ? 'blur(2px)' : 'none'}
          transition='filter 0.25s ease-in-out'
        >
          <GridItem
            p={3}
            borderRight='1px solid'
            borderColor='gray.300'
            maxW='50%'
          >
            <Text fontSize='xs' fontWeight='semibold'>
              Domestic Payments
            </Text>
            <Text fontSize='xs' whiteSpace='pre-wrap'>
              {domesticBankingDetails}
            </Text>
          </GridItem>
          <GridItem p={3}>
            <Text fontSize='xs' fontWeight='semibold'>
              International Payments
            </Text>
            <Text fontSize='xs'>{internationalBankingDetails}</Text>
          </GridItem>
        </Grid>
      )}
    </Flex>
  );
}
