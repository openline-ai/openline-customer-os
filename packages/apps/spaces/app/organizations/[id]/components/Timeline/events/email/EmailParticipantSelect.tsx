'use client';
import React, { FC } from 'react';
import { Text } from '@ui/typography/Text';
import { Flex } from '@chakra-ui/react';
import { FormMultiCreatableSelect } from '@ui/form/SyncSelect';
import { ComparisonOperator, Contact } from '@graphql/types';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { GetContactsEmailListDocument } from '@organization/graphql/getContactsEmailList.generated';
import { emailRegex } from '@organization/components/Timeline/events/email/utils';
import { OptionsOrGroups } from 'react-select';

interface EmailParticipantSelect {
  entryType: string;
  fieldName: string;
  formId: string;
}

export const EmailParticipantSelect: FC<EmailParticipantSelect> = ({
  entryType,
  fieldName,
  formId,
}) => {
  const client = getGraphQLClient();

  const getFilteredSuggestions = async (
    filterString: string,
    callback: (options: OptionsOrGroups<any, any>) => void,
  ) => {
    try {
      const results = await client.request<any>(GetContactsEmailListDocument, {
        pagination: {
          page: 1,
          limit: 5,
        },
        where: {
          OR: [
            {
              filter: {
                property: 'FIRST_NAME',
                value: filterString,
                operation: ComparisonOperator.Contains,
              },
            },
            {
              filter: {
                property: 'LAST_NAME',
                value: filterString,
                operation: ComparisonOperator.Contains,
              },
            },
            {
              filter: {
                property: 'NAME',
                value: filterString,
                operation: ComparisonOperator.Contains,
              },
            },
          ],
        },
      });
      const options: OptionsOrGroups<string, any> = (results?.contacts?.content || [])
        .filter((e: Contact) => e.emails.length)
        .map((e: Contact) =>
          e.emails.map((email) => ({
            value: email.email,
            label: `${e.firstName} ${e.lastName}`,
          })),
        )
        .flat();
      callback(options);
    } catch (error) {
      callback([]);
    }
  };

  return (
    <Flex alignItems='center' flex={1} marginBottom={-2} marginTop={-1}>
      <Text as={'span'} color='#344054' fontWeight={600} mr={1}>
        {entryType}:
      </Text>
      <FormMultiCreatableSelect
        name={fieldName}
        formId={formId}
        placeholder='Enter name or email...'
        noOptionsMessage={() => {
          return (
            <>
              No suggestions available, input name or email to search for
              suggestions{' '}
            </>
          );
        }}
        loadOptions={(inputValue: string, callback) => {
          getFilteredSuggestions(inputValue, callback);
        }}
        allowCreateWhileLoading={false}
        formatCreateLabel={(input) => {
          return input;
        }}
        isValidNewOption={(input) => emailRegex.test(input)}
        getOptionLabel={(d) => {
          if (d?.__isNew__) {
            return `${d.label}`;
          }
          return `${d.label} - ${d.value}`;
        }}
      />
    </Flex>
  );
};
