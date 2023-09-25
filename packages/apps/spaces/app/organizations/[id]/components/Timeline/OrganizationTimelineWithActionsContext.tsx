'use client';
import { TimelineActionContextContextProvider } from '@organization/components/Timeline/TimelineActions/context/TimelineActionContext';
import { OrganizationTimeline } from './OrganizationTimeline';

export const OrganizationTimelineWithActionsContext = () => {
  return (
    <>
      <TimelineActionContextContextProvider>
        <OrganizationTimeline />
      </TimelineActionContextContextProvider>
    </>
  );
};
