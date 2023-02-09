import * as React from 'react';
import { SVGProps } from 'react';
const SvgComment = (props: SVGProps<SVGSVGElement>) => (
  <svg
    width={24}
    height={24}
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <path
      d='M4.5 20.25a.75.75 0 0 1-.72-1L5.38 14a7.76 7.76 0 0 1-.52-2.83 8 8 0 0 1 .62-3.1 8.12 8.12 0 0 1 1.7-2.52 7.83 7.83 0 0 1 2.53-1.7 7.92 7.92 0 0 1 6.19 0 8 8 0 0 1 4.85 7.32 8 8 0 0 1-2.33 5.62 8.121 8.121 0 0 1-2.52 1.7 8 8 0 0 1-5.93.1l-5.25 1.6a.832.832 0 0 1-.22.06Zm8.3-15.5a6.49 6.49 0 0 0-5.94 3.94 6.55 6.55 0 0 0 0 5 .75.75 0 0 1 0 .51l-1.23 4.17 4.15-1.26a.75.75 0 0 1 .51 0 6.52 6.52 0 0 0 5 0 6.44 6.44 0 0 0 3.43-8.45 6.45 6.45 0 0 0-5.92-3.91Z'
      fill='currentColor'
    />
  </svg>
);
export default SvgComment;
