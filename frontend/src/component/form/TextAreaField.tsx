import React, { ChangeEvent, Component } from 'react';
import './TextAreaField.scss';

interface IProps {
  value?: string;
  placeholder?: string;
  onChange?: (value: string) => void;
  onBlur?: () => void;
}

export class TextAreaField extends Component<IProps, any> {
  render() {
    return (
      <textarea
        className={'text-area'}
        value={this.props.value}
        placeholder={this.props.placeholder}
        onChange={this.handleChange}
        onBlur={this.props.onBlur}
      />
    );
  }

  private handleChange = (event: ChangeEvent<HTMLTextAreaElement>) => {
    if (!this.props.onChange) {
      return;
    }
    this.props.onChange(event.target.value);
  };
}
