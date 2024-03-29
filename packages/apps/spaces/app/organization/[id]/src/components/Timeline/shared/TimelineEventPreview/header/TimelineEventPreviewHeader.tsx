import React from 'react';

import { escapeForSlackWithMarkdown } from 'slack-to-html';

import { Flex } from '@ui/layout/Flex';
import { Text } from '@ui/typography/Text';
import { Link03 } from '@ui/media/icons/Link03';
import { XClose } from '@ui/media/icons/XClose';
import { IconButton } from '@ui/form/IconButton';
import { CardHeader } from '@ui/presentation/Card';
import { DateTimeUtils } from '@spaces/utils/date';
import { Tooltip } from '@ui/presentation/Tooltip';
import { useCopyToClipboard } from '@shared/hooks/useCopyToClipboard';

interface TimelineEventPreviewHeaderProps {
  name: string;
  date?: string;
  parse?: 'slack';
  copyLabel: string;
  onClose: () => void;
  children?: React.ReactNode;
}

export const TimelineEventPreviewHeader: React.FC<
  TimelineEventPreviewHeaderProps
> = ({ date, name, onClose, copyLabel, children, parse }) => {
  const [_, copy] = useCopyToClipboard();

  const parsedName =
    parse === 'slack' ? escapeForSlackWithMarkdown(name) : name;

  return (
    <CardHeader
      py='4'
      px='6'
      pb='1'
      position='sticky'
      background='white'
      top={0}
      borderRadius='xl'
      onClick={(e) => e.stopPropagation()}
    >
      <Flex
        direction='row'
        justifyContent='space-between'
        alignItems='flex-start'
      >
        <div>
          <Text
            fontSize='lg'
            fontWeight='semibold'
            dangerouslySetInnerHTML={
              parse === 'slack' ? { __html: parsedName } : undefined
            }
          >
            {parse !== 'slack' ? name : null}
          </Text>
          {date && (
            <Text size='2xs' color='gray.500' fontSize='12px'>
              {DateTimeUtils.format(date, DateTimeUtils.dateWithHour)}
            </Text>
          )}
        </div>
        <Flex direction='row' justifyContent='flex-end' alignItems='center'>
          {children}
          <Tooltip label={copyLabel} placement='bottom'>
            <IconButton
              variant='ghost'
              aria-label={copyLabel}
              color='gray.500'
              size='sm'
              mr={1}
              icon={<Link03 color='gray.500' height='18px' />}
              onClick={() => copy(window.location.href)}
            />
          </Tooltip>
          <Tooltip label='Close' aria-label='close' placement='bottom'>
            <IconButton
              variant='ghost'
              aria-label='Close preview'
              color='gray.500'
              size='sm'
              icon={<XClose color='gray.500' height='24px' />}
              onClick={onClose}
            />
          </Tooltip>
        </Flex>
      </Flex>
    </CardHeader>
  );
};
