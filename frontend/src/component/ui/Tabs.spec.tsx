import React from 'react';
import { render } from '@testing-library/react';
import { Tabs } from './Tabs';

describe('Tabs component', () => {
  test('should render without crash', () => {
    render(<Tabs />);
  });

  test('should not render anything when there is no tab', () => {
    const { container } = render(<Tabs />);
    expect(container.innerHTML).toBeFalsy();
  });

  test('should render first tab by default', () => {
    const tabs = ['Content 1', 'Content 2', 'Content 3'];

    const { container } = render(<Tabs>{tabs}</Tabs>);

    expect(container.textContent).toContain(tabs[0]);
    expect(container.textContent).not.toContain(tabs[1]);
    expect(container.textContent).not.toContain(tabs[2]);
  });

  test('should render correct tab when default index is given', () => {
    const tabs = ['Content 1', 'Content 2', 'Content 3'];

    const { container } = render(<Tabs defaultTabIdx={1}>{tabs}</Tabs>);

    expect(container.textContent).not.toContain(tabs[0]);
    expect(container.textContent).toContain(tabs[1]);
    expect(container.textContent).not.toContain(tabs[2]);
  });

  test('should render correct tab content when showTab is invoked', () => {
    const tabs = ['Content 1', 'Content 2', 'Content 3'];
    const tabRef = React.createRef<Tabs>();

    const { container } = render(<Tabs ref={tabRef}>{tabs}</Tabs>);

    expect(container.textContent).not.toContain(tabs[1]);
    tabRef.current!.showTab(1);
    expect(container.textContent).toContain(tabs[1]);
  });

  test('should render nothing when negative index is given', () => {
    const tabs = ['Content 1', 'Content 2', 'Content 3'];
    const tabRef = React.createRef<Tabs>();

    const { container } = render(<Tabs ref={tabRef}>{tabs}</Tabs>);

    tabRef.current!.showTab(-1);
    expect(container.innerHTML).toBeFalsy();
  });

  test('should not render anything when index grater than tabs count is given', () => {
    const tabs = ['Content 1', 'Content 2', 'Content 3'];
    const tabRef = React.createRef<Tabs>();

    const { container } = render(<Tabs ref={tabRef}>{tabs}</Tabs>);

    tabRef.current!.showTab(tabs.length + 1);
    expect(container.innerHTML).toBeFalsy();
  });
});
