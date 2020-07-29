import React, { Component } from 'react';
import { Table } from '../../../ui/Table';
import { Change } from '../../../../entity/Change';
import { Icon, IconID } from '../../../ui/Icon';
import styles from './ManageChangeLogsSection.module.scss';
import classNames from 'classnames';

interface IProps {
  changes: Change[];
  onChangeDelete: (changeId: string) => void;
}

interface IState {
  expandedChangeIds: Set<string>;
}

export class ManageChangeLogsSection extends Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);

    this.state = {
      expandedChangeIds: new Set<string>()
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
        <div
          className={classNames(styles.expandableButton, {
            [styles.expand]: this.isChangeRowExpanded(change.id)
          })}
        >
          <Icon
            iconID={IconID.RightArrow}
            onClick={this.handleOnExpandableButtonClick(change.id)}
          />
        </div>
      ];
    });
  };

  private handleOnDeleteChange = (changeId: string) => () => {
    this.props.onChangeDelete(changeId);
  };

  private handleOnExpandableButtonClick = (changeId: string) => () => {
    if (this.isChangeRowExpanded(changeId)) {
      this.collapseChangeRow(changeId);
      return;
    }
    this.expandChangeRow(changeId);
  };

  private isChangeRowExpanded(changeId: string) {
    return this.state.expandedChangeIds.has(changeId);
  }

  private collapseChangeRow(changeId: string) {
    this.setState(({ expandedChangeIds }) => {
      expandedChangeIds.delete(changeId);
      return { expandedChangeIds };
    });
  }

  private expandChangeRow(changeId: string) {
    this.setState(({ expandedChangeIds }) => {
      expandedChangeIds.add(changeId);
      return { expandedChangeIds };
    });
  }
}
