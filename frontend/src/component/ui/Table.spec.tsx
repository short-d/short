import React from 'react';
import { render } from '@testing-library/react';
import { Table } from './Table';

describe('Table component', () => {
  test('should render without crash', () => {
    render(<Table />);
  });

  test('should have no children for thead without passing headers', () => {
    const tableRef = React.createRef<Table>();
    const { container } = render(<Table ref={tableRef} />);

    expect(container.querySelector('thead')).not.toBeNull();
    expect(container.querySelector('thead')!.childElementCount).toBe(0);
  });

  test('should have no children for thead with empty headers', () => {
    const tableRef = React.createRef<Table>();
    const { container } = render(<Table ref={tableRef} headers={[]} />);

    expect(container.querySelector('thead')).not.toBeNull();
    expect(container.querySelector('thead')!.childElementCount).toBe(0);
  });

  test('should have no children for tbody without passing rows', () => {
    const tableRef = React.createRef<Table>();
    const { container } = render(<Table ref={tableRef} />);

    expect(container.querySelector('tbody')).not.toBeNull();
    expect(container.querySelector('tbody')!.childElementCount).toBe(0);
  });

  test('should have no children for tbody with empty rows', () => {
    const tableRef = React.createRef<Table>();
    const { container } = render(<Table ref={tableRef} rows={[]} />);

    expect(container.querySelector('tbody')).not.toBeNull();
    expect(container.querySelector('tbody')!.childElementCount).toBe(0);
  });

  test('should render correctly with headers', () => {
    const sampleHeaders = ['Header1', 'Header2', 'Header3'];
    const tableRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={tableRef} headers={sampleHeaders} />
    );

    const tableHeaders = container.querySelectorAll('thead > tr > th');
    expect(tableHeaders).toHaveLength(sampleHeaders.length);

    for (let column = 0; column < sampleHeaders.length; column++) {
      expect(tableHeaders[column].innerHTML).toMatch(sampleHeaders[column]);
    }
  });

  test('should render correctly with rows', () => {
    const sampleRows = [
      ['row1-cell1', 'row1-cell2', 'row1-cell3'],
      ['row2-cell1', 'row2-cell2', 'row2-cell3'],
      ['row3-cell1', 'row3-cell2', 'row3-cell3']
    ];
    const tableRef = React.createRef<Table>();
    const { container } = render(<Table ref={tableRef} rows={sampleRows} />);

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

  test('should render correctly with both headers and rows', () => {
    const sampleHeaders = ['Header1', 'Header2', 'Header3'];
    const sampleRows = [
      ['row1-cell1', 'row1-cell2', 'row1-cell3'],
      ['row2-cell1', 'row2-cell2', 'row2-cell3'],
      ['row3-cell1', 'row3-cell2', 'row3-cell3']
    ];
    const tableRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={tableRef} headers={sampleHeaders} rows={sampleRows} />
    );

    const tableHeaders = container.querySelectorAll('thead > tr > th');
    expect(tableHeaders).toHaveLength(sampleHeaders.length);

    for (let column = 0; column < sampleHeaders.length; column++) {
      expect(tableHeaders[column].innerHTML).toMatch(sampleHeaders[column]);
    }

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

  test('should render correctly with unequal number of cells in rows', () => {
    const unequalRows = [
      ['row1-cell1', 'row1-cell2', 'row1-cell3'],
      ['row2-cell1', 'row2-cell2'],
      ['row3-cell1', 'row3-cell2', 'row3-cell3', 'row3-cell4', 'row3-cell5']
    ];
    const tableRef = React.createRef<Table>();
    const { container } = render(<Table ref={tableRef} rows={unequalRows} />);

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

  test('should have no class name defined for cells if colClassName not specified', () => {
    const sampleHeaders = ['Header1', 'Header2'];
    const sampleRows = [
      ['data1', 'data2'],
      ['data3', 'data4']
    ];
    const tableRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={tableRef} headers={sampleHeaders} rows={sampleRows} />
    );

    const tableHeaders = container.querySelectorAll('thead > tr > th');
    for (let column = 0; column < sampleHeaders.length; column++) {
      expect(tableHeaders[column].className).toEqual('');
    }

    const tableRows = container.querySelectorAll('tbody > tr');
    for (let row = 0; row < sampleRows.length; row++) {
      const curRowData = tableRows[row].querySelectorAll('td');
      for (let column = 0; column < sampleRows[row].length; column++) {
        expect(curRowData[column].className).toEqual('');
      }
    }
  });

  test('should have colClassNames set to the appropriate columns', () => {
    const sampleClassNames = ['className1', 'className2'];
    const sampleHeaders = ['Header1', 'Header2'];
    const sampleRows = [
      ['data1', 'data2'],
      ['data3', 'data4']
    ];
    const tableRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={tableRef} headers={sampleHeaders} rows={sampleRows} colClassNames={sampleClassNames} />
    );

    const tableHeaders = container.querySelectorAll('thead > tr > th');
    for (let column = 0; column < sampleHeaders.length; column++) {
      expect(tableHeaders[column].className).toEqual(sampleClassNames[column]);
    }

    const tableRows = container.querySelectorAll('tbody > tr');
    for (let row = 0; row < sampleRows.length; row++) {
      const curRowData = tableRows[row].querySelectorAll('td');
      for (let column = 0; column < sampleRows[row].length; column++) {
        expect(curRowData[column].className).toEqual(sampleClassNames[column]);
      }
    }
  });
});
