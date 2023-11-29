'use client';
import dynamic from 'next/dynamic';

import { ChartCard } from '@customerMap/components/ChartCard';
import ParentSize from '@visx/responsive/lib/components/ParentSize';
import { useMrrPerCustomerQuery } from '@customerMap/graphql/mrrPerCustomer.generated';

import { Skeleton } from '@ui/presentation/Skeleton';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { formatCurrency } from '@spaces/utils/getFormattedCurrencyNumber';

import { PercentageTrend } from '../../PercentageTrend';
import { MrrPerCustomerDatum } from './MrrPerCustomer.chart';

const MrrPerCustomerChart = dynamic(() => import('./MrrPerCustomer.chart'), {
  ssr: false,
});

export const MrrPerCustomer = () => {
  const client = getGraphQLClient();
  const { data, isLoading } = useMrrPerCustomerQuery(client);

  const chartData = (data?.dashboard_MRRPerCustomer?.perMonth ?? []).map(
    (d) => ({
      month: d?.month,
      value: d?.value,
    }),
  ) as MrrPerCustomerDatum[];
  const stat = formatCurrency(
    data?.dashboard_MRRPerCustomer?.mrrPerCustomer ?? 0,
  );
  const percentage = data?.dashboard_MRRPerCustomer?.increasePercentage ?? 0;

  return (
    <ChartCard
      flex='1'
      stat={stat}
      title='MRR per customer'
      renderSubStat={() => <PercentageTrend percentage={percentage} />}
    >
      <ParentSize>
        {({ width }) => (
          <Skeleton
            w='full'
            h='200px'
            endColor='gray.300'
            startColor='gray.300'
            isLoaded={!isLoading}
          >
            <MrrPerCustomerChart width={width} data={chartData} />
          </Skeleton>
        )}
      </ParentSize>
    </ChartCard>
  );
};
