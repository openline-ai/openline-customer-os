import React, { forwardRef, cloneElement } from 'react';

import { twMerge } from 'tailwind-merge';
import { cva, type VariantProps } from 'class-variance-authority';

import {
  linkButton,
  ghostButton,
  solidButton,
  iconVariant,
  outlineButton,
} from './Button.variants';

export const buttonSize = cva([], {
  variants: {
    size: {
      sm: ['px-3', 'py-1', 'rounded-lg'],
      md: ['px-4', 'py-2.5', 'rounded-lg'],
      lg: ['px-[1.125rem]', 'py-2.5', 'rounded-lg', 'text-base'],
      xl: ['px-5', 'py-3', 'rounded-lg', 'text-base'],
      '2xl': ['px-7', 'py-4', 'gap-3', 'rounded-lg', 'text-lg'],
    },
  },
  defaultVariants: {
    size: 'md',
  },
});

export interface ButtonProps
  extends React.HTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof solidButton>,
    VariantProps<typeof buttonSize> {
  asChild?: boolean;
  isLoading?: boolean;
  isDisabled?: boolean;
  spinner?: React.ReactElement;
  leftIcon?: React.ReactElement;
  rightIcon?: React.ReactElement;
  variant?: 'link' | 'ghost' | 'solid' | 'outline';
}

export const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  (
    {
      leftIcon,
      children,
      className,
      rightIcon,
      colorScheme,
      spinner,
      variant,
      isLoading = false,
      isDisabled = false,
      size,
      ...props
    },
    ref,
  ) => {
    const buttonVariant = (() => {
      switch (variant) {
        case 'link':
          return linkButton;
        case 'ghost':
          return ghostButton;
        case 'solid':
          return solidButton;
        case 'outline':
          return outlineButton;
        default:
          return solidButton;
      }
    })();

    return (
      <button
        ref={ref}
        {...props}
        className={twMerge(
          buttonVariant({ colorScheme, className }),
          buttonSize({ className, size }),
          isLoading ? 'opacity-50 cursor-not-allowed' : '',
        )}
        disabled={isLoading || isDisabled}
      >
        {isLoading && spinner && (
          <span className='relative inline-flex'>{spinner}</span>
        )}

        {!isLoading && leftIcon && (
          <>
            {cloneElement(leftIcon, {
              className: twMerge(
                iconVariant({
                  size,
                  variant,
                  colorScheme,
                  className: leftIcon.props.className,
                }),
              ),
            })}
          </>
        )}

        {!isLoading && children}
        {!isLoading && rightIcon && (
          <>
            {cloneElement(rightIcon, {
              className: twMerge(
                iconVariant({
                  size,
                  variant,
                  colorScheme,
                  className: rightIcon.props.className,
                }),
              ),
            })}
          </>
        )}
      </button>
    );
  },
);
