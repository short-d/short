import React from 'react';
import { fireEvent, render } from '@testing-library/react';
import { Icon, IconID } from './Icon';

describe('Icon component', () => {
  test('should render without crash', () => {
    render(<Icon iconID={IconID.Close} />);
  });

  test('should render the icon correctly', () => {
    const { container } = render(<Icon iconID={IconID.Search} />);

    expect(container.querySelector('.icon-search')).toBeTruthy();
  });

  test('should call click handler when clicked on icon', () => {
    const onClickHandler = jest.fn();
    const { container } = render(
      <Icon iconID={IconID.Close} onClick={onClickHandler} />
    );

    const icon = container.querySelector('.icon');
    expect(icon).toBeTruthy();

    expect(onClickHandler).not.toHaveBeenCalled();
    fireEvent.click(icon!);
    expect(onClickHandler).toHaveBeenCalled();
  });

  test('should not crash when clicked without onClick callback', () => {
    const { container } = render(<Icon iconID={IconID.Close} />);

    const icon = container.querySelector('.icon');

    expect(icon).toBeTruthy();
    fireEvent.click(icon!);
  });

  test('should change icon when passed a different icon as prop', () => {
    const { rerender, container } = render(<Icon iconID={IconID.Search} />);

    expect(container.querySelector('.icon-search')).toBeTruthy();

    rerender(<Icon iconID={IconID.Close} />);
    expect(container.querySelector('.icon-search')).toBeFalsy();
    expect(container.querySelector('.icon-close')).toBeTruthy();
  });
});
