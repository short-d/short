import React, { Component } from 'react';
import GitHubButton from 'react-github-btn';

import './Footer.scss';

interface Props {
  authorName: string;
  authorPortfolio: string;
  version: string;
}

export class Footer extends Component<Props> {
  render() {
    return (
      <footer>
        <div className={'center'}>
          <div className={'row'}>
            Made with
            <i className={'heart'}>
              <div />
            </i>
            by&nbsp;
            <a href={this.props.authorPortfolio}>{this.props.authorName}</a>
          </div>
          <div className={'row app-version'}>
            App version: {this.props.version}
          </div>
          <div className={'github-button'}>
            <GitHubButton
              href={'https://github.com/byliuyang/short'}
              data-size={'large'}
              data-show-count={true}
              aria-label="Star byliuyang/short on GitHub"
            >
              Star
            </GitHubButton>
          </div>
        </div>
      </footer>
    );
  }
}
