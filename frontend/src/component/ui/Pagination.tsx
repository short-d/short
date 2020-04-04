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
  RIGHT = 'RIGHT'
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

  constructor(props: IProps) {
    super(props);

    this.state = {
      currentPage: 1
    };
  }

  private fetchBlocks(currentPage: number, totalPages: number) {
    const blocksCount = 5;
    if (totalPages <= blocksCount) {
      return range(1, totalPages);
    }

    return [1, ...this.fetchMiddleBlocks(currentPage, totalPages), totalPages];
  }

  private fetchMiddleBlocks(currentPage: number, lastPage: number) {
    // has hidden pages to the left
    const hasLeftSpill = currentPage > 2;
    // has hidden pages to the right
    const hasRightSpill = lastPage - currentPage > 1;

    let middleBlocks: ReactText[];
    switch (true) {
      // case: 1 < 8 9 10
      case hasLeftSpill && !hasRightSpill: {
        middleBlocks = [PageNavigator.LEFT, lastPage - 2, lastPage - 1];
        break;
      }

      // case: 1 2 3 > 10
      case !hasLeftSpill && hasRightSpill: {
        middleBlocks = [2, 3, PageNavigator.RIGHT];
        break;
      }

      // case: 1 < 5 > 10
      case hasLeftSpill && hasRightSpill:
      default: {
        middleBlocks = [PageNavigator.LEFT, currentPage, PageNavigator.RIGHT];
        break;
      }
    }

    return middleBlocks;
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
    const blocks = this.fetchBlocks(currentPage, totalPages);

    return (
      <div className="pagination">
        {blocks.map((block: ReactText, index: number) => {
          if (block === PageNavigator.LEFT)
            return (
              <button key={`${index}`} onClick={this.handleLeftNav}>
                &laquo;
              </button>
            );

          if (block === PageNavigator.RIGHT)
            return (
              <button key={`${index}`} onClick={this.handleRightNav}>
                &raquo;
              </button>
            );

          if (typeof block === 'number')
            return (
              <button
                key={`${index}`}
                className={`${currentPage === block ? 'active' : ''}`}
                onClick={this.handlePageClick(block)}
              >
                {block}
              </button>
            );

          return null;
        })}
      </div>
    );
  }
}
