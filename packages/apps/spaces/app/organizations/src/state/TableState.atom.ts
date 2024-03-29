import { selector, useRecoilValue } from 'recoil';

import { OwnerFilterSelector } from '@organizations/components/Columns/Filters/Owner/OwnerFilter.atom';
import { WebsiteFilterSelector } from '@organizations/components/Columns/Filters/Website/WebsiteFilter.atom';
import { ForecastFilterSelector } from '@organizations/components/Columns/Filters/Forecast/ForecastFilter.atom';
import { OnboardingFilterSelector } from '@organizations/components/Columns/Filters/Onboarding/OnboardingFilter.atom';
import { OrganizationFilterSelector } from '@organizations/components/Columns/Filters/Organization/OrganizationFilter.atom';
import { RelationshipFilterSelector } from '@organizations/components/Columns/Filters/Relationship/RelationshipFilter.atom';
import { LastTouchpointSelector } from '@organizations/components/Columns/Filters/LastTouchpoint/LastTouchpointFilter.atom';
import { TimeToRenewalFilterSelector } from '@organizations/components/Columns/Filters/TimeToRenewal/TimeToRenewalFilter.atom';
import { RenewalLikelihoodFilterSelector } from '@organizations/components/Columns/Filters/RenewalLikelihood/RenewalLikelihoodFilter.atom';

const tableStateSelector = selector({
  key: 'tableState',
  get: ({ get }) => {
    const owner = get(OwnerFilterSelector);
    const website = get(WebsiteFilterSelector);
    const forecast = get(ForecastFilterSelector);
    const onboarding = get(OnboardingFilterSelector);
    const organization = get(OrganizationFilterSelector);
    const relationship = get(RelationshipFilterSelector);
    const renewalLikelihood = get(RenewalLikelihoodFilterSelector);

    const timeToRenewal = (() => {
      const state = get(TimeToRenewalFilterSelector);
      const value = new Date(state.value).toISOString();

      return {
        ...state,
        value,
      };
    })();

    const lastTouchpoint = (() => {
      const state = get(LastTouchpointSelector);
      const after = state.after
        ? new Date(state.after).toISOString()
        : undefined;

      return {
        ...state,
        after,
      };
    })();

    return {
      columnFilters: {
        owner,
        website,
        forecast,
        onboarding,
        organization,
        relationship,
        timeToRenewal,
        lastTouchpoint,
        renewalLikelihood,
      },
    };
  },
});

export const useTableState = () => {
  return useRecoilValue(tableStateSelector);
};
