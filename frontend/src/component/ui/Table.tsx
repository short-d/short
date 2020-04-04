import React, { Component, ReactChild } from 'react';

import './Table.scss';

interface IProps {
  headers?: ReactChild[];
  rows?: ReactChild[][];
}

export class Table extends Component<IProps> {
  private createHeaders(headers: ReactChild[] | undefined) {
    if (!headers || headers.length === 0) {
      return null;
    }
    return (
      <tr key={`header`}>
        {headers.map((cell: ReactChild, cellIndex: number) => {
          return (
            <th key={`cell-${cellIndex}`} className="table-cell">
              {cell}
            </th>
          );
        })}
      </tr>
    );
  }

  private createBody(rows: ReactChild[][] | undefined) {
    if (!rows || rows.length === 0) {
      return null;
    }
    return rows.map((row: ReactChild[], rowIndex: number) => {
      return <tr key={`row-${rowIndex}`}>{this.createBodyRow(row)}</tr>;
    });
  }

  private createBodyRow(row: ReactChild[]) {
    return row.map((cell: ReactChild, cellIndex: number) => {
      return (
        <td key={`cell-${cellIndex}`} className="table-cell">
          {cell}
        </td>
      );
    });
  }

  render() {
    const { headers, rows } = this.props;

    return (
      <div className="table-container">
        <table className="table">
          <thead>{this.createHeaders(headers)}</thead>
          <tbody>{this.createBody(rows)}</tbody>
        </table>
      </div>
    );
  }
}
