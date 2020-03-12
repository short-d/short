import React, { Component } from 'react';

import './Footer.scss';
import { Update } from '../../../entity/Update';
import { UIFactory } from '../../UIFactory';

interface Props {
  uiFactory: UIFactory;
  authorName: string;
  authorPortfolio: string;
  version: string;
  changeLog?: Array<Update>;
  shouldShowChangeLogModal?: boolean;
  handleShowChangeLogModal: () => void;
  handleHideChangeLogModal: () => void;
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
            changeLog: this.props.changeLog,
            openModal: this.props.handleShowChangeLogModal,
            closeModal: this.props.handleHideChangeLogModal,
            shouldShowModal: this.props.shouldShowChangeLogModal
          })}
        </div>
      </footer>
    );
  }
}
