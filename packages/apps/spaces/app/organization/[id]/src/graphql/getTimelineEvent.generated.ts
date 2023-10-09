// @ts-nocheck remove this when typscript-react-query plugin is fixed
import * as Types from '../../../../src/types/__generated__/graphql.types';

import { GraphQLClient } from 'graphql-request';
import { RequestInit } from 'graphql-request/dist/types.dom';
import {
  InteractionEventParticipantFragmentFragmentDoc,
  MeetingParticipantFragmentFragmentDoc,
} from './participantsFragment.generated';
import {
  useQuery,
  useInfiniteQuery,
  UseQueryOptions,
  UseInfiniteQueryOptions,
} from '@tanstack/react-query';

function fetcher<TData, TVariables extends { [key: string]: any }>(
  client: GraphQLClient,
  query: string,
  variables?: TVariables,
  requestHeaders?: RequestInit['headers'],
) {
  return async (): Promise<TData> =>
    client.request({
      document: query,
      variables,
      requestHeaders,
    });
}
export type GetTimelineEventQueryVariables = Types.Exact<{
  ids: Array<Types.Scalars['ID']> | Types.Scalars['ID'];
}>;

export type GetTimelineEventQuery = {
  __typename?: 'Query';
  timelineEvents: Array<
    | {
        __typename: 'Action';
        id: string;
        actionType: Types.ActionType;
        appSource: string;
        createdAt: any;
        metadata?: string | null;
        content?: string | null;
        actionCreatedBy?: {
          __typename: 'User';
          id: string;
          firstName: string;
          lastName: string;
          profilePhotoUrl?: string | null;
        } | null;
      }
    | { __typename: 'Analysis' }
    | {
        __typename: 'InteractionEvent';
        id: string;
        channel?: string | null;
        content?: string | null;
        contentType?: string | null;
        source: Types.DataSource;
        date: any;
        includes: Array<{
          __typename?: 'Attachment';
          id: string;
          mimeType: string;
          name: string;
          extension: string;
        }>;
        issue?: {
          __typename?: 'Issue';
          externalLinks: Array<{
            __typename?: 'ExternalSystem';
            type: Types.ExternalSystemType;
            externalId?: string | null;
            externalUrl?: string | null;
          }>;
        } | null;
        externalLinks: Array<{
          __typename?: 'ExternalSystem';
          externalUrl?: string | null;
          type: Types.ExternalSystemType;
        }>;
        repliesTo?: { __typename?: 'InteractionEvent'; id: string } | null;
        summary?: {
          __typename?: 'Analysis';
          id: string;
          content?: string | null;
          contentType?: string | null;
        } | null;
        actionItems?: Array<{
          __typename?: 'ActionItem';
          id: string;
          content: string;
        }> | null;
        sentBy: Array<
          | {
              __typename: 'ContactParticipant';
              contactParticipant: {
                __typename?: 'Contact';
                id: string;
                name?: string | null;
                firstName?: string | null;
                lastName?: string | null;
                profilePhotoUrl?: string | null;
              };
            }
          | {
              __typename: 'EmailParticipant';
              type?: string | null;
              emailParticipant: {
                __typename?: 'Email';
                email?: string | null;
                id: string;
                contacts: Array<{
                  __typename?: 'Contact';
                  id: string;
                  name?: string | null;
                  firstName?: string | null;
                  lastName?: string | null;
                }>;
                users: Array<{
                  __typename?: 'User';
                  id: string;
                  firstName: string;
                  lastName: string;
                }>;
                organizations: Array<{
                  __typename?: 'Organization';
                  id: string;
                  name: string;
                }>;
              };
            }
          | {
              __typename: 'JobRoleParticipant';
              jobRoleParticipant: {
                __typename?: 'JobRole';
                id: string;
                contact?: {
                  __typename?: 'Contact';
                  id: string;
                  name?: string | null;
                  firstName?: string | null;
                  lastName?: string | null;
                  profilePhotoUrl?: string | null;
                } | null;
              };
            }
          | {
              __typename: 'OrganizationParticipant';
              organizationParticipant: {
                __typename?: 'Organization';
                id: string;
                name: string;
              };
            }
          | { __typename: 'PhoneNumberParticipant' }
          | {
              __typename: 'UserParticipant';
              userParticipant: {
                __typename?: 'User';
                id: string;
                firstName: string;
                lastName: string;
                profilePhotoUrl?: string | null;
              };
            }
        >;
        sentTo: Array<
          | {
              __typename: 'ContactParticipant';
              contactParticipant: {
                __typename?: 'Contact';
                id: string;
                name?: string | null;
                firstName?: string | null;
                lastName?: string | null;
                profilePhotoUrl?: string | null;
              };
            }
          | {
              __typename: 'EmailParticipant';
              type?: string | null;
              emailParticipant: {
                __typename?: 'Email';
                email?: string | null;
                id: string;
                contacts: Array<{
                  __typename?: 'Contact';
                  id: string;
                  name?: string | null;
                  firstName?: string | null;
                  lastName?: string | null;
                }>;
                users: Array<{
                  __typename?: 'User';
                  id: string;
                  firstName: string;
                  lastName: string;
                }>;
                organizations: Array<{
                  __typename?: 'Organization';
                  id: string;
                  name: string;
                }>;
              };
            }
          | {
              __typename: 'JobRoleParticipant';
              jobRoleParticipant: {
                __typename?: 'JobRole';
                id: string;
                contact?: {
                  __typename?: 'Contact';
                  id: string;
                  name?: string | null;
                  firstName?: string | null;
                  lastName?: string | null;
                  profilePhotoUrl?: string | null;
                } | null;
              };
            }
          | {
              __typename: 'OrganizationParticipant';
              organizationParticipant: {
                __typename?: 'Organization';
                id: string;
                name: string;
              };
            }
          | { __typename: 'PhoneNumberParticipant' }
          | {
              __typename: 'UserParticipant';
              userParticipant: {
                __typename?: 'User';
                id: string;
                firstName: string;
                lastName: string;
                profilePhotoUrl?: string | null;
              };
            }
        >;
        interactionSession?: {
          __typename?: 'InteractionSession';
          name: string;
          events: Array<{
            __typename?: 'InteractionEvent';
            id: string;
            channel?: string | null;
            date: any;
            sentBy: Array<
              | {
                  __typename: 'ContactParticipant';
                  contactParticipant: {
                    __typename?: 'Contact';
                    id: string;
                    name?: string | null;
                    firstName?: string | null;
                    lastName?: string | null;
                    profilePhotoUrl?: string | null;
                  };
                }
              | {
                  __typename: 'EmailParticipant';
                  type?: string | null;
                  emailParticipant: {
                    __typename?: 'Email';
                    email?: string | null;
                    id: string;
                    contacts: Array<{
                      __typename?: 'Contact';
                      id: string;
                      name?: string | null;
                      firstName?: string | null;
                      lastName?: string | null;
                    }>;
                    users: Array<{
                      __typename?: 'User';
                      id: string;
                      firstName: string;
                      lastName: string;
                    }>;
                    organizations: Array<{
                      __typename?: 'Organization';
                      id: string;
                      name: string;
                    }>;
                  };
                }
              | {
                  __typename: 'JobRoleParticipant';
                  jobRoleParticipant: {
                    __typename?: 'JobRole';
                    id: string;
                    contact?: {
                      __typename?: 'Contact';
                      id: string;
                      name?: string | null;
                      firstName?: string | null;
                      lastName?: string | null;
                      profilePhotoUrl?: string | null;
                    } | null;
                  };
                }
              | {
                  __typename: 'OrganizationParticipant';
                  organizationParticipant: {
                    __typename?: 'Organization';
                    id: string;
                    name: string;
                  };
                }
              | { __typename?: 'PhoneNumberParticipant' }
              | {
                  __typename: 'UserParticipant';
                  userParticipant: {
                    __typename?: 'User';
                    id: string;
                    firstName: string;
                    lastName: string;
                    profilePhotoUrl?: string | null;
                  };
                }
            >;
          }>;
        } | null;
      }
    | { __typename: 'InteractionSession' }
    | { __typename: 'Issue' }
    | {
        __typename: 'LogEntry';
        id: string;
        createdAt: any;
        updatedAt: any;
        source: Types.DataSource;
        content?: string | null;
        contentType?: string | null;
        logEntryStartedAt: any;
        logEntryCreatedBy?: {
          __typename: 'User';
          id: string;
          firstName: string;
          lastName: string;
          profilePhotoUrl?: string | null;
          emails?: Array<{
            __typename?: 'Email';
            email?: string | null;
          }> | null;
        } | null;
        tags: Array<{ __typename?: 'Tag'; id: string; name: string }>;
        externalLinks: Array<{
          __typename?: 'ExternalSystem';
          type: Types.ExternalSystemType;
          externalUrl?: string | null;
          externalSource?: string | null;
        }>;
      }
    | {
        __typename: 'Meeting';
        id: string;
        name?: string | null;
        createdAt: any;
        updatedAt: any;
        startedAt?: any | null;
        endedAt?: any | null;
        agenda?: string | null;
        status: Types.MeetingStatus;
        attendedBy: Array<
          | {
              __typename: 'ContactParticipant';
              contactParticipant: {
                __typename?: 'Contact';
                id: string;
                name?: string | null;
                firstName?: string | null;
                lastName?: string | null;
                profilePhotoUrl?: string | null;
                timezone?: string | null;
                emails: Array<{
                  __typename?: 'Email';
                  id: string;
                  email?: string | null;
                  rawEmail?: string | null;
                  primary: boolean;
                }>;
              };
            }
          | { __typename?: 'EmailParticipant' }
          | {
              __typename: 'OrganizationParticipant';
              organizationParticipant: {
                __typename?: 'Organization';
                id: string;
                name: string;
                emails: Array<{
                  __typename?: 'Email';
                  id: string;
                  email?: string | null;
                  rawEmail?: string | null;
                  primary: boolean;
                }>;
              };
            }
          | {
              __typename: 'UserParticipant';
              userParticipant: {
                __typename?: 'User';
                id: string;
                firstName: string;
                lastName: string;
                profilePhotoUrl?: string | null;
                emails?: Array<{
                  __typename?: 'Email';
                  id: string;
                  email?: string | null;
                  rawEmail?: string | null;
                  primary: boolean;
                }> | null;
              };
            }
        >;
        createdBy: Array<
          | {
              __typename: 'ContactParticipant';
              contactParticipant: {
                __typename?: 'Contact';
                id: string;
                name?: string | null;
                firstName?: string | null;
                lastName?: string | null;
                profilePhotoUrl?: string | null;
                timezone?: string | null;
                emails: Array<{
                  __typename?: 'Email';
                  id: string;
                  email?: string | null;
                  rawEmail?: string | null;
                  primary: boolean;
                }>;
              };
            }
          | { __typename?: 'EmailParticipant' }
          | {
              __typename: 'OrganizationParticipant';
              organizationParticipant: {
                __typename?: 'Organization';
                id: string;
                name: string;
                emails: Array<{
                  __typename?: 'Email';
                  id: string;
                  email?: string | null;
                  rawEmail?: string | null;
                  primary: boolean;
                }>;
              };
            }
          | {
              __typename: 'UserParticipant';
              userParticipant: {
                __typename?: 'User';
                id: string;
                firstName: string;
                lastName: string;
                profilePhotoUrl?: string | null;
                emails?: Array<{
                  __typename?: 'Email';
                  id: string;
                  email?: string | null;
                  rawEmail?: string | null;
                  primary: boolean;
                }> | null;
              };
            }
        >;
        note: Array<{
          __typename?: 'Note';
          id: string;
          content?: string | null;
        }>;
      }
    | { __typename: 'Note' }
    | { __typename: 'PageView' }
  >;
};

export const GetTimelineEventDocument = `
    query GetTimelineEvent($ids: [ID!]!) {
  timelineEvents(ids: $ids) {
    __typename
    ... on Action {
      __typename
      id
      actionType
      appSource
      createdAt
      metadata
      actionCreatedBy: createdBy {
        ... on User {
          __typename
          id
          firstName
          lastName
          profilePhotoUrl
        }
      }
      content
    }
    ... on InteractionEvent {
      id
      date: createdAt
      channel
      content
      contentType
      includes {
        id
        mimeType
        name
        extension
      }
      issue {
        externalLinks {
          type
          externalId
          externalUrl
        }
      }
      externalLinks {
        externalUrl
        type
      }
      repliesTo {
        id
      }
      summary {
        id
        content
        contentType
      }
      actionItems {
        id
        content
      }
      sentBy {
        __typename
        ...InteractionEventParticipantFragment
      }
      sentTo {
        __typename
        ...InteractionEventParticipantFragment
      }
      interactionSession {
        name
        events {
          ... on InteractionEvent {
            id
            date: createdAt
            channel
            sentBy {
              ...InteractionEventParticipantFragment
            }
          }
        }
      }
      source
    }
    ... on Meeting {
      id
      name
      createdAt
      updatedAt
      startedAt
      endedAt
      attendedBy {
        ...MeetingParticipantFragment
      }
      createdBy {
        ...MeetingParticipantFragment
      }
      note {
        id
        content
      }
      agenda
      status
    }
    ... on LogEntry {
      id
      createdAt
      updatedAt
      logEntryStartedAt: startedAt
      logEntryCreatedBy: createdBy {
        ... on User {
          __typename
          id
          firstName
          lastName
          profilePhotoUrl
          emails {
            email
          }
        }
      }
      tags {
        id
        name
      }
      source
      content
      contentType
      externalLinks {
        type
        externalUrl
        externalSource
      }
    }
  }
}
    ${InteractionEventParticipantFragmentFragmentDoc}
${MeetingParticipantFragmentFragmentDoc}`;
export const useGetTimelineEventQuery = <
  TData = GetTimelineEventQuery,
  TError = unknown,
>(
  client: GraphQLClient,
  variables: GetTimelineEventQueryVariables,
  options?: UseQueryOptions<GetTimelineEventQuery, TError, TData>,
  headers?: RequestInit['headers'],
) =>
  useQuery<GetTimelineEventQuery, TError, TData>(
    ['GetTimelineEvent', variables],
    fetcher<GetTimelineEventQuery, GetTimelineEventQueryVariables>(
      client,
      GetTimelineEventDocument,
      variables,
      headers,
    ),
    options,
  );
useGetTimelineEventQuery.document = GetTimelineEventDocument;

useGetTimelineEventQuery.getKey = (
  variables: GetTimelineEventQueryVariables,
) => ['GetTimelineEvent', variables];
export const useInfiniteGetTimelineEventQuery = <
  TData = GetTimelineEventQuery,
  TError = unknown,
>(
  pageParamKey: keyof GetTimelineEventQueryVariables,
  client: GraphQLClient,
  variables: GetTimelineEventQueryVariables,
  options?: UseInfiniteQueryOptions<GetTimelineEventQuery, TError, TData>,
  headers?: RequestInit['headers'],
) =>
  useInfiniteQuery<GetTimelineEventQuery, TError, TData>(
    ['GetTimelineEvent.infinite', variables],
    (metaData) =>
      fetcher<GetTimelineEventQuery, GetTimelineEventQueryVariables>(
        client,
        GetTimelineEventDocument,
        { ...variables, ...(metaData.pageParam ?? {}) },
        headers,
      )(),
    options,
  );

useInfiniteGetTimelineEventQuery.getKey = (
  variables: GetTimelineEventQueryVariables,
) => ['GetTimelineEvent.infinite', variables];
useGetTimelineEventQuery.fetcher = (
  client: GraphQLClient,
  variables: GetTimelineEventQueryVariables,
  headers?: RequestInit['headers'],
) =>
  fetcher<GetTimelineEventQuery, GetTimelineEventQueryVariables>(
    client,
    GetTimelineEventDocument,
    variables,
    headers,
  );
