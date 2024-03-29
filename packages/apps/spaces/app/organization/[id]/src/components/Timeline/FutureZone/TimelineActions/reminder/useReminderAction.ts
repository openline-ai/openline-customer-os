import { useParams } from 'next/navigation';

import set from 'date-fns/set';
import { produce } from 'immer';
import addDays from 'date-fns/addDays';
import { useQueryClient } from '@tanstack/react-query';

import { toastError } from '@ui/presentation/Toast';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { useGlobalCacheQuery } from '@shared/graphql/global_Cache.generated';
import { useTimelineMeta } from '@organization/src/components/Timeline/state';
import { useRemindersQuery } from '@organization/src/graphql/reminders.generated';
import { useCreateReminderMutation } from '@organization/src/graphql/createReminder.generated';
import { useTimelineRefContext } from '@organization/src/components/Timeline/context/TimelineRefContext';

export const useReminderAction = () => {
  const organizationId = useParams()?.id as string;
  const client = getGraphQLClient();
  const queryClient = useQueryClient();
  const { virtuosoRef } = useTimelineRefContext();
  const [timelineMeta, setTimelineMeta] = useTimelineMeta();
  const { data: globalCacheData } = useGlobalCacheQuery(client);

  const remindersQueryKey = useRemindersQuery.getKey({ organizationId });
  const createReminder = useCreateReminderMutation(client, {
    onMutate: (values) => {
      queryClient.cancelQueries({ queryKey: remindersQueryKey });

      const previousEntries = useRemindersQuery.mutateCacheEntry(queryClient, {
        organizationId,
      })((cache) =>
        produce(cache, (draft) => {
          if (!draft) return;

          draft.remindersForOrganization.push({
            metadata: {
              id: 'TEMP',
              created: new Date().toISOString(),
              lastUpdated: new Date().toISOString(),
            },
            dueDate: values.input.dueDate,
            content: values.input.content,
            owner: {
              id: globalCacheData?.global_Cache?.user?.id ?? '',
              firstName: globalCacheData?.global_Cache?.user?.firstName ?? '',
              lastName: globalCacheData?.global_Cache?.user?.lastName ?? '',
            },
            dismissed: false,
          });
        }),
      );

      return { previousEntries };
    },
    onError: (_, __, context) => {
      if (context?.previousEntries) {
        queryClient.setQueryData(remindersQueryKey, context.previousEntries);
      }
      toastError(`We couldn't create the reminder`, 'create-reminder-error');
    },
    onSettled: (data) => {
      setTimeout(() => {
        queryClient.invalidateQueries({ queryKey: remindersQueryKey });
      }, 500);

      setTimelineMeta((prev) =>
        produce(prev, (draft) => {
          draft.reminders.recentlyCreatedId =
            data?.reminder_Create?.metadata?.id ?? '';
          draft.reminders.recentlyUpdatedId = '';
        }),
      );
    },
  });

  const handleCreateReminder = (defaultDate?: string) => {
    const targetDate = defaultDate ? new Date(defaultDate) : new Date();
    const dueDate = set(addDays(targetDate, 1), {
      hours: 9,
      minutes: 0,
      seconds: 0,
      milliseconds: 0,
    }).toISOString();

    createReminder.mutate({
      input: {
        content: '',
        dueDate,
        organizationId,
        userId: globalCacheData?.global_Cache?.user?.id ?? '',
      },
    });

    setTimeout(() => {
      virtuosoRef?.current?.scrollToIndex(
        timelineMeta.itemCount + timelineMeta.remindersCount + 1,
      );
    }, 0);
  };

  return {
    handleCreateReminder,
  };
};
