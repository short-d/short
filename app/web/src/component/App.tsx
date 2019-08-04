import React, {Component} from 'react';
import './App.scss';
import {Header} from "./Header";
import {Section} from "./Section";
import {TextField} from "./form/TextField";
import {Button} from "./Button";
import {Url} from "../entity/Url";
import {UrlService} from "../service/url.service";

interface Props {
}

interface State {
    url: Url
}

export class App extends Component<Props, State> {
    urlService = new UrlService();

    constructor(props: Props) {
        super(props);
        this.state = {
            url: {
                originalUrl: "",
                alias: ""
            }
        };
    }

    handlerOriginalUrlChange = (newValue: string) => {
        this.setState({
            url: Object.assign({}, this.state.url, {
                originalUrl: newValue
            })
        });

    };

    handleAliasChange = (newValue: string) => {
        this.setState({
            url: Object.assign({}, this.state.url, {
                alias: newValue
            })
        });
    };

    handleCreateShortLinkClick = () => {
        this.urlService.createShortLink(this.state.url)
    };

    render = () => {
        return (
            <div className='App'>
                <Header/>
                <Section title={'New Short Link'}>
                    <div className={'control create-short-link'}>
                        <div className={'text-field-wrapper'}>
                            <TextField text={this.state.url.originalUrl} placeHolder={'Long Link'}
                                       onChange={this.handlerOriginalUrlChange}/>
                        </div>
                        <div className={'text-field-wrapper'}>
                            <TextField text={this.state.url.alias} placeHolder={'Custom Short Link ( Optional )'}
                                       onChange={this.handleAliasChange}/>
                        </div>
                        <Button onClick={this.handleCreateShortLinkClick}>Create Short Link</Button>
                    </div>
                </Section>
            </div>
        );
    };
}

export default App;
