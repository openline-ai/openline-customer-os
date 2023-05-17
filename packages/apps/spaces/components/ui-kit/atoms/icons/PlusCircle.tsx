import * as React from 'react';
import { SVGProps } from 'react';
const SvgPlusCircle = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='M12 21A9 9 0 0 1 5.636 5.636 9 9 0 0 1 21 12a9 9 0 0 1-9 9zm0-16.5a7.5 7.5 0 0 0-5.303 12.803A7.5 7.5 0 0 0 19.5 12 7.5 7.5 0 0 0 12 4.5zm0 12.25a.76.76 0 0 1-.75-.75V8a.75.75 0 1 1 1.5 0v8a.76.76 0 0 1-.75.75z' />
    <path d='M16 12.75H8a.75.75 0 0 1-.75-.75.75.75 0 0 1 .75-.75h8a.75.75 0 0 1 .75.75.75.75 0 0 1-.75.75z' />
  </svg>
);
export default SvgPlusCircle;
