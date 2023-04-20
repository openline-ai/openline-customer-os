import React, { memo } from 'react';
import { Avatar, User } from '../../atoms';
import { useContactNameFromId } from '../../../../hooks/useContact';
import { getContactDisplayName } from '../../../../utils';

interface Props {
  contactId: string;
  size?: number;
  showName?: boolean;
  onlyName?: boolean;
}

export const ContactAvatar: React.FC<Props> = memo(
  function ContactAvatarComponent({
    contactId,
    showName = false,
    onlyName = false,
    size = 30,
  }) {
    const { loading, error, data } = useContactNameFromId({ id: contactId });
    if (loading || error) {
      return <div />;
    }
    const name = getContactDisplayName(data).split(' ');
    return (
      <>
        {!onlyName && (
          <Avatar
            name={name?.[0] || ''}
            surname={name.length === 2 ? name[1] : name[2]}
            size={size}
            image={name.length === 1 && <User />}
          />
        )}

        {(showName || onlyName) && <div>{name}</div>}
      </>
    );
  },
  (prevProps, nextProps) =>
    prevProps.contactId === nextProps.contactId &&
    nextProps.size === prevProps.size,
);
