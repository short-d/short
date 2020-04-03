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

    expect(container.querySelector('thead')!.childElementCount).toBe(0);
  });

  test('should have no children for thead when Headers prop passed empty', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} rows={sampleRows} headers={[]} />
    );

    expect(container.querySelector('thead')!.childElementCount).toBe(0);
  });

  test('should have no children for tbody when rows prop not passed', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} headers={sampleHeaders} />
    );

    expect(container.querySelector('tbody')!.childElementCount).toBe(0);
  });

  test('should have no children for tbody when rows prop passed empty', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} rows={[]} headers={sampleHeaders} />
    );

    expect(container.querySelector('tbody')!.childElementCount).toBe(0);
  });

  test('should render passed Headers correctly', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} rows={sampleRows} headers={sampleHeaders} />
    );

    expect(container.querySelector('thead')!.childElementCount).toBe(1);
    for (let column = 0; column < sampleHeaders.length; column++) {
      expect(
        container.querySelectorAll('thead > tr > th')[column].innerHTML
      ).toMatch(sampleHeaders[column]);
    }
  });

  test('should render passed rows correctly', () => {
    const toastRef = React.createRef<Table>();
    const { container } = render(
      <Table ref={toastRef} rows={sampleRows} headers={sampleHeaders} />
    );

    expect(container.querySelector('tbody')!.childElementCount).toBe(
      sampleRows.length
    );
    for (let row = 0; row < sampleRows.length; row++) {
      for (let column = 0; column < sampleRows[row].length; column++) {
        expect(
          container.querySelectorAll('tbody > tr')[row].querySelectorAll('td')[
            column
          ].innerHTML
        ).toMatch(sampleRows[row][column]);
      }
    }
  });
});
