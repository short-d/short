import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { SearchBar } from './SearchBar';

const DEBOUNCE_DURATION: number = 300;

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

  await new Promise(r => setTimeout(r, DEBOUNCE_DURATION));

  expect(changeHandler).toBeCalledTimes(1);
  expect(changeHandler).toBeCalledWith('Lorem ipsum');
});
