import React, {Component} from 'react';

import './Popup.scss';
import classNames from 'classnames';

interface Props {
    shown: boolean
}

export class Popup extends Component<Props> {

    render() {
        return (
            <div className={classNames('popup', {'shown': this.props.shown})}>
                <div className={'card'}>{this.props.children}</div>
                <div className={'mask'}></div>
            </div>
        );
    }
}