import React from 'react';
import { fireEvent, render } from '@testing-library/react';
import { Navigation } from './Navigation';

describe('Navigation component', () => {
  test('should render without crash', () => {
    render(<Navigation menuItems={[]} onMenuItemSelected={jest.fn} />);
  });

  test('should render empty navigation component when there are no menu items', () => {
    const { container } = render(
      <Navigation menuItems={[]} onMenuItemSelected={jest.fn} />
    );

    expect(container.querySelectorAll('.navigation')).toHaveLength(1);
    expect(container.querySelectorAll('.menu')).toHaveLength(1);
    expect(container.querySelectorAll('.menu > li')).toHaveLength(0);
  });

  test('should render correctly when at least one menu item given', () => {
    const menuItemContents = ['item1', 'item2', 'item3'];
    const { container } = render(
      <Navigation menuItems={menuItemContents} onMenuItemSelected={jest.fn} />
    );

    expect(container.querySelectorAll('.navigation')).toHaveLength(1);
    expect(container.querySelectorAll('.menu')).toHaveLength(1);
    expect(container.querySelectorAll('.menu > li')).toHaveLength(
      menuItemContents.length
    );
    menuItemContents.forEach(item => {
      expect(container.textContent).toContain(item);
    });
  });

  test('should select the first menu item when default menu item index not overridden', () => {
    const menuItemContents = ['item1', 'item2', 'item3'];
    const { container } = render(
      <Navigation menuItems={menuItemContents} onMenuItemSelected={jest.fn} />
    );

    const menuItems = container.querySelectorAll('.menu > li');
    expect(menuItems[0]).toBeTruthy();
    expect(menuItems[0].className).toContain('active');
  });

  test('should select correct menu item when default menu item index overridden', () => {
    const menuItemContents = ['item1', 'item2', 'item3'];
    const defaultMenuItemIdx = 1;
    const { container } = render(
      <Navigation
        menuItems={menuItemContents}
        defaultMenuItemIdx={defaultMenuItemIdx}
        onMenuItemSelected={jest.fn}
      />
    );

    const menuItems = container.querySelectorAll('.menu > li');
    expect(menuItems[defaultMenuItemIdx]).toBeTruthy();
    expect(menuItems[defaultMenuItemIdx].className).toContain('active');
  });

  test('should select clicked menu item', () => {
    const menuItemContents = ['item1', 'item2', 'item3'];
    const { container } = render(
      <Navigation menuItems={menuItemContents} onMenuItemSelected={jest.fn()} />
    );

    const menuItems = container.querySelectorAll('.menu > li');
    expect(menuItems[0].className).toContain('active');
    expect(menuItems[1].className).not.toContain('active');

    fireEvent.click(menuItems[1]);
    expect(menuItems[0].className).not.toContain('active');
    expect(menuItems[1].className).toContain('active');
  });

  test('should call correct menu item select handler', () => {
    const menuItemContents = ['item1', 'item2', 'item3'];
    const handleMenuItemSelect = jest.fn();
    const { container } = render(
      <Navigation
        menuItems={menuItemContents}
        onMenuItemSelected={handleMenuItemSelect}
      />
    );

    const menuItems = container.querySelectorAll('.menu > li');
    fireEvent.click(menuItems[1]);

    expect(handleMenuItemSelect).toHaveBeenLastCalledWith(1);
  });
});
