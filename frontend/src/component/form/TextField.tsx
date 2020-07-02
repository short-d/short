import React, { ChangeEvent, Component, createRef } from 'react';
import './TextField.scss';

interface IProps {
  defaultText?: string;
  placeHolder?: string;
  onChange?: (text: string) => void;
  onBlur?: () => void;
}

interface IState {
  text: string;
}

export class TextField extends Component<IProps, IState> {
  textInput = createRef<HTMLInputElement>();

  constructor(props: IProps) {
    super(props);
    this.state = {
      text: props.defaultText || ''
    };
  }

  handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const text = event.target.value;
    this.setState({ text: text });
    if (!this.props.onChange) {
      return;
    }
    this.props.onChange(text);
  };

  updateValue(text: string) {
    this.setState({
      text: text
    });
  }

  handleBlur = () => {
    if (!this.props.onBlur) {
      return;
    }
    this.props.onBlur();
  };

  focus = () => {
    this.textInput.current!.focus();
  };

  render = () => {
    return (
      <input
        ref={this.textInput}
        className={'text-field'}
        type={'text'}
        value={this.state.text}
        onChange={this.handleChange}
        onBlur={this.handleBlur}
        placeholder={this.props.placeHolder}
      />
    );
  };
}
