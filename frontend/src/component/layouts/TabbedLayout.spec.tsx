import React from 'react';
import { fireEvent, render } from '@testing-library/react';
import { Tab, TabbedLayout } from './TabbedLayout';

describe('TabbedLayout component', () => {
  test('should render without crash', () => {
    render(<TabbedLayout tabs={[]} />);
  });

  test('should render empty container when there are no tabs', () => {
    const { container } = render(<TabbedLayout tabs={[]} />);

    expect(container.querySelector('.tab-layout')).toBeTruthy();
    expect(container.querySelector('.tab-headers')).toBeTruthy();

    const tabContent = container.querySelector('.tab-content');
    expect(tabContent).toBeTruthy();
    expect(tabContent!.childNodes).toHaveLength(0);
  });

  test('should render all headers correctly', () => {
    const tabs: Tab[] = [
      {
        header: 'header 1',
        content: 'content 1'
      },
      {
        header: 'header 2',
        content: 'content 2'
      }
    ];

    const { container } = render(<TabbedLayout tabs={tabs} />);

    const tabHeaders = container.querySelector('.tab-headers');
    expect(tabHeaders).toBeTruthy();
    tabs.forEach((tab: Tab) => {
      expect(tabHeaders!.textContent).toContain(tab.header);
    });
  });

  test('should render first tab content by default', () => {
    const tabs: Tab[] = [
      { header: 'header 1', content: 'content 1' },
      { header: 'header 2', content: 'content 2' },
      { header: 'header 3', content: 'content 3' }
    ];

    const { container } = render(<TabbedLayout tabs={tabs} />);

    const tabContent = container.querySelector('.tab-content');
    expect(tabContent).toBeTruthy();
    expect(tabContent!.textContent).toContain(tabs[0].content);
  });

  test('should render correct tab content when navigated to different tab', () => {
    const tabs: Tab[] = [
      { header: 'header 1', content: 'content 1' },
      { header: 'header 2', content: 'content 2' },
      { header: 'header 3', content: 'content 3' }
    ];

    const { container } = render(<TabbedLayout tabs={tabs} />);
    const tabHeaders = container.querySelectorAll('.tab-headers .menu li');
    const tabContent = container.querySelector('.tab-content');

    expect(tabHeaders).toHaveLength(tabs.length);
    for (let headerIdx = 0; headerIdx < tabs.length; headerIdx++) {
      fireEvent.click(tabHeaders[headerIdx]);
      expect(tabContent).toBeTruthy();
      expect(tabContent!.textContent).toContain(tabs[headerIdx].content);
    }
  });

  test('should hide tab headers when triggered to hide headers', () => {
    const tabs: Tab[] = [
      { header: 'header 1', content: 'content 1' },
      { header: 'header 2', content: 'content 2' }
    ];
    const tabRef = React.createRef<TabbedLayout>();

    const { container } = render(<TabbedLayout tabs={tabs} ref={tabRef} />);

    const tabHeaders = container.querySelector('.tab-headers .drawer');
    expect(tabHeaders).toBeTruthy();
    expect(tabHeaders!.className).toContain('open');

    tabRef.current!.hideHeaders();
    expect(tabHeaders!.className).not.toContain('open');
  });

  test('should show tab headers when triggered to show headers', () => {
    const tabs: Tab[] = [
      { header: 'header 1', content: 'content 1' },
      { header: 'header 2', content: 'content 2' }
    ];
    const tabRef = React.createRef<TabbedLayout>();

    const { container } = render(<TabbedLayout tabs={tabs} ref={tabRef} />);

    const tabHeaders = container.querySelector('.tab-headers .drawer');
    tabRef.current!.hideHeaders();
    expect(tabHeaders!.className).not.toContain('open');

    tabRef.current!.showHeaders();
    expect(tabHeaders!.className).toContain('open');
  });
});
