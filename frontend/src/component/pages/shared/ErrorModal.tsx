import React, { Component } from 'react';
import './ErrorModal.scss';

import { IErr } from '../../../entity/Err';
import { Icon, IconID } from '../../ui/Icon';
import { Modal } from '../../ui/Modal';
import { Store, Unsubscribe } from 'redux';
import { IAppState } from '../../../state/reducers';
import { clearError } from '../../../state/actions';

interface IProps {
  store: Store<IAppState>;
}

interface IState {
  error?: IErr;
}

export class ErrorModal extends Component<IProps, IState> {
  private modalRef = React.createRef<Modal>();
  private unsubscribeStateChange: Unsubscribe | undefined;

  constructor(props: IProps) {
    super(props);

    this.state = {
      error: props.store.getState().err
    };
  }

  componentDidMount(): void {
    this.subscribeErrorChange();
  }

  private subscribeErrorChange() {
    this.unsubscribeStateChange = this.props.store.subscribe(() => {
      const { err } = this.props.store.getState();
      this.setState({ error: err });
      if (err) {
        this.showError();
      }
    });
  }

  componentWillUnmount(): void {
    if (this.unsubscribeStateChange) {
      this.unsubscribeStateChange();
    }
  }

  render() {
    const { error } = this.state;

    return (
      <Modal
        canClose={true}
        onModalClose={this.handleModalClose}
        ref={this.modalRef}
      >
        <div className={'error-modal-content'}>
          <div className={'close-icon'}>
            <Icon
              defaultIconID={IconID.Close}
              onClick={this.handleCloseClick}
            />
          </div>
          <div className={'title'}>{error?.name}</div>
          <div className={'description'}>{error?.description}</div>
        </div>
      </Modal>
    );
  }

  private handleModalClose = () => {
    this.props.store.dispatch(clearError());
  };

  private handleCloseClick = () => {
    this.modalRef.current!.close();
  };

  private showError = () => {
    this.modalRef.current!.open();
  };
}
