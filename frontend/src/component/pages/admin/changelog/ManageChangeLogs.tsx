import React, { Component } from 'react';
import { Section } from '../../../ui/Section';
import { Button } from '../../../ui/Button';
import { Store } from 'redux';
import { IAppState } from '../../../../state/reducers';
import { ChangeLogService } from '../../../../service/ChangeLog.service';
import { Modal } from '../../../ui/Modal';
import './ManageChangeLogs.scss';
import { CreateChangeSection } from './CreateChangeSection';
import { Icon, IconID } from '../../../ui/Icon';

interface IProps {
  changeLogService: ChangeLogService;
  notifyToast: (message: string, duration?: number) => void;
  store: Store<IAppState>;
  onAuthenticationFailed: () => void;
}

export class ManageChangeLogs extends Component<IProps> {
  private createChangeModalRef = React.createRef<Modal>();

  render() {
    return (
      <div className="manage-change-logs">
        <Section
          title={'Manage Change Logs'}
          options={[this.renderCreateButton()]}
        />
        <Modal ref={this.createChangeModalRef} canClose={true}>
          <div className={'create-change-modal-close-icon'}>
            <Icon
              defaultIconID={IconID.Close}
              onClick={this.handleOnCreateChangeModalCloseClick}
            />
          </div>
          <div className={'create-change-section-container'}>
            <CreateChangeSection
              onChangeCreated={this.handleOnChangeCreated}
              changeLogService={this.props.changeLogService}
              onAuthenticationFailed={this.props.onAuthenticationFailed}
              store={this.props.store}
            />
          </div>
        </Modal>
      </div>
    );
  }

  private renderCreateButton() {
    return (
      <div className="create-change-button">
        <Button
          onClick={this.handleOnCreateChangeClick}
          styles={['black', 'full-width', 'full-height']}
        >
          + Change
        </Button>
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

  private handleOnCreateChangeModalCloseClick = () => {
    this.createChangeModalRef.current!.close();
  };

  private refreshChanges = () => {
    // call to refresh the changes list while displaying the changes
  };
}
