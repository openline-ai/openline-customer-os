import {
  CreateOrganizationMutation,
  OrganizationInput,
  useCreateOrganizationMutation,
} from './types';
import { toast } from 'react-toastify';
import { ApolloCache } from 'apollo-cache';
import { GetOrganizationsOptionsDocument } from '../../graphQL/__generated__/generated';
import client from '../../apollo-client';

interface Result {
  onCreateOrganization: (
    input: OrganizationInput,
  ) => Promise<CreateOrganizationMutation['organization_Create'] | null>;
}
export const useCreateOrganization = (): Result => {
  const [createOrganizationMutation, { loading, error, data }] =
    useCreateOrganizationMutation();
  const handleUpdateCacheAfterAddingOrg = (
    cache: ApolloCache<any>,
    { data: { organization_Create } }: any,
  ) => {
    const data: any | null = client.readQuery({
      query: GetOrganizationsOptionsDocument,
    });

    if (data === null) {
      client.writeQuery({
        query: GetOrganizationsOptionsDocument,
        data: {
          organizations: {
            content: [organization_Create],
          },
        },
      });
      return;
    }

    client.writeQuery({
      query: GetOrganizationsOptionsDocument,
      data: {
        organizations: {
          content: {
            organization_Create,
            ...data.organizations?.content,
          },
        },
      },
    });
  };

  const handleCreateOrganization: Result['onCreateOrganization'] = async (
    input: OrganizationInput,
  ) => {
    try {
      const response = await createOrganizationMutation({
        variables: { input },
        // @ts-expect-error fixme
        update: handleUpdateCacheAfterAddingOrg,
      });
      if (response.data?.organization_Create) {
        toast.success('Organization was successfully created!', {
          toastId: `organization-create-success-${response.data.organization_Create.id}`,
        });
      }
      return response.data?.organization_Create ?? null;
    } catch (err) {
      console.error(err);
      toast.error('Something went wrong while adding organization');
      return null;
    }
  };

  return {
    onCreateOrganization: handleCreateOrganization,
  };
};
