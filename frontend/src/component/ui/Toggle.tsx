import React, { Component } from 'react';

import './Toggle.scss';
import classNames from 'classnames';

interface Props {
    onClick?: (enabled: boolean) => void;
}

interface State {
    enabled: boolean,
    toggleClassName: string,
    toggleBackClassName: string
}

export class Toggle extends Component<Props, State> {
    
    constructor(props: Props) {
        super(props);
        this.state = {
            enabled: false,
            toggleClassName: classNames('toggle'),
            toggleBackClassName: classNames('toggle-back')
        };
    }

    handleClick = () => {
        this.setState({
            enabled: !this.state.enabled
        }, () => {
            const { enabled } = this.state;
            if (!this.props.onClick) {
                return;
            }
            this.props.onClick(enabled);
            if (enabled) {
                this.setState({
                    toggleClassName: classNames('toggle', 'active'),
                    toggleBackClassName: classNames('toggle-back', 'active')
                });
            } else {
                this.setState({
                    toggleClassName: classNames('toggle'),
                    toggleBackClassName: classNames('toggle-back')
                });
            }
        });
    }

    render() {
        return <div className={'toggle-component'}>
            <p className={'toggle-label'}>{this.props.children}</p>
            <div className={this.state.toggleBackClassName} onClick={this.handleClick}>
                <div className={this.state.toggleClassName}>
                </div>
            </div>
        </div>
    }
}