import React from 'react';
import { Table } from '../../frontend/src/component/ui/Table';
import { array, number, select } from '@storybook/addon-knobs';

export default {
  title: 'UI/Table',
  component: Table
};

function alignSelector(label: string) {
  return select(
    label,
    {
      center: 'center',
      end: 'end',
      justify: 'justify',
      left: 'left',
      'match-parent': 'match-parent',
      right: 'right',
      start: 'start'
    },
    'left'
  );
}

export const table = () => {
  const headers = array('Header Columns', ['Long Link', 'Alias']);
  return (
    <Table
      headers={headers}
      rows={new Array(number('Number of rows', 2))
        .fill(0)
        .map((_, i) => headers.map((_, j) => `data ${i}-${j}`))}
      headerFontSize={`${number('Header Font Size (px)', 16)}px`}
      widths={headers.map(
        (_, i) => number(`Column ${i + 1} Width (%)`, 50) + '%'
      )}
      alignHeaders={headers.map((_, i) =>
        alignSelector(`Column ${i + 1} Header Alignment (%)`)
      )}
      alignBodyColumns={headers.map((_, i) =>
        alignSelector(`Column ${i + 1} Body Column Alignment (%)`)
      )}
    />
  );
};
