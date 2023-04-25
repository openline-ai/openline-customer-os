export {
  useCreateMeetingMutation,
  useUpdateMeetingMutation,
  GetOrganizationTimelineDocument,
  GetContactTimelineDocument,
  useMeetingLinkAttachmentMutation,
  useMeetingUnlinkAttachmentMutation,
} from '../../graphQL/__generated__/generated';
export type {
  MeetingInput,
  MeetingParticipant,
  CreateMeetingMutation,
  GetContactTimelineQuery,
  UpdateMeetingMutation,
  MeetingUnlinkAttachmentMutation,
  Meeting,
  GetOrganizationTimelineQuery,
} from '../../graphQL/__generated__/generated';
import type {
  MeetingInput,
  MeetingParticipant,
  CreateMeetingMutation,
} from '../../graphQL/__generated__/generated';

export const NOW_DATE = new Date().toISOString();

export interface Result {
  onCreateMeeting: (
    userId: string,
  ) => Promise<CreateMeetingMutation['meeting_Create'] | null>;
}
