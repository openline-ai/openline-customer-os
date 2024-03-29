'use client';

import { Flex } from '@ui/layout/Flex';

import { CustomerMap } from './src/components/charts/CustomerMap';
import { ARRBreakdown } from './src/components/charts/ARRBreakdown';
import { NewCustomers } from './src/components/charts/NewCustomers';
import { RevenueAtRisk } from './src/components/charts/RevenueAtRisk';
import { RetentionRate } from './src/components/charts/RetentionRate';
import { TimeToOnboard } from './src/components/charts/TimeToOnboard';
import { MrrPerCustomer } from './src/components/charts/MrrPerCustomer';
import { OnboardingCompletion } from './src/components/charts/OnboardingCompletion';
import { GrossRevenueRetention } from './src/components/charts/GrossRevenueRetention';

export default function DashboardPage() {
  return (
    <Flex flexDir='column' pl='3' pt='4'>
      <Flex mb='6'>
        <CustomerMap />
      </Flex>

      <Flex gap='3' mb='3'>
        <MrrPerCustomer />
        <GrossRevenueRetention />
      </Flex>

      <Flex gap='3' mb='3'>
        <ARRBreakdown />
        <RevenueAtRisk />
      </Flex>

      <Flex gap='3' mb='3'>
        <NewCustomers />
        <RetentionRate />
      </Flex>

      <Flex gap='3'>
        <TimeToOnboard />
        <OnboardingCompletion />
      </Flex>
    </Flex>
  );
}
