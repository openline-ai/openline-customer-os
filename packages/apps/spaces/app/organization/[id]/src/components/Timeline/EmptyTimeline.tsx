import React from 'react';
import { useParams } from 'next/navigation';

import { useFeatureIsOn } from '@growthbook/growthbook-react';

import { Flex } from '@ui/layout/Flex';
import { Text } from '@ui/typography/Text';
import { useOrganization } from '@organization/src/hooks/useOrganization';
import { TimelineActions } from '@organization/src/components/Timeline/FutureZone/TimelineActions/TimelineActions';

import { FutureZone } from './FutureZone/FutureZone';
import EmptyTimelineIllustration from './assets/EmptyTimelineIllustration';

interface EmptyTimelineProps {
  invalidateQuery: () => void;
}
export const EmptyTimeline: React.FC<EmptyTimelineProps> = ({
  invalidateQuery,
}) => {
  const id = useParams()?.id as string;
  const isRemindersEnabled = useFeatureIsOn('reminders');

  const { data } = useOrganization({ id });

  return (
    <Flex direction='column' height='calc(100vh - 5rem)' overflow='auto'>
      <Flex
        direction='column'
        alignItems='center'
        flex={1}
        backgroundImage='/backgrounds/organization/dotted-bg-pattern.svg'
        backgroundRepeat='no-repeat'
        backgroundSize='contain'
        backgroundPosition='center'
        maxH='50%'
        as='article'
      >
        <Flex
          direction='column'
          alignItems='center'
          justifyContent='center'
          height='100%'
          maxWidth='390px'
        >
          <EmptyTimelineIllustration />
          <Text
            color='gray.900'
            fontSize='lg'
            as='h1'
            fontWeight={600}
            mt={3}
            mb={2}
          >
            {data?.organization?.name || 'Unknown'} has no events yet
          </Text>
          <Text color='gray.600' size='xs' textAlign='center'>
            This organization’s events will show up here once a data source has
            been linked
          </Text>
        </Flex>
      </Flex>
      <Flex bg='#F9F9FB' direction='column' flex={1}>
        <div>
          <TimelineActions invalidateQuery={invalidateQuery} />
        </div>
        <Flex flex={1} height='100%' bg='#F9F9FB'>
          {isRemindersEnabled && <FutureZone />}
        </Flex>
      </Flex>
    </Flex>
  );
};
