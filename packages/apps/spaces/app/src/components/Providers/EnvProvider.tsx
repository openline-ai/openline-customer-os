'use client';

import { createContext } from 'react';

export type Env = {
  PRODUCTION: string;
  NOTIFICATION_PROD_APP_IDENTIFIER: string;
  NOTIFICATION_TEST_APP_IDENTIFIER: string;
};

export const EnvContext = createContext<Env>({
  PRODUCTION: '',
  NOTIFICATION_PROD_APP_IDENTIFIER: '',
  NOTIFICATION_TEST_APP_IDENTIFIER: '',
});

interface EnvProviderProps {
  env: Env;
  children: React.ReactNode;
}

export const EnvProvider = ({ children, env }: EnvProviderProps) => {
  return <EnvContext.Provider value={env}>{children}</EnvContext.Provider>;
};
