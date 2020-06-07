import React from 'react';
import { render, fireEvent, Matcher } from '@testing-library/react';
import { SearchBar } from './SearchBar';

describe('Searchbar', () => {
  beforeAll(() => {
    jest.useFakeTimers();
  });

  afterEach(() => {
    jest.clearAllTimers();
  });

  test('should render without auto complete suggestions', () => {
    const changeHandler = jest.fn();
    render(<SearchBar onChange={changeHandler} />);
  });

  test('should trigger change events successfully after debounce time', async () => {
    const changeHandler = jest.fn();
    const TOTAL_DURATION = 300;
    const { getByPlaceholderText } = render(
      <SearchBar
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

    const input = getByPlaceholderText(
      'Search short links'
    ) as HTMLInputElement;

    fireEvent.change(input, {
      target: { value: 'Lorem ipsum' }
    });

    expect(changeHandler).toBeCalledTimes(0);

    jest.advanceTimersByTime(TOTAL_DURATION);

    expect(changeHandler).toBeCalledTimes(1);
    expect(changeHandler).toBeCalledWith('Lorem ipsum');
  });

  test('should show autocomplete box on focus', () => {
    const changeHandler = jest.fn();
    const { getByPlaceholderText, container } = render(
      <SearchBar
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

    const input = getByPlaceholderText(
      'Search short links'
    ) as HTMLInputElement;

    expect(container.querySelector('.suggestions.show')).toBeFalsy();
    fireEvent.focus(input);
    expect(container.querySelector('.suggestions.show')).toBeTruthy();
  });

  test('should hide autocomplete box on blur', async () => {
    const changeHandler = jest.fn();
    const TOTAL_DURATION = 300;
    const { getByPlaceholderText, container } = render(
      <SearchBar
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

    const input = getByPlaceholderText(
      'Search short links'
    ) as HTMLInputElement;

    expect(container.querySelector('.suggestions.show')).toBeFalsy();
    fireEvent.focus(input);
    expect(container.querySelector('.suggestions.show')).toBeTruthy();
    fireEvent.blur(input);

    jest.advanceTimersByTime(TOTAL_DURATION);

    expect(container.querySelector('.suggestions.show')).toBeFalsy();
  });
});
