import React, {Component} from 'react';
import {Notice} from '../../../ui/Notice';

import './ExtPromo.scss';

export class ExtPromo extends Component {
  render() {
    return (
      <Notice>
        <div className={'ext-promo'}>
          <a
            target={'_blank'}
            title={'Get Chrome Extension'}
            href={'https://s.time4hacks.com/r/ext'}
          >
            Get Chrome Extension
          </a>
        </div>
      </Notice>
    );
  }
}
