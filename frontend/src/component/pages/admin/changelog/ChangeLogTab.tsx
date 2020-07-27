import React, { Component } from 'react';
import { Button } from '../../../ui/Button';
import { Store } from 'redux';
import { IAppState } from '../../../../state/reducers';
import { ChangeLogService } from '../../../../service/ChangeLog.service';
import styles from './ChangeLogTab.module.scss';
import { CreateChangeModal } from './CreateChangeModal';
import { Change } from '../../../../entity/Change';
import {
  raiseCreateChangeError,
  raiseGetAllChangesError
} from '../../../../state/actions';
import { ManageChangeLogsSection } from './ManageChangeLogsSection';

interface IProps {
  changeLogService: ChangeLogService;
  notifyToast: (message: string, duration?: number) => void;
  store: Store<IAppState>;
  onAuthenticationFailed: () => void;
}

interface IStates {
  changes: Change[];
}

export class ChangeLogTab extends Component<IProps, IStates> {
  private createChangeModalRef = React.createRef<CreateChangeModal>();

  constructor(props: IProps) {
    super(props);

    this.state = {
      changes: []
    };
  }

  componentDidMount() {
    this.getChanges();
  }

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
        <div className={styles.viewChangeLogsContainer}>
          <ManageChangeLogsSection
            changes={this.state.changes}
            onChangeDelete={this.handleOnChangeDelete}
          />
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

    const changeCreatedMessage = 'Changelog created successfully';
    this.props.notifyToast(changeCreatedMessage);
  };

  private handleOnChangeDelete = (changeId: string) => {
    this.props.changeLogService
      .deleteChange(changeId)
      .then(_ => {
        this.refreshChanges();

        const changeDeletedMessage = 'Changelog deleted successfully';
        this.props.notifyToast(changeDeletedMessage);
      })
      .catch(({ authenticationErr, changeErr }) => {
        if (authenticationErr) {
          this.props.onAuthenticationFailed();
        }
        this.props.store.dispatch(raiseCreateChangeError(changeErr));
      });
  };

  private handleOnCreateChangeClick = () => {
    this.createChangeModalRef.current!.open();
  };

  private refreshChanges = () => {
    this.getChanges();
  };

  private getChanges = () => {
    this.props.changeLogService
      .getAllChanges()
      .then(changes => {
        this.setState({ changes });
      })
      .catch(({ authenticationErr, changeErr }) => {
        if (authenticationErr) {
          this.props.onAuthenticationFailed();
        }
        this.props.store.dispatch(raiseGetAllChangesError(changeErr));
      });
  };
}
