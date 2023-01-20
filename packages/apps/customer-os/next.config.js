/** @type {import('next').NextConfig} */

const withPWA = require('next-pwa')({
  dest: 'public',
})

const config = {
  reactStrictMode: true,
  swcMinify: true,
  env: {
    CUSTOMER_OS_API_PATH: process.env.CUSTOMER_OS_API_PATH,
    CUSTOMER_OS_API_KEY: process.env.CUSTOMER_OS_API_KEY,

    SETTINGS_API_PATH: process.env.SETTINGS_API_PATH,
    SETTINGS_API_KEY: process.env.SETTINGS_API_KEY,

    OASIS_GUI_PATH: process.env.OASIS_GUI_PATH,

    NEXTAUTH_URL: process.env.NEXTAUTH_URL,
    NEXTAUTH_OAUTH_CLIENT_ID: process.env.NEXTAUTH_OAUTH_CLIENT_ID,
    NEXTAUTH_OAUTH_CLIENT_SECRET: process.env.NEXTAUTH_OAUTH_CLIENT_SECRET,
    NEXTAUTH_OAUTH_TENANT_ID: process.env.NEXTAUTH_OAUTH_TENANT_ID,
    NEXTAUTH_OAUTH_SERVER_URL: process.env.NEXTAUTH_OAUTH_SERVER_URL,
    NEXTAUTH_SECRET: process.env.NEXTAUTH_SECRET,

    WEB_CHAT_API_KEY:process.env.WEB_CHAT_API_KEY,
    WEB_CHAT_HTTP_PATH:process.env.WEB_CHAT_HTTP_PATH,
    WEB_CHAT_WS_PATH: process.env.WEB_CHAT_WS_PATH,

    WEB_CHAT_TRACKER_ENABLED: process.env.WEB_CHAT_TRACKER_ENABLED,
    WEB_CHAT_TRACKER_APP_ID: process.env.WEB_CHAT_TRACKER_APP_ID,
    WEB_CHAT_TRACKER_ID: process.env.WEB_CHAT_TRACKER_ID,
    WEB_CHAT_TRACKER_COLLECTOR_URL: process.env.WEB_CHAT_TRACKER_COLLECTOR_URL,
    WEB_CHAT_TRACKER_BUFFER_SIZE: process.env.WEB_CHAT_TRACKER_BUFFER_SIZE,
    WEB_CHAT_TRACKER_MINIMUM_VISIT_LENGTH: process.env.WEB_CHAT_TRACKER_MINIMUM_VISIT_LENGTH,
    WEB_CHAT_TRACKER_HEARTBEAT_DELAY: process.env.WEB_CHAT_TRACKER_HEARTBEAT_DELAY,
  },
  output: 'standalone'
}

module.exports = withPWA({
  // config
})
