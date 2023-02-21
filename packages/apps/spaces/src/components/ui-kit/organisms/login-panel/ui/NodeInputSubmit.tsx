import { UiNode, UiNodeInputAttributes } from '@ory/client';
import { getNodeLabel } from '@ory/integrations/ui';

import { FormDispatcher, NodeInputProps, ValueSetter } from './helpers';
import { Button } from '../../../atoms';

export function NodeInputSubmit<T>({
  node,
  attributes,
  setValue,
  disabled,
  dispatchSubmit,
}: NodeInputProps) {
  return (
    <Button
      mode='primary'
      name={attributes.name}
      onClick={(e: any) => {
        // On click, we set this value, and once set, dispatch the submission!
        setValue(attributes.value).then(() => dispatchSubmit(e));
      }}
      value={attributes.value || ''}
      disabled={attributes.disabled || disabled}
    >
      {getNodeLabel(node)}
    </Button>
  );
}
