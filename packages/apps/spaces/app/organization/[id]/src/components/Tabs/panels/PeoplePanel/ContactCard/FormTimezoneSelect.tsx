import { chakraComponents, SingleValueProps } from 'chakra-react-select';

import { Icons } from '@ui/media/Icon';
import { Text } from '@ui/typography/Text';
import { DateTimeUtils } from '@spaces/utils/date';
import { FormSelect, SelectProps } from '@ui/form/SyncSelect';

const SingleValue = (props: SingleValueProps) => {
  const rawTimezone = props.children as string;
  const timezone = rawTimezone.includes('UTC')
    ? rawTimezone.split(' ')[0]
    : rawTimezone;

  const time = DateTimeUtils.convertToTimeZone(
    new Date(),
    DateTimeUtils.defaultTimeFormatString,
    timezone,
  );
  const value = `${time} local time`;

  return (
    <chakraComponents.SingleValue {...props}>
      <Text color='gray.700' isTruncated>
        {value}
        {` `}
        <Text as='span' color='gray.500'>
          • {timezone}
        </Text>
      </Text>
    </chakraComponents.SingleValue>
  );
};

const components = {
  SingleValue,
};

interface FormTimezoneSelectProps extends SelectProps {
  name: string;
  formId: string;
}

export const FormTimezoneSelect = ({ ...props }: FormTimezoneSelectProps) => {
  return (
    <FormSelect
      components={components}
      leftElement={<Icons.Clock color='gray.500' mr='3' />}
      {...props}
    />
  );
};
