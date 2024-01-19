import { useField } from 'react-inverted-form';

import { produce } from 'immer';

import { Button } from '@ui/form/Button';
import { VStack } from '@ui/layout/Stack';
import { Plus } from '@ui/media/icons/Plus';

import { Task } from './Task';

interface TasksProps {
  formId: string;
}

export const Tasks = ({ formId }: TasksProps) => {
  const { getInputProps } = useField('items', formId);
  const { value, onChange, onBlur } = getInputProps();

  const handleAddTask = () => {
    const nextItems = produce<string[]>(value, (draft) => {
      draft.push('');
    });

    onChange?.(nextItems);
    onBlur?.(nextItems);
  };

  return (
    <VStack align='flex-start' spacing='1'>
      {(value as string[])?.map((item, idx, arr) => (
        <Task
          key={idx}
          index={idx}
          value={item}
          formId={formId}
          isLast={idx === arr.length - 1}
        />
      ))}
      <Button
        ml='-12px'
        size='sm'
        variant='ghost'
        color='gray.500'
        fontWeight='normal'
        onClick={handleAddTask}
        leftIcon={<Plus color='gray.400' />}
      >
        Add task
      </Button>
    </VStack>
  );
};
