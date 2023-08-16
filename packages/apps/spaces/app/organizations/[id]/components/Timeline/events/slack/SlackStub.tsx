'use client';
import React, { FC } from 'react';
import { Card, CardBody } from '@ui/presentation/Card';
import {
  ContactParticipant,
  InteractionEvent,
  UserParticipant,
} from '@graphql/types';
import { useTimelineEventPreviewContext } from '@organization/components/Timeline/preview/TimelineEventsPreviewContext/TimelineEventPreviewContext';
import { Avatar } from '@ui/media/Avatar';
import { Flex, Text } from '@chakra-ui/react';
import { getParticipantName } from '@spaces/utils/getParticipantsName';
import User from '@spaces/atoms/icons/User';
import Slack from '@spaces/atoms/icons/Slack';
import { Button } from '@ui/form/Button';
import { ViewInSlackButton } from '@organization/components/Timeline/events/slack/ViewInSlackButton';

export const SlackStub: FC<{ slackEvent: InteractionEvent }> = ({
  slackEvent,
}) => {
  const { openModal } = useTimelineEventPreviewContext();

  const slackSender =
    (slackEvent?.sentBy?.[0] as ContactParticipant)?.contactParticipant ||
    (slackEvent?.sentBy?.[0] as UserParticipant)?.userParticipant;

  if (!slackSender) {
    return (
      <Card
        variant='outline'
        size='md'
        fontSize='14px'
        background='white'
        flexDirection='row'
        maxWidth={549}
        position='unset'
        cursor='pointer'
        boxShadow='xs'
        borderColor='gray.100'
      >
        <CardBody p={3} overflow={'hidden'}>
          <Flex gap={3} flex={1} alignItems='center'>
            <Avatar
              variant='roundedSquare'
              size='sm'
              colorScheme='gray'
              bg='white'
              border='1px solid var(--chakra-colors-gray-200)'
              icon={<Slack height='1.8rem' />}
            />
            <Text color='gray.700' as='span' fontWeight={600}>
              {slackEvent?.content}
            </Text>
          </Flex>
        </CardBody>
      </Card>
    );
  }

  const slackEventReplies = slackEvent.interactionSession?.events?.filter(
    (e) => e.id !== slackEvent.id,
  );

  const uniqThreadParticipants = slackEventReplies
    ?.map((e) => {
      const threadSender =
        (e?.sentBy?.[0] as ContactParticipant)?.contactParticipant ||
        (e?.sentBy?.[0] as UserParticipant)?.userParticipant;

      return threadSender;
    })
    ?.filter((v, i, a) => a.findIndex((t) => t.id === v.id) === i);

  return (
    <>
      <Card
        variant='outline'
        size='md'
        fontSize='14px'
        background='white'
        flexDirection='row'
        maxWidth={549}
        position='unset'
        // aspectRatio='10/2'
        cursor='pointer'
        boxShadow='xs'
        borderColor='gray.100'
        // onClick={() => openModal(email)}
      >
        <CardBody p={3} overflow={'hidden'}>
          <Flex gap={3} flex={1}>
            <Avatar
              name={getParticipantName(slackEvent?.sentBy?.[0])}
              variant='roundedSquare'
              size='lg'
              icon={
                <User color={'var(--chakra-colors-gray-500)'} height='1.8rem' />
              }
              border={
                slackSender?.profilePhotoUrl
                  ? 'none'
                  : '2px solid var(--chakra-colors-primary-200)'
              }
              src={slackSender?.profilePhotoUrl || undefined}
            />
            <Flex direction='column' flex={1}>
              <Flex justifyContent='space-between' flex={1}>
                <Text color='gray.700' fontWeight={600}>
                  {getParticipantName(slackEvent?.sentBy?.[0])}
                </Text>
                <ViewInSlackButton url='' />
              </Flex>

              <Text noOfLines={3}>{slackEvent?.content}</Text>
              {!!slackEventReplies?.length && (
                <Flex mt={1}>
                  <Flex columnGap={1} mr={1}>
                    {uniqThreadParticipants?.map(
                      ({ profilePhotoUrl, id, name, firstName, lastName }) => (
                        <Avatar
                          name={name || `${firstName} ${lastName}`}
                          key={`uniq-slack-thread-participant-${slackEvent.id}-${id}`}
                          variant='roundedSquareSmall'
                          size='xs'
                          src={profilePhotoUrl || undefined}
                        />
                      ),
                    )}
                  </Flex>
                  <Button variant='link'>
                    {slackEventReplies.length}{' '}
                    {slackEventReplies.length === 1 ? 'reply' : 'replies'}
                  </Button>
                </Flex>
              )}
            </Flex>
          </Flex>
        </CardBody>
      </Card>
    </>
  );
};
