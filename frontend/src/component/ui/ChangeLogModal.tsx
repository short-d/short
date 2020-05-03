import React, { Component } from 'react';
import moment from 'moment';

import './ChangeLogModal.scss';
import { Update } from '../../entity/Update';
import { Button } from './Button';
import { Modal } from './Modal';
import { Icon, IconID } from './Icon';

interface State {
  shouldShowFullChangeLog: boolean;
}

interface Props {
  changeLog?: Array<Update>;
  defaultVisibleLogs: number;
}

export class ChangeLogModal extends Component<Props, State> {
  static defaultProps = {
    defaultVisibleLogs: 3
  };

  state = {
    shouldShowFullChangeLog: false
  };

  private modalRef = React.createRef<Modal>();

  showFullChangeLog = () => {
    this.setState({
      shouldShowFullChangeLog: true
    });
  };

  createChangeLog = () => {
    let changeLog = this.props.changeLog;
    if (!changeLog) {
      return <div />;
    }
    if (!this.state.shouldShowFullChangeLog) {
      changeLog = changeLog!.slice(0, this.props.defaultVisibleLogs);
    }
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

  open = () => this.modalRef.current && this.modalRef.current.open();

  close = () => this.modalRef.current && this.modalRef.current.close();

  render() {
    return (
      <Modal ref={this.modalRef} canClose={true}>
        <div className={'modal-body'}>
          <div className={'modal-header'}>
            Since You've Been Gone
            <span className={'close-button'}>
              <Icon iconID={IconID.Close} onClick={this.close} />
            </span>
          </div>
          {this.createChangeLog()}
          {this.createShowCompleteChangeLogButton()}
        </div>
      </Modal>
    );
  }
}
