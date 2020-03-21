import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { SearchBar } from './SearchBar';

function getSearchBarUtil() {
  const changeHandler = jest.fn();
  const searchBarRef = React.createRef<SearchBar>();
  const { getByPlaceholderText } = render(
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
  );

  const input = getByPlaceholderText('Search short links') as HTMLInputElement;
  return {
    searchBarRef,
    input
  };
}

describe('Searchbar', () => {
  test('should render without crash', () => {
    const changeHandler = jest.fn();
    render(<SearchBar onChange={changeHandler} />);
  });

  test('should trigger change events successfully', async () => {
    const changeHandler = jest.fn();
    const { getByPlaceholderText } = render(
      <SearchBar onChange={changeHandler} />
    );
    const input = getByPlaceholderText(
      'Search short links'
    ) as HTMLInputElement;

    fireEvent.change(input, {
      target: { value: 'Lorem ipsum' }
    });

    expect(changeHandler).toBeCalledTimes(0);

    await new Promise(r => setTimeout(r, 300));

    expect(changeHandler).toBeCalledTimes(1);
    expect(changeHandler).toBeCalledWith('Lorem ipsum');
  });

  test('should show autocomplete box on focus', () => {
    const { searchBarRef, input } = getSearchBarUtil();

    expect(searchBarRef.current.state.showAutoCompleteBox).toBe(false);
    input.focus();
    expect(searchBarRef.current.state.showAutoCompleteBox).toBe(true);
  });

  test('should hide autocomplete box on blur', async () => {
    const { searchBarRef, input } = getSearchBarUtil();

    expect(searchBarRef.current.state.showAutoCompleteBox).toBe(false);
    input.focus();
    expect(searchBarRef.current.state.showAutoCompleteBox).toBe(true);
    input.blur();

    await new Promise(r => setTimeout(r, 300));

    expect(searchBarRef.current.state.showAutoCompleteBox).toBe(false);
  });
});
