import React, { ChangeEvent, Component } from 'react';
import classNames from 'classnames';

import './SearchBar.scss';
import { ShortLink } from '../../entity/ShortLink';
import { Subject } from 'rxjs';
import { debounceTime } from 'rxjs/operators';
import { Icon, IconID } from './Icon';

interface State {
  showAutoCompleteBox: boolean;
}

interface Props {
  onChange: (text: String) => void;
  autoCompleteSuggestions?: Array<ShortLink>;
}

const DEBOUNCE_DURATION: number = 300;

export class SearchBar extends Component<Props, State> {
  state = {
    showAutoCompleteBox: false
  };

  private onSearch$: any = new Subject();
  private subscription: any = null;

  componentDidMount() {
    this.subscription = this.onSearch$
      .pipe(debounceTime(DEBOUNCE_DURATION))
      .subscribe((debouncedValue: string) => {
        this.props.onChange(debouncedValue);
      });
  }

  componentWillUnmount() {
    if (this.subscription) {
      this.subscription.unsubscribe();
    }
  }

  handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    this.onSearch$.next(event.target.value);
  };

  createAutoCompleteBox() {
    if (!this.props.autoCompleteSuggestions) {
      return <div />;
    }

    return (
      <ul
        className={classNames('suggestions', {
          show: this.state.showAutoCompleteBox
        })}
      >
        {this.props.autoCompleteSuggestions.map(e => (
          <li key={e.alias}>
            <a href={e.longLink}>{e.alias}</a>
          </li>
        ))}
      </ul>
    );
  }

  hideAutoCompleteBox = () => {
    setTimeout(() => {
      this.setState({
        showAutoCompleteBox: false
      });
    }, 300);
  };

  showAutoCompleteBox = () => {
    this.setState({
      showAutoCompleteBox: true
    });
  };

  render() {
    return (
      <div className="search-box">
        <div className="search-input">
          <input
            minLength={2}
            maxLength={50}
            placeholder={'Search short links'}
            onChange={this.handleChange}
            onFocus={this.showAutoCompleteBox}
            onBlur={this.hideAutoCompleteBox}
          />
          <div className={'search-icon'}>
            <Icon iconID={IconID.Search} />
          </div>
        </div>
        {this.createAutoCompleteBox()}
      </div>
    );
  }
}
