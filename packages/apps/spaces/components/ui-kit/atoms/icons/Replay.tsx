import * as React from 'react';
import { SVGProps } from 'react';
const SvgReplay = (props: SVGProps<SVGSVGElement>) => (
  <svg
    xmlns='http://www.w3.org/2000/svg'
    viewBox='0 0 24 24'
    fill='currentColor'
    {...props}
  >
    <path d='M12 20.75a7.26 7.26 0 0 1-7.25-7.25.75.75 0 0 1 .75-.75.75.75 0 0 1 .75.75 5.75 5.75 0 0 0 3.55 5.312 5.75 5.75 0 0 0 6.266-1.246 5.75 5.75 0 0 0 1.246-6.266A5.75 5.75 0 0 0 12 7.75H9.5A.75.75 0 0 1 8.75 7a.75.75 0 0 1 .75-.75H12a7.25 7.25 0 0 1 7.25 7.25A7.25 7.25 0 0 1 12 20.75z' />
    <path d='M12 10.75a.74.74 0 0 1-.53-.22l-3-3a.75.75 0 0 1 0-1.06l3-3a.75.75 0 0 1 .535-.239.75.75 0 0 1 .543.22.75.75 0 0 1 .22.544.75.75 0 0 1-.239.535L10.06 7l2.47 2.47a.75.75 0 0 1 0 1.06.74.74 0 0 1-.53.22z' />
  </svg>
);
export default SvgReplay;
