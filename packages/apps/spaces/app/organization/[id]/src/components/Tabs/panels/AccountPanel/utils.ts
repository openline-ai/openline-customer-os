import { QueryClient } from '@tanstack/react-query';

import { SelectOption } from '@shared/types/SelectOptions';
import { useOrganizationAccountDetailsQuery } from '@organization/src/graphql/getAccountPanelDetails.generated';
import {
  Maybe,
  BilledType,
  ContractRenewalCycle,
  ContractBillingCycle,
  OpportunityRenewalLikelihood,
} from '@graphql/types';

export const invalidateAccountDetailsQuery = (
  queryClient: QueryClient,
  id: string,
) =>
  queryClient.invalidateQueries({
    queryKey: useOrganizationAccountDetailsQuery.getKey({ id }),
  });

export function getRenewalLikelihoodColor(
  renewalLikelihood?: Maybe<OpportunityRenewalLikelihood> | undefined,
) {
  switch (renewalLikelihood) {
    case OpportunityRenewalLikelihood.HighRenewal:
      return 'success';
    case OpportunityRenewalLikelihood.MediumRenewal:
      return 'warning';
    case OpportunityRenewalLikelihood.LowRenewal:
      return 'error';
    case OpportunityRenewalLikelihood.ZeroRenewal:
      return 'gray';
    default:
      return 'gray';
  }
}
export function getRenewalLikelihoodLabel(
  renewalLikelihood?: Maybe<OpportunityRenewalLikelihood> | undefined,
) {
  switch (renewalLikelihood) {
    case OpportunityRenewalLikelihood.HighRenewal:
      return 'High';
    case OpportunityRenewalLikelihood.MediumRenewal:
      return 'Medium';
    case OpportunityRenewalLikelihood.LowRenewal:
      return 'Low';
    case OpportunityRenewalLikelihood.ZeroRenewal:
      return 'Zero';
    default:
      return '';
  }
}
export const billingFrequencyOptions: SelectOption<
  ContractRenewalCycle | 'MULTI_YEAR'
>[] = [
  { label: 'Monthly', value: ContractRenewalCycle.MonthlyRenewal },
  { label: 'Quarterly', value: ContractRenewalCycle.QuarterlyRenewal },
  { label: 'Annual', value: ContractRenewalCycle.AnnualRenewal },
  { label: 'Multi-year', value: 'MULTI_YEAR' },
];
export const contractBillingCycleOptions: SelectOption<ContractBillingCycle>[] =
  [
    { label: 'Monthly', value: ContractBillingCycle.MonthlyBilling },
    { label: 'Quarterly', value: ContractBillingCycle.QuarterlyBilling },
    { label: 'Annually', value: ContractBillingCycle.AnnualBilling },
  ];

export const billedTypeOptions: SelectOption<BilledType>[] = [
  { label: 'Once', value: BilledType.Once },
  { label: 'Usage', value: BilledType.Usage },
  { label: 'Monthly', value: BilledType.Monthly },
  { label: 'Quarterly', value: BilledType.Quarterly },
  { label: 'Annually', value: BilledType.Annually },
];

export const currencyOptions: SelectOption<string>[] = [
  { label: 'United States Dollar', value: 'USD' },
];

export const autorenewalOptions = [
  { label: 'Does auto-renew', value: true },
  { label: 'Does not auto-renew', value: false },
];
export const paymentDueOptions: SelectOption<number>[] = [
  { label: 'On receipt', value: 0 },
  { label: 'Net 15', value: 15 },
  { label: 'Net 30', value: 30 },
  { label: 'Net 45', value: 45 },
  { label: 'Net 60', value: 60 },
  { label: 'Net 90', value: 90 },
];
