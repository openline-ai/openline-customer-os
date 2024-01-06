'use client';
import { useParams } from 'next/navigation';

import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { useOrganizationQuery } from '@organization/src/graphql/organization.generated';

import { PanelContainer } from './PanelContainer';
import { OnboardingStatus } from './OnboardingStatus';

export const SuccessPanel = () => {
  const client = getGraphQLClient();
  const id = useParams()?.id as string;
  const { data } = useOrganizationQuery(client, { id });

  return (
    <PanelContainer title='Success'>
      <OnboardingStatus data={data?.organization?.accountDetails?.onboarding} />
    </PanelContainer>
  );
};
