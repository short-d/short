import React, { Component } from 'react';

import './Footer.scss';
import { UIFactory } from '../../UIFactory';

interface Props {
  uiFactory: UIFactory;
  authorName: string;
  authorPortfolio: string;
  version: string;
  handleShowChangeLogModal: () => void;
}

export class Footer extends Component<Props> {
  render() {
    return (
      <footer>
        <div className={'center'}>
          <div className={'row'}>
            Made with
            <i className={'heart'}>
              <div />
            </i>
            by <a href={this.props.authorPortfolio}>{this.props.authorName}</a>
          </div>
          <div className={'row app-version'}>
            App version: {this.props.version}
          </div>
          {this.props.uiFactory.createViewChangeLogButton({
            onClick: this.props.handleShowChangeLogModal
          })}
        </div>
      </footer>
    );
  }
}
