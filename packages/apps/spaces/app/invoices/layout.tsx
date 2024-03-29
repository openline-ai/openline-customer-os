'use client';
import { GridItem } from '@ui/layout/Grid';
import { PageLayout } from '@shared/components/PageLayout';
import { RootSidenav } from '@shared/components/RootSidenav/RootSidenav';

export default function InvoicesLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <PageLayout>
      <RootSidenav />
      <GridItem h='100%' area='content' overflowX='hidden' overflowY='auto'>
        {children}
      </GridItem>
    </PageLayout>
  );
}
