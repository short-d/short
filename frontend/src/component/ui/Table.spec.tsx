import React from 'react';
import { render } from '@testing-library/react';
import { Table } from './Table';

describe('Table component', () => {
  const sampleHeadings = ['Heading1', 'Heading2', 'Heading3'];
  const sampleRows = [
    ['row1-cell1', 'row1-cell2', 'row1-cell3'],
    ['row2-cell1', 'row2-cell2', 'row2-cell3'],
    ['row3-cell1', 'row3-cell2', 'row3-cell3']
  ];

  test('should render without crash', () => {
    render(<Table />);
  });

  test('should have no children for thead when headings prop not passed', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(<Table ref={toastRef} rows={sampleRows} />);

    expect(container.getElementsByTagName('thead')[0].childElementCount).toBe(
      0
    );
  });

  test('should have no children for thead when headings prop passed empty', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} rows={sampleRows} headings={[]} />
    );

    expect(container.getElementsByTagName('thead')[0].childElementCount).toBe(
      0
    );
  });

  test('should have no children for tbody when rows prop not passed', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} headings={sampleHeadings} />
    );

    expect(container.getElementsByTagName('tbody')[0].childElementCount).toBe(
      0
    );
  });

  test('should have no children for tbody when rows prop passed empty', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} rows={[]} headings={sampleHeadings} />
    );

    expect(container.getElementsByTagName('tbody')[0].childElementCount).toBe(
      0
    );
  });

  test('should render passed headings correctly', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} rows={sampleRows} headings={sampleHeadings} />
    );

    expect(container.getElementsByTagName('thead')[0].childElementCount).toBe(
      1
    );
    for (let i = 0; i < sampleHeadings.length; i++) {
      expect(
        container
          .getElementsByTagName('thead')[0]
          .getElementsByTagName('tr')[0]
          .getElementsByTagName('th')[i].innerHTML
      ).toMatch(sampleHeadings[i]);
    }
  });

  test('should render passed rows correctly', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} rows={sampleRows} headings={sampleHeadings} />
    );

    expect(container.getElementsByTagName('tbody')[0].childElementCount).toBe(
      sampleRows.length
    );
    for (let i = 0; i < sampleRows.length; i++) {
      for (let j = 0; j < sampleRows[i].length; j++) {
        expect(
          container
            .getElementsByTagName('tbody')[0]
            .getElementsByTagName('tr')
            [i].getElementsByTagName('td')[j].innerHTML
        ).toMatch(sampleRows[i][j]);
      }
    }
  });
});
