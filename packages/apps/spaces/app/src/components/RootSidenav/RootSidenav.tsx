'use client';
import React, { useMemo, useEffect } from 'react';
import { useRouter, usePathname, useSearchParams } from 'next/navigation';

import { produce } from 'immer';
import { signOut } from 'next-auth/react';
import { useLocalStorage } from 'usehooks-ts';
import { useFeatureIsOn } from '@growthbook/growthbook-react';
import { useTenantSettingsQuery } from '@settings/graphql/getTenantSettings.generated';

import { Flex } from '@ui/layout/Flex';
import { Icons } from '@ui/media/Icon';
import { VStack } from '@ui/layout/Stack';
import { Text } from '@ui/typography/Text';
import { GridItem } from '@ui/layout/Grid';
import { Button } from '@ui/form/Button/Button';
import { Receipt } from '@ui/media/icons/Receipt';
import { Bubbles } from '@ui/media/icons/Bubbles';
import { LogOut01 } from '@ui/media/icons/LogOut01';
import { Skeleton } from '@ui/presentation/Skeleton';
import { Image as ChakraImage } from '@ui/media/Image';
import { mockedTableDefs } from '@shared/util/tableDefs.mock';
import { getGraphQLClient } from '@shared/util/getGraphQLClient';
import { ClockFastForward } from '@ui/media/icons/ClockFastForward';
import { useOrganizationsMeta } from '@shared/state/OrganizationsMeta.atom';
import { useGlobalCacheQuery } from '@shared/graphql/global_Cache.generated';
import { useTableViewDefsQuery } from '@shared/graphql/tableViewDefs.generated';
import { NotificationCenter } from '@shared/components/Notifications/NotificationCenter';
import { useGetAllInvoicesCountQuery } from '@shared/graphql/getAllInvoicesCount.generated';

import { SidenavItem } from './components/SidenavItem';
import logoCustomerOs from './assets/logo-customeros.png';
import { GoogleSidebarNotification } from './components/GoogleSidebarNotification';

export const RootSidenav = () => {
  const client = getGraphQLClient();
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const [_, setOrganizationsMeta] = useOrganizationsMeta();
  const showMyViewsItems = useFeatureIsOn('my-views-nav-item');
  const [lastActivePosition, setLastActivePosition] = useLocalStorage(
    `customeros-player-last-position`,
    { root: 'organization' },
  );

  const { data: tenantSettingsData } = useTenantSettingsQuery(client);
  const { data: totalInvoices } = useGetAllInvoicesCountQuery(client);

  const { data: tableViewDefsData } = useTableViewDefsQuery(
    client,
    {
      pagination: { limit: 100, page: 1 },
    },
    {
      enabled: false,
      placeholderData: { tableViewDefs: { content: mockedTableDefs } },
    },
  );
  const { data, isLoading } = useGlobalCacheQuery(client);
  const globalCache = data?.global_Cache;
  const myViews = tableViewDefsData?.tableViewDefs?.content ?? [];

  const handleItemClick = (path: string) => {
    setLastActivePosition({ ...lastActivePosition, root: path });
    setOrganizationsMeta((prev) =>
      produce(prev, (draft) => {
        draft.getOrganization.pagination.page = 1;
      }),
    );

    router.push(`/${path}`);
  };
  const checkIsActive = (path: string, options?: { preset: string }) => {
    const [_pathName, _searchParams] = path.split('?');
    const presetParam = new URLSearchParams(searchParams?.toString()).get(
      'preset',
    );

    if (options?.preset) {
      return (
        pathname?.startsWith(`/${_pathName}`) && presetParam === options.preset
      );
    } else {
      return pathname?.startsWith(`/${_pathName}`) && !presetParam;
    }
  };

  const handleSignOutClick = () => {
    signOut();
  };

  const showInvoices = useMemo(() => {
    return (
      tenantSettingsData?.tenantSettings?.billingEnabled ||
      totalInvoices?.invoices?.totalElements > 0
    );
  }, [tenantSettingsData?.tenantSettings?.billingEnabled]);

  useEffect(() => {
    [
      '/organizations',
      '/organizations?preset=customer',
      '/organizations?preset=portfolio',
      '/renewals?preset=1',
      '/renewals?preset=2',
      '/renewals?preset=3',
    ].forEach((path) => {
      router.prefetch(path);
    });
  }, []);

  return (
    <GridItem
      px='2'
      pt='2.5'
      pb='4'
      h='full'
      w='200px'
      bg='white'
      display='flex'
      flexDir='column'
      gridArea='sidebar'
      position='relative'
      borderRight='1px solid'
      borderColor='gray.200'
    >
      <Flex
        tabIndex={0}
        role='button'
        cursor='pointer'
        justify='flex-start'
        overflow='hidden'
        position='relative'
      >
        {isLoading && <Skeleton w='180px' h='30px' mb='4' ml='3' />}
        {data?.global_Cache?.cdnLogoUrl && (
          <Flex
            position='relative'
            maxHeight={120}
            width='full'
            justifyContent='center'
          >
            <ChakraImage
              alt='CustomerOS'
              pointerEvents='none'
              src={data.global_Cache.cdnLogoUrl}
              transition='opacity 0.25s ease-in-out'
              width={136}
              height={40}
              style={{
                objectFit: 'contain',
                maxHeight: '40px',
                maxWidth: 'fit-content',
              }}
            />
          </Flex>
        )}
        {!data?.global_Cache?.cdnLogoUrl && !isLoading && (
          <ChakraImage
            width={136}
            height={30}
            mb='4'
            ml='3'
            w='136px'
            h='30px'
            alt='CustomerOS'
            pointerEvents='none'
            src={logoCustomerOs}
            transition='opacity 0.25s ease-in-out'
          />
        )}
      </Flex>

      <VStack spacing='2' w='full' mb='4'>
        <SidenavItem
          label='Customer map'
          isActive={checkIsActive('customer-map')}
          onClick={() => handleItemClick('customer-map')}
          icon={(isActive) => (
            <Bubbles boxSize='5' color={isActive ? 'gray.700' : 'gray.500'} />
          )}
        />
        <SidenavItem
          label='Organizations'
          isActive={checkIsActive('organizations')}
          onClick={() => handleItemClick('organizations')}
          icon={(isActive) => (
            <Icons.Building7
              boxSize='5'
              color={isActive ? 'gray.700' : 'gray.500'}
            />
          )}
        />
        <SidenavItem
          label='Customers'
          isActive={checkIsActive('organizations', { preset: 'customer' })}
          onClick={() => handleItemClick('organizations?preset=customer')}
          icon={(isActive) => (
            <Icons.CheckHeart
              boxSize='5'
              color={isActive ? 'gray.700' : 'gray.500'}
            />
          )}
        />

        {showInvoices && (
          <SidenavItem
            label='Invoices'
            isActive={checkIsActive('invoices')}
            onClick={() => handleItemClick('invoices')}
            icon={
              <Receipt
                color={checkIsActive('invoices') ? 'gray.700' : 'gray.500'}
                boxSize='5'
              />
            }
          />
        )}
      </VStack>

      <VStack spacing='2' w='full'>
        {(globalCache?.isOwner || showMyViewsItems) && (
          <Flex w='full' justify='flex-start' pl='3.5'>
            <Text color='gray.500' fontSize='sm'>
              My views
            </Text>
          </Flex>
        )}

        {globalCache?.isOwner && (
          <SidenavItem
            label='My portfolio'
            isActive={checkIsActive('organizations', {
              preset: 'portfolio',
            })}
            onClick={() => handleItemClick('organizations?preset=portfolio')}
            icon={(isActive) => (
              <Icons.Briefcase1
                boxSize='5'
                color={isActive ? 'gray.700' : 'gray.500'}
              />
            )}
          />
        )}
        {showMyViewsItems &&
          myViews.map((view) => (
            <SidenavItem
              key={view.id}
              label={view.name}
              isActive={checkIsActive('renewals', { preset: view.id })}
              onClick={() => handleItemClick(`renewals?preset=${view.id}`)}
              icon={(isActive) => (
                <ClockFastForward
                  boxSize='5'
                  color={isActive ? 'gray.700' : 'gray.500'}
                />
              )}
            />
          ))}
      </VStack>
      <Button
        isLoading
        loadingText='proba'
        colorScheme='blueDark'
        variant='solid'
      >
        plm
      </Button>

      <VStack
        spacing='1'
        flexDir='column'
        flexWrap='initial'
        flexGrow='1'
        justifyContent='flex-end'
        sx={{
          '& > span': {
            width: '100%',
          },
        }}
      >
        <NotificationCenter />
        <GoogleSidebarNotification />

        <SidenavItem
          label='Settings'
          isActive={checkIsActive('settings')}
          onClick={() => router.push('/settings')}
          icon={(isActive) => (
            <Icons.Settings
              boxSize='5'
              color={isActive ? 'gray.700' : 'gray.500'}
            />
          )}
        />
        <SidenavItem
          label='Sign out'
          isActive={false}
          onClick={handleSignOutClick}
          icon={() => <LogOut01 boxSize='5' color='gray.500' />}
        />
      </VStack>

      <Flex h='64px' />
    </GridItem>
  );
};
