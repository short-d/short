import React, { Component } from 'react';
import { Modal } from '../../../ui/Modal';
import styles from './CreateChangeModal.module.scss';
import { Icon, IconID } from '../../../ui/Icon';
import { CreateChangeSection } from './CreateChangeSection';
import { Store } from 'redux';
import { IAppState } from '../../../../state/reducers';
import { ChangeLogService } from '../../../../service/ChangeLog.service';

interface IProps {
  store: Store<IAppState>;
  changeLogService: ChangeLogService;
  onAuthenticationFailed: () => void;
  handleOnChangeCreated: () => void;
}

export class CreateChangeModal extends Component<IProps> {
  private modalRef = React.createRef<Modal>();

  render() {
    return (
      <Modal ref={this.modalRef} canClose={true}>
        <div className={styles.modalCloseIcon}>
          <Icon
            defaultIconID={IconID.Close}
            onClick={this.handleOnModalCloseClick}
          />
        </div>
        <div className={styles.content}>
          <CreateChangeSection
            onChangeCreated={this.props.handleOnChangeCreated}
            changeLogService={this.props.changeLogService}
            onAuthenticationFailed={this.props.onAuthenticationFailed}
            store={this.props.store}
          />
        </div>
      </Modal>
    );
  }

  private handleOnModalCloseClick = () => {
    this.close();
  };

  open = () => {
    this.modalRef.current!.open();
  };

  close = () => {
    this.modalRef.current!.close();
  };
}
