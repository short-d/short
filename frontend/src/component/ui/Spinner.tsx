import React, { Component } from 'react';

import './Spinner.scss';

export class Spinner extends Component {
  render() {
    return (
      <div className={'spinner'}>
        <div className={'solar-system'}>
          <div className={'earth-orbit orbit'}>
            <div className={'planet earth'} />
            <div className={'venus-orbit orbit'}>
              <div className={'planet venus'} />
              <div className={'mercury-orbit orbit'}>
                <div className={'planet mercury'} />
                <div className={'sun'} />
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}
