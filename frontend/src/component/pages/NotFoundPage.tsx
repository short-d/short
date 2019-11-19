import React, { Component } from 'react';
import './NotFoundPage.scss';

export class NotFoundPage extends Component {
  render() {
    return (
      <div className={'page-404'}>
        <div className="code">404</div>
        <div className="to-home">
          Take me back to <a href="/">homepage</a>.
        </div>
      </div>
    );
  }
}
