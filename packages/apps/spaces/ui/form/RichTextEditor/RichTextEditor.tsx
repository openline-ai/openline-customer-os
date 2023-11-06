import { useField } from 'react-inverted-form';
import React, {
  FC,
  forwardRef,
  PropsWithChildren,
  useImperativeHandle,
} from 'react';

import { prosemirrorNodeToHtml } from 'remirror';
import { Toolbar, Remirror, ThemeProvider } from '@remirror/react';

import { RemirrorProps } from '@ui/form/RichTextEditor/types';
import { FloatingLinkToolbar } from '@ui/form/RichTextEditor/floatingMenu/FloatingLinkMenu';

export const RichTextEditor: FC<
  {
    name: string;
    formId: string;
    placeholder?: string;
    showToolbar: boolean;
    // exhaustively typing this is not really necessary for us at the moment
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
  } & RemirrorProps<any> &
    PropsWithChildren
> = forwardRef(
  (
    {
      children,
      name,
      showToolbar,
      formId,
      manager,
      getContext,
      state,
      setState,
      placeholder = '',
    },
    ref,
  ) => {
    const { getInputProps } = useField(name, formId);
    const { onChange, value } = getInputProps();
    useImperativeHandle(ref, () => getContext(), [getContext]);

    return (
      <ThemeProvider>
        <Remirror
          manager={manager}
          placeholder={placeholder}
          onChange={(parameter) => {
            const nextState = parameter.state;
            const htmlValue = prosemirrorNodeToHtml(nextState?.doc);
            // first update is happening before form store is initialized this change prevents error
            if (value !== undefined) {
              onChange(htmlValue);
            }
            setState(nextState);
          }}
          initialContent={state}
          autoRender='end'
        >
          <FloatingLinkToolbar />
          {showToolbar ? (
            <Toolbar
              height={'var(--chakra-sizes-8)'}
              style={{ overflowX: 'visible' }}
            >
              {children}
            </Toolbar>
          ) : (
            children
          )}
        </Remirror>
      </ThemeProvider>
    );
  },
);
