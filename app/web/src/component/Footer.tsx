import React, {Component} from 'react';

import './Footer.scss';
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome';
import {faHeart} from '@fortawesome/free-solid-svg-icons';

interface Props {
    authorName: string
    authorPortfolio: string
}

export class Footer extends Component<Props> {
    render() {
        return (
            <footer>
                <div className={'center'}>
                    Made with&nbsp;<FontAwesomeIcon icon={faHeart}/>&nbsp;by&nbsp;<a
                    href={this.props.authorPortfolio}>{this.props.authorName}</a>
                </div>
            </footer>
        );
    }
}