import React, { useContext } from 'react';
import Image from 'next/image';
import { IconButton } from '../../ui-kit/atoms';
import styles from './contact-details.module.scss';
import { ContactPersonalDetails } from './ContactPersonalDetails';
import { WebRTCContext } from '../../../context/web-rtc';
import { useContactCommunicationChannelsDetails } from '../../../hooks/useContact';
import { toast } from 'react-toastify';
import { useRecoilState, useSetRecoilState } from 'recoil';
import { callParticipant, contactDetailsEdit } from '../../../state';
import { getContactDisplayName } from '../../../utils';
import { Contact } from '../../../graphQL/__generated__/generated';
import classNames from 'classnames';

export const ContactDetails = ({ id }: { id: string }) => {
  const webRtc = useContext(WebRTCContext) as any;
  const { data, loading, error } = useContactCommunicationChannelsDetails({
    id,
  });
  const [{ isEditMode }, setContactDetailsEdit] =
    useRecoilState(contactDetailsEdit);
  const setCallParticipant = useSetRecoilState(callParticipant);

  const handleStartPhoneCall = () => {
    const number =
      data?.phoneNumbers.find((pn) => pn.primary)?.e164 ||
      data?.phoneNumbers[0].e164;

    if (!number) {
      toast.error('Error! Number missing!', {
        toastId: `${id}-missing-phone-number`,
      });
      return;
    }
    setCallParticipant({ identity: getContactDisplayName(data as Contact) });
    webRtc?.makeCall(number);
  };

  return (
    <div className={styles.contactDetails}>
      <ContactPersonalDetails id={id} />

      <div className={styles.details}>
        <div className={styles.section}>
          <IconButton
            disabled={loading || error !== null || !data?.phoneNumbers.length}
            aria-describedby='phone-icon-label'
            mode='secondary'
            className={styles.icon}
            onClick={handleStartPhoneCall}
            icon={
              <Image alt={''} src='/icons/phone.svg' width={20} height={20} />
            }
          />

          <div className={styles.label} id='phone-icon-label'>
            Phone
          </div>
        </div>
        <div className={classNames(styles.section, styles.disabled)}>
          <IconButton
            disabled={true}
            aria-describedby='email-icon-label'
            mode='secondary'
            className={styles.icon}
            onClick={() => null}
            icon={
              <Image
                alt={''}
                src='/icons/envelope.svg'
                width={20}
                height={20}
              />
            }
          />

          <div className={styles.label} id={'email-icon-label'}>
            Email
          </div>
        </div>
        <div className={classNames(styles.section, styles.disabled)}>
          <IconButton
            disabled={true}
            aria-describedby='message-icon-label'
            mode='secondary'
            className={styles.icon}
            onClick={() => null}
            icon={
              <Image
                alt={''}
                src='/icons/whatsapp.svg'
                width={20}
                height={20}
              />
            }
          />
          <div className={styles.label} id='message-icon-label'>
            Message
          </div>
        </div>
        <div className={classNames(styles.section, styles.disabled)}>
          <IconButton
            disabled={true}
            aria-describedby='message-icon-label'
            mode='secondary'
            className={styles.icon}
            onClick={() => null}
            icon={
              <Image
                alt={''}
                src='/icons/share-alt.svg'
                width={20}
                height={20}
              />
            }
          />
          <div className={styles.label} id='message-icon-label'>
            Share
          </div>
        </div>
        {isEditMode ? (
          <div className={classNames(styles.section)}>
            <IconButton
              aria-describedby='message-icon-label'
              mode='secondary'
              className={styles.icon}
              onClick={() => setContactDetailsEdit({ isEditMode: !isEditMode })}
              icon={
                <Image alt={''} src='/icons/check.svg' width={20} height={20} />
              }
            />
            <div className={styles.label} id='message-icon-label'>
              Done
            </div>
          </div>
        ) : (
          <div className={styles.section}>
            <IconButton
              aria-describedby='message-icon-label'
              mode='secondary'
              className={styles.icon}
              onClick={() => setContactDetailsEdit({ isEditMode: !isEditMode })}
              icon={
                <Image
                  alt={''}
                  src='/icons/pencil.svg'
                  width={20}
                  height={20}
                />
              }
            />
            <div className={styles.label} id='message-icon-label'>
              Edit
            </div>
          </div>
        )}
      </div>
    </div>
  );
};
