import React, {Component} from 'react';
import './App.scss';
import {Header} from './Header';
import {Section} from './Section';
import {TextField} from './form/TextField';
import {Button} from './Button';
import {Url} from '../entity/Url';
import {UrlService} from '../service/Url.service';
import {Footer} from './Footer';
import {QrcodeService} from '../service/Qrcode.service';
import {ShortLinkUsage} from './ShortLinkUsage';

interface Props {
}

interface State {
    editingUrl: Url
    createdUrl?: Url
    qrCodeUrl?: string
}

export class App extends Component<Props, State> {
    urlService = new UrlService();

    constructor(props: Props) {
        super(props);
        this.state = {
            editingUrl: {
                originalUrl: '',
                alias: ''
            }
        };
    }

    handlerOriginalUrlChange = (newValue: string) => {
        this.setState({
            editingUrl: Object.assign({}, this.state.editingUrl, {
                originalUrl: newValue
            })
        });

    };

    handleAliasChange = (newValue: string) => {
        this.setState({
            editingUrl: Object.assign({}, this.state.editingUrl, {
                alias: newValue
            })
        });
    };

    handleCreateShortLinkClick = () => {
        this.urlService
            .createShortLink(this.state.editingUrl)
            .then((url: Url) => {
                this.setState({
                    createdUrl: url
                });

                if (url.alias) {
                    QrcodeService.newQrCode(this.urlService.aliasToLink(url.alias))
                        .then((qrCodeUrl: string) => {
                            this.setState({
                                qrCodeUrl: qrCodeUrl
                            });
                        });
                }
            });
    };

    render = () => {
        return (
            <div className='app'>
                <Header/>
                <div className={'main'}>
                    <Section title={'New Short Link'}>
                        <div className={'control create-short-link'}>
                            <div className={'text-field-wrapper'}>
                                <TextField text={this.state.editingUrl.originalUrl} placeHolder={'Long Link'}
                                           onChange={this.handlerOriginalUrlChange}/>
                            </div>
                            <div className={'text-field-wrapper'}>
                                <TextField text={this.state.editingUrl.alias}
                                           placeHolder={'Custom Short Link ( Optional )'}
                                           onChange={this.handleAliasChange}/>
                            </div>
                            <Button onClick={this.handleCreateShortLinkClick}>Create Short Link</Button>
                        </div>
                        {this.state.createdUrl && this.state.qrCodeUrl ?
                            <div className={'short-link-usage-wrapper'}>
                                <ShortLinkUsage
                                shortLink={this.urlService.aliasToLink(this.state.createdUrl.alias!)}
                                originalUrl={this.state.createdUrl.originalUrl!}
                                qrCodeUrl={this.state.qrCodeUrl}/>
                            </div>
                            :
                            false
                        }
                    </Section>
                </div>
                <Footer authorName={'Harry'} authorPortfolio={'https://github.com/byliuyang'}/>
            </div>
        );
    };
}

export default App;
