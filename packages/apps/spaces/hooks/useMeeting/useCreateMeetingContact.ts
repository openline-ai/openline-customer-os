import {
  GetContactTimelineDocument,
  NOW_DATE,
  Result,
  useCreateMeetingMutation,
} from './types';
import { toast } from 'react-toastify';
import { ApolloCache } from 'apollo-cache';
import { GetContactTimelineQuery } from '../../graphQL/__generated__/generated';
import client from '../../apollo-client';
import { useRecoilValue } from 'recoil';
import { userData } from '../../state';

export interface Props {
  contactId?: string;
}
export const useCreateMeetingFromContact = ({ contactId }: Props): Result => {
  const [createMeetingMutation, { loading, error, data }] =
    useCreateMeetingMutation();
  const loggedInUserData = useRecoilValue(userData);

  const handleUpdateCacheAfterAddingMeeting = (
    cache: ApolloCache<any>,
    { data: { meeting_Create } }: any,
  ) => {
    const data: GetContactTimelineQuery | null = client.readQuery({
      query: GetContactTimelineDocument,
      variables: {
        contactId,
        from: NOW_DATE,
        size: 10,
      },
    });

    if (data === null) {
      client.writeQuery({
        query: GetContactTimelineDocument,
        data: {
          contact: {
            contactId,
            timelineEvents: [meeting_Create],
          },
          variables: { contactId, from: NOW_DATE, size: 10 },
        },
      });
      return;
    }

    const newData = {
      contact: {
        ...data.contact,
        timelineEvents: [
          ...(data.contact?.timelineEvents || []),
          meeting_Create,
        ],
      },
    };

    client.writeQuery({
      query: GetContactTimelineDocument,
      data: newData,
      variables: {
        contactId,
        from: NOW_DATE,
        size: 10,
      },
    });
  };

  const handleCreateMeetingFromContact: Result['onCreateMeeting'] =
    async () => {
      try {
        const response = await createMeetingMutation({
          variables: {
            meeting: {
              createdBy: [{ userID: loggedInUserData.id, type: 'user' }],
              attendedBy: [{ contactID: contactId, type: 'contact' }],
              appSource: 'OPENLINE',
            },
          },

          //@ts-expect-error fixme
          update: handleUpdateCacheAfterAddingMeeting,
        });

        toast.success(`Added draft meeting to the timeline`);
        return response.data?.meeting_Create ?? null;
      } catch (err) {
        console.error(err);
        toast.error(
          `Something went wrong while adding draft meeting to the timeline`,
        );
        return null;
      }
    };

  return {
    onCreateMeeting: handleCreateMeetingFromContact,
  };
};
