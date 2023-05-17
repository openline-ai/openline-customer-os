import * as React from 'react';
import { SVGProps } from 'react';
const SvgCalendarTimes = (props: SVGProps<SVGSVGElement>) => (
  <svg
    viewBox='0 0 24 24'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <g fill='currentColor'>
      <path d='M17 4.75h-1.25V3.5a.75.75 0 1 0-1.5 0v1.25h-4.5V3.5a.75.75 0 0 0-1.5 0v1.25H7A2.75 2.75 0 0 0 4.25 7.5v11A2.75 2.75 0 0 0 7 21.25h10a2.75 2.75 0 0 0 2.75-2.75v-11A2.75 2.75 0 0 0 17 4.75zM7 6.25h1.25V7.5a.75.75 0 0 0 1.5 0V6.25h4.5V7.5a.75.75 0 1 0 1.5 0V6.25H17a1.25 1.25 0 0 1 1.25 1.25v2.75H5.75V7.5A1.25 1.25 0 0 1 7 6.25zm10 13.5H7a1.25 1.25 0 0 1-1.25-1.25v-6.75h12.5v6.75A1.25 1.25 0 0 1 17 19.75z' />
      <path d='M13.94 14.06a.74.74 0 0 0-1.06 0l-.88.88-.88-.88a.75.75 0 0 0-1.06 1.06l.88.88-.88.88a.74.74 0 0 0 0 1.06.71.71 0 0 0 .53.22.74.74 0 0 0 .53-.22l.88-.88.88.88a.74.74 0 0 0 .53.22.71.71 0 0 0 .53-.22.74.74 0 0 0 0-1.06l-.88-.88.88-.88a.74.74 0 0 0 0-1.06z' />
    </g>
  </svg>
);
export default SvgCalendarTimes;
