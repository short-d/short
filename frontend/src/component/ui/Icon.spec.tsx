import React from 'react';
import { fireEvent, render } from '@testing-library/react';
import { Icon, IconID } from './Icon';

describe('Navigation component', () => {
  test('should render without crash', () => {
    render(<Icon defaultIconID={IconID.Close} />);
  });

  test('should render the icon correctly', () => {
    const iconRef = React.createRef<Icon>();
    const { container } = render(
      <Icon ref={iconRef} defaultIconID={IconID.Search} />
    );

    const expectedSVG = render(iconRef.current!.renderSVG(IconID.Search));
    expect(container.querySelector('.icon')!.innerHTML).toContain(
      expectedSVG.container.innerHTML
    );
  });

  test('should call click handler when clicked on icon', () => {
    const onClickHandler = jest.fn();
    const { container } = render(
      <Icon defaultIconID={IconID.Close} onClick={onClickHandler} />
    );

    const icon = container.querySelector('.icon');
    expect(icon).toBeTruthy();

    expect(onClickHandler).not.toHaveBeenCalled();
    fireEvent.click(icon!);
    expect(onClickHandler).toHaveBeenCalled();
  });

  test('should change icon when icon setter is triggered', () => {
    const iconRef = React.createRef<Icon>();
    const { container } = render(
      <Icon ref={iconRef} defaultIconID={IconID.Search} />
    );

    const SVGBeforeChange = render(iconRef.current!.renderSVG(IconID.Search));
    expect(container.querySelector('.icon')!.innerHTML).toContain(
      SVGBeforeChange.container.innerHTML
    );

    iconRef.current!.setIcon(IconID.Close);
    const SVGAfterChange = render(iconRef.current!.renderSVG(IconID.Close));
    expect(container.querySelector('.icon')!.innerHTML).toContain(
      SVGAfterChange.container.innerHTML
    );
  });
});
