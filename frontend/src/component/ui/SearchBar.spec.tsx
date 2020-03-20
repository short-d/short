import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { SearchBar } from './SearchBar';

it('renders without crashing', () => {
  const changeHandler = jest.fn();
  render(<SearchBar onChange={changeHandler} />);
});

it('triggers change events successfully', async () => {
  const changeHandler = jest.fn();
  const { getByPlaceholderText } = render(
    <SearchBar onChange={changeHandler} />
  );
  const input = getByPlaceholderText('Search short links') as HTMLInputElement;

  fireEvent.change(input, {
    target: { value: 'Lorem ipsum' }
  });

  expect(changeHandler).toBeCalledTimes(0);

  await new Promise(r => setTimeout(r, 300));

  expect(changeHandler).toBeCalledTimes(1);
  expect(changeHandler).toBeCalledWith('Lorem ipsum');
});

it('shows autocomplete box on focus', () => {
    const changeHandler = jest.fn();
    const searchBarRef = React.createRef<SearchBar>();
    const { getByPlaceholderText } = render(
      <SearchBar ref={searchBarRef} onChange={changeHandler} autoCompleteSuggestions={[
        {
          originalUrl: 'https://www.google.com/',
          alias: 'google'
        },
        {
          originalUrl: 'https://github.com/short-d/short/',
          alias: 'short'
        },
        {
          originalUrl: 'https://developer.mozilla.org/en-US/',
          alias: 'mozilla'
        }
      ]}/>
    );

    const input = getByPlaceholderText('Search short links') as HTMLInputElement;

    expect(searchBarRef).toBeTruthy();
    expect(searchBarRef.current).toBeTruthy();

    expect(searchBarRef.current!.state.showAutoCompleteBox).toBe(false);
    input.focus();
    expect(searchBarRef.current!.state.showAutoCompleteBox).toBe(true);
});

it('hides autocomplete box on blur', async () => {
    const changeHandler = jest.fn();
    const searchBarRef = React.createRef<SearchBar>();
    const { getByPlaceholderText } = render(
      <SearchBar ref={searchBarRef} onChange={changeHandler} autoCompleteSuggestions={[
        {
          originalUrl: 'https://www.google.com/',
          alias: 'google'
        },
        {
          originalUrl: 'https://github.com/short-d/short/',
          alias: 'short'
        },
        {
          originalUrl: 'https://developer.mozilla.org/en-US/',
          alias: 'mozilla'
        }
      ]}/>
    );

    const input = getByPlaceholderText('Search short links') as HTMLInputElement;

    expect(searchBarRef).toBeTruthy();
    expect(searchBarRef.current).toBeTruthy();

    expect(searchBarRef.current!.state.showAutoCompleteBox).toBe(false);
    input.focus();
    expect(searchBarRef.current!.state.showAutoCompleteBox).toBe(true);
    input.blur();

    await new Promise(r => setTimeout(r, 300));

    expect(searchBarRef.current!.state.showAutoCompleteBox).toBe(false);
});