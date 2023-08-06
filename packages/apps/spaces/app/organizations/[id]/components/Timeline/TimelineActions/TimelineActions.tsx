import React, { useEffect, useRef, useState } from 'react';
import { SlideFade } from '@ui/transitions/SlideFade';
import { Box } from '@ui/layout/Box';
import { Button } from '@ui/form/Button';
import { ButtonGroup } from '@ui/form/ButtonGroup';
import { ComposeEmail } from '@organization/components/Timeline/events/email/compose-email/ComposeEmail';
import Envelope from '@spaces/atoms/icons/Envelope';
import { useForm } from 'react-inverted-form';
import {
  ComposeEmailDto,
  ComposeEmailDtoI,
} from '@organization/components/Timeline/events/email/compose-email/ComposeEmail.dto';
import { handleSendEmail } from '@organization/components/Timeline/events/email/compose-email/utils';
import { useSearchParams } from 'next/navigation';
import { useSession } from 'next-auth/react';

interface TimelineActionsProps {
  onScrollBottom: () => void;
}

export const TimelineActions: React.FC<TimelineActionsProps> = ({
  onScrollBottom,
}) => {
  const [show, setShow] = React.useState(false);
  const [isSending, setIsSending] = useState(false);
  const virtuoso = useRef(null);
  const searchParams = useSearchParams();
  const { data: session } = useSession();
  const formId = 'compose-email-timeline-footer';
  const defaultValues: ComposeEmailDtoI = new ComposeEmailDto({
    to: [],
    cc: [],
    bcc: [],
    subject: '',
    content: '',
  });
  useEffect(() => {
    if (show) {
      onScrollBottom();
    }
  }, [show]);
  const { state, reset } = useForm<ComposeEmailDtoI>({
    formId,
    defaultValues,

    stateReducer: (state, action, next) => {
      return next;
    },
  });

  const handleEmailSendSuccess = () => {
    setIsSending(false);
    reset();
    setShow(false)
  };
  const handleEmailSendError = () => {
    setIsSending(false);
  };

  const handleSubmit = () => {
    const destination = [
      ...state.values.to,
      ...state.values.cc,
      ...state.values.bcc,
    ].map(({ value }) => value);
    const params = new URLSearchParams(searchParams ?? '');

    setIsSending(true);
    const id = params.get('events');
    return handleSendEmail(
      state.values.content,
      destination,
      id,
      state.values.subject,
      handleEmailSendSuccess,
      handleEmailSendError,
      session?.user?.email,
    );
  };

  const handleToggle = () => setShow(!show);
  return (
    <Box bg='gray.25'>
      <ButtonGroup
        // mt={6}
        position='sticky'
        py={2}
        border='1px dashed var(--gray-200, #EAECF0)'
        p={2}
        borderRadius={30}
        bg='white'
        top='0'
        left={6}
        zIndex={1}
        translateY='6px'
      >
        <Button
          variant='outline'
          onClick={() => handleToggle()}
          borderRadius='3xl'
          size='xs'
          leftIcon={<Envelope color='inherit' height={16} width={16} />}
        >
          Email
        </Button>
      </ButtonGroup>
      <Box
        bg={'#F9F9FB'}
        borderTop='1px dashed var(--gray-200, #EAECF0)'
        pt={show ? 6 : 0}
        pb={show ? 2 : 8}
        mt={-4}
      >
        {show && (
          <SlideFade in={true}>
            <Box
              ref={virtuoso}
              borderRadius={'md'}
              boxShadow={'lg'}
              m={6}
              mt={0}
              bg={'white'}
              border='1px solid var(--gray-100, #F2F4F7)'
            >
              <ComposeEmail
                formId={formId}
                modal={false}
                to={state.values.to}
                cc={state.values.cc}
                bcc={state.values.bcc}
                onSubmit={handleSubmit}
                isSending={isSending}
              />
            </Box>
          </SlideFade>
        )}
      </Box>
    </Box>
  );
};
