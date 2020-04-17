import React, { Component } from 'react';

import { Section } from '../../ui/Section';
import { PageControl } from '../../ui/PageControl';
import { Table } from '../../ui/Table';
import { Url } from '../../../entity/Url';
import { IPagedShortLinks } from '../../../service/ShortLink.service';

interface IProps {
  pagedShortLinks?: IPagedShortLinks;
  pageSize: number;
  onPageLoad: (offset: number, pageSize: number) => void;
}

const DEFAULT_PROPS = {
  pageSize: 10
};

export class UserShortLinksSection extends Component<IProps> {
  static defaultProps: Partial<IProps> = DEFAULT_PROPS;

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
      <div>
        <Section title={'Created Short Links'}>
          <Table
            headers={['Long URL', 'Alias']}
            rows={this.createTableRows()}
          />
          <PageControl
            totalPages={this.calculateTotalPages()}
            onPageChanged={this.onPageChanged}
          />
        </Section>
      </div>
    );
  }

  private createTableRows = () => {
    const { shortLinks } = this.props.pagedShortLinks!;
    return shortLinks.map((shortLink: Url) => {
      return [
        this.renderLongLink(shortLink.originalUrl),
        this.renderAlias(shortLink.alias)
      ];
    });
  };

  private renderLongLink = (longLink: string) => {
    return longLink;
  };

  private renderAlias = (alias: string) => {
    return alias;
  };

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
