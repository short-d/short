import { Component } from 'react';
import { Styling } from './styling';
interface Props extends Styling {
    onClick?: () => void;
}
export declare class Button extends Component<Props> {
    static defaultProps: Props;
    handleClick: () => void;
    render(): JSX.Element;
}
export {};
