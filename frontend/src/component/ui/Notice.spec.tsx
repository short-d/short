import React from 'react';
import { render } from '@testing-library/react';
import { Notice } from './Notice';

it('renders content correctly', () => {
  const { container } = render(
    <Notice>
      <div>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
        tempor incididunt ut labore et dolore magna aliqua
      </div>
    </Notice>
  );
  expect(container.textContent).toMatch(
    'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua'
  );
});
