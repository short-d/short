import React, {Component} from 'react';

import './Footer.scss';
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome';
import {faHeart} from '@fortawesome/free-solid-svg-icons';

interface Props {
    authorName: string
    authorPortfolio: string
    version: string
}

export class Footer extends Component<Props> {
    render() {
        return (
            <footer>
                <div className={'center'}>
                    <div className={'row'}>
                        Made with&nbsp;<FontAwesomeIcon className={'heart'} icon={faHeart}/>&nbsp;by&nbsp;<a
                        href={this.props.authorPortfolio}>{this.props.authorName}</a>
                    </div>
                    <div className={'row app-version'}>App version: {this.props.version}</div>
                </div>
            </footer>
        );
    }
}