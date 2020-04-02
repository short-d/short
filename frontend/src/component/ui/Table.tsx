import React, { Component, ReactChild, ReactText } from 'react';

import './Table.scss';

interface IProps {
  headings?: ReactChild[];
  rows?: ReactChild[][];
}

export class Table extends Component<IProps> {
  private constructHeadIfExists(headings: ReactChild[] | undefined) {
    if (!headings || headings.length === 0) {
      return null;
    }
    return <tr>{this.constructRow(headings, true)}</tr>;
  }

  private constructRow(row: ReactChild[], isHeader: boolean) {
    return row.map((cell: ReactChild, cellIndex: number) => {
      return isHeader
        ? this.constructHeadingCell(cellIndex, cell)
        : this.constructBodyCell(cellIndex, cell);
    });
  }

  private constructHeadingCell(key: ReactText, content: ReactChild) {
    return (
      <th key={`${key}`} className="table-cell">
        {content}
      </th>
    );
  }

  private constructBodyCell(key: ReactText, content: ReactChild) {
    return (
      <td key={`${key}`} className="table-cell">
        {content}
      </td>
    );
  }

  private constructBodyIfExists(rows: ReactChild[][] | undefined) {
    if (!rows || rows.length === 0) {
      return null;
    }
    return rows.map((row: ReactChild[], rowIndex: number) => {
      return <tr key={`row-${rowIndex}`}>{this.constructRow(row, false)}</tr>;
    });
  }

  render() {
    const { headings, rows } = this.props;

    const theadMarkup = this.constructHeadIfExists(headings);
    const tbodyMarkup = this.constructBodyIfExists(rows);

    return (
      <div className="table-container">
        <table className="table">
          <thead>{theadMarkup}</thead>
          <tbody>{tbodyMarkup}</tbody>
        </table>
      </div>
    );
  }
}
