import set from 'date-fns/set';
import differenceInDays from 'date-fns/differenceInDays';
import differenceInWeeks from 'date-fns/differenceInWeeks';
import differenceInHours from 'date-fns/differenceInHours';
import differenceInMonths from 'date-fns/differenceInMonths';
import differenceInMinutes from 'date-fns/differenceInMinutes';

export function getDifferenceFromNow(targetDate: string) {
  const now = set(new Date(), { hours: 0, minutes: 0, seconds: 0 });
  const next = set(new Date(targetDate), { hours: 0, minutes: 0, seconds: 1 });

  const months = differenceInMonths(next, now);
  const weeks = differenceInWeeks(next, now);
  const days = differenceInDays(next, now);

  if (days === 0) return ['0', 'days'];

  if (days === 1) return [days, 'day'];
  if (days < 7 && days !== 1) return [days, 'days'];

  if (weeks === 1) return [weeks, 'week'];
  if (weeks <= 4 && weeks !== 1 && months === 0) return [weeks, 'weeks'];
  if (weeks % 4 === 0 && weeks / 4 !== 1) return [weeks / 4, 'months'];

  if (months === 1 && weeks % 4 === 0) return [months, 'month'];

  const roundedMonths = weeks % 4 > 2 ? months + 1 : months;

  return [roundedMonths, 'months'];
}

export function getDifferenceInMinutesOrHours(targetDate: string) {
  const now = new Date();
  const next = new Date(targetDate);

  const minutes = Math.abs(differenceInMinutes(next, now));
  const hours = Math.abs(differenceInHours(next, now));

  if (minutes === 0) return ['1', 'minute'];
  if (minutes === 1) return [minutes, 'minute'];
  if (minutes < 60 && minutes > 1) return [minutes, 'minutes'];

  if (hours === 1) return [hours, 'hour'];
  if (hours <= 24 && hours > 1) return [hours, 'hours'];

  return [hours, 'hours'];
}
