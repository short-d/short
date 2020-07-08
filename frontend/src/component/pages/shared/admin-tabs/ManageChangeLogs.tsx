import React, { Component } from 'react';
import { Section } from '../../../ui/Section';
import { Button } from '../../../ui/Button';
import { Store } from 'redux';
import { IAppState } from '../../../../state/reducers';
import { ChangeLogService } from '../../../../service/ChangeLog.service';
import { Modal } from '../../../ui/Modal';
import './ManageChangeLogs.scss';
import { CreateChangeSection } from '../CreateChangeSection';

interface IProps {
  changeLogService: ChangeLogService;
  notifyToast: (message: string, duration?: number) => void;
  store: Store<IAppState>;
  onAuthenticationFailed: () => void;
}

export class ManageChangeLogs extends Component<IProps> {
  private createModalRef = React.createRef<Modal>();

  constructor(props: IProps) {
    super(props);
  }

  render() {
    return (
      <div className="manage-change-logs">
        <Section
          title={'Manage Change Logs'}
          options={[this.renderCreateButton()]}
        />
        <Modal ref={this.createModalRef} canClose={true}>
          <CreateChangeSection
            onChangeCreated={this.handleChangeCreated}
            changeLogService={this.props.changeLogService}
            onAuthenticationFailed={this.props.onAuthenticationFailed}
            store={this.props.store}
          />
        </Modal>
      </div>
    );
  }

  private renderCreateButton() {
    return (
      <div className="create-change-button">
        <Button onClick={this.handleCreateChangeClick}>+ Change</Button>
      </div>
    );
  }

  private handleChangeCreated = () => {
    this.createModalRef.current!.close();
    this.refreshChanges();

    const changeCreatedMessage = 'Change log created successfully';
    this.props.notifyToast(changeCreatedMessage);
  };

  private handleCreateChangeClick = () => {
    this.createModalRef.current!.open();
  };

  private refreshChanges = () => {
    // call to refresh the changes list while displaying the changes
  };
}
