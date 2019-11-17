import React, { Component } from 'react';

import './Toggle.scss';
import classNames from 'classnames';

interface Props {
    onClick?: (enabled: boolean) => void;
}

interface State {
    enabled: boolean,
    toggleClassName: string
}

export class Toggle extends Component<Props, State> {
    
    constructor(props: Props) {
        super(props);
        this.state = {
            enabled: false,
            toggleClassName: classNames("toggle")
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
                    toggleClassName: classNames("toggle", "active")
                });
            } else {
                this.setState({
                    toggleClassName: classNames("toggle")
                });
            }
        });
    }

    render() {
        return <div className={this.state.toggleClassName} onClick={this.handleClick}>
        </div>
    }
}