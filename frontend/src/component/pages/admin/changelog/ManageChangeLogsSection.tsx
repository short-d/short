import React, { Component } from 'react';
import { Table } from '../../../ui/Table';
import { Change } from '../../../../entity/Change';
import { Icon, IconID } from '../../../ui/Icon';
import styles from './ManageChangeLogsSection.module.scss';

interface IProps {
  changes: Change[];
  onChangeDelete: (changeId: string) => void;
}

interface IState {
  expandedChangeId: string | undefined;
}

export class ManageChangeLogsSection extends Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);

    this.state = {
      expandedChangeId: undefined
    };
  }

  render() {
    return (
      <div>
        <Table
          headers={['ID', 'Title', 'Released At', 'Actions', '']}
          widths={['10%', '35%', '35%', '15%', '5%']}
          rows={this.renderChangeLogRows()}
          alignHeaders={['center', 'center', 'center', 'center', 'initial']}
          alignBodyColumns={['center', 'center', 'center', 'center', 'center']}
          headerFontSize={'26px'}
        />
      </div>
    );
  }

  private renderChangeLogRows = () => {
    return this.props.changes.map(change => {
      return [
        change.id,
        change.title,
        change.releasedAt.toUTCString(),
        <div className={styles.deleteButton}>
          <Icon
            iconID={IconID.Delete}
            onClick={this.handleOnDeleteChange(change.id)}
          />
        </div>,
        <div className={styles.collapsibleButton}>
          <Icon
            iconID={
              this.isChangeRowExpanded(change.id)
                ? IconID.DownArrow
                : IconID.RightArrow
            }
            onClick={this.handleOnCollapsibleButtonClick(change.id)}
          />
        </div>
      ];
    });
  };

  private handleOnDeleteChange = (changeId: string) => () => {
    this.props.onChangeDelete(changeId);
  };

  private handleOnCollapsibleButtonClick = (changeId: string) => () => {
    if (this.isChangeRowExpanded(changeId)) {
      this.collapseChangeRow();
      return;
    }
    this.expandChangeRow(changeId);
  };

  private isChangeRowExpanded(changeId: string) {
    return this.state.expandedChangeId === changeId;
  }

  private collapseChangeRow() {
    this.setState({ expandedChangeId: undefined });
  }

  private expandChangeRow(changeId: string) {
    this.setState({ expandedChangeId: changeId });
  }
}
