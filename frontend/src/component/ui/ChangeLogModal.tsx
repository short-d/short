import React, { Component } from 'react';
import moment from 'moment';
import classnames from 'classnames';

import './ChangeLogModal.scss';
import { Update } from '../../entity/Update';
import { Button } from './Button';

interface State {
  shouldShowCompleteChangeLog: boolean;
}

interface Props {
  changeLog?: Array<Update>;
  closeModal: () => void;
  shouldShowModal: boolean;
}

export class ChangeLogModal extends Component<Props, State> {
  state = {
    shouldShowCompleteChangeLog: false
  };

  showCompleteChangeLog = () => {
    this.setState({
      shouldShowCompleteChangeLog: true
    });
  };

  createChangeLog = () => {
    let changeLog = this.props.changeLog;
    if (!this.state.shouldShowCompleteChangeLog) {
      changeLog = changeLog!.slice(0, 3);
    }
    if (changeLog) {
      return (
        <ul
          className={classnames({
            'complete-list': this.state.shouldShowCompleteChangeLog
          })}
        >
          {changeLog.map((update: Update) => (
            <li key={update.publishedAt}>
              <div className={'title'}>{update.title}</div>
              <div>{update.excerpt}</div>
              <br />
              <div className={'published-date'}>
                {moment(update.publishedAt).fromNow()} -{' '}
                {moment(update.publishedAt).format('MMMM Do YYYY, h:mm:ss a')}
              </div>
            </li>
          ))}
        </ul>
      );
    }

    return <div />;
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
          <div className={'view-all-updates'}>
            {!this.state.shouldShowCompleteChangeLog && (
              <Button onClick={this.showCompleteChangeLog}>
                View All Updates
              </Button>
            )}
          </div>
        </div>
        <div className={'modal-backdrop'} onClick={this.props.closeModal} />
      </div>
    );
  }
}
