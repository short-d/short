import React, { Component } from 'react';

import { Section } from '../../ui/Section';
import { PageControl } from '../../ui/PageControl';
import { Table } from '../../ui/Table';
import { Url } from '../../../entity/Url';
import { IQueryUrlData } from '../../../service/ShortLink.service';

interface IProps {
  urlData?: IQueryUrlData;
  updateUrlData: (offset: number, limit: number) => void;
}

export class UserShortLinksSection extends Component<IProps> {
  private PAGE_SIZE = 10;
  private TABLE_HEADERS = ['Long URL', 'Alias'];

  componentDidMount(): void {
    this.displayPage(0);
  }

  render() {
    if (!this.props.urlData) {
      return false;
    }

    if (this.props.urlData.total <= 0) {
      return false;
    }

    return (
      <div>
        <Section title={'Created Short Links'}>
          <Table
            headers={this.TABLE_HEADERS}
            rows={this.createTableRows(this.props.urlData.urls)}
          />
          <PageControl
            totalPages={this.calculateTotalPages(this.props.urlData.total)}
            onPageChanged={this.onPageChanged}
          />
        </Section>
      </div>
    );
  }

  private createTableRows = (urls: Url[]) => {
    return urls.map((url: Url) => {
      return [url.originalUrl, url.alias];
    });
  };

  public onPageChanged = (currentPageIdx: number) => {
    this.displayPage(currentPageIdx);
  };

  private displayPage = (pageIdx: number) => {
    const offset = pageIdx * this.PAGE_SIZE;
    const limit = this.PAGE_SIZE;

    this.props.updateUrlData(offset, limit);
  };

  private calculateTotalPages = (totalUrls: number) => {
    return Math.ceil(totalUrls / this.PAGE_SIZE);
  };
}
