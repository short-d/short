import React, {Component} from 'react';

import './Footer.scss';

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
                        Made with
                        <i className={'heart'}>
                            <div/>
                        </i> by&nbsp;<a
                        href={this.props.authorPortfolio}>{this.props.authorName}</a>
                    </div>
                    <div className={'row app-version'}>App version: {this.props.version}</div>
                </div>
            </footer>
        );
    }
}