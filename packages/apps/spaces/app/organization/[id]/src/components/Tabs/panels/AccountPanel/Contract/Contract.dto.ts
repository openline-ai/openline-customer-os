import { SelectOption } from '@shared/types/SelectOptions';
import { UpdateContractMutationVariables } from '@organization/src/graphql/updateContract.generated';
import {
  Contract,
  ContractUpdateInput,
  ContractRenewalCycle,
  ContractBillingCycle,
} from '@graphql/types';
import {
  billingFrequencyOptions,
  contractBillingCycleOptions,
} from '@organization/src/components/Tabs/panels/AccountPanel/utils';

export interface TimeToRenewalData {
  name?: string;
  endedAt?: Date;
  serviceStarted?: Date;
  invoicingStartDate?: Date;
  organizationName?: string;
  contractUrl?: string | null;

  billingEnabled?: boolean | null;
  committedPeriods?: number | null;
  billingCycle?: ContractBillingCycle | null;
  contractRenewalCycle?: ContractRenewalCycle;
}
export interface TimeToRenewalForm {
  name?: string;
  endedAt?: Date;
  serviceStarted?: Date;
  invoicingStartDate?: Date;
  contractUrl?: string | null;
  committedPeriods?: string | null;
  country?: SelectOption<string> | null;
  organizationLegalName?: string | null;
  billingEnabled?: SelectOption<boolean> | null;
  billingCycle?: SelectOption<ContractBillingCycle> | null;
  contractRenewalCycle?: SelectOption<
    ContractRenewalCycle | 'MULTI_YEAR'
  > | null;
}

export class ContractDTO implements TimeToRenewalForm {
  endedAt?: Date;
  invoicingStartDate?: Date;
  serviceStarted?: Date;
  contractRenewalCycle?: SelectOption<
    ContractRenewalCycle | 'MULTI_YEAR'
  > | null;
  name?: string;
  contractUrl?: string | null;
  renewalPeriods?: string | null;
  billingCycle?: SelectOption<ContractBillingCycle> | null;
  billingEnabled?: SelectOption<boolean> | null;

  constructor(data?: Contract | null) {
    this.contractRenewalCycle =
      [...billingFrequencyOptions].find(({ value }) =>
        (data?.committedPeriods ?? 0) > 1
          ? value === 'MULTI_YEAR'
          : value === data?.contractRenewalCycle,
      ) ?? undefined;
    this.billingEnabled = data?.billingEnabled
      ? { label: 'Enabled', value: true }
      : { label: 'Disabled', value: false };
    this.billingCycle =
      [...contractBillingCycleOptions].find(
        ({ value }) => value === data?.billingDetails?.billingCycle,
      ) ?? undefined;
    this.endedAt = data?.contractEnded && new Date(data.contractEnded);
    this.invoicingStartDate =
      data?.billingDetails?.invoicingStarted &&
      new Date(data.billingDetails?.invoicingStarted);
    this.serviceStarted = data?.serviceStarted && new Date(data.serviceStarted);
    this.name = data?.contractName?.length
      ? data?.contractName
      : `${
          data?.billingDetails?.organizationLegalName?.length
            ? `${data?.billingDetails?.organizationLegalName}'s`
            : "Unnamed's"
        } contract`;
    this.contractUrl = data?.contractUrl ?? '';
    this.renewalPeriods = String(data?.committedPeriods ?? 2);
  }

  static toForm(data?: Contract | null): TimeToRenewalForm {
    const formData = new ContractDTO(data);

    return {
      ...formData,
    };
  }

  static toPayload(
    data: Partial<ContractUpdateInput> & { contractId: string },
  ): UpdateContractMutationVariables {
    return {
      input: {
        patch: true,
        ...data,
      },
    };
  }
}
