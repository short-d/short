import React, { Component, ChangeEvent } from 'react';

import './SearchBar.scss';
import { Url } from '../../entity/Url';
import { DebounceInput } from 'react-debounce-input';

interface Props {
  onChange: (arg0: String) => void;
  autoCompleteSuggestions?: Array<Url>;
}

export class SearchBar extends Component<Props> {
  handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    this.props.onChange(event.target.value);
  };

  render() {
    return (
      <div className="search-box">
        <DebounceInput
          minLength={2}
          placeholder={'Search short links'}
          debounceTimeout={300}
          onChange={this.handleChange} />
        <img className="image" src={"https://images-na.ssl-images-amazon.com/images/I/41gYkruZM2L.png"} alt="Magnifying Glass" />
        <ul className="suggestions">
          {this.props.autoCompleteSuggestions &&
            this.props.autoCompleteSuggestions.map(e => (
              <li key={e.alias}>
                <a href={e.originalUrl}>{e.alias}</a>
              </li>
            ))}
        </ul>
      </div>
    );
  }
}
