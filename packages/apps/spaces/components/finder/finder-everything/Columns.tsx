import React from 'react';
import { Column } from '../../ui-kit/atoms/table/types';
import { FinderMergeItemTableHeader } from '../finder-table/FinderMergeItemTableHeader';
import { ContactTableCell } from '../finder-table/ContactTableCell';
import { EmailTableCell } from '../finder-table/EmailTableCell';
import { AddressTableCell } from '../finder-table/AddressTableCell';
import { OrganizationTableCell } from '../finder-table/OrganizationTableCell';
import { ActionColumn } from '../finder-table/ActionTableHeader';

export const columns: Array<Column> = [
  {
    id: 'finder-table-column-contact',
    width: '25%',
    label: (
      <FinderMergeItemTableHeader
        mergeMode='MERGE_CONTACT'
        label='Name'
        subLabel='Role'
      />
    ),

    template: (c: any) => (
      <ContactTableCell contact={c?.contact} organization={c?.organization} />
    ),
  },
  {
    id: 'finder-table-column-email',
    width: '25%',
    label: 'Email',
    template: (c: any) => {
      if (!c?.contact) {
        return <span>-</span>;
      }
      return <EmailTableCell emails={c.contact?.emails} />;
    },
  },
  {
    id: 'finder-table-column-address',
    width: '25%',
    label: 'Location',
    subLabel: 'City, State, Country',
    template: (c: any) => {
      return <AddressTableCell locations={c?.contact?.locations} />;
    },
  },
  {
    id: 'finder-table-column-org',
    width: '25%',
    label: (
      <FinderMergeItemTableHeader
        mergeMode='MERGE_ORG'
        label='Organization'
        subLabel='Industry'
      />
    ),
    template: (c: any) => {
      return <OrganizationTableCell organization={c?.organization} />;
    },
  },
  {
    id: 'finder-table-column-actions',
    width: '12%',
    label: <ActionColumn />,
    subLabel: '',
    template: () => {
      return <div style={{ display: 'none' }} />;
    },
  },
];
