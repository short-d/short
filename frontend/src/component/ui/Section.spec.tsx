import React from 'react';
import { render } from '@testing-library/react';
import { Section } from './Section';

describe('Section', () => {
  it('renders content correctly', () => {
    const { container } = render(
      <Section title="Section Title">
        <div>Content</div>
      </Section>
    );
    expect(container.textContent).toMatch('Section Title');
    expect(container.textContent).toMatch('Content');
  });

  it('renders content correctly', () => {
    let options = [<div>Option1</div>, <div>Option2</div>];
    const { rerender, container } = render(
      <Section title="Section Title" options={options}>
        <div>Content</div>
      </Section>
    );
    expect(container.textContent).toMatch('Section Title');
    expect(container.textContent).toMatch('Content');
    expect(container.textContent).toMatch('Option1');
    expect(container.textContent).toMatch('Option2');
  });
});
