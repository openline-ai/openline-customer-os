'use client';

import type { TabProps } from '@ui/disclosure/Tabs';
import { Tabs, TabList, TabPanels, Tab, TabPanel } from '@ui/disclosure/Tabs';

import { OrganizationDetails } from './OrganizationDetails';

const CustomTab = (props: TabProps) => (
  <Tab
    mr='1'
    w='90px'
    fontSize='14px'
    bg='gray.100'
    _selected={{ bg: 'white', fontWeight: 'bold' }}
    {...props}
  />
);

export const OrganizationTabs = () => {
  return (
    <Tabs
      mt='-38px'
      zIndex='1'
      variant='enclosed'
      defaultIndex={4}
      h='full'
      display='flex'
      flexDir='column'
    >
      <TabList>
        <CustomTab>Up Next</CustomTab>
        <CustomTab>Account</CustomTab>
        <CustomTab>Success</CustomTab>
        <CustomTab>People</CustomTab>
        <CustomTab>About</CustomTab>
      </TabList>

      <TabPanels h='full' position='relative' overflowY='auto' flex='1'>
        <TabPanel>
          <p>Up next</p>
        </TabPanel>
        <TabPanel>
          <p>Account</p>
        </TabPanel>
        <TabPanel>
          <p>Success</p>
        </TabPanel>
        <TabPanel>
          <p>People</p>
        </TabPanel>
        <TabPanel h='100%' overflowY='auto' flex='1'>
          <OrganizationDetails />
        </TabPanel>
      </TabPanels>
    </Tabs>
  );
};
