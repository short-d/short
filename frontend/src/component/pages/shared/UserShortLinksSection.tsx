import React, { Component } from 'react';
import { PageControl } from '../../ui/PageControl';
import { Table } from '../../ui/Table';
import { Url } from '../../../entity/Url';
import { IPagedShortLinks } from '../../../service/ShortLink.service';
import './UserShortLinksSection.scss';
import { Section } from '../../ui/Section';
import { TextField } from '../../form/TextField';
import { Icon, IconID } from '../../ui/Icon';

export interface IUpdatedShortLinks {
  [alias: string]: Partial<Url>;
}

interface IProps {
  pagedShortLinks?: IPagedShortLinks;
  pageSize: number;
  onPageLoad: (offset: number, pageSize: number) => void;
  onShortLinksUpdated?: (updatedShortLinks: IUpdatedShortLinks) => void;
}

interface IState {
  isEditing: boolean;
  updatedShortLinks: IUpdatedShortLinks;
}

const DEFAULT_PROPS = {
  pageSize: 10
};

export class UserShortLinksSection extends Component<IProps, IState> {
  static defaultProps: Partial<IProps> = DEFAULT_PROPS;

  constructor(props: IProps) {
    super(props);
    this.state = {
      isEditing: false,
      updatedShortLinks: {}
    };
  }

  componentDidMount(): void {
    this.showPage(0);
  }

  render() {
    if (!this.props.pagedShortLinks) {
      return false;
    }

    if (this.props.pagedShortLinks.totalCount < 1) {
      return false;
    }

    return (
      <div className={'UserShortLinksSection'}>
        <Section title={'Favorites'} options={this.renderTableOptions()}>
          <Table
            headers={['Long Link', 'Alias']}
            rows={this.createTableRows()}
            widths={['70%', '30%']}
          />
          <div className={'page-control-wrapper'}>
            <PageControl
              totalPages={this.calculateTotalPages()}
              onPageChanged={this.onPageChanged}
            />
          </div>
        </Section>
      </div>
    );
  }

  private renderTableOptions(): React.ReactElement[] {
    const { isEditing } = this.state;
    if (isEditing) {
      return [this.renderSaveOption()];
    }
    return [this.renderEditOption()];
  }

  private renderEditOption(): React.ReactElement {
    return (
      <div key="edit" className={'option edit'} onClick={this.handleEditClick}>
        <Icon defaultIconID={IconID.Edit} />
      </div>
    );
  }

  private renderSaveOption(): React.ReactElement {
    return (
      <div key="done" className="option done" onClick={this.handleSaveClick}>
        <Icon defaultIconID={IconID.Check} />
      </div>
    );
  }

  private handleSaveClick = () => {
    const { updatedShortLinks } = this.state;
    if (Object.keys(updatedShortLinks).length > 0) {
      if (this.props.onShortLinksUpdated) {
        this.props.onShortLinksUpdated(updatedShortLinks);
      }
    }
    this.setState({
      isEditing: false,
      updatedShortLinks: {}
    });
  };

  private handleEditClick = () => {
    this.setState({
      isEditing: !this.state.isEditing
    });
  };

  private createTableRows = () => {
    const { shortLinks } = this.props.pagedShortLinks!;
    return shortLinks.map((shortLink: Url) => {
      return [
        this.renderLongLink(shortLink),
        this.renderAlias(shortLink.alias)
      ];
    });
  };

  private renderLongLink = (shortLink: Url) => {
    const { isEditing } = this.state;
    if (isEditing) {
      return (
        <TextField
          defaultText={shortLink.originalUrl}
          placeHolder={'Long link'}
          onChange={this.handleLongLinkChange(shortLink.alias)}
        />
      );
    }
    return (
      <a
        className={'long-link'}
        href={shortLink.originalUrl}
        target="_blank"
        rel="noopener noreferrer"
      >
        {shortLink.originalUrl}
      </a>
    );
  };

  private handleLongLinkChange = (
    alias: string
  ): ((longLink: string) => void) => {
    return (longLink: string) => {
      this.updateShortLink(alias, {
        originalUrl: longLink
      });
    };
  };

  private renderAlias = (alias: string) => {
    const { isEditing } = this.state;
    if (isEditing) {
      return (
        <TextField
          defaultText={alias}
          placeHolder={'Alias'}
          onChange={this.handleAliasChange(alias)}
        />
      );
    }
    return <span className={'alias'}>{alias}</span>;
  };

  private handleAliasChange = (alias: string): ((newAlias: string) => void) => {
    return (newAlias: string) => {
      this.updateShortLink(alias, {
        alias: newAlias
      });
    };
  };

  private updateShortLink(alias: string, update: Partial<Url>) {
    const shortLink = this.state.updatedShortLinks[alias] || {};
    const updatedShortLink = Object.assign<any, Partial<Url>, Partial<Url>>(
      {},
      shortLink,
      update
    );
    this.setState({
      updatedShortLinks: Object.assign({}, this.state.updatedShortLinks, {
        [alias]: updatedShortLink
      })
    });
  }

  public onPageChanged = (currentPageIdx: number) => {
    this.showPage(currentPageIdx);
  };

  private showPage = (pageIdx: number) => {
    const { pageSize } = this.props;
    const offset = pageIdx * pageSize;

    this.props.onPageLoad(offset, pageSize);
  };

  private calculateTotalPages = () => {
    const totalShortLinksCount = this.props.pagedShortLinks!.totalCount;
    return Math.ceil(totalShortLinksCount / this.props.pageSize);
  };
}
