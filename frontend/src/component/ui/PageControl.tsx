import React, { Component, ReactText } from 'react';

import './PageControl.scss';

interface IProps {
  pageSize: number;
  totalItems: number;
  onPageChanged: (currentPage: number, pageLimit: number) => void;
}

interface IStates {
  currentPage: number;
}

const DEFAULT_PAGE_SIZE: number = 10;
const ELLIPSES = 'ELLIPSES';

export class PageControl extends Component<IProps, IStates> {
  public static defaultProps: Partial<IProps> = {
    pageSize: DEFAULT_PAGE_SIZE,
    totalItems: 0,
    onPageChanged: () => {}
  };
  private PAGE_CONTROL_ITEMS_COUNT = 5;

  constructor(props: IProps) {
    super(props);

    this.state = {
      currentPage: 1
    };
  }

  componentDidMount(): void {
    this.gotoPage(1);
  }

  render() {
    const { totalItems } = this.props;
    if (!totalItems) return null;

    return <div className="page-control">{this.renderPageControlItems()}</div>;
  }

  private renderPageControlItems = () => {
    const pageContents = this.fetchPageContents();
    return [
      this.createPreviousPageNavComponent(),
      pageContents.map(this.createPageItem),
      this.createNextPageNavComponent()
    ];
  };

  private createPageItem = (page: ReactText, index: number) => {
    const attrs = { key: index };
    if (page === ELLIPSES) {
      return this.createEllipsesComponent(attrs);
    }

    return this.createPageNumberComponent(page, attrs);
  };

  private createPreviousPageNavComponent = () => {
    const { currentPage } = this.state;

    return (
      <button
        key={`previous`}
        onClick={this.handlePreviousPageNavClick}
        disabled={currentPage === 1}
      >
        &lt; Previous
      </button>
    );
  };

  private handlePreviousPageNavClick = () => {
    this.gotoPage(this.state.currentPage - 1);
  };

  private createNextPageNavComponent = () => {
    const { currentPage } = this.state;
    const lastPage = this.calculateTotalPages();

    return (
      <button
        key={`next`}
        onClick={this.handleNextPageNavClick}
        disabled={currentPage === lastPage}
      >
        Next &gt;
      </button>
    );
  };

  private handleNextPageNavClick = () => {
    this.gotoPage(this.state.currentPage + 1);
  };

  private createEllipsesComponent = (attrs: any) => {
    const { key } = attrs;
    return (
      <button key={key} disabled={true}>
        &hellip;
      </button>
    );
  };

  private createPageNumberComponent = (page: ReactText, attrs: any) => {
    const { currentPage } = this.state;
    const { key } = attrs;

    return (
      <button
        key={key}
        className={`${currentPage === page ? 'active' : ''}`}
        onClick={this.handlePageClick(page)}
      >
        {page}
      </button>
    );
  };

  private handlePageClick = (page: ReactText) => () => {
    if (typeof page === 'number') {
      this.gotoPage(page);
    }
  };

  private gotoPage = (page: number) => {
    const { pageSize, onPageChanged } = this.props;
    const totalPages = this.calculateTotalPages();
    const currentPage = Math.max(1, Math.min(page, totalPages));

    this.setState({ currentPage }, () => onPageChanged(currentPage, pageSize));
  };

  private fetchPageContents = () => {
    const lastPage = this.calculateTotalPages();
    const { currentPage } = this.state;

    // show all page numbers when number of pages are less than minimum pages
    // required to construct ellipses based page controls
    if (!this.hasEnoughPages()) {
      return range(1, lastPage);
    }

    const hasLeftHiddenPages = currentPage > 2;
    const hasRightHiddenPages = lastPage - currentPage > 1;

    // case: 1 2 3 .. 10
    if (!hasLeftHiddenPages && hasRightHiddenPages) {
      return [1, 2, 3, ELLIPSES, lastPage];
    }

    // case: 1 .. 8 9 10
    if (hasLeftHiddenPages && !hasRightHiddenPages) {
      return [1, ELLIPSES, lastPage - 2, lastPage - 1, lastPage];
    }

    // case: 1 .. 5 .. 10
    return [1, ELLIPSES, currentPage, ELLIPSES, lastPage];
  };

  private hasEnoughPages = () => {
    const totalPages = this.calculateTotalPages();
    return totalPages > this.PAGE_CONTROL_ITEMS_COUNT;
  };

  private calculateTotalPages = () => {
    const { pageSize, totalItems } = this.props;
    return Math.ceil(totalItems / pageSize);
  };
}

const range = (from: number, to: number) => {
  const result = [];
  for (let i = from; i <= to; i++) {
    result.push(i);
  }
  return result;
};
