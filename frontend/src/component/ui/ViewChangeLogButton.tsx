import React, { Component, Fragment } from 'react';
import { ChangeLogModal } from './ChangeLogModal';
import { Update } from '../../entity/Update';

interface Props {
  onClick?: () => void;
  changeLog?: Array<Update>;
  closeModal: () => void;
  shouldShowModal: boolean;
  defaultVisibleLogs: number;
}

export class ViewChangeLogButton extends Component<Props> {
  render() {
    return (
      <Fragment>
        <div className={'row view-changelog'} onClick={this.props.onClick}>
          <a href={'/#'}>View Changelog</a>
        </div>
        <ChangeLogModal {...this.props} />
      </Fragment>
    );
  }
}
