import * as React from 'react';
import { SVGProps } from 'react';
const SvgAlignCenter = (props: SVGProps<SVGSVGElement>) => (
  <svg
    viewBox='0 0 24 24'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <path
      d='M17 10.75H7a.75.75 0 1 1 0-1.5h10a.75.75 0 1 1 0 1.5zm3-4H4a.75.75 0 0 1 0-1.5h16a.75.75 0 1 1 0 1.5zm0 8H4a.75.75 0 1 1 0-1.5h16a.75.75 0 1 1 0 1.5zm-3 4H7a.75.75 0 1 1 0-1.5h10a.75.75 0 1 1 0 1.5z'
      fill='currentColor'
    />
  </svg>
);
export default SvgAlignCenter;
