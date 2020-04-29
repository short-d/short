import React, { Component, ReactChild } from 'react';

import './Table.scss';
import classNames from 'classnames';

interface IProps {
  headers?: ReactChild[];
  rows?: ReactChild[][];
  colNames?: string[];
}

export class Table extends Component<IProps> {
  private createHeaders(headers: ReactChild[] | undefined, colNames: string[] | undefined) {
    if (!headers || headers.length === 0) {
      return null;
    }
    return (
      <tr key={`header`}>
        {headers.map((cell: ReactChild, cellIndex: number) => {
          return (
            <th key={`cell-${cellIndex}`} className={!colNames ? "table-cell" : classNames("table-cell", colNames[cellIndex])}>
              {cell}
            </th>
          );
        })}
      </tr>
    );
  }

  private createBody(rows: ReactChild[][] | undefined, colNames: string[] | undefined) {
    if (!rows || rows.length === 0) {
      return null;
    }
    return rows.map((row: ReactChild[], rowIndex: number) => {
      return <tr key={`row-${rowIndex}`}>{this.createBodyRow(row, colNames)}</tr>;
    });
  }

  private createBodyRow(row: ReactChild[], colNames: string[] | undefined) {
    return row.map((cell: ReactChild, cellIndex: number) => {
      return (
        <td key={`cell-${cellIndex}`} className={!colNames ? "table-cell" : classNames("table-cell", colNames[cellIndex])}>
          {cell}
        </td>
      );
    });
  }

  render() {
    const { headers, rows, colNames } = this.props;

    return (
      <div className="table-container">
        <table className="table">
          <thead>{this.createHeaders(headers, colNames)}</thead>
          <tbody>{this.createBody(rows, colNames)}</tbody>
        </table>
      </div>
    );
  }
}
