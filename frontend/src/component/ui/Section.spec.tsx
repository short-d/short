import React from 'react';
import { render } from '@testing-library/react';
import { Section } from './Section';

it('renders without crashing', () => {
  const { container } = render(
    <Section title="Section Title">
        <div>
        Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua
        </div>
    </Section>
  );
  expect(container.textContent).toMatch('Section Title');
  expect(container.textContent).toMatch('Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua');
});
