import React from 'react';
import { render, fireEvent, Matcher } from '@testing-library/react';
import { SearchBar } from './SearchBar';

describe('Searchbar', () => {
  let changeHandler : () => void;
  let searchBarRef: any;
  let getByPlaceholderText: (id: Matcher) => HTMLElement;
  let input: HTMLInputElement;
  beforeEach(() => {
    changeHandler = jest.fn();
    searchBarRef = React.createRef<SearchBar>();
    getByPlaceholderText = render(
      <SearchBar
        ref={searchBarRef}
        onChange={changeHandler}
        autoCompleteSuggestions={[
          {
            originalUrl: 'https://www.google.com/',
            alias: 'google'
          },
          {
            originalUrl: 'https://github.com/short-d/short/',
            alias: 'short'
          }
        ]}
      />
    ).getByPlaceholderText;
    
    input = getByPlaceholderText(
      'Search short links'
    ) as HTMLInputElement;
  });

  test('should render without auto complete suggestions', () => {
    render(<SearchBar onChange={changeHandler} />);
  });

  test('should trigger change events successfully', async () => {
    fireEvent.change(input, {
      target: { value: 'Lorem ipsum' }
    });

    expect(changeHandler).toBeCalledTimes(0);

    await new Promise(r => setTimeout(r, 300));

    expect(changeHandler).toBeCalledTimes(1);
    expect(changeHandler).toBeCalledWith('Lorem ipsum');
  });

  test('should show autocomplete box on focus', () => {
    expect(searchBarRef.current.state.showAutoCompleteBox).toBe(false);
    input.focus();
    expect(searchBarRef.current.state.showAutoCompleteBox).toBe(true);
  });

  test('should hide autocomplete box on blur', async () => {
    expect(searchBarRef.current.state.showAutoCompleteBox).toBe(false);
    input.focus();
    expect(searchBarRef.current.state.showAutoCompleteBox).toBe(true);
    input.blur();

    await new Promise(r => setTimeout(r, 300));

    expect(searchBarRef.current.state.showAutoCompleteBox).toBe(false);
  });
});
