import React from 'react';
import { render } from '@testing-library/react';
import { Drawer } from './Drawer';

describe('Drawer component', () => {
  test('should render without crash', () => {
    render(<Drawer />);
  });

  test('should render in open state by default', () => {
    const { container } = render(<Drawer />);

    const drawer = container.querySelector('.drawer');
    expect(drawer).toBeTruthy();
    expect(drawer!.className).toContain('open');
  });

  test('should close the drawer when triggered to close', () => {
    const drawerRef = React.createRef<Drawer>();
    const { container } = render(<Drawer ref={drawerRef} />);

    const drawer = container.querySelector('.drawer');
    expect(drawer).toBeTruthy();
    expect(drawer!.className).toContain('open');

    drawerRef.current!.close();
    expect(drawer!.className).not.toContain('open');
  });

  test('should open the drawer when triggered to open', () => {
    const drawerRef = React.createRef<Drawer>();
    const { container } = render(<Drawer ref={drawerRef} />);

    const drawer = container.querySelector('.drawer');
    expect(drawer).toBeTruthy();
    drawerRef.current!.close();
    expect(drawer!.className).not.toContain('open');

    drawerRef.current!.open();
    expect(drawer!.className).toContain('open');
  });

  test('should render component children correctly', () => {
    const childContent = 'children of drawer component';
    const { container } = render(<Drawer> {childContent} </Drawer>);

    expect(container.textContent).toContain(childContent);
  });
});
