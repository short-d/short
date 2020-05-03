import React from 'react';
import { fireEvent, render } from '@testing-library/react';
import { Icon, IconID } from './Icon';

describe('Icon component', () => {
  test('should render without crash', () => {
    render(<Icon iconID={IconID.Close} />);
  });

  test('should render the icon correctly', () => {
    const iconRef = React.createRef<Icon>();
    const { container } = render(<Icon ref={iconRef} iconID={IconID.Search} />);

    const expectedSVG = render(iconRef.current!.renderSVG(IconID.Search));
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
});
