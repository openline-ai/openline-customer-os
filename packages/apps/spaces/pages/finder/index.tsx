import type { NextPage } from 'next';
import React, { useState } from 'react';
import { Finder } from '../../src/components/finder';
import { PageContentLayout } from '../../src/components/ui-kit/layouts';
import { SidePanel } from '../../src/components/ui-kit/organisms';
import { WebChat } from '@openline-ai/openline-web-chat';
import { useRouter } from 'next/router';

const FinderComponent: NextPage = () => {
  const router = useRouter();
  const [isSidePanelVisible, setSidePanelVisible] = useState(false);

  return (
    <PageContentLayout
      isPanelOpen={isSidePanelVisible}
      isSideBarShown={router.pathname === '/'}
    >
      {router.pathname === '/' && (
        <SidePanel
          onPanelToggle={setSidePanelVisible}
          isPanelOpen={isSidePanelVisible}
          logoutUrl={undefined}
        >
          <WebChat
            apikey={`${process.env.WEB_CHAT_API_KEY}`}
            httpServerPath={`${process.env.WEB_CHAT_HTTP_PATH}`}
            wsServerPath={`${process.env.WEB_CHAT_WS_PATH}`}
            location='left'
            trackerEnabled={
              `${process.env.WEB_CHAT_TRACKER_ENABLED}` === 'true'
            }
            trackerAppId={`${process.env.WEB_CHAT_TRACKER_APP_ID}`}
            trackerId={`${process.env.WEB_CHAT_TRACKER_ID}`}
            trackerCollectorUrl={`${process.env.WEB_CHAT_TRACKER_COLLECTOR_URL}`}
            trackerBufferSize={`${process.env.WEB_CHAT_TRACKER_BUFFER_SIZE}`}
            trackerMinimumVisitLength={`${process.env.WEB_CHAT_TRACKER_MINIMUM_VISIT_LENGTH}`}
            trackerHeartbeatDelay={`${process.env.WEB_CHAT_TRACKER_HEARTBEAT_DELAY}`}
            userEmail={'userEmail' || ''}
          />
        </SidePanel>
      )}
      <article style={{ gridArea: 'content' }}>
        <Finder />
      </article>
    </PageContentLayout>
  );
};

export default FinderComponent;
