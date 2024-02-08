import { useField } from 'react-inverted-form';
import { memo, useRef, useState, useCallback, ChangeEvent } from 'react';

import { useKey } from 'rooks';
import { produce } from 'immer';

import { Flex } from '@ui/layout/Flex';
import { Input } from '@ui/form/Input';
import { Text } from '@ui/typography/Text';
import { Tooltip } from '@ui/overlay/Tooltip';
import { IconButton } from '@ui/form/IconButton';
import { DateTimeUtils } from '@spaces/utils/date';
import { SkipForward } from '@ui/media/icons/SkipForward';
import { OnboardingPlanMilestoneItemStatus } from '@graphql/types';

import { TaskDatum } from '../../../../types';
import { StatusCheckbox } from '../StatusCheckbox';

interface TaskProps {
  index: number;
  formId: string;
  isLast?: boolean;
  defaultValue?: string;
  shouldFocusRef?: React.MutableRefObject<number | null>;
}

export const Task = memo(
  ({ index, formId, shouldFocusRef, isLast, defaultValue }: TaskProps) => {
    const ref = useRef<HTMLInputElement>(null);

    const [showSkip, setShowSkip] = useState(false);
    const [isFocused, setIsFocused] = useState(false);

    const { getInputProps } = useField('items', formId);
    const { value, onChange, onBlur } = getInputProps();
    const itemValue = value?.[index] as TaskDatum;

    const milestoneDueDate = useField('dueDate', formId).getInputProps()
      ?.value as string;

    const taskStatus =
      itemValue?.status ?? OnboardingPlanMilestoneItemStatus.NotDone;
    const taskUpdatedAt = new Date(itemValue?.updatedAt).valueOf();
    const taskUpdatedAtDate = DateTimeUtils.format(
      itemValue?.updatedAt,
      DateTimeUtils.dateWithShortYear,
    );
    const milestoneDueAt = new Date(milestoneDueDate).valueOf();

    const isDone = [
      OnboardingPlanMilestoneItemStatus.Done,
      OnboardingPlanMilestoneItemStatus.DoneLate,
    ].includes(taskStatus);
    const isSkipped = [
      OnboardingPlanMilestoneItemStatus.Skipped,
      OnboardingPlanMilestoneItemStatus.SkippedLate,
    ].includes(taskStatus);

    const colorScheme = (() => {
      switch (taskStatus) {
        case OnboardingPlanMilestoneItemStatus.NotDone:
        case OnboardingPlanMilestoneItemStatus.Skipped:
        case OnboardingPlanMilestoneItemStatus.SkippedLate:
          return 'gray';
        case OnboardingPlanMilestoneItemStatus.Done:
          return 'success';
        case OnboardingPlanMilestoneItemStatus.DoneLate:
        case OnboardingPlanMilestoneItemStatus.NotDoneLate:
          return 'warning';
        default:
          return 'gray';
      }
    })();

    const handleChange = useCallback(
      (e: ChangeEvent<HTMLInputElement>) => {
        if (!itemValue) return;
        const isChecked = e.target.checked;

        const updatedItem = produce<TaskDatum>(itemValue, (draft) => {
          const isLate = new Date().valueOf() > milestoneDueAt;

          if (isChecked) {
            draft.status = isLate
              ? OnboardingPlanMilestoneItemStatus.DoneLate
              : OnboardingPlanMilestoneItemStatus.Done;
          } else {
            draft.status = isLate
              ? OnboardingPlanMilestoneItemStatus.NotDoneLate
              : OnboardingPlanMilestoneItemStatus.NotDone;
          }

          draft.updatedAt = new Date().toISOString();
        });

        const next = (value as TaskDatum[]).map((v, i) =>
          i === index ? updatedItem : v,
        );

        onChange?.(next);
      },
      [onChange, index, value, taskUpdatedAt, milestoneDueAt],
    );

    const handleInputChange = useCallback(
      (e: ChangeEvent<HTMLInputElement>) => {
        const nextItems = produce<TaskDatum[]>(value, (draft) => {
          const item = draft?.[index];
          if (!item) return;

          item.text = e.target.value;
        });

        onChange?.(nextItems);

        if (shouldFocusRef) {
          shouldFocusRef.current = index;
        }
      },
      [onChange, index, value],
    );

    const handleInputBlur = useCallback(
      (e: ChangeEvent<HTMLInputElement>) => {
        setIsFocused(false);

        const nextItems = produce<TaskDatum[]>(value, (draft) => {
          const item = draft?.[index];
          if (!item) return;

          if (!e.target.value) {
            draft.splice(index, 1);

            return;
          }

          item.text = e.target.value;
        });

        onBlur?.(nextItems);
      },
      [onBlur, onChange, index, value],
    );

    const handleRemove = () => {
      const nextItems = produce<TaskDatum[]>(value, (draft) => {
        draft.splice(index, 1);
      });

      onChange?.(nextItems);
    };

    const handleSkip = useCallback(() => {
      const nextItems = produce<TaskDatum[]>(value, (draft) => {
        const item = draft?.[index];
        if (!item) return;

        const isPastDueDate = taskUpdatedAt > milestoneDueAt;

        item.status = (() => {
          switch (item.status) {
            case OnboardingPlanMilestoneItemStatus.NotDone:
              return isPastDueDate
                ? OnboardingPlanMilestoneItemStatus.SkippedLate
                : OnboardingPlanMilestoneItemStatus.Skipped;
            case OnboardingPlanMilestoneItemStatus.Skipped:
            case OnboardingPlanMilestoneItemStatus.SkippedLate:
              return isPastDueDate
                ? OnboardingPlanMilestoneItemStatus.NotDoneLate
                : OnboardingPlanMilestoneItemStatus.NotDone;
            default:
              return item.status;
          }
        })();

        item.updatedAt = new Date().toISOString();
      });

      onChange?.(nextItems);
    }, [onChange, index, value, taskUpdatedAt, milestoneDueAt]);

    const handleAdd = () => {
      setIsFocused(false);

      const nextItems = produce<TaskDatum[]>(value, (draft) => {
        const isPastDueDate = taskUpdatedAt > milestoneDueAt;

        draft.push({
          text: '',
          updatedAt: new Date().toISOString(),
          status: isPastDueDate
            ? OnboardingPlanMilestoneItemStatus.NotDoneLate
            : OnboardingPlanMilestoneItemStatus.NotDone,
        });
      });

      if (shouldFocusRef) {
        shouldFocusRef.current = value.length;
      }

      onChange?.(nextItems);
    };

    useKey('Enter', () => {
      if (isFocused) {
        handleAdd();
      }
    });
    useKey('Backspace', () => {
      if (isFocused && !ref.current?.value) {
        handleRemove();
      }
    });

    return (
      <Flex
        w='full'
        onMouseEnter={() => (!isSkipped ? setShowSkip(true) : undefined)}
        onMouseLeave={() => (!isSkipped ? setShowSkip(false) : undefined)}
      >
        <StatusCheckbox
          mr='2'
          size='md'
          isChecked={isDone}
          onChange={handleChange}
          colorScheme={colorScheme}
        />
        <MemoizedInput
          w='full'
          ref={ref}
          fontSize='sm'
          variant='unstyled'
          borderRadius='unset'
          placeholder='Task name'
          onBlur={handleInputBlur}
          onChange={handleInputChange}
          value={value?.[index]?.text ?? defaultValue}
          autoFocus={shouldFocusRef?.current === index}
          onFocus={() => {
            setIsFocused(true);
          }}
          fontStyle={isSkipped ? 'italic' : 'normal'}
        />
        {isDone && (
          <Text fontSize='sm' color='gray.500' whiteSpace='nowrap'>
            {taskUpdatedAtDate}
          </Text>
        )}
        {!isDone && (
          <Tooltip label={isSkipped ? 'Skipped' : 'Skip this'}>
            <IconButton
              size='xs'
              variant='ghost'
              onClick={handleSkip}
              opacity={showSkip || isSkipped ? 1 : 0}
              aria-label='Skip Milestone Task'
              icon={<SkipForward color='gray.400' />}
            />
          </Tooltip>
        )}
      </Flex>
    );
  },
  (prev, next) => prev.defaultValue === next.defaultValue,
);

const MemoizedInput = memo(Input);
