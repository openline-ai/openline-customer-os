import React, { ChangeEventHandler, ReactNode, useRef } from 'react';
import { DebounceInput, DebounceInputProps } from 'react-debounce-input';
import styles from './input.module.scss';
import classNames from 'classnames';
interface DebouncedInputProps
  extends Partial<
    Omit<DebounceInputProps<HTMLInputElement, HTMLInputElement>, 'children'>
  > {
  onChange: ChangeEventHandler<HTMLInputElement>;
  placeholder?: string;
  minLength?: number;
  debounceTimeout?: number;
  inputSize?: 'xxxs' | 'xxs' | 'xs' | 'sm' | 'md' | 'lg';
  children?: ReactNode;
  inlineMode?: boolean;
}

export const DebouncedInput = ({
  onChange,
  placeholder = '',
  minLength = 3,
  children,
  inputSize = 'md',
  debounceTimeout = 300,
  inlineMode,
  className,
  ...rest
}: DebouncedInputProps) => {
  const inputRef = useRef(null);

  return (
    <div
      className={classNames(styles.wrapper, {
        [styles.inlineMode]: inlineMode,
      })}
    >
      <DebounceInput
        {...rest}
        size={rest?.value?.length || placeholder?.length}
        inputRef={inputRef}
        className={classNames(styles.input, {
          [styles?.[inputSize]]: inputSize,
          [styles.xxxs]: inlineMode,
          [`${className}`]: className,
        })}
        minLength={minLength}
        debounceTimeout={debounceTimeout}
        onChange={onChange}
        placeholder={placeholder}
      />

      {children && <span className={styles.icon}>{children}</span>}
    </div>
  );
};
