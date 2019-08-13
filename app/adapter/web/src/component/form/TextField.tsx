import React, {ChangeEvent, Component} from 'react';
import './TextField.scss';

interface Props {
    text?: string
    placeHolder?: string
    onChange?: (text: string) => void
    onBlur?: () => void
}


interface State {
    value: string
}

export class TextField extends Component<Props, State> {
    handleChange = (event: ChangeEvent<HTMLInputElement>) => {
        if (!this.props.onChange) {
            return;
        }
        this.props.onChange(event.target.value);
    };

    handleBlur = () => {
        if (!this.props.onBlur) {
            return;
        }
        this.props.onBlur();
    };

    render = () => {
        return (
            <input
                className={'text-field'}
                type={'text'}
                value={this.props.text}
                onChange={this.handleChange}
                onBlur={this.handleBlur}
                placeholder={this.props.placeHolder}/>
        );
    };
}