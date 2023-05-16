import * as React from 'react';
import { SVGProps } from 'react';
const SvgSend = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='none'
    {...props}
  >
    <path
      d='M20.33 3.67a1.45 1.45 0 0 0-1.47-.35L4.23 8.2a1.44 1.44 0 0 0-1 1.248A1.44 1.44 0 0 0 4 10.85l6.07 3 3 6.09a1.44 1.44 0 0 0 1.29.79h.1a1.43 1.43 0 0 0 1.26-1l4.95-14.59a1.41 1.41 0 0 0-.34-1.47zM4.85 9.58l12.77-4.26-7.09 7.09-5.68-2.83zm9.58 9.57-2.84-5.68 7.09-7.09-4.25 12.77z'
      fill='currentColor'
    />
  </svg>
);
export default SvgSend;
