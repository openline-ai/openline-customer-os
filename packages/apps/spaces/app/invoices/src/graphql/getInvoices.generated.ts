// @ts-nocheck remove this when typscript-react-query plugin is fixed
import * as Types from '../../../src/types/__generated__/graphql.types';

import { GraphQLClient } from 'graphql-request';
import { RequestInit } from 'graphql-request/dist/types.dom';
import {
  useQuery,
  useInfiniteQuery,
  UseQueryOptions,
  UseInfiniteQueryOptions,
  InfiniteData,
} from '@tanstack/react-query';

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
export type GetInvoicesQueryVariables = Types.Exact<{
  pagination: Types.Pagination;
}>;

export type GetInvoicesQuery = {
  __typename?: 'Query';
  invoices: {
    __typename?: 'InvoicesPage';
    totalPages: number;
    totalElements: any;
    content: Array<{
      __typename?: 'Invoice';
      id: string;
      number: string;
      date: any;
      dueDate: any;
      total: number;
      currency: string;
      dryRun: boolean;
      invoiceLines: Array<{
        __typename?: 'InvoiceLine';
        id: string;
        name: string;
        price: number;
        quantity: number;
        amount: number;
        vat: number;
        total: number;
      }>;
    }>;
  };
};

export const GetInvoicesDocument = `
    query getInvoices($pagination: Pagination!) {
  invoices(pagination: $pagination) {
    content {
      id
      number
      date
      dueDate
      total
      currency
      dryRun
      invoiceLines {
        id
        name
        price
        quantity
        amount
        vat
        total
      }
    }
    totalPages
    totalElements
  }
}
    `;

export const useGetInvoicesQuery = <TData = GetInvoicesQuery, TError = unknown>(
  client: GraphQLClient,
  variables: GetInvoicesQueryVariables,
  options?: Omit<
    UseQueryOptions<GetInvoicesQuery, TError, TData>,
    'queryKey'
  > & {
    queryKey?: UseQueryOptions<GetInvoicesQuery, TError, TData>['queryKey'];
  },
  headers?: RequestInit['headers'],
) => {
  return useQuery<GetInvoicesQuery, TError, TData>({
    queryKey: ['getInvoices', variables],
    queryFn: fetcher<GetInvoicesQuery, GetInvoicesQueryVariables>(
      client,
      GetInvoicesDocument,
      variables,
      headers,
    ),
    ...options,
  });
};

useGetInvoicesQuery.document = GetInvoicesDocument;

useGetInvoicesQuery.getKey = (variables: GetInvoicesQueryVariables) => [
  'getInvoices',
  variables,
];

export const useInfiniteGetInvoicesQuery = <
  TData = InfiniteData<GetInvoicesQuery>,
  TError = unknown,
>(
  client: GraphQLClient,
  variables: GetInvoicesQueryVariables,
  options: Omit<
    UseInfiniteQueryOptions<GetInvoicesQuery, TError, TData>,
    'queryKey'
  > & {
    queryKey?: UseInfiniteQueryOptions<
      GetInvoicesQuery,
      TError,
      TData
    >['queryKey'];
  },
  headers?: RequestInit['headers'],
) => {
  return useInfiniteQuery<GetInvoicesQuery, TError, TData>(
    (() => {
      const { queryKey: optionsQueryKey, ...restOptions } = options;
      return {
        queryKey: optionsQueryKey ?? ['getInvoices.infinite', variables],
        queryFn: (metaData) =>
          fetcher<GetInvoicesQuery, GetInvoicesQueryVariables>(
            client,
            GetInvoicesDocument,
            { ...variables, ...(metaData.pageParam ?? {}) },
            headers,
          )(),
        ...restOptions,
      };
    })(),
  );
};

useInfiniteGetInvoicesQuery.getKey = (variables: GetInvoicesQueryVariables) => [
  'getInvoices.infinite',
  variables,
];

useGetInvoicesQuery.fetcher = (
  client: GraphQLClient,
  variables: GetInvoicesQueryVariables,
  headers?: RequestInit['headers'],
) =>
  fetcher<GetInvoicesQuery, GetInvoicesQueryVariables>(
    client,
    GetInvoicesDocument,
    variables,
    headers,
  );

useGetInvoicesQuery.mutateCacheEntry =
  (queryClient: QueryClient, variables: GetInvoicesQueryVariables) =>
  (mutator: (cacheEntry: GetInvoicesQuery) => GetInvoicesQuery) => {
    const cacheKey = useGetInvoicesQuery.getKey(variables);
    const previousEntries =
      queryClient.getQueryData<GetInvoicesQuery>(cacheKey);
    if (previousEntries) {
      queryClient.setQueryData<GetInvoicesQuery>(cacheKey, mutator);
    }
    return { previousEntries };
  };
useInfiniteGetInvoicesQuery.mutateCacheEntry =
  (queryClient: QueryClient, variables: GetInvoicesQueryVariables) =>
  (
    mutator: (
      cacheEntry: InfiniteData<GetInvoicesQuery>,
    ) => InfiniteData<GetInvoicesQuery>,
  ) => {
    const cacheKey = useInfiniteGetInvoicesQuery.getKey(variables);
    const previousEntries =
      queryClient.getQueryData<InfiniteData<GetInvoicesQuery>>(cacheKey);
    if (previousEntries) {
      queryClient.setQueryData<InfiniteData<GetInvoicesQuery>>(
        cacheKey,
        mutator,
      );
    }
    return { previousEntries };
  };