import React, { Component } from 'react';

import './Section.scss';

interface Props {
  title: string;
  options?: React.ReactElement[];
}

export class Section extends Component<Props> {
  render() {
    const { options } = this.props;
    return (
      <div className={'section'}>
        <div className={'center'}>
          <div className={'header'}>
            <div className={'title'}>{this.props.title}</div>
            {options && (
              <ul className={'options'}>{this.renderOptions(options)}</ul>
            )}
          </div>
          {this.props.children}
        </div>
      </div>
    );
  }

  private renderOptions(options: React.ReactElement[]) {
    return options.map((option, index) => <li key={index}>{option}</li>);
  }
}
