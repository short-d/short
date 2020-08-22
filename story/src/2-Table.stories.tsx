import React from 'react';
import { Table } from '../../frontend/src/component/ui/Table';
import { array, number } from '@storybook/addon-knobs';

export default {
  title: 'UI/Table',
  component: Table
};

export const table = () => {
  const headers = array('Header Columns', ['Long Link', 'Alias']);
  return (
    <Table
      headers={headers}
      rows={new Array(number('Number of rows', 2))
        .fill(0)
        .map((_, i) => headers.map((_, j) => `data ${i}-${j}`))}
    />
  );
};
