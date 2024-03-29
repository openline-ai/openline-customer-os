import { useForm } from 'react-inverted-form';
import { useRef, useState, FormEvent, useEffect } from 'react';

import { produce } from 'immer';
import { useDidMount, useDebounceFn } from 'rooks';

import { Flex } from '@ui/layout/Flex';
import { Button } from '@ui/form/Button';
import { FormAutoresizeTextarea } from '@ui/form/Textarea';
import { useTimelineMeta } from '@organization/src/components/Timeline/state';

import { ReminderEditForm } from './types';
import { ReminderPostit, ReminderDueDatePicker } from '../../shared';

interface ReminderItem {
  index: number;
  currentOwner: string;
  data: ReminderEditForm;
  onDismiss: (id: string) => void;
  onChange: (value: ReminderEditForm) => void;
}

export const ReminderItem = ({
  data,
  index,
  onChange,
  onDismiss,
  currentOwner,
}: ReminderItem) => {
  const ref = useRef<HTMLTextAreaElement>(null);
  const containerRef = useRef<HTMLDivElement>(null);
  const formId = `reminder-edit-form-${data.id}`;
  const [timelineMeta, setTimelineMeta] = useTimelineMeta();
  const [debouncedOnChange] = useDebounceFn(
    (arg) => onChange(arg as ReminderEditForm),
    1000,
  );
  const { recentlyCreatedId, recentlyUpdatedId } = timelineMeta.reminders;
  const isMutating = data.id === 'TEMP';
  const [isFocused, setIsFocused] = useState(false);

  const { handleSubmit, setDefaultValues } = useForm<ReminderEditForm>({
    formId,
    defaultValues: data,
    onSubmit: async (values) => onChange(values),
    stateReducer: (_, action, next) => {
      if (action.type === 'FIELD_CHANGE') {
        switch (action.payload.name) {
          case 'date': {
            onChange(next.values);
            break;
          }
          default: {
            debouncedOnChange(next.values);
            break;
          }
        }
      }

      return next;
    },
  });

  const updateReminder = () => {
    setIsFocused(false);
    handleSubmit({} as FormEvent<HTMLFormElement>);
  };

  useEffect(() => {
    setDefaultValues(data);
  }, [currentOwner, data.id]);

  useDidMount(() => {
    if (['TEMP', recentlyCreatedId, recentlyUpdatedId].includes(data.id)) {
      ref.current?.focus();

      if (data.id === recentlyCreatedId) {
        setTimelineMeta((prev) =>
          produce(prev, (draft) => {
            draft.reminders.recentlyCreatedId = '';
            draft.reminders.recentlyUpdatedId = '';
          }),
        );
      }
    }
  });

  useEffect(() => {
    if (
      data.id === recentlyUpdatedId ||
      data.id === 'TEMP' ||
      data.id === recentlyCreatedId
    ) {
      containerRef.current && containerRef.current.scrollIntoView();
    }
  }, [recentlyUpdatedId, data.id, index]);

  return (
    <ReminderPostit
      ref={containerRef}
      owner={data?.owner === currentOwner ? undefined : data?.owner}
      isFocused={isFocused}
      isMutating={isMutating}
      boxShadow={data.id === recentlyUpdatedId ? 'ringWarning' : 'unset'}
      onClickOutside={() => {
        setTimelineMeta((prev) =>
          produce(prev, (draft) => {
            draft.reminders.recentlyUpdatedId = '';
          }),
        );
      }}
    >
      <FormAutoresizeTextarea
        px='4'
        pb='0'
        ref={ref}
        isReadOnly={isMutating}
        fontFamily='sticky'
        fontWeight='300'
        fontSize='sm'
        name='content'
        formId={formId}
        lineHeight='inherit'
        onBlur={updateReminder}
        onFocus={() => setIsFocused(true)}
        cacheMeasurements
        maxRows={isFocused ? undefined : 3}
        placeholder='What should we remind you about?'
        borderBottom='unset'
        _hover={{
          borderBottom: 'unset',
        }}
        _focus={{
          borderBottom: 'unset',
        }}
      />
      <Flex align='center' px='4' w='full' justify='space-between' mb='2'>
        <ReminderDueDatePicker name='date' formId={formId} />

        <Button
          size='sm'
          variant='ghost'
          colorScheme='yellow'
          _hover={{
            bg: 'transparent',
            color: 'yellow.900',
          }}
          _focus={{
            boxShadow: 'ringWarning',
          }}
          onClick={() => onDismiss(data.id)}
        >
          Dismiss
        </Button>
      </Flex>
    </ReminderPostit>
  );
};
