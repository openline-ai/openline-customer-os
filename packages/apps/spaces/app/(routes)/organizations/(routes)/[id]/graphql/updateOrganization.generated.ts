// @ts-nocheck remove this when typscript-react-query plugin is fixed
import * as Types from '../../../../../types/__generated__/graphql.types';

import { GraphQLClient } from 'graphql-request';
import { RequestInit } from 'graphql-request/dist/types.dom';
import { useMutation, UseMutationOptions } from '@tanstack/react-query';

function fetcher<TData, TVariables extends { [key: string]: any }>(
  client: GraphQLClient,
  query: string,
  variables?: TVariables,
  requestHeaders?: RequestInit['headers'],
) {
  return async (): Promise<TData> =>
    client.request({
      document: query,
      variables,
      requestHeaders,
    });
}
export type UpdateOrganizationMutationVariables = Types.Exact<{
  input: Types.OrganizationUpdateInput;
}>;

export type UpdateOrganizationMutation = {
  __typename?: 'Mutation';
  organization_Update: {
    __typename?: 'Organization';
    id: string;
    name: string;
    description?: string | null;
    domain?: string | null;
    domains: Array<string>;
    website?: string | null;
    industry?: string | null;
    isPublic?: boolean | null;
    market?: Types.Market | null;
    employees?: any | null;
  };
};

export const UpdateOrganizationDocument = `
    mutation updateOrganization($input: OrganizationUpdateInput!) {
  organization_Update(input: $input) {
    id
    name
    description
    domain
    domains
    website
    industry
    isPublic
    market
    employees
  }
}
    `;
export const useUpdateOrganizationMutation = <
  TError = unknown,
  TContext = unknown,
>(
  client: GraphQLClient,
  options?: UseMutationOptions<
    UpdateOrganizationMutation,
    TError,
    UpdateOrganizationMutationVariables,
    TContext
  >,
  headers?: RequestInit['headers'],
) =>
  useMutation<
    UpdateOrganizationMutation,
    TError,
    UpdateOrganizationMutationVariables,
    TContext
  >(
    ['updateOrganization'],
    (variables?: UpdateOrganizationMutationVariables) =>
      fetcher<UpdateOrganizationMutation, UpdateOrganizationMutationVariables>(
        client,
        UpdateOrganizationDocument,
        variables,
        headers,
      )(),
    options,
  );
useUpdateOrganizationMutation.getKey = () => ['updateOrganization'];

useUpdateOrganizationMutation.fetcher = (
  client: GraphQLClient,
  variables: UpdateOrganizationMutationVariables,
  headers?: RequestInit['headers'],
) =>
  fetcher<UpdateOrganizationMutation, UpdateOrganizationMutationVariables>(
    client,
    UpdateOrganizationDocument,
    variables,
    headers,
  );
