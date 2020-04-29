import React, { Component, ReactChild } from 'react';

import './Table.scss';
import classNames from 'classnames';

interface IProps {
  headers?: ReactChild[];
  rows?: ReactChild[][];
  colClassNames?: string[];
}

export class Table extends Component<IProps> {
  private createHeaders(headers: ReactChild[] | undefined, colClassNames: string[] | undefined) {
    if (!headers || headers.length === 0) {
      return null;
    }
    return (
      <tr key={`header`}>
        {headers.map((cell: ReactChild, cellIndex: number) => {
          return (
            <th key={`cell-${cellIndex}`} className={colClassNames?.[cellIndex]}>
              {cell}
            </th>
          );
        })}
      </tr>
    );
  }

  private createBody(rows: ReactChild[][] | undefined, colClassNames: string[] | undefined) {
    if (!rows || rows.length === 0) {
      return null;
    }
    return rows.map((row: ReactChild[], rowIndex: number) => {
      return <tr key={`row-${rowIndex}`}>{this.createBodyRow(row, colClassNames)}</tr>;
    });
  }

  private createBodyRow(row: ReactChild[], colClassNames: string[] | undefined) {
    return row.map((cell: ReactChild, cellIndex: number) => {
      return (
        <td key={`cell-${cellIndex}`} className={colClassNames?.[cellIndex]}>
          {cell}
        </td>
      );
    });
  }

  render() {
    const { headers, rows, colClassNames } = this.props;

    return (
      <div className="table-container">
        <table className="table">
          <thead>{this.createHeaders(headers, colClassNames)}</thead>
          <tbody>{this.createBody(rows, colClassNames)}</tbody>
        </table>
      </div>
    );
  }
}
