import React from 'react';

import { Button } from './Button';
import { render, fireEvent } from '@testing-library/react';

it('renders without crashing', () => {
  const { getByText } = render(<Button>Click Me</Button>);
  fireEvent.click(getByText('Click Me'));
});

it('handles click on button', () => {
  const clickHandler = jest.fn();
  const { getByText } = render(
    <Button onClick={clickHandler}>Click Me</Button>
  );
  fireEvent.click(getByText('Click Me'));
  expect(clickHandler).toHaveBeenCalledTimes(1);
});
