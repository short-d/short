import React, { Component } from 'react';
import moment from 'moment';

import './ChangeLogModal.scss';
import { Change } from '../../entity/Change';
import { Button } from './Button';
import { Modal } from './Modal';
import { Icon, IconID } from './Icon';

interface State {
  shouldShowFullChangeLog: boolean;
}

interface Props {
  changeLog?: Array<Change>;
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
          {changeLog.map((change: Change) => (
            <li key={change.id}>
              <div className={'title'}>{change.title}</div>
              <div className={'summary'}>{change.summaryMarkdown}</div>
              <div className={'released-date'}>
                {moment(change.releasedAt).fromNow()}
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
              <Icon defaultIconID={IconID.Close} onClick={this.close} />
            </span>
          </div>
          {this.createChangeLog()}
          {this.createShowCompleteChangeLogButton()}
        </div>
      </Modal>
    );
  }
}
