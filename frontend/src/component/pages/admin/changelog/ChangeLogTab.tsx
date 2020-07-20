import React, { Component } from 'react';
import { Button } from '../../../ui/Button';
import { Store } from 'redux';
import { IAppState } from '../../../../state/reducers';
import { ChangeLogService } from '../../../../service/ChangeLog.service';
import styles from './ChangeLogTab.module.scss';
import { CreateChangeModal } from './CreateChangeModal';

interface IProps {
  changeLogService: ChangeLogService;
  notifyToast: (message: string, duration?: number) => void;
  store: Store<IAppState>;
  onAuthenticationFailed: () => void;
}

export class ChangeLogTab extends Component<IProps> {
  private createChangeModalRef = React.createRef<CreateChangeModal>();

  render() {
    return (
      <div className={`${styles.changelogTab}`}>
        <div className={styles.header}>
          <div className={styles.title}>Changelog</div>
          <div className={styles.createChangeButton}>
            <Button
              onClick={this.handleOnCreateChangeClick}
              styles={['black', 'full-width', 'full-height', 'round-edged']}
            >
              + Change
            </Button>
          </div>
        </div>
        <CreateChangeModal
          ref={this.createChangeModalRef}
          onAuthenticationFailed={this.props.onAuthenticationFailed}
          changeLogService={this.props.changeLogService}
          store={this.props.store}
          handleOnChangeCreated={this.handleOnChangeCreated}
        />
      </div>
    );
  }

  private handleOnChangeCreated = () => {
    this.createChangeModalRef.current!.close();
    this.refreshChanges();

    const changeCreatedMessage = 'Change log created successfully';
    this.props.notifyToast(changeCreatedMessage);
  };

  private handleOnCreateChangeClick = () => {
    this.createChangeModalRef.current!.open();
  };

  private refreshChanges = () => {
    // call to refresh the changes list while displaying the changes
    this.props.changeLogService.getAllChanges().then(console.log);
  };
}
