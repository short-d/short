import React, { ChangeEvent, Component, createRef } from 'react';
import './TextField.scss';

interface IProps {
  text?: string;
  placeHolder?: string;
  onChange?: (text: string) => void;
  onBlur?: () => void;
}

export class TextField extends Component<IProps, any> {
  textInput = createRef<HTMLInputElement>();

  render = () => {
    return (
      <input
        ref={this.textInput}
        className={'text-field'}
        type={'text'}
        value={this.props.text}
        onChange={this.handleChange}
        onBlur={this.handleBlur}
        placeholder={this.props.placeHolder}
      />
    );
  };

  handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const text = event.target.value;
    if (!this.props.onChange) {
      return;
    }
    this.props.onChange(text);
  };

  handleBlur = () => {
    if (!this.props.onBlur) {
      return;
    }
    this.props.onBlur();
  };

  focus = () => {
    this.textInput.current!.focus();
  };
}
