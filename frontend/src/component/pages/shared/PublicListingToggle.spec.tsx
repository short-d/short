import React from 'react';

import { PublicListingToggle } from './PublicListingToggle';
import { render } from '@testing-library/react';

describe('PublicListingToggle component', () => {
  beforeAll(() => {
    jest.useFakeTimers();
  });

  test('should render without crash', () => {
    render(<PublicListingToggle onToggleClick={jest.fn} />);
  });

  test('should render the public listing label', () => {
    const { container } = render(
      <PublicListingToggle onToggleClick={jest.fn} />
    );

    expect(container.textContent).toContain("Share on public feed");
  });

  test('should render a toggle', () => {
    const { container } = render(<PublicListingToggle onToggleClick={jest.fn} />);

    expect(container.querySelector('.toggle')).not.toBeNull();
  });

  test('should render an active toggle if enabled by default', () => {
    const { container } = render(
      <PublicListingToggle defaultIsEnabled={true} onToggleClick={jest.fn} />
    );

    expect(container.querySelector('.toggle > .active')).not.toBeNull();
  });

  test('should render an inactive toggle if explicitly disabled by default', () => {
    const { container } = render(
      <PublicListingToggle defaultIsEnabled={false} onToggleClick={jest.fn} />
    );

    expect(container.querySelector('.toggle')).not.toBeNull();
  });
});
