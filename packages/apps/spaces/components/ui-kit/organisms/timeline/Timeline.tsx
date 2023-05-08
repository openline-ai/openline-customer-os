import React, { useEffect, useRef } from 'react';

import { ConversationTimelineItem } from '@spaces/molecules/conversation-timeline-item';
import { EmailTimelineItem } from '@spaces/molecules/email-timeline-item';
import { LiveEventTimelineItem } from '@spaces/molecules/live-event-timeline-item';
import { NoteTimelineItem } from '@spaces/molecules/note-timeline-item';
import { WebActionTimelineItem } from '@spaces/molecules/web-action-timeline-item';
import { IssueTimelineItem } from '@spaces/molecules/issue-timeline-item';
import { EmailTimelineItemTemp } from '@spaces/molecules/conversation-timeline-item/EmailTimelineItemTemp';
import { PhoneConversationTimelineItem } from '@spaces/molecules/conversation-timeline-item/PhoneConversationTimelineItem';
import { MeetingTimelineItem } from '@spaces/molecules//meeting-timeline-item';
import { InteractionTimelineItem } from '@spaces/molecules/interaction-timeline-item';
import { TimelineItem } from '@spaces/atoms/timeline-item';
import { useInfiniteScroll } from './useInfiniteScroll';
import { Skeleton } from '@spaces/atoms/skeleton';
import { TimelineStatus } from './timeline-status';
import classNames from 'classnames';


import styles from './timeline.module.scss';

interface Props {
  loading: boolean;
  noActivity: boolean;
  id?: string;
  loggedActivities: Array<any>;
  notifyChange?: (id: any) => void;
  onLoadMore: (ref: any) => void;
  contactName?: string;
  mode: 'CONTACT' | 'ORGANIZATION';
}

export const Timeline = ({
  loading,
  noActivity,
  loggedActivities,
  id,
  onLoadMore,
  contactName = '',
  mode,
}: Props) => {
  const timelineContainerRef = useRef<HTMLDivElement>(null);
  const containerRef = useRef(null);

  const lastItemRef = useRef<Array<HTMLDivElement>>([]);

  const infiniteScrollElementRef = useRef(null);
  useInfiniteScroll({
    element: infiniteScrollElementRef,
    isFetching: loading,
    callback: () => {
      if (loggedActivities.length > 10) {
        onLoadMore(containerRef);
      }
    },
  });
  useEffect(() => {
    if (timelineContainerRef?.current) {
      timelineContainerRef.current.scrollTop =
        timelineContainerRef.current.scrollHeight;
    }
  }, []);

  useEffect(() => {
    if (
      loading &&
      lastItemRef.current.length > 0 &&
      timelineContainerRef.current
    ) {
      timelineContainerRef.current.scrollTop = 400;
    }
  }, [loading]);

  const getTimelineItemByType = (type: string, data: any, index: number) => {
    switch (type) {
      case 'Note':
        console.log('------------data', data);
        return (
          <TimelineItem source={data.source || data.appSource} first={index == 0} createdAt={data?.createdAt}>
            <NoteTimelineItem
              note={data}
            />
          </TimelineItem>
        );
      case 'Conversation':
        // TODO move to interaction event once we have the data in backend
        // if (data.channel === 'WEB_CHAT') {
        //   return (
        //     <ChatTimelineItem
        //       first={index == 0}
        //       feedId={data.id}
        //       source={data.source}
        //       createdAt={data?.startedAt}
        //       feedInitiator={{
        //         firstName: data.initiatorFirstName,
        //         lastName: data.initiatorLastName,
        //         phoneNumber: data.initiatorUsername.identifier,
        //         lastTimestamp: data.lastTimestamp,
        //       }}
        //     />
        //   );
        // }
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
        // TODO move to interaction event once we have the data in backend
        if (data.channel === 'VOICE') {
          return (
            <PhoneConversationTimelineItem
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

      case 'Analysis': {
        if (data.describes.find((e: any) => e.__typename === 'Meeting')) {
          return null;
        }
        const decodeContent = (content: string) => {
          let response;
          try {
            response = JSON.parse(content);
          } catch (e) {
            response = {
              dialog: {
                type: 'MESSAGE',
                mimetype: 'text/plain',
                body: content,
              },
            };
          }
          return response;
        };
        if (data.analysisType === 'transcript') {
          return null;
        }

        const transcriptForSummary = loggedActivities
          .filter((e) => e.__typename === 'Analysis')
          .filter((e) => e.analysisType !== 'summary')
          .find((e) => e.describes[0].id === data.describes[0].id);

        if (!transcriptForSummary?.content) {
          return;
        }

        return (
          <ConversationTimelineItem
            id={data.id}
            source={data.source}
            content={decodeContent(data.content)}
            transcript={decodeContent(transcriptForSummary.content)}
            type={data.analysisType}
            createdAt={data?.createdAt}
            contentType={transcriptForSummary.contentType}
            mode='PHONE_CALL' // fixme - mode will be assessed from data inside the component (on message base)
          />
        );
      }
      case 'PageView':
        return (
          <TimelineItem source={data.source} first={index == 0} createdAt={data?.startedAt}>
            <WebActionTimelineItem {...data} contactName={contactName} />
          </TimelineItem>
        );
      case 'InteractionSession':
        return (
          <TimelineItem source={data.source} first={index == 0} createdAt={data?.startedAt}>
            <InteractionTimelineItem
              {...data}
              contactId={contactName && id}
              organizationId={!contactName && id}
            />
          </TimelineItem>
        );
      case 'Issue':
        return (
          <TimelineItem source={data.source} first={index == 0} createdAt={data?.createdAt}>
            <IssueTimelineItem {...data} />
          </TimelineItem>
        );

      case 'InteractionEvent':
        if (data.channel === 'EMAIL') {
          return (
            <TimelineItem source={data.source} first={index == 0} createdAt={data?.createdAt}>
              <EmailTimelineItem
                {...data}
                contactId={mode === 'CONTACT' && id}
              />
            </TimelineItem>
          );
        }
        if (data.channel === 'VOICE') {
          const from =
            data.sentBy && data.sentBy.length > 0
              ? data.sentBy
                  .map((p: any) => {
                    if (
                      p.__typename === 'EmailParticipant' &&
                      p.emailParticipant
                    ) {
                      return p.emailParticipant.email;
                    }
                    return '';
                  })
                  .join('; ')
              : '';

          const to =
            data.sentTo && data.sentTo.length > 0
              ? data.sentTo
                  .map((p: any) => {
                    if (
                      p.__typename === 'EmailParticipant' &&
                      p.emailParticipant
                    ) {
                      return p.emailParticipant.email;
                    } else if (
                      p.__typename === 'ContactParticipant' &&
                      p.contactParticipant
                    ) {
                      if (
                        p.contactParticipant.name &&
                        p.contactParticipant.name !== ''
                      ) {
                        return p.contactParticipant.name;
                      } else {
                        return (
                          p.contactParticipant.firstName +
                          ' ' +
                          p.contactParticipant.lastName
                        );
                      }
                    }
                    return '';
                  })
                  .join('; ')
              : '';

          //we are using this to render the phone calls manually created by the user
          return (
            <ConversationTimelineItem
              id={data.id}
              content={undefined}
              transcript={[
                {
                  text: data.content,
                  party: {
                    tel: from,
                    mailto: to,
                  },
                },
              ]}
              source={data.source}
              type={'summary'} //fixme: this is used to get the same style as the summary of  phone call
              createdAt={data?.createdAt}
              mode='PHONE_CALL' // fixme - mode will be assessed from data inside the component (on message base)
            />
          );
        } else {
          return null;
        }
      case 'LiveEventTimelineItem':
        return (
          <LiveEventTimelineItem
            first={index == 0}
            contactId={id}
            source={data.source}
          />
        );
      case 'Meeting':
        return (
          <TimelineItem source={data.source}
            first={index == 0}
            createdAt={data?.createdAt || new Date()}
            hideTimeTick
          >
            <MeetingTimelineItem meeting={data} />
          </TimelineItem>
        );
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
    <div ref={timelineContainerRef} className={styles.timeline}>
      <div
        className={classNames(styles.timelineContent, styles.scrollable, {
          [styles.scrollable]: !noActivity,
        })}
        ref={containerRef}
      >
        {!!loggedActivities.length && (
          <div
            ref={infiniteScrollElementRef}
            style={{
              height: '1px',
              width: '1px',
            }}
          />
        )}
        {loading && (
          <div className='flex flex-column mt-4'>
            <Skeleton height={'40px'} className='mb-3' />
            <Skeleton height={'40px'} className='mb-3' />
          </div>
        )}
        {!loading && noActivity && <TimelineStatus status='no-activity' />}

        {loggedActivities.map((e: any, index) => {
          return (
            <div
              key={`${e.__typename}-${e.id}`}
              //@ts-expect-error ts issue fix later
              ref={(el) => (lastItemRef.current[index] = el)}
            >
              {getTimelineItemByType(e.__typename, e, index)}
            </div>
          );
        })}
      </div>

      <div id={styles.scrollAnchor} />
    </div>
  );
};
