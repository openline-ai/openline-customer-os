'use client';
import React, { useState } from 'react';
import { useParams } from 'next/navigation';

import { useConnections } from '@integration-app/react';

import { Flex } from '@ui/layout/Flex';
import { Issue } from '@graphql/types';
import { VStack } from '@ui/layout/Stack';
import { Link } from '@ui/navigation/Link';
import { Fade } from '@ui/transitions/Fade';
import { Heading } from '@ui/typography/Heading';
import { Collapse } from '@ui/transitions/Collapse';
import { ChevronUp } from '@ui/media/icons/ChevronUp';
import { ChevronDown } from '@ui/media/icons/ChevronDown';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { useGetIssuesQuery } from '@organization/src/graphql/getIssues.generated';
import { IssueCard } from '@organization/src/components/Tabs/panels/IssuesPanel/IssueCard/IssueCard';
import { IssuesPanelSkeleton } from '@organization/src/components/Tabs/panels/IssuesPanel/IssuesPanelSkeleton';
import { OrganizationPanel } from '@organization/src/components/Tabs/panels/OrganizationPanel/OrganizationPanel';
import { EmptyIssueMessage } from '@organization/src/components/Tabs/panels/IssuesPanel/EmptyIssueMessage/EmptyIssueMessage';

import { ChannelLinkSelect } from './ChannelLinkSelect';

export const NEW_DATE = new Date(new Date().setDate(new Date().getDate() + 1));

function filterIssues(issues: Array<Issue>): {
  open: Array<Issue>;
  closed: Array<Issue>;
} {
  return issues.reduce(
    (
      acc: {
        open: Array<Issue>;
        closed: Array<Issue>;
      },
      issue,
    ) => {
      if (['closed', 'solved'].includes(issue.status.toLowerCase())) {
        acc.closed.push(issue);
      } else {
        acc.open.push(issue);
      }

      return acc;
    },
    { open: [], closed: [] },
  );
}

export const IssuesPanel = () => {
  const id = useParams()?.id as string;
  const client = getGraphQLClient();
  const [isExpanded, setIsExpanded] = useState(true);
  const { data, isLoading } = useGetIssuesQuery(client, {
    organizationId: id,
    from: NEW_DATE,
    size: 50,
  });
  const issues: Array<Issue> =
    (data?.organization?.timelineEvents as Array<Issue>) ?? [];
  const { open: openIssues, closed: closedIssues } = filterIssues(issues);
  const { items, loading } = useConnections();
  const connections = items
    .map((item) => item.integration?.key)
    .filter((item) => ['unthread', 'zendesk'].includes(item ?? ''));

  if (loading || isLoading) {
    return <IssuesPanelSkeleton />;
  }

  if (!connections.length) {
    return (
      <OrganizationPanel title='Issues' withFade>
        <EmptyIssueMessage title='Connect your customer support app'>
          To see your customers support issues here,{' '}
          <Link color='primary.600' as='span' href='/settings?tab=integrations'>
            Go to settings
          </Link>{' '}
          and connect an app like Zendesk or Unthread.
        </EmptyIssueMessage>
      </OrganizationPanel>
    );
  }

  if (connections?.[0] === 'unthread' && !issues.length) {
    return (
      <OrganizationPanel
        title='Issues'
        withFade
        actionItem={<ChannelLinkSelect from={NEW_DATE} />}
      >
        <EmptyIssueMessage title='Link an Unthread Slack channel'>
          Show your Unthread support issues here by linking a Slack channel.
        </EmptyIssueMessage>
      </OrganizationPanel>
    );
  }

  if (!issues.length) {
    return (
      <OrganizationPanel
        title='Issues'
        withFade
        actionItem={<ChannelLinkSelect from={NEW_DATE} />}
      >
        <EmptyIssueMessage
          title='No issues detected'
          description={`It looks like ${
            data?.organization?.name ?? '[Unknown]'
          } has had a smooth journey thus far. Or
      perhaps they’ve been shy about reporting issues. Stay proactive and keep
      monitoring for optimal support.`}
        />
      </OrganizationPanel>
    );
  }

  return (
    <OrganizationPanel
      title='Issues'
      withFade
      actionItem={<ChannelLinkSelect from={NEW_DATE} />}
    >
      <Flex as='article' w='full' direction='column'>
        <Heading fontWeight='semibold' fontSize='md' mb={2}>
          Open
        </Heading>
        <VStack>
          {!!openIssues?.length &&
            openIssues.map((issue, index) => (
              <Fade
                key={`issue-panel-${issue.id}`}
                in
                style={{ width: '100%' }}
              >
                <IssueCard issue={issue} />
              </Fade>
            ))}
        </VStack>
      </Flex>
      {!openIssues.length && (
        <EmptyIssueMessage
          description={`It looks like ${
            data?.organization?.name ?? '[Unknown]'
          } has no open issues at the moment`}
        />
      )}
      {!!closedIssues.length && (
        <Flex as='article' w='full' direction='column' mt={2}>
          <Flex
            justifyContent='space-between'
            alignItems='center'
            w='full'
            as='button'
            pb={2}
            onClick={() => setIsExpanded((prev) => !prev)}
          >
            <Heading fontWeight='semibold' fontSize='md'>
              Closed
            </Heading>
            {isExpanded ? <ChevronDown /> : <ChevronUp />}
          </Flex>

          <Collapse
            in={isExpanded}
            style={{ overflow: 'unset' }}
            delay={{
              exit: 2,
            }}
          >
            <Fade
              in={isExpanded}
              delay={{
                enter: 0.2,
              }}
            >
              {!closedIssues.length && (
                <EmptyIssueMessage
                  description={`It looks like ${
                    data?.organization?.name ?? '[Unknown]'
                  } has no closed issues at the moment`}
                />
              )}
              {!!closedIssues?.length && (
                <VStack>
                  {closedIssues.map((issue, index) => (
                    <IssueCard issue={issue} key={`issue-panel-${issue.id}`} />
                  ))}
                </VStack>
              )}
            </Fade>
          </Collapse>
        </Flex>
      )}
    </OrganizationPanel>
  );
};
