import React from 'react';
import Image from 'next/image';
import { useField } from 'react-inverted-form';

import { Button } from '@ui/form/Button';
import { Box, Flex } from '@chakra-ui/react';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { TagSuggestor } from '@ui/form/RichTextEditor/TagSuggestor';
import { RichTextEditor } from '@ui/form/RichTextEditor/RichTextEditor';
import { useGetTagsQuery } from '@organization/src/graphql/getTags.generated';
import { useTimelineActionLogEntryContext } from '@organization/src/components/Timeline/TimelineActions/context/TimelineActionLogEntryContext';

import { Keymapper } from './Keymapper';
import { TagsSelect } from './TagSelect';
import noteIcon from 'public/images/event-ill-log.png';

export const Logger = () => {
  const { onCreateLogEntry, remirrorProps, isSaving } =
    useTimelineActionLogEntryContext();
  const client = getGraphQLClient();
  const { getInputProps } = useField(
    'content',
    'organization-create-log-entry',
  );
  const { value } = getInputProps();
  const { data } = useGetTagsQuery(client);
  const isLogEmpty = !value?.length || value === `<p style=""></p>`;

  return (
    <Flex
      flexDirection='column'
      position='relative'
      className='customeros-logger'
    >
      <Box position='absolute' top={-4} right={-6}>
        <Image src={noteIcon} alt='' height={123} width={174} />
      </Box>

      <RichTextEditor
        {...remirrorProps}
        placeholder='Log a conversation you had with a customer'
        formId='organization-create-log-entry'
        name='content'
        showToolbar={false}
      >
        <TagSuggestor
          tags={data?.tags?.map((e: { label: string; value: string }) => ({
            label: e.label,
            id: e.value,
          }))}
        />
        <Keymapper />
      </RichTextEditor>
      <Flex justifyContent='space-between' zIndex={8} fontSize='md'>
        <TagsSelect
          formId='organization-create-log-entry'
          name='tags'
          tags={data?.tags}
        />
        <Button
          variant='outline'
          colorScheme='gray'
          fontWeight={600}
          borderRadius='lg'
          pt={1}
          pb={1}
          pl={3}
          pr={3}
          size='sm'
          fontSize='sm'
          isDisabled={isSaving || isLogEmpty}
          isLoading={isSaving}
          loadingText='Sending'
          onClick={() => onCreateLogEntry()}
        >
          Log
        </Button>
      </Flex>
    </Flex>
  );
};
