import React, {ChangeEvent, Component} from 'react';
import './TextField.scss';

interface Props {
    text?: string
    placeHolder?: string
    onChange?: (text: string) => void
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

    render = () => {
        return (
            <input className={'text-field'}
                   type={'text'}
                   value={this.props.text}
                   onChange={this.handleChange}
                   placeholder={this.props.placeHolder}/>
        );
    };
}