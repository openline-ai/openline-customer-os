import * as React from 'react';
import { SVGProps } from 'react';
const SvgNoteTimeline = (props: SVGProps<SVGSVGElement>) => (
  <svg
    viewBox='0 0 144 58'
    fill='none'
    xmlns='http://www.w3.org/2000/svg'
    {...props}
  >
    <g filter='url(#note-timeline_svg__a)'>
      <rect
        x={4}
        width={136}
        height={50}
        rx={4}
        fill='#fff'
        shapeRendering='crispEdges'
      />
      <path
        d='M4.5 6V4A3.5 3.5 0 0 1 8 .5h128a3.5 3.5 0 0 1 3.5 3.5v2H4.5Z'
        fill='#F4E69C'
      />
      <path
        d='M89.665 26.602V33h-.853l-3.221-4.935V33h-.848v-6.398h.848l3.234 4.948v-4.948h.84Zm1.164 4.073v-.1c0-.344.05-.661.15-.954.1-.296.243-.553.43-.77.188-.22.415-.39.681-.51.267-.122.566-.184.897-.184.334 0 .634.062.9.185.27.12.499.29.686.51.19.216.336.473.435.769.1.293.15.61.15.953v.101c0 .343-.05.66-.15.954-.1.293-.244.55-.435.769a1.993 1.993 0 0 1-.68.51 2.14 2.14 0 0 1-.897.18 2.17 2.17 0 0 1-.901-.18 2.037 2.037 0 0 1-.686-.51 2.322 2.322 0 0 1-.43-.77c-.1-.292-.15-.61-.15-.953Zm.813-.1v.1c0 .238.028.462.084.673.055.208.139.392.25.553.114.162.256.289.426.383.17.09.368.136.594.136.222 0 .417-.046.584-.136.17-.094.31-.221.422-.383a1.77 1.77 0 0 0 .25-.553c.059-.211.088-.435.088-.673v-.1c0-.235-.029-.456-.088-.664a1.71 1.71 0 0 0-.254-.559 1.213 1.213 0 0 0-.422-.386 1.184 1.184 0 0 0-.59-.14 1.2 1.2 0 0 0-.588.14 1.266 1.266 0 0 0-.422.386 1.763 1.763 0 0 0-.25.559c-.056.208-.084.429-.084.663Zm6.526-2.33v.624h-2.57v-.624h2.57Zm-1.7-1.156h.812v4.733c0 .161.025.283.075.365.05.082.114.136.193.163a.8.8 0 0 0 .255.04c.068 0 .138-.007.211-.018l.172-.036.004.664c-.064.02-.15.04-.255.057a1.94 1.94 0 0 1-.374.03c-.199 0-.382-.039-.549-.118a.883.883 0 0 1-.4-.395c-.096-.188-.145-.44-.145-.756v-4.729Zm4.627 5.999c-.331 0-.632-.056-.901-.167a2.042 2.042 0 0 1-.69-.48 2.13 2.13 0 0 1-.44-.729 2.674 2.674 0 0 1-.153-.922v-.185c0-.387.057-.731.171-1.033.114-.304.27-.562.466-.773.196-.211.419-.37.668-.48.249-.108.507-.162.773-.162.34 0 .633.059.879.176.249.117.453.281.611.492.158.208.275.454.352.739.076.28.114.588.114.922v.365h-3.551v-.663h2.738v-.062a2.042 2.042 0 0 0-.132-.615c-.073-.2-.19-.363-.352-.492-.161-.13-.38-.194-.659-.194a1.111 1.111 0 0 0-.909.462 1.719 1.719 0 0 0-.26.558 2.83 2.83 0 0 0-.092.76v.185c0 .225.03.438.092.637.065.196.157.369.277.518.123.15.271.267.444.352.176.085.375.127.598.127.287 0 .53-.058.729-.175.199-.118.374-.274.523-.47l.492.39a2.262 2.262 0 0 1-.391.444c-.158.14-.353.255-.584.343-.229.088-.5.132-.813.132Z'
        fill='#1D1D1D'
      />
      <rect
        x={4.25}
        y={0.25}
        width={135.5}
        height={49.5}
        rx={3.75}
        stroke='#9B9B9B'
        strokeWidth={0.5}
        shapeRendering='crispEdges'
      />
    </g>
    <defs>
      <filter
        id='note-timeline_svg__a'
        x={0}
        y={0}
        width={144}
        height={58}
        filterUnits='userSpaceOnUse'
        colorInterpolationFilters='sRGB'
      >
        <feFlood floodOpacity={0} result='BackgroundImageFix' />
        <feColorMatrix
          in='SourceAlpha'
          values='0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0'
          result='hardAlpha'
        />
        <feOffset dy={4} />
        <feGaussianBlur stdDeviation={2} />
        <feComposite in2='hardAlpha' operator='out' />
        <feColorMatrix values='0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.25 0' />
        <feBlend
          in2='BackgroundImageFix'
          result='effect1_dropShadow_737_8734'
        />
        <feBlend
          in='SourceGraphic'
          in2='effect1_dropShadow_737_8734'
          result='shape'
        />
      </filter>
    </defs>
  </svg>
);
export default SvgNoteTimeline;
