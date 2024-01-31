import { Checkbox, CheckboxProps } from '@ui/form/Checkbox';
import { ArrowNarrowDownRight } from '@ui/media/icons/ArrowNarrowDownRight';

interface StatusCheckboxProps extends CheckboxProps {
  readOnly?: boolean;
  showCustomIcon?: boolean;
}

export const StatusCheckbox = (props: StatusCheckboxProps) => {
  return (
    <Checkbox
      mr='2'
      size='md'
      icon={
        props?.showCustomIcon ? (
          <ArrowNarrowDownRight color={`${props.colorScheme}.400`} />
        ) : undefined
      }
      {...props}
      sx={{
        '& > span': {
          bg: `${props.colorScheme}.100`,
          borderColor: `${props.colorScheme}.300`,
          _focus: {
            bg: `${props.colorScheme}.100`,
          },
        },
      }}
    />
  );
};