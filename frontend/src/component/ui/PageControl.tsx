import React, { Component } from 'react';

import './PageControl.scss';
import classNames from 'classnames';

interface IProps {
  totalPages: number;
  onPageChanged?: (currentPageIdx: number) => void;
}

interface IStates {
  currentPageIdx: number;
}

export class PageControl extends Component<IProps, IStates> {
  constructor(props: IProps) {
    super(props);

    this.state = {
      currentPageIdx: 0
    };
  }

  componentDidMount(): void {
    this.showPage(0);
  }

  render() {
    const { totalPages } = this.props;

    if (!totalPages) {
      return false;
    }

    if (totalPages < 1) {
      return false;
    }

    return (
      <div className="page-control">
        {this.renderPreviousButton()}
        <div className={'page-numbers'}>{this.renderPageNumberButtons()}</div>
        {this.renderNextButton()}
      </div>
    );
  }

  private renderPreviousButton = () => {
    const { currentPageIdx } = this.state;

    const hasItemBefore = currentPageIdx > 0;

    return (
      <button
        key={`previous`}
        onClick={this.handlePreviousButtonClick}
        disabled={!hasItemBefore}
        className={'nav previous'}
      >
        &lt; Previous
      </button>
    );
  };

  private renderNextButton = () => {
    const { currentPageIdx } = this.state;
    const { totalPages } = this.props;

    const hasPageAfter = currentPageIdx < totalPages - 1;

    return (
      <button
        key={`next`}
        onClick={this.handleNextButtonClick}
        disabled={!hasPageAfter}
        className={'nav next'}
      >
        Next &gt;
      </button>
    );
  };

  private renderPageNumberButtons = () => {
    const { totalPages } = this.props;
    let pageNumberButtons = [];

    for (let idx = 0; idx < totalPages; idx++) {
      let button = this.renderPageNumberButton(idx);
      pageNumberButtons.push(button);
    }
    return pageNumberButtons;
  };

  private renderPageNumberButton = (pageIdx: number) => {
    const { currentPageIdx } = this.state;

    const isHighLighted = pageIdx === currentPageIdx;

    return (
      <button
        key={pageIdx}
        className={classNames({
          'page-number': true,
          active: isHighLighted
        })}
        onClick={this.handlePageNumberButtonClick(pageIdx)}
      >
        {pageIdx + 1}
      </button>
    );
  };

  private handlePreviousButtonClick = () => {
    this.showPage(this.state.currentPageIdx - 1);
  };

  private handleNextButtonClick = () => {
    this.showPage(this.state.currentPageIdx + 1);
  };

  private handlePageNumberButtonClick = (pageIdx: number) => () => {
    this.showPage(pageIdx);
  };

  private showPage(pageIdx: number) {
    const { onPageChanged } = this.props;
    this.setState({ currentPageIdx: pageIdx });
    if (!onPageChanged) {
      return;
    }
    onPageChanged(pageIdx);
  }
}
