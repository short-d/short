import React, { Component, ReactChild } from 'react';

import './Table.scss';
import { TextAlignProperty } from 'csstype';

interface IProps {
  headers?: ReactChild[];
  rows?: ReactChild[][];
  widths?: string[];
  alignHeaders?: TextAlignProperty[];
  alignBody?: TextAlignProperty[];
  headerFontSize?: string;
}

export class Table extends Component<IProps> {
  private createHeaders() {
    const { headers, widths, alignHeaders, headerFontSize } = this.props;
    if (!headers || headers.length === 0) {
      return null;
    }
    return (
      <tr key={`header`}>
        {headers.map((cell: ReactChild, cellIndex: number) => {
          return (
            <th
              key={`cell-${cellIndex}`}
              style={{
                width: widths?.[cellIndex],
                textAlign: alignHeaders?.[cellIndex],
                fontSize: headerFontSize
              }}
            >
              {cell}
            </th>
          );
        })}
      </tr>
    );
  }

  private createBody() {
    const { rows } = this.props;
    if (!rows || rows.length === 0) {
      return null;
    }
    return rows.map((row: ReactChild[], rowIndex: number) => {
      return <tr key={`row-${rowIndex}`}>{this.createBodyRow(row)}</tr>;
    });
  }

  private createBodyRow(row: ReactChild[]) {
    const { widths, alignBody } = this.props;
    return row.map((cell: ReactChild, cellIndex: number) => {
      return (
        <td
          key={`cell-${cellIndex}`}
          style={{
            width: widths?.[cellIndex],
            textAlign: alignBody?.[cellIndex]
          }}
        >
          {cell}
        </td>
      );
    });
  }

  render() {
    return (
      <div className="table-container">
        <table className="table">
          <thead>{this.createHeaders()}</thead>
          <tbody>{this.createBody()}</tbody>
        </table>
      </div>
    );
  }
}
