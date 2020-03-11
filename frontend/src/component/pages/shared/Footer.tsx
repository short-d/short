import React, { Component } from 'react';

import './Footer.scss';
import { Update } from '../../../entity/Update';
import { UIFactory } from '../../UIFactory';

interface State {
  showChangeLogModal: boolean;
}

interface Props {
  uiFactory: UIFactory;
  authorName: string;
  authorPortfolio: string;
  version: string;
  changeLog?: Array<Update>;
  newUpdateReleased?: boolean;
  updateLastSeenChangeLog?: () => void;
}

export class Footer extends Component<Props, State> {
  state = {
    showChangeLogModal: false
  };

  componentDidMount() {
    this.setState({
      showChangeLogModal: this.props.newUpdateReleased || false
    });
  }

  componentDidUpdate(prevProps: Props) {
    if (prevProps.newUpdateReleased !== this.props.newUpdateReleased) {
      this.setState({
        showChangeLogModal: this.props.newUpdateReleased || false
      });
    }
  }

  handleShowChangeLog = () => {
    this.setState({
      showChangeLogModal: true
    });
  };

  handleHideChangeLog = () => {
    if (this.props.updateLastSeenChangeLog) {
      this.props.updateLastSeenChangeLog();
    }
    this.setState({
      showChangeLogModal: false
    });
  };

  render() {
    return (
      <footer>
        <div className={'center'}>
          <div className={'row'}>
            Made with
            <i className={'heart'}>
              <div />
            </i>
            by&nbsp;
            <a href={this.props.authorPortfolio}>{this.props.authorName}</a>
          </div>
          <div className={'row app-version'}>
            App version: {this.props.version}
          </div>
          {this.props.uiFactory.createViewChangeLogButton({
            changeLog: this.props.changeLog,
            openModal: this.handleShowChangeLog,
            closeModal: this.handleHideChangeLog,
            shouldShowModal: this.state.showChangeLogModal
          })}
        </div>
      </footer>
    );
  }
}
