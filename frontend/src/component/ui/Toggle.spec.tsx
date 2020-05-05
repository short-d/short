import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { Toggle } from './Toggle';

describe('Toggle component', () => {
  test('should render without crash', () => {
    render(<Toggle />);
  });

  test('should render an inactive toggle when disabled by default', () => {
    const { container } = render(<Toggle defaultIsEnabled={false} />);

    expect(container.querySelector('.background')).toBeTruthy();
    expect(container.querySelector('.knob')).toBeTruthy();
    expect(container.querySelector('.background.active')).toBeNull();
    expect(container.querySelector('.knob.active')).toBeNull();
  });

  test('should render an active toggle when enabled by default', () => {
    const { container } = render(<Toggle defaultIsEnabled={true} />);

    expect(container.querySelector('.background.active')).toBeTruthy();
    expect(container.querySelector('.knob.active')).toBeTruthy();
  });

  test('should toggle from disabled to enabled', () => {
    const { container } = render(<Toggle defaultIsEnabled={false} />);

    const toggle = container.querySelector('.background');
    expect(toggle).toBeTruthy();

    expect(container.querySelector('.knob')).toBeTruthy();
    expect(container.querySelector('.background.active')).toBeNull();
    expect(container.querySelector('.knob.active')).toBeNull();

    fireEvent.click(toggle!);
    expect(container.querySelector('.background.active')).toBeTruthy();
    expect(container.querySelector('.knob.active')).toBeTruthy();
  });

  test('should toggle from enabled to disabled', () => {
    const { container } = render(<Toggle defaultIsEnabled={true} />);

    const toggle = container.querySelector('.background');
    expect(toggle).toBeTruthy();

    expect(container.querySelector('.background.active')).toBeTruthy();
    expect(container.querySelector('.knob.active')).toBeTruthy();

    fireEvent.click(toggle!);
    expect(container.querySelector('.background.active')).toBeNull();
    expect(container.querySelector('.knob.active')).toBeNull();
    expect(container.querySelector('.background')).toBeTruthy();
    expect(container.querySelector('.knob')).toBeTruthy();
  });

  test('should trigger onClick callback when toggle clicked', () => {
    const onClick = jest.fn();

    const { container } = render(<Toggle onClick={onClick} />);

    const toggle = container.querySelector('.background');
    expect(toggle).toBeTruthy();

    expect(onClick).not.toBeCalled();
    fireEvent.click(toggle!);
    expect(onClick).toBeCalled();
  });

  test('should receive new toggle state when clicked', () => {
    const onClick = jest.fn();

    const { container } = render(
      <Toggle onClick={onClick} defaultIsEnabled={true} />
    );

    const toggle = container.querySelector('.background');
    expect(toggle).toBeTruthy();

    fireEvent.click(toggle!);
    expect(onClick).toHaveBeenLastCalledWith(false);
  });

  test('should not crash if toggled without onClick callback', () => {
    const { container } = render(<Toggle />);

    const toggle = container.querySelector('.background');
    expect(toggle).toBeTruthy();
    
    fireEvent.click(toggle!);
  });
});
