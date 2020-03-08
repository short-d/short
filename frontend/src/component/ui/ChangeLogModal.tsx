import React, { Component } from 'react';
import moment from 'moment';

import './ChangeLogModal.scss';
import { Update } from '../../entity/Update';
import { Button } from './Button';

interface State {
  shouldShowFullChangeLog: boolean;
}

interface Props {
  changeLog?: Array<Update>;
  closeModal: () => void;
  shouldShowModal: boolean;
  defaultVisibleLogs: number;
}

export class ChangeLogModal extends Component<Props, State> {
  static defaultProps = {
    defaultVisibleLogs: 3
  };

  state = {
    shouldShowFullChangeLog: false
  };

  showFullChangeLog = () => {
    this.setState({
      shouldShowFullChangeLog: true
    });
  };

  createChangeLog = () => {
    let changeLog = this.props.changeLog;
    if (!this.state.shouldShowFullChangeLog) {
      changeLog = changeLog!.slice(0, this.props.defaultVisibleLogs);
    }
    if (changeLog) {
      return (
        <div className={'changelog'}>
          <ul>
            {changeLog.map((update: Update) => (
              <li key={update.releasedAt}>
                <div className={'title'}>{update.title}</div>
                <div className={'summary'}>{update.summary}</div>
                <div className={'released-date'}>
                  {moment(update.releasedAt).fromNow()}
                </div>
              </li>
            ))}
          </ul>
        </div>
      );
    }

    return <div />;
  };

  createShowCompleteChangeLogButton = () => {
    if (this.state.shouldShowFullChangeLog) {
      return <div />;
    }
    return (
      <div className={'view-all-updates'}>
        <Button onClick={this.showFullChangeLog}>View All Updates</Button>
      </div>
    );
  };

  render() {
    if (!this.props.shouldShowModal) {
      return <div />;
    }
    return (
      <div className={'modal-wrapper'}>
        <div className={'modal-body'}>
          <div className={'modal-header'}>
            Since You've Been Gone
            <i
              className={'material-icons clear'}
              onClick={this.props.closeModal}
            >
              clear
            </i>
          </div>
          {this.createChangeLog()}
          {this.createShowCompleteChangeLogButton()}
        </div>
        <div className={'modal-backdrop'} onClick={this.props.closeModal} />
      </div>
    );
  }
}
