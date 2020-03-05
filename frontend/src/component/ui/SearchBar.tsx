import React, { Component, ChangeEvent } from 'react';

import './SearchBar.scss';
import { Url } from '../../entity/Url';
import { DebounceInput } from 'react-debounce-input';

interface Props {
  onChange: (text: String) => void;
  autoCompleteSuggestions?: Array<Url>;
}

export class SearchBar extends Component<Props> {
  handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    this.props.onChange(event.target.value);
  };

  createAutoCompleteBox() {
    if (!this.props.autoCompleteSuggestions) {
      return <div />;
    }

    return this.props.autoCompleteSuggestions.map(e => (
      <li key={e.alias}>
        <a href={e.originalUrl}>{e.alias}</a>
      </li>
    ));
  }

  render() {
    return (
      <div className="search-box">
        <div className="search-input">
          <DebounceInput
            minLength={2}
            maxLength={50}
            placeholder={'Search short links'}
            debounceTimeout={300}
            onChange={this.handleChange}
          />
          <i className="material-icons search">search</i>
        </div>
        <ul className="suggestions">{this.createAutoCompleteBox()}</ul>
      </div>
    );
  }
}
