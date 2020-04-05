import React, { Component, ReactText } from 'react';

import './Pagination.scss';

interface IProps {
  pageLimit?: number;
  totalRecords?: number;
  onPageChanged?: (paginationData: any) => void;
}

interface IStates {
  currentPage: number;
}

enum PageNavigator {
  LEFT = 'LEFT',
  RIGHT = 'RIGHT',
  ELLIPSES = 'ELLIPSES'
}

const range = (from: number, to: number) => {
  const result = [];
  for (let i = from; i <= to; i++) {
    result.push(i);
  }
  return result;
};

export class Pagination extends Component<IProps, IStates> {
  private DEFAULT_PAGE_LIMIT: number = 10;
  private MIN_PAGE_ITEMS_COUNT = 5;

  constructor(props: IProps) {
    super(props);

    this.state = {
      currentPage: 1
    };
  }

  private hasEnoughPages(totalPages: number) {
    return totalPages > this.MIN_PAGE_ITEMS_COUNT;
  }

  private createPaginationItems(currentPage: number, totalPages: number) {
    let pages: ReactText[];
    if (this.hasEnoughPages(totalPages)) {
      pages = this.createPageItems(currentPage, totalPages);
    } else {
      pages = range(1, totalPages);
    }

    return [PageNavigator.LEFT, ...pages, PageNavigator.RIGHT];
  }

  private createPageItems(currentPage: number, lastPage: number) {
    // has hidden pages to the left
    const hasLeftSpill = currentPage > 2;
    // has hidden pages to the right
    const hasRightSpill = lastPage - currentPage > 1;

    let middlePageItems: ReactText[];
    switch (true) {
      // case: 1 [.. 8 9] 10
      case hasLeftSpill && !hasRightSpill: {
        middlePageItems = [PageNavigator.ELLIPSES, lastPage - 2, lastPage - 1];
        break;
      }

      // case: 1 [2 3 ..] 10
      case !hasLeftSpill && hasRightSpill: {
        middlePageItems = [2, 3, PageNavigator.ELLIPSES];
        break;
      }

      // case: 1 [.. 5 ..] 10
      case hasLeftSpill && hasRightSpill:
      default: {
        middlePageItems = [
          PageNavigator.ELLIPSES,
          currentPage,
          PageNavigator.ELLIPSES
        ];
        break;
      }
    }

    return [1, ...middlePageItems, lastPage];
  }

  private gotoPage = (page: number) => {
    const {
      pageLimit = this.DEFAULT_PAGE_LIMIT,
      totalRecords = 0,
      onPageChanged = () => {}
    } = this.props;

    const totalPages = Math.ceil(totalRecords / pageLimit);
    const currentPage = Math.max(1, Math.min(page, totalPages));

    const paginationData = { currentPage, pageLimit };
    this.setState({ currentPage }, () => onPageChanged(paginationData));
  };

  private handlePageClick = (page: number) => () => {
    this.gotoPage(page);
  };

  private handleLeftNav = () => {
    this.gotoPage(this.state.currentPage - 1);
  };

  private handleRightNav = () => {
    this.gotoPage(this.state.currentPage + 1);
  };

  render() {
    const {
      pageLimit = this.DEFAULT_PAGE_LIMIT,
      totalRecords = 0
    } = this.props;

    if (!totalRecords) return null;

    const { currentPage } = this.state;
    const totalPages = Math.ceil(totalRecords / pageLimit);
    const pages = this.createPaginationItems(currentPage, totalPages);

    return (
      <div className="pagination">
        {pages.map((page: ReactText, index: number) => {
          if (page === PageNavigator.LEFT)
            return (
              <button
                key={`${index}`}
                onClick={this.handleLeftNav}
                disabled={currentPage === 1}
              >
                &lt; Previous
              </button>
            );

          if (page === PageNavigator.RIGHT)
            return (
              <button
                key={`${index}`}
                onClick={this.handleRightNav}
                disabled={currentPage === totalPages}
              >
                Next &gt;
              </button>
            );

          if (page === PageNavigator.ELLIPSES)
            return (
              <button key={`${index}`} disabled={true}>
                &hellip;
              </button>
            );

          if (typeof page === 'number')
            return (
              <button
                key={`${index}`}
                className={`${currentPage === page ? 'active' : ''}`}
                onClick={this.handlePageClick(page)}
              >
                {page}
              </button>
            );

          return null;
        })}
      </div>
    );
  }
}
