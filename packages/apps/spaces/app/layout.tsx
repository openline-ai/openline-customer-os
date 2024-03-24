import { Metadata } from 'next';
import Script from 'next/script';
import localFont from 'next/font/local';
import { getServerSession } from 'next-auth';
import { ToastContainer } from 'react-toastify';
import { Merriweather } from 'next/font/google';

import { HighlightInit } from '@highlight-run/next/client';

import { ThemeProvider } from '@ui/theme/ThemeProvider';

import { Providers } from './src/components/Providers/Providers';

import './../styles/globals.scss';
import './../styles/date-picker.scss';
import './../styles/remirror-editor.scss';
import 'filepond-plugin-image-preview/dist/filepond-plugin-image-preview.min.css';
import 'filepond/dist/filepond.min.css';
import 'react-toastify/dist/ReactToastify.css';
import 'react-datepicker/dist/react-datepicker.css';
import './../styles/filepond.scss';

const barlow = localFont({
  src: [
    { path: './src/fonts/Barlow-Regular.woff', weight: '400', style: 'normal' },
    { path: './src/fonts/Barlow-Medium.woff', weight: '500', style: 'normal' },
    {
      path: './src/fonts/Barlow-SemiBold.woff',
      weight: '600',
      style: 'normal',
    },
  ],
  preload: false,
  display: 'swap',
  variable: '--font-barlow',
});

const merriweather = Merriweather({
  weight: ['300', '400', '700', '900'],
  preload: false,
  display: 'swap',
  variable: '--font-merriweather',
});

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const session = await getServerSession();

  const env = {
    PRODUCTION: process.env.NEXT_PUBLIC_PRODUCTION ?? '',
    NOTIFICATION_PROD_APP_IDENTIFIER:
      process.env.NEXT_PUBLIC_NOTIFICATION_PROD_APP_IDENTIFIER ?? '',
    NOTIFICATION_TEST_APP_IDENTIFIER:
      process.env.NEXT_PUBLIC_NOTIFICATION_TEST_APP_IDENTIFIER ?? '',
    REALTIME_WS_PATH: process.env.REALTIME_WS_PATH ?? '',
    REALTIME_WS_API_KEY: process.env.REALTIME_WS_API_KEY ?? '',
  };

  const isProduction = env.PRODUCTION === 'true';

  return (
    <>
      {isProduction && (
        <HighlightInit
          environment={isProduction ? 'production' : 'development'}
          projectId={'ldwno7wd'}
          serviceName='customer-os'
          tracingOrigins
          networkRecording={{
            enabled: isProduction,
            recordHeadersAndBody: isProduction,
            urlBlocklist: [],
          }}
        />
      )}

      <html
        lang='en'
        className={`${barlow.variable} ${merriweather.variable}`}
        data-theme='light'
      >
        <Script
          async
          strategy='afterInteractive'
          id='openline-customer-os-clarity-script'
          dangerouslySetInnerHTML={{
            __html: `(function(c,l,a,r,i,t,y){
                        c[a]=c[a]||function(){(c[a].q=c[a].q||[]).push(arguments)};
                        t=l.createElement(r);t.async=1;t.src="https://www.clarity.ms/tag/"+i;
                        y=l.getElementsByTagName(r)[0];y.parentNode.insertBefore(t,y);
                    })(window, document, "clarity", "script", "fryzkewrjw");`,
          }}
        />
        {isProduction && (
          <Script
            async
            strategy='afterInteractive'
            id='openline-customer-os-prod-heap-script'
            dangerouslySetInnerHTML={{
              __html: `
            window.heap=window.heap||[],heap.load=function(e,t){window.heap.appid=e,window.heap.config=t=t||{};var r=document.createElement("script");r.type="text/javascript",r.async=!0,r.src="https://cdn.heapanalytics.com/js/heap-"+e+".js";var a=document.getElementsByTagName("script")[0];a.parentNode.insertBefore(r,a);for(var n=function(e){return function(){heap.push([e].concat(Array.prototype.slice.call(arguments,0)))}},p=["addEventProperties","addUserProperties","clearEventProperties","identify","resetIdentity","removeEventProperty","setEventProperties","track","unsetEventProperty"],o=0;o<p.length;o++)heap[p[o]]=n(p[o])};
            heap.load("1078792267");
            `,
            }}
          />
        )}

        {!isProduction && (
          <Script
            async
            strategy='afterInteractive'
            id='openline-customer-os-dev-heap-script'
            dangerouslySetInnerHTML={{
              __html: `
            window.heap=window.heap||[],heap.load=function(e,t){window.heap.appid=e,window.heap.config=t=t||{};var r=document.createElement("script");r.type="text/javascript",r.async=!0,r.src="https://cdn.heapanalytics.com/js/heap-"+e+".js";var a=document.getElementsByTagName("script")[0];a.parentNode.insertBefore(r,a);for(var n=function(e){return function(){heap.push([e].concat(Array.prototype.slice.call(arguments,0)))}},p=["addEventProperties","addUserProperties","clearEventProperties","identify","resetIdentity","removeEventProperty","setEventProperties","track","unsetEventProperty"],o=0;o<p.length;o++)heap[p[o]]=n(p[o])};
  heap.load("3563674186");
            `,
            }}
          />
        )}

        <Script
          async
          strategy='afterInteractive'
          id='openline-customer-os-intercom-user-data-script'
          dangerouslySetInnerHTML={{
            __html: `
        window.intercomSettings = {
          api_base: "https://api-iam.intercom.io",
          app_id: "pqerb2dx",
          alignment: "left",
          horizontal_padding: 28,
          vertical_padding: 28,
          name: "${session?.user.name}",
          email: "${session?.user.email}",
          created_at: ${new Date().valueOf()} // Signup date as a Unix timestamp
        };
        `,
          }}
        />

        <Script
          async
          strategy='afterInteractive'
          id='openline-customer-os-intercom-script'
          dangerouslySetInnerHTML={{
            __html: `
          (function(){var w=window;var ic=w.Intercom;if(typeof ic==="function"){ic('reattach_activator');ic('update',w.intercomSettings);}else{var d=document;var i=function(){i.c(arguments);};i.q=[];i.c=function(args){i.q.push(args);};w.Intercom=i;var l=function(){var s=d.createElement('script');s.type='text/javascript';s.async=true;s.src='https://widget.intercom.io/widget/pqerb2dx';var x=d.getElementsByTagName('script')[0];x.parentNode.insertBefore(s,x);};if(document.readyState==='complete'){l();}else if(w.attachEvent){w.attachEvent('onload',l);}else{w.addEventListener('load',l,false);}}})();
        `,
          }}
        />

        <body className='scrollbar'>
          <ThemeProvider>
            <Providers isProduction={isProduction} env={env}>
              {children}
              <ToastContainer
                position='bottom-right'
                autoClose={8000}
                limit={3}
                closeOnClick={true}
                hideProgressBar={true}
                theme='colored'
              />
            </Providers>
          </ThemeProvider>
          <div id='tooltip-portal' className='absolute z-[1000]' />
        </body>
      </html>
    </>
  );
}

export const metadata: Metadata = {
  title: 'Customer OS',
  description: 'Customer OS',
  keywords: ['CustomerOS', 'Spaces', 'Openline'],
  viewport: 'width=device-width,initial-scale=1',
};
