import React from 'react';

import { Flex } from '@ui/layout/Flex';
import { Text } from '@ui/typography/Text';
import { FeaturedIcon } from '@ui/media/Icon';
import { Heading } from '@ui/typography/Heading';
import { FileX02 } from '@ui/media/icons/FileX02';
import { Invoice } from '@shared/components/Invoice/Invoice';
import { GetInvoiceQuery } from '@shared/graphql/getInvoice.generated';
import { InvoiceSkeleton } from '@shared/components/Invoice/InvoiceSkeleton';
import { InvoiceLine, InvoiceCustomer, InvoiceProvider } from '@graphql/types';
interface InvoicePreviewModalProps {
  isError: boolean;
  isFetching: boolean;
  data: GetInvoiceQuery | undefined;
}

const extractAddressData = (invoiceData: InvoiceCustomer | InvoiceProvider) => {
  return {
    zip: invoiceData?.addressZip ?? '',
    email: (invoiceData as InvoiceCustomer)?.email ?? '',
    name: invoiceData?.name ?? '',
    country: invoiceData?.addressCountry ?? '',
    locality: invoiceData?.addressLocality ?? '',
    addressLine1: invoiceData?.addressLine1 ?? '',
    addressLine2: invoiceData?.addressLine2 ?? '',
  };
};

export const InvoicePreviewModalContent: React.FC<InvoicePreviewModalProps> = ({
  isFetching,
  isError,
  data,
}) => {
  if (isFetching) {
    return <InvoiceSkeleton />;
  }

  if (!data?.invoice || isError) {
    return (
      <Flex
        as='article'
        position='relative'
        flexDirection='column'
        alignItems='center'
        px={4}
        py={4}
        mt={5}
        overflow='hidden'
      >
        <FeaturedIcon colorScheme='warning'>
          <FileX02 boxSize='7' />
        </FeaturedIcon>
        <Heading fontSize='md' mt={4} mb={1}>
          Preview not available
        </Heading>
        <Text textAlign='center' fontSize='sm' color='gray.500'>
          Sorry, selected invoice cannot be previewed at this moment
        </Text>
      </Flex>
    );
  }

  const customerAddressData = extractAddressData(data?.invoice?.customer);
  const providerAddressData = extractAddressData(data?.invoice?.provider);

  return (
    <Invoice
      tax={data?.invoice?.taxDue}
      note={data?.invoice?.note}
      logoUrl={data?.invoice?.provider?.logoUrl}
      from={providerAddressData}
      total={data?.invoice.amountDue}
      dueDate={data?.invoice.due}
      subtotal={data?.invoice.subtotal}
      issueDate={data?.invoice?.metadata?.created}
      billedTo={customerAddressData}
      invoiceNumber={data?.invoice?.invoiceNumber ?? ''}
      lines={(data?.invoice?.invoiceLineItems as Array<InvoiceLine>) ?? []}
      currency={data?.invoice?.currency || 'USD'}
    />
  );
};
