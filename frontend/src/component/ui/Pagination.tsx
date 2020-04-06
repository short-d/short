import React, { Component, ReactText } from 'react';

import './Pagination.scss';

interface IProps {
  pageLimit: number;
  totalRecords: number;
  onPageChanged: (currentPage: number, pageLimit: number) => void;
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

const DEFAULT_PAGE_LIMIT: number = 10;

export class Pagination extends Component<IProps, IStates> {
  public static defaultProps: Partial<IProps> = {
    pageLimit: DEFAULT_PAGE_LIMIT,
    totalRecords: 0,
    onPageChanged: () => {}
  };
  private MIN_PAGE_ITEMS_COUNT = 5;

  constructor(props: IProps) {
    super(props);

    this.state = {
      currentPage: 1
    };
  }

  componentDidMount(): void {
    this.gotoPage(1);
  }

  private hasEnoughPages = (totalPages: number) => {
    return totalPages > this.MIN_PAGE_ITEMS_COUNT;
  };

  private calculateTotalPages = () => {
    const { pageLimit, totalRecords } = this.props;
    return Math.ceil(totalRecords / pageLimit);
  };

  private createPaginationItems = () => {
    const { currentPage } = this.state;
    const totalPages = this.calculateTotalPages();

    let pages: ReactText[];
    if (this.hasEnoughPages(totalPages)) {
      pages = this.createPageItems(currentPage, totalPages);
    } else {
      pages = range(1, totalPages);
    }

    return [PageNavigator.LEFT, ...pages, PageNavigator.RIGHT];
  };

  private createPageItems = (currentPage: number, lastPage: number) => {
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
  };

  private gotoPage = (page: number) => {
    const { pageLimit, onPageChanged } = this.props;
    const totalPages = this.calculateTotalPages();
    const currentPage = Math.max(1, Math.min(page, totalPages));

    this.setState({ currentPage }, () => onPageChanged(currentPage, pageLimit));
  };

  private handlePageClick = (page: ReactText) => () => {
    if (typeof page === 'number') {
      this.gotoPage(page);
    }
  };

  private handleLeftNav = () => {
    this.gotoPage(this.state.currentPage - 1);
  };

  private handleRightNav = () => {
    this.gotoPage(this.state.currentPage + 1);
  };

  private renderPaginationItem = (page: ReactText, index: number) => {
    const { currentPage } = this.state;
    const lastPage = this.calculateTotalPages();

    let paginationItem;
    switch (page) {
      case PageNavigator.LEFT: {
        paginationItem = (
          <button
            key={`${index}`}
            onClick={this.handleLeftNav}
            disabled={currentPage === 1}
          >
            &lt; Previous
          </button>
        );
        break;
      }

      case PageNavigator.RIGHT: {
        paginationItem = (
          <button
            key={`${index}`}
            onClick={this.handleRightNav}
            disabled={currentPage === lastPage}
          >
            Next &gt;
          </button>
        );
        break;
      }

      case PageNavigator.ELLIPSES: {
        paginationItem = (
          <button key={`${index}`} disabled={true}>
            &hellip;
          </button>
        );
        break;
      }

      default: {
        paginationItem = (
          <button
            key={`${index}`}
            className={`${currentPage === page ? 'active' : ''}`}
            onClick={this.handlePageClick(page)}
          >
            {page}
          </button>
        );
      }
    }
    return paginationItem;
  };

  render() {
    const { totalRecords } = this.props;
    if (!totalRecords) return null;

    const pages = this.createPaginationItems();
    return (
      <div className="pagination">{pages.map(this.renderPaginationItem)}</div>
    );
  }
}
