import React, { Component, ReactChild } from 'react';

import './Table.scss';

interface IProps {
  headers?: ReactChild[];
  rows?: ReactChild[][];
  widths?: string[];
}

export class Table extends Component<IProps> {
  private createHeaders(
    headers: ReactChild[] | undefined,
    widths: string[] | undefined
  ) {
    if (!headers || headers.length === 0) {
      return null;
    }
    return (
      <tr key={`header`}>
        {headers.map((cell: ReactChild, cellIndex: number) => {
          return (
            <th
              key={`cell-${cellIndex}`}
              style={{ width: widths?.[cellIndex] }}
            >
              {cell}
            </th>
          );
        })}
      </tr>
    );
  }

  private createBody(
    rows: ReactChild[][] | undefined,
    widths: string[] | undefined
  ) {
    if (!rows || rows.length === 0) {
      return null;
    }
    return rows.map((row: ReactChild[], rowIndex: number) => {
      return <tr key={`row-${rowIndex}`}>{this.createBodyRow(row, widths)}</tr>;
    });
  }

  private createBodyRow(row: ReactChild[], widths: string[] | undefined) {
    return row.map((cell: ReactChild, cellIndex: number) => {
      return (
        <td key={`cell-${cellIndex}`} style={{ width: widths?.[cellIndex] }}>
          {cell}
        </td>
      );
    });
  }

  render() {
    const { headers, rows, widths } = this.props;

    return (
      <div className="table-container">
        <table className="table">
          <thead>{this.createHeaders(headers, widths)}</thead>
          <tbody>{this.createBody(rows, widths)}</tbody>
        </table>
      </div>
    );
  }
}
