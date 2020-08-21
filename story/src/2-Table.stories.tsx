import React from 'react';
import { Table } from '../../frontend/src/component/ui/Table';
import { action } from '@storybook/addon-actions';
import { text } from '@storybook/addon-knobs';

export default {
  title: 'UI/Table',
  component: Table
};

export const table = () => {
  return (
    <Table
      headers={['Long Link', 'Alias']}
      rows={[
        ['http://www.google.com', 'google'],
        ['http://www.facebook.com', 'facebook']
      ]}
    />
  );
};
