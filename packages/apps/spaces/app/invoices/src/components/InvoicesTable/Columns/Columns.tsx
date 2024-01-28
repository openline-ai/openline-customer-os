import React from 'react';

import { Invoice } from '@graphql/types';
import { Text } from '@ui/typography/Text';
import { DateTimeUtils } from '@spaces/utils/date';
import { Skeleton } from '@ui/presentation/Skeleton';
import { createColumnHelper } from '@ui/presentation/Table';
import { StatusCell } from '@shared/components/Invoice/Cells';
import THead, { getTHeadProps } from '@ui/presentation/Table/THead';
import { formatCurrency } from '@spaces/utils/getFormattedCurrencyNumber';

const columnHelper = createColumnHelper<Invoice>();

export const columns = [
  columnHelper.accessor('number', {
    id: 'NUMBER',
    minSize: 150,
    maxSize: 150,
    enableSorting: false,
    enableColumnFilter: false,
    cell: (props) => <Text overflow='hidden'>{props?.getValue()}</Text>,
    header: (props) => (
      <THead id='number' title='N°' {...getTHeadProps<Invoice>(props)} py='2' />
    ),
    skeleton: () => (
      <Skeleton
        width='50%'
        height='18px'
        startColor='gray.300'
        endColor='gray.300'
      />
    ),
  }),
  columnHelper.accessor('status', {
    id: 'STATUS',
    minSize: 60,
    maxSize: 60,

    header: (props) => (
      <THead
        id='status'
        title='Status'
        {...getTHeadProps<Invoice>(props)}
        py='2'
      />
    ),
    cell: (props) => {
      return <StatusCell status={props.getValue()} />;
    },
    skeleton: () => (
      <Skeleton
        width='100%'
        height='18px'
        startColor='gray.300'
        endColor='gray.300'
      />
    ),
  }),
  columnHelper.accessor('dueDate', {
    id: 'DUE_DATE',
    minSize: 40,
    maxSize: 40,

    header: (props) => (
      <THead
        id='issued'
        title='Issued'
        {...getTHeadProps<Invoice>(props)}
        py='2'
      />
    ),
    cell: (props) => {
      return (
        <Text>
          {DateTimeUtils.format(
            props.getValue(),
            DateTimeUtils.defaultFormatShortString,
          )}
        </Text>
      );
    },
    skeleton: () => (
      <Skeleton
        width='100%'
        height='18px'
        startColor='gray.300'
        endColor='gray.300'
      />
    ),
  }),
  columnHelper.accessor('totalAmount', {
    id: 'AMOUNT_DUE',
    minSize: 40,
    header: (props) => (
      <THead
        id='amount'
        title='Amount'
        {...getTHeadProps<Invoice>(props)}
        py='2'
      />
    ),
    cell: (props) => {
      return (
        <Text>
          {formatCurrency(props.getValue(), 2, props.row.original.currency)}
        </Text>
      );
    },
    skeleton: () => (
      <Skeleton
        width='100%'
        height='18px'
        startColor='gray.300'
        endColor='gray.300'
      />
    ),
  }),
];
