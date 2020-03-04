import React, { Component, ChangeEvent } from 'react';

import './SearchBar.scss';
import { Url } from '../../entity/Url';

interface Props {
  onChange: (arg0: String) => void;
  autoCompleteSuggestions?: Array<Url>
}

export class SearchBar extends Component<Props> {
  handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    this.props.onChange(event.target.value);
  }
  render() {
    return <div className="search-box">
        <input type={'text'} placeholder={'Search short links'} onChange={this.handleChange} />
        <ul className="suggestions">
            {this.props.autoCompleteSuggestions && this.props.autoCompleteSuggestions.map(e => <li key={e.alias}><a href={e.originalUrl}>{e.alias}</a></li>)}
        </ul>
    </div>;
  }
}
