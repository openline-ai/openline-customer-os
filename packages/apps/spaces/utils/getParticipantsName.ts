import {
  ContactParticipant,
  EmailParticipant,
  PhoneNumberParticipant,
  UserParticipant,
} from '@spaces/graphql';
import { InteractionEventParticipant } from '@graphql/types';

type Participant =
  | EmailParticipant
  | PhoneNumberParticipant
  | ContactParticipant
  | UserParticipant;

export const getParticipantNames = (participants: Participant[]): string[] => {
  return participants.map((participant) => {
    if (participant.__typename === 'EmailParticipant') {
      const { emailParticipant } = participant;
      const { contacts, users } = emailParticipant;
      if (contacts.length) {
        return contacts
          .map((c) => (c?.name ? c.name : `${c.firstName} ${c.lastName}`))
          .join(' ');
      }
      if (users.length) {
        return users.map((c) => `${c.firstName} ${c.lastName}`).join(' ');
      }

      const participantName =
        contacts?.[0]?.name ||
        users?.[0]?.firstName + ' ' + users?.[0]?.lastName;
      return participantName || 'Unnamed';
    } else if (participant.__typename === 'PhoneNumberParticipant') {
      const { phoneNumberParticipant } = participant;
      const { contacts, users } = phoneNumberParticipant;
      const participantName =
        contacts?.[0]?.name ||
        users?.[0]?.firstName + ' ' + users?.[0]?.lastName;
      return participantName || 'name';
    } else if (participant.__typename === 'ContactParticipant') {
      const { contactParticipant } = participant;
      const { name, firstName, lastName } = contactParticipant;
      return firstName + ' ' + lastName || name || 'Unnamed';
    } else if (participant.__typename === 'UserParticipant') {
      const { userParticipant } = participant;
      const { firstName, lastName } = userParticipant;
      return firstName + ' ' + lastName || 'Unnamed';
    }
    return 'Unnamed';
  });
};

export const getParticipant = (
  participant: InteractionEventParticipant,
): string => {
  if (participant?.__typename === 'EmailParticipant') {
    const { emailParticipant } = participant;
    const { contacts, users } = emailParticipant;
    if (contacts.length) {
      return contacts
        .map((c) => (c?.name ? c.name : `${c.firstName} ${c.lastName}`))
        .join(' ');
    }
    if (users.length) {
      return users.map((c) => `${c.firstName} ${c.lastName}`).join(' ');
    }

    const participantName =
      contacts?.[0]?.name || users?.[0]?.firstName + ' ' + users?.[0]?.lastName;
    return participantName || emailParticipant?.email || '';
  } else if (participant?.__typename === 'ContactParticipant') {
    const { contactParticipant } = participant;
    const { name, firstName, lastName } = contactParticipant;
    return firstName + ' ' + lastName || name || 'Unnamed';
  } else if (participant?.__typename === 'UserParticipant') {
    const { userParticipant } = participant;
    const { firstName, lastName } = userParticipant;
    return firstName + ' ' + lastName || 'Unnamed';
  }
  return '';
};
export const getEmailParticipantsName = (
  participants: InteractionEventParticipant[],
): string[] => {
  return participants?.map((participant) => getParticipant(participant));
};

export const getParticipantNameAndEmail = (
  participant: EmailParticipant,
): { email: string | null; label: string } => {
  const { emailParticipant } = participant;
  const { contacts, users, email, rawEmail } = emailParticipant;
  console.log('🏷️ ----- contacts: '
      , contacts);
  console.log('🏷️ ----- users: '
      , users);
  if (contacts.length) {
    const label = contacts
      .map((c) => (c?.name ? c.name : `${c.firstName} ${c.lastName}`))
      .join(' ');
    return {
      label,
      email: email || rawEmail || '',
    };
  }
  if (users.length) {
    const label = users.map((c) => `${c.firstName} ${c.lastName}`).join(' ');
    return {
      label,
      email: email || rawEmail || '',
    };
  }

  return {
    label: '',
    email: email || rawEmail || '',
  };
};

export const getEmailParticipantsNameAndEmail = (
  participants: InteractionEventParticipant[],
): Array<{ email: string | null; label: string }> => {

  return participants?.map((participant) =>
    getParticipantNameAndEmail(participant as EmailParticipant),
  );
};
