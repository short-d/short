import React, {Component} from 'react';
import './App.scss';
import {Header} from './Header';
import {Section} from './Section';
import {TextField} from './form/TextField';
import {Button} from './Button';
import {Url} from '../entity/Url';
import {ErrUrl, UrlService} from '../service/Url.service';
import {Footer} from './Footer';
import {QrcodeService} from '../service/Qrcode.service';
import {ShortLinkUsage} from './ShortLinkUsage';
import {VersionService} from '../service/Version.service';
import {Modal} from './ui/Modal';
import {Notice} from "./ui/Notice";

interface Props {
}

interface State {
    editingUrl: Url
    createdUrl?: Url
    qrCodeUrl?: string
    err: Err
}

interface Err {
    name: string,
    description: string
}

function getErr(errCode: ErrUrl): Err {
    switch (errCode) {
        case ErrUrl.AliasAlreadyExist:
            return ({
                name: 'Alias not available',
                description: `
                The alias you choose is not available, please choose a different one. 
                Leaving custom alias field empty will automatically generate a available alias.
                `
            });
        default:
            return ({
                name: 'Unknown error',
                description: `
                I am not aware of this error. 
                Please email byliuyang11@gmail.com the screenshots and detailed steps to reproduce it so that I can investigate.
                `
            });
    }
}

export class App extends Component<Props, State> {
    urlService = new UrlService();
    appVersion = VersionService.getAppVersion();

    errModal = React.createRef<Modal>();

    constructor(props: Props) {
        super(props);
        this.state = {
            editingUrl: {
                originalUrl: '',
                alias: '',
            },
            err: {
                name: '',
                description: ''
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

    handleOnErrModalCloseClick = () => {
        this.errModal.current!.close();
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
            })
            .catch((errCodes: ErrUrl[]) => {
                for (const errCode of errCodes) {
                    this.setState({
                        err: getErr(errCode)
                    });
                    this.errModal.current!.open();
                }
            });
    };

    render = () => {
        return (
            <div className='app'>
                <Notice>
                    <div className={'ext-promo'}>
                        Type less with <a target={'_blank'}
                                          title={'Get s/'}
                                          href={'https://github.com/byliuyang/short-ext'}>s/</a>.
                    </div>
                </Notice>
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
                <Footer
                    authorName={'Harry'}
                    authorPortfolio={'https://github.com/byliuyang'}
                    version={this.appVersion}/>
                <Modal ref={this.errModal}>
                    <div className={'err'}>
                        <i className={'material-icons close'}
                           title={'close'}
                           onClick={this.handleOnErrModalCloseClick}>close</i>
                        <div className={'title'}>
                            {this.state.err.name}
                        </div>
                        <div className={'description'}>
                            {this.state.err.description}
                        </div>
                    </div>
                </Modal>
            </div>
        );
    };
}

export default App;
