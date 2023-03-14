import React, { useRef } from 'react';
import { useStickyScroll } from '../../../../hooks/useStickyScroll';

import {
  ConversationTimelineItem,
  LiveConversationTimelineItem,
  NoteTimelineItem,
  WebActionTimelineItem,
} from '../../molecules';
import { TimelineItem } from '../../atoms/timeline-item';
import { TicketTimelineItem } from '../../molecules/ticket-timeline-item';
import styles from './timeline.module.scss';
import { InteractionTimelineItem } from '../../molecules/interaction-timeline-item';
import { EmailTimelineItemTemp } from '../../molecules/conversation-timeline-item/EmailTimelineItemTemp';
import { ChatTimelineItem } from '../../molecules/conversation-timeline-item/ChatTimelineItem';
import { useInfiniteScroll } from './useInfiniteScroll';
import { Skeleton } from '../../atoms/skeleton';
import { ConversationTimelineItemDeprecated } from '../../molecules/conversation-timeline-item/ConversationTimelineItemDeprecated';

interface Props {
  loading: boolean;
  noActivity: boolean;
  notPaginated?: boolean; // todo remove when org list is paginated
  contactId?: string;
  loggedActivities: Array<any>;
  notifyChange?: (id: any) => void;
  notifyContactNotesUpdate?: (id: any) => void;
  onLoadMore: () => void;
}

export const Timeline = ({
  loading,
  noActivity,
  loggedActivities,
  contactId,
  notifyChange = () => null,
  notifyContactNotesUpdate = () => null,
  onLoadMore,
  notPaginated = false,
}: Props) => {
  const timelineContainerRef = useRef(null);
  const containerRef = useRef(null);
  const infiniteScrollElementRef = useRef(null);
  // @ts-expect-error revisit later
  useStickyScroll(containerRef, loggedActivities || []);
  useInfiniteScroll({
    element: infiniteScrollElementRef,
    isFetching: loading,
    callback: () => {
      if (loggedActivities.length > 10 && !notPaginated) {
        // @ts-expect-error this code will be removed after total items count is provided and we switch to virtualized list
        containerRef.current.scrollTop = 100;
        onLoadMore();
      }
    },
  });
  if (!loading && noActivity) {
    return (
      <p className='text-gray-600 font-italic mt-4'>No activity logged yet</p>
    );
  }

  const getTimelineItemByType = (type: string, data: any, index: number) => {
    switch (type) {
      case 'Note':
        return (
          <TimelineItem first={index == 0} createdAt={data?.createdAt}>
            <NoteTimelineItem
              noteContent={data.html}
              createdAt={data.createdAt}
              createdBy={data?.createdBy}
              id={data.id}
              source={data?.source}
              refreshNoteData={
                data?.contact ? notifyContactNotesUpdate : notifyChange
              }
              contactId={contactId}
            />
          </TimelineItem>
        );
      case 'Conversation':
        if (notPaginated) {
          // TODO remove when org timeline is paginated
          return (
            <ConversationTimelineItemDeprecated
              first={index == 0}
              feedId={data.id}
              source={data.source}
              createdAt={data?.startedAt}
            />
          );
        }
        if (data.channel === 'WEB_CHAT') {
          return (
            <ChatTimelineItem
              first={index == 0}
              feedId={data.id}
              source={data.source}
              createdAt={data?.startedAt}
              feedInitiator={{
                firstName: data.initiatorFirstName,
                lastName: data.initiatorLastName,
                phoneNumber: data.initiatorUsername.identifier,
                lastTimestamp: data.lastTimestamp,
              }}
            />
          );
        }
        if (data.channel === 'EMAIL') {
          return (
            <EmailTimelineItemTemp
              first={index == 0}
              feedId={data.id}
              source={data.source}
              createdAt={data?.startedAt}
              feedInitiator={{
                firstName: data.initiatorFirstName,
                lastName: data.initiatorLastName,
                phoneNumber: data.initiatorUsername.identifier,
                lastTimestamp: data.lastTimestamp,
              }}
            />
          );
        }
        if (data.channel === 'VOICE') {
          return (
            <ConversationTimelineItem
              first={index == 0}
              feedId={data.id}
              source={data.source}
              createdAt={data?.startedAt}
              feedInitiator={{
                firstName: data.initiatorFirstName,
                lastName: data.initiatorLastName,
                phoneNumber: data.initiatorUsername.identifier,
                lastTimestamp: data.lastTimestamp,
              }}
            />
          );
        }
        return null;

      case 'LiveConversation':
        return (
          <LiveConversationTimelineItem
            first={index == 0}
            contactId={contactId}
            source={data.source}
          />
        );
      case 'PageView':
        return (
          <TimelineItem first={index == 0} createdAt={data?.startedAt}>
            <WebActionTimelineItem {...data} />
          </TimelineItem>
        );
      case 'InteractionSession':
        return (
          <TimelineItem
            first={index == 0}
            createdAt={data?.startedAt}
            contentClassName={'interactionTimeLineItemClass'}
          >
            <InteractionTimelineItem {...data} />
          </TimelineItem>
        );
      case 'Ticket':
        return (
          <TimelineItem first={index == 0} createdAt={data?.createdAt}>
            <TicketTimelineItem {...data} />
          </TimelineItem>
        );
      // case "CALL":
      //     return <PhoneCallTimelineItem phoneCallParties={data} duration={}/>
      default:
        return type ? (
          <div>
            Sorry, looks like &apos;{type}&apos; activity type is not supported
            yet{' '}
          </div>
        ) : (
          ''
        );
    }
  };

  return (
    <article ref={timelineContainerRef} className={styles.timeline}>
      <div className={styles.timelineContent} ref={containerRef}>
        {!!loggedActivities.length && (
          <div
            ref={infiniteScrollElementRef}
            style={{
              height: '6px',
              width: '6px',
            }}
          />
        )}
        {loading && (
          <div className='flex flex-column mt-4'>
            <Skeleton height={'40px'} className='mb-3' />
            <Skeleton height={'40px'} className='mb-3' />
            <Skeleton height={'40px'} className='mb-3' />
            <Skeleton height={'40px'} className='mb-3' />
            <Skeleton height={'40px'} className='mb-3' />
          </div>
        )}
        {loggedActivities.map((e: any, index) => (
          <React.Fragment key={e.id}>
            {getTimelineItemByType(e.__typename, e, index)}
          </React.Fragment>
        ))}
      </div>
    </article>
  );
};
