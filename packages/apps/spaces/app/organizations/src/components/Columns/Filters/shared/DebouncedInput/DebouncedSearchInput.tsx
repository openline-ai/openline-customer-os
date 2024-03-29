'use client';
import {
  memo,
  useRef,
  useEffect,
  forwardRef,
  ChangeEvent,
  useImperativeHandle,
} from 'react';

import { Input } from '@ui/form/Input';
import { Delete } from '@ui/media/icons/Delete';
import { IconButton } from '@ui/form/IconButton';
import { SearchSm } from '@ui/media/icons/SearchSm';
import {
  InputGroup,
  InputLeftElement,
  InputRightElement,
} from '@ui/form/InputGroup';

interface DebouncedInputProps {
  value: string;
  onChange: (value: string) => void;
  onDisplayChange?: (value: string) => void;
}

export const DebouncedSearchInput = memo(
  forwardRef<HTMLInputElement, DebouncedInputProps>(
    ({ value, onChange, onDisplayChange }, ref) => {
      const timeout = useRef<NodeJS.Timeout>();
      const innerRef = useRef<HTMLInputElement>(null);

      const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        onDisplayChange?.(e.target.value);

        if (timeout.current) {
          clearTimeout(timeout.current);
        }

        timeout.current = setTimeout(() => {
          onChange(e.target.value);
        }, 250);
      };

      const handleClear = () => {
        onChange('');
        onDisplayChange?.('');
        innerRef?.current?.focus();
      };

      useImperativeHandle(ref, () => innerRef.current!, []);

      useEffect(() => {
        return () => {
          if (timeout.current) {
            clearTimeout(timeout.current);
          }
        };
      }, []);

      return (
        <InputGroup>
          <InputLeftElement w='fit-content'>
            <SearchSm color='gray.500' />
          </InputLeftElement>
          <Input
            pl='6'
            value={value}
            ref={innerRef}
            variant='flushed'
            autoComplete='off'
            placeholder='Search'
            onChange={handleChange}
          />
          {value.length && (
            <InputRightElement w='fit-content'>
              <IconButton
                size='xs'
                variant='ghost'
                onClick={handleClear}
                aria-label='search organization'
                icon={<Delete color='gray.500' />}
              />
            </InputRightElement>
          )}
        </InputGroup>
      );
    },
  ),
);
