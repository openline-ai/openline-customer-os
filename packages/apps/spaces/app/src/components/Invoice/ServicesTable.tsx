import React from 'react';

import { Invoice } from '@graphql/types';
import { formatCurrency } from '@spaces/utils/getFormattedCurrencyNumber';
import {
  Tr,
  Th,
  Td,
  Table,
  Tbody,
  Thead,
  TableContainer,
} from '@ui/presentation/SimpleTable';

type ServicesTableProps = {
  currency: string;
  services: Partial<Invoice['invoiceLineItems']>;
};

export function ServicesTable({ services, currency }: ServicesTableProps) {
  return (
    <TableContainer width='100%'>
      <Table variant='simple' size='md' width='100%'>
        <Thead>
          <Tr>
            <Th
              pl={0}
              minW='250px'
              textTransform='capitalize'
              fontSize='sm'
              borderColor='gray.300'
            >
              Service
            </Th>
            <Th
              w={'10%'}
              isNumeric
              textTransform='capitalize'
              fontSize='sm'
              borderColor='gray.300'
            >
              Qty
            </Th>
            <Th
              w={'10%'}
              isNumeric
              textTransform='capitalize'
              fontSize='sm'
              borderColor='gray.300'
            >
              Unit price
            </Th>
            <Th
              w={'10%'}
              isNumeric
              textTransform='capitalize'
              fontSize='sm'
              borderColor='gray.300'
              pr={0}
            >
              Amount
            </Th>
          </Tr>
        </Thead>
        <Tbody>
          {services.map((service, index) => (
            <Tr key={index}>
              <Td
                fontSize='sm'
                pl={0}
                fontWeight='medium'
                borderColor='gray.300'
              >
                {service?.description ?? 'Unnamed'}
              </Td>
              <Td
                fontSize='sm'
                isNumeric
                borderColor='gray.300'
                color='gray.500'
              >
                {service?.quantity}
              </Td>
              <Td
                fontSize='sm'
                isNumeric
                borderColor='gray.300'
                color='gray.500'
              >
                {formatCurrency(service?.price ?? 0, 2, currency)}
              </Td>
              <Td
                fontSize='sm'
                isNumeric
                borderColor='gray.300'
                color='gray.500'
                pr={0}
              >
                {formatCurrency(service?.total ?? 0, 2, currency)}
              </Td>
            </Tr>
          ))}
        </Tbody>
      </Table>
    </TableContainer>
  );
}
