import React, { ChangeEvent, Component } from 'react';
import './TextArea.scss';

interface IProps {
  rows?: number;
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
        rows={this.props.rows}
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
