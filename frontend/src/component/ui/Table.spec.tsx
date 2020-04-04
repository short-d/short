import React from 'react';
import { render } from '@testing-library/react';
import { Table } from './Table';

describe('Table component', () => {
  const sampleHeaders = ['Header1', 'Header2', 'Header3'];
  const sampleRows = [
    ['row1-cell1', 'row1-cell2', 'row1-cell3'],
    ['row2-cell1', 'row2-cell2', 'row2-cell3'],
    ['row3-cell1', 'row3-cell2', 'row3-cell3']
  ];

  test('should render without crash', () => {
    render(<Table />);
  });

  test('should have no children for thead when Headers prop not passed', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(<Table ref={toastRef} rows={sampleRows} />);

    expect(container.querySelector('thead')).not.toBeNull();
    expect(container.querySelector('thead')!.childElementCount).toBe(0);
  });

  test('should have no children for thead when Headers prop passed empty', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} rows={sampleRows} headers={[]} />
    );

    expect(container.querySelector('thead')).not.toBeNull();
    expect(container.querySelector('thead')!.childElementCount).toBe(0);
  });

  test('should have no children for tbody when rows prop not passed', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} headers={sampleHeaders} />
    );

    expect(container.querySelector('tbody')).not.toBeNull();
    expect(container.querySelector('tbody')!.childElementCount).toBe(0);
  });

  test('should have no children for tbody when rows prop passed empty', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} rows={[]} headers={sampleHeaders} />
    );

    expect(container.querySelector('tbody')).not.toBeNull();
    expect(container.querySelector('tbody')!.childElementCount).toBe(0);
  });

  test('should render when passed Headers correctly', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} rows={sampleRows} headers={sampleHeaders} />
    );

    const tableHeaders = container.querySelectorAll('thead > tr > th');
    expect(tableHeaders).toHaveLength(sampleHeaders.length);
    for (let column = 0; column < sampleHeaders.length; column++) {
      expect(tableHeaders[column].innerHTML).toMatch(sampleHeaders[column]);
    }
  });

  test('should render when passed rows correctly', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} rows={sampleRows} headers={sampleHeaders} />
    );

    const tableRows = container.querySelectorAll('tbody > tr');
    expect(tableRows).toHaveLength(sampleRows.length);
    for (let row = 0; row < sampleRows.length; row++) {
      const curRowData = tableRows[row].querySelectorAll('td');
      expect(curRowData).toHaveLength(sampleRows[row].length);
      for (let column = 0; column < sampleRows[row].length; column++) {
        expect(curRowData[column].innerHTML).toMatch(sampleRows[row][column]);
      }
    }
  });

  test('should render when unequal row widths are passed', () => {
    const unequalRows = [
      ['row1-cell1', 'row1-cell2', 'row1-cell3'],
      ['row2-cell1', 'row2-cell2'],
      ['row3-cell1', 'row3-cell2', 'row3-cell3', 'row3-cell4', 'row3-cell5']
    ];
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} rows={unequalRows} headers={sampleHeaders} />
    );

    const tableRows = container.querySelectorAll('tbody > tr');
    expect(tableRows).toHaveLength(unequalRows.length);
    for (let row = 0; row < unequalRows.length; row++) {
      const curRowData = tableRows[row].querySelectorAll('td');
      expect(curRowData).toHaveLength(unequalRows[row].length);
      for (let column = 0; column < unequalRows[row].length; column++) {
        expect(curRowData[column].innerHTML).toMatch(unequalRows[row][column]);
      }
    }
  });
});
