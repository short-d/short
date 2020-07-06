import './MarkdownViewer.scss';
import React, { Component } from 'react';
import marked from 'marked';

interface IProps {
  markdown: string;
}

export class MarkdownViewer extends Component<IProps> {
  render() {
    return (
      <div
        className={'markdown-viewer'}
        dangerouslySetInnerHTML={{ __html: marked(this.props.markdown) }}
      />
    );
  }
}
