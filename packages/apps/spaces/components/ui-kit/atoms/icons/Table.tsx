import * as React from 'react';
import { SVGProps } from 'react';
const SvgTable = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M18 4.25H6A2.75 2.75 0 0 0 3.25 7v10A2.75 2.75 0 0 0 6 19.75h12A2.75 2.75 0 0 0 20.75 17V7A2.75 2.75 0 0 0 18 4.25zM19.25 7v4.25h-6.5v-5.5H18A1.25 1.25 0 0 1 19.25 7zM6 5.75h5.25v5.5h-6.5V7A1.25 1.25 0 0 1 6 5.75zM4.75 17v-4.25h6.5v5.5H6A1.25 1.25 0 0 1 4.75 17zM18 18.25h-5.25v-5.5h6.5V17A1.25 1.25 0 0 1 18 18.25z'
      fill='currentColor'
    />
  </svg>
);
export default SvgTable;
