import React from 'react';
import { Box } from '@ui/layout/Box';
import { Flex } from '@ui/layout/Flex';
import { AlertSquare } from '@ui/media/icons/AlertSquare';
import { Tooltip } from '@ui/overlay/Tooltip';

type Priority = 'low' | 'medium' | 'high';

interface PriorityBadgeProps {
  priority: Priority | 'urgent';
}

const boxWidth = '3px';
const boxMarginRight = '1px';
const colorMap: Record<Priority, string[]> = {
  low: ['gray.500', 'gray.300', 'gray.300'],
  medium: ['gray.500', 'gray.500', 'gray.300'],
  high: ['gray.500', 'gray.500', 'gray.500'],
};

export const PriorityBadge: React.FC<PriorityBadgeProps> = ({ priority }) => {
  if (priority === 'urgent') {
    return (
      <Tooltip label={priority}>
        <AlertSquare
          color='red.600'
          aria-label={priority}
          role='presentation'
        />
      </Tooltip>
    );
  }

  return (
    <Tooltip label={priority}>
      <Flex alignItems='flex-end' role='presentation' aria-label={priority}>
        {colorMap[`${priority}`]?.map((color, i) => (
          <Box
            key={i}
            width={boxWidth}
            borderRadius='sm'
            height={`${3 * (i + 1)}px`}
            mr={boxMarginRight}
            bgColor={color}
          />
        ))}
      </Flex>
    </Tooltip>
  );
};
