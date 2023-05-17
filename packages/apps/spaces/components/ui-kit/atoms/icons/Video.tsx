import * as React from 'react';
import { SVGProps } from 'react';
const SvgVideo = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='M13 18.75H6A2.75 2.75 0 0 1 3.25 16V8A2.75 2.75 0 0 1 6 5.25h7A2.75 2.75 0 0 1 15.75 8v8A2.75 2.75 0 0 1 13 18.75zm-7-12A1.25 1.25 0 0 0 4.75 8v8A1.25 1.25 0 0 0 6 17.25h7A1.25 1.25 0 0 0 14.25 16V8A1.25 1.25 0 0 0 13 6.75H6z' />
    <path d='M20 16.75a.79.79 0 0 1-.39-.11l-5-3a.75.75 0 0 1-.36-.64v-2a.75.75 0 0 1 .36-.64l5-3a.74.74 0 0 1 .76 0 .75.75 0 0 1 .38.65v8a.75.75 0 0 1-.38.65.71.71 0 0 1-.37.09zm-4.25-4.17 3.5 2.1V9.32l-3.5 2.1v1.16z' />
  </svg>
);
export default SvgVideo;
