import React from 'react';
import { render } from '@testing-library/react';
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
    const toggleRef = React.createRef<Toggle>();
    const { container } = render(
      <Toggle ref={toggleRef} defaultIsEnabled={false} />
    );

    expect(container.querySelector('.background')).toBeTruthy();
    expect(container.querySelector('.knob')).toBeTruthy();
    expect(container.querySelector('.background.active')).toBeNull();
    expect(container.querySelector('.knob.active')).toBeNull();
    toggleRef.current?.handleClick();
    expect(container.querySelector('.background.active')).toBeTruthy();
    expect(container.querySelector('.knob.active')).toBeTruthy();
  });

  test('should toggle from enabled to disabled', () => {
    const toggleRef = React.createRef<Toggle>();
    const { container } = render(
      <Toggle ref={toggleRef} defaultIsEnabled={true} />
    );

    expect(container.querySelector('.background.active')).toBeTruthy();
    expect(container.querySelector('.knob.active')).toBeTruthy();
    toggleRef.current?.handleClick();
    expect(container.querySelector('.background.active')).toBeNull();
    expect(container.querySelector('.knob.active')).toBeNull();
    expect(container.querySelector('.background')).toBeTruthy();
    expect(container.querySelector('.knob')).toBeTruthy();
  });

  test('should trigger onClick callback when toggle clicked', () => {
    const onClick = jest.fn();

    const toggleRef = React.createRef<Toggle>();
    render(<Toggle ref={toggleRef} onClick={onClick} />);

    expect(onClick).not.toBeCalled();
    toggleRef.current?.handleClick();
    expect(onClick).toBeCalled();
  });

  test('should contain the new enabled state of the toggle when clicked', () => {
    const onClick = jest.fn();

    const toggleRef = React.createRef<Toggle>();
    render(
      <Toggle ref={toggleRef} onClick={onClick} defaultIsEnabled={true} />
    );

    toggleRef.current?.handleClick();
    expect(onClick).toHaveBeenLastCalledWith(false);
  });

  test('should not crash if toggled without onClick callback', () => {
    const toggleRef = React.createRef<Toggle>();
    render(<Toggle ref={toggleRef} />);

    toggleRef.current?.handleClick();
  });
});
