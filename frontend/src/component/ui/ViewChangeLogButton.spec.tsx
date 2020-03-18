import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { ViewChangeLogButton } from './ViewChangeLogButton';

it('renders without crashing', () => {
  render(<ViewChangeLogButton />);
});

it('handles click events correctly', () => {
  const clickHandler = jest.fn();
  const { getByText } = render(<ViewChangeLogButton onClick={clickHandler} />);
  fireEvent.click(getByText('View Changelog'));
  expect(clickHandler).toHaveBeenCalledTimes(1);
});
