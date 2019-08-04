import React, {Component} from 'react';

import './ShortLinkUsage.scss';

interface Props {
    shortLink: string
    originalUrl: string
    qrCodeUrl: string
}

export class ShortLinkUsage extends Component<Props> {
    render() {
        return (
            <div className={'short-link-usage'}>
                <div>
                    You can now paste&nbsp;
                    <a target={'_blank'}
                       href={this.props.shortLink}>
                        {/*{this.urlService.aliasToLink(this.state.createdUrl.alias!)}*/}
                        {this.props.shortLink}
                    </a>
                    &nbsp;in your browser to visit {this.props.originalUrl}.
                </div>
                <div className={'qr-code'}>
                    <img src={this.props.qrCodeUrl}/>
                </div>
            </div>
        );
    }
}