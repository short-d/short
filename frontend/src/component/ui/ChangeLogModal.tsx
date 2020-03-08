import React, { Component } from 'react';
import moment from 'moment';

import './ChangeLogModal.scss';
import { Update } from '../../entity/Update';
import { Button } from './Button';
import { Modal } from './Modal';

interface State {
  shouldShowFullChangeLog: boolean;
}

interface Props {
  changeLog?: Array<Update>;
  closeModal: () => void;
  shouldShowModal: boolean;
  defaultVisibleLogs: number;
}

enum ModalState {
  Open = 'open',
  Close = 'close'
}

export class ChangeLogModal extends Component<Props, State> {
  static defaultProps = {
    defaultVisibleLogs: 3
  };

  state = {
    shouldShowFullChangeLog: false
  };

  private modalRef = React.createRef<Modal>();

  componentDidMount() {
    if (this.props.shouldShowModal) {
      this.open();
    }
  }

  componentDidUpdate(prevProps: Props) {
    if (
      this.props.shouldShowModal !== prevProps.shouldShowModal &&
      this.props.shouldShowModal
    ) {
      this.open();
    }
  }

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

  open = () => this.updateModalState(ModalState.Open);

  close = () => this.updateModalState(ModalState.Close);

  private updateModalState = (state: ModalState) => {
    if (!this.modalRef.current) {
      return;
    }
    this.modalRef.current[state]();
  };

  render() {
    return (
      <Modal
        ref={this.modalRef}
        onClose={this.props.closeModal}
        canClose={true}
      >
        <div className={'modal-body'}>
          <div className={'modal-header'}>
            Since You've Been Gone
            <i className={'material-icons clear'} onClick={this.close}>
              clear
            </i>
          </div>
          {this.createChangeLog()}
          {this.createShowCompleteChangeLogButton()}
        </div>
      </Modal>
    );
  }
}
