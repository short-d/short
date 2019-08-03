import React, {Component} from "react";

import './Button.scss';

interface Props {
    onClick?: () => void
}


export class Button extends Component<Props> {
    handleClick = () => {
        if (!this.props.onClick) {
            return;
        }

        this.props.onClick();
    };

    render() {
        return (
            <button onClick={this.handleClick}>{this.props.children}</button>
        );
    }
}