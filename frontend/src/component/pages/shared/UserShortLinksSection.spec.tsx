import React from 'react';

import { UserShortLinksSection } from './UserShortLinksSection';
import { fireEvent, render } from '@testing-library/react';
import { IPagedShortLinks } from '../../../service/ShortLink.service';

describe('UserShortLinksSection component', () => {
  beforeAll(() => {
    jest.useFakeTimers();
  });

  test('should render without crash', () => {
    render(<UserShortLinksSection onPageLoad={jest.fn} />);
  });

  test('should render nothing when there is no short link', () => {
    const { container } = render(
      <UserShortLinksSection onPageLoad={jest.fn} />
    );

    expect(container.textContent).not.toContain('Favorites');
    expect(container.textContent).not.toContain('Long Link');
    expect(container.textContent).not.toContain('Alias');
  });

  test('should render nothing when there is 0 total pages', () => {
    const pagedShortLinks: IPagedShortLinks = {
      shortLinks: [],
      totalCount: 0,
      offset: 0,
      pageSize: 0
    };

    const { container } = render(
      <UserShortLinksSection
        onPageLoad={jest.fn}
        pagedShortLinks={pagedShortLinks}
      />
    );

    expect(container.textContent).not.toContain('Favorites');
    expect(container.textContent).not.toContain('Long Link');
    expect(container.textContent).not.toContain('Alias');
  });

  test('should render correctly when given at least 1 page', () => {
    const pagedShortLinks: IPagedShortLinks = {
      shortLinks: [],
      totalCount: 1,
      offset: 0,
      pageSize: 0
    };

    const { container } = render(
      <UserShortLinksSection
        onPageLoad={jest.fn}
        pagedShortLinks={pagedShortLinks}
      />
    );

    expect(container.textContent).toContain('Favorites');
    expect(container.textContent).toContain('Long Link');
    expect(container.textContent).toContain('Alias');
  });

  test('should render short links correctly', () => {
    const pagedShortLinks: IPagedShortLinks = {
      shortLinks: [
        { originalUrl: 'https://longurl.com/1', alias: 'alias1' },
        { originalUrl: 'https://longurl.com/2', alias: 'alias2' },
        { originalUrl: 'https://longurl.com/3', alias: 'alias3' }
      ],
      totalCount: 12,
      offset: 0,
      pageSize: 0
    };

    const { container } = render(
      <UserShortLinksSection
        onPageLoad={jest.fn}
        pagedShortLinks={pagedShortLinks}
      />
    );

    expect(container.textContent).toContain('Favorites');
    expect(container.textContent).toContain('Long Link');
    expect(container.textContent).toContain('Alias');

    for (let urlIdx = 0; urlIdx < pagedShortLinks.shortLinks.length; urlIdx++) {
      const shortLink = pagedShortLinks.shortLinks[urlIdx];
      expect(container.textContent).toContain(shortLink.originalUrl);
      expect(container.textContent).toContain(shortLink.alias);
    }
  });

  test('should render short links correctly when the short links does not fill the full page', () => {
    const pagedShortLinks: IPagedShortLinks = {
      shortLinks: [
        { originalUrl: 'https://longurl.com/1', alias: 'alias1' },
        { originalUrl: 'https://longurl.com/2', alias: 'alias2' },
        { originalUrl: 'https://longurl.com/3', alias: 'alias3' }
      ],
      totalCount: 3,
      offset: 0,
      pageSize: 0
    };
    const pageSize = 5;

    const { container } = render(
      <UserShortLinksSection
        onPageLoad={jest.fn}
        pagedShortLinks={pagedShortLinks}
        pageSize={pageSize}
      />
    );

    expect(container.querySelectorAll('.page-number')).toHaveLength(1);
  });

  test('should allow users to navigate between pages when there are multiple pages', () => {
    const pageSize = 5;
    const pagedShortLinks: IPagedShortLinks = {
      shortLinks: [
        { originalUrl: 'https://longurl.com/1', alias: 'alias1' },
        { originalUrl: 'https://longurl.com/2', alias: 'alias2' },
        { originalUrl: 'https://longurl.com/3', alias: 'alias3' },
        { originalUrl: 'https://longurl.com/4', alias: 'alias4' },
        { originalUrl: 'https://longurl.com/5', alias: 'alias5' }
      ],
      totalCount: 5,
      offset: 0,
      pageSize: 0
    };

    const { container } = render(
      <UserShortLinksSection
        onPageLoad={jest.fn}
        pagedShortLinks={pagedShortLinks}
        pageSize={pageSize}
      />
    );

    const expectedPagesCount = Math.ceil(pagedShortLinks.totalCount / pageSize);
    expect(container.querySelectorAll('.page-number')).toHaveLength(
      expectedPagesCount
    );
  });

  test('should load short links with custom page size', () => {
    const onPageLoad = jest.fn();
    const pageSize = 5;

    render(
      <UserShortLinksSection onPageLoad={onPageLoad} pageSize={pageSize} />
    );

    expect(onPageLoad).toBeCalledTimes(1);
    expect(onPageLoad).toBeCalledWith(0, pageSize);
  });

  test('should load short links for initial page correctly', () => {
    const pageSize = 3;
    const initialPagedShortLinks: IPagedShortLinks = {
      shortLinks: [
        { originalUrl: 'https://longurl.com/1', alias: 'alias1' },
        { originalUrl: 'https://longurl.com/2', alias: 'alias2' },
        { originalUrl: 'https://longurl.com/3', alias: 'alias3' }
      ],
      totalCount: 12,
      offset: 0,
      pageSize: 0
    };

    let rerender: any;
    const onPageLoad = jest.fn((offset: number, pageSize: number) => {
      setTimeout(() => {
        rerender(
          <UserShortLinksSection
            onPageLoad={jest.fn}
            pageSize={pageSize}
            pagedShortLinks={initialPagedShortLinks}
          />
        );
      }, 10);
    });

    const component = render(
      <UserShortLinksSection onPageLoad={onPageLoad} pageSize={pageSize} />
    );
    const container = component.container;
    rerender = component.rerender;
    jest.advanceTimersByTime(10);

    expect(onPageLoad).toHaveBeenLastCalledWith(0, pageSize);
    initialPagedShortLinks.shortLinks.forEach(shortLink => {
      expect(container.textContent).toContain(shortLink.originalUrl);
      expect(container.textContent).toContain(shortLink.originalUrl);
    });
    jest.clearAllTimers();
  });

  test('should load correct short links when page changed', () => {
    const pageSize = 3;
    let shortLinks = [
      { originalUrl: 'https://longurl.com/1', alias: 'alias1' },
      { originalUrl: 'https://longurl.com/2', alias: 'alias2' },
      { originalUrl: 'https://longurl.com/3', alias: 'alias3' },
      { originalUrl: 'https://longurl.com/4', alias: 'alias4' },
      { originalUrl: 'https://longurl.com/5', alias: 'alias5' },
      { originalUrl: 'https://longurl.com/6', alias: 'alias6' },
      { originalUrl: 'https://longurl.com/7', alias: 'alias7' },
      { originalUrl: 'https://longurl.com/8', alias: 'alias8' },
      { originalUrl: 'https://longurl.com/9', alias: 'alias9' },
      { originalUrl: 'https://longurl.com/10', alias: 'alias10' },
      { originalUrl: 'https://longurl.com/11', alias: 'alias11' },
      { originalUrl: 'https://longurl.com/12', alias: 'alias12' }
    ];

    let rerender: any;
    const onPageLoad = jest.fn((offset: number, pageSize: number) => {
      setTimeout(() => {
        rerender(
          <UserShortLinksSection
            onPageLoad={jest.fn}
            pageSize={pageSize}
            pagedShortLinks={{
              shortLinks: shortLinks.slice(offset, pageSize),
              totalCount: shortLinks.length,
              offset: 0,
              pageSize: 0
            }}
          />
        );
      }, 10);
    });

    const component = render(
      <UserShortLinksSection
        onPageLoad={onPageLoad}
        pageSize={pageSize}
        pagedShortLinks={{
          shortLinks: shortLinks.slice(0, pageSize),
          totalCount: shortLinks.length,
          offset: 0,
          pageSize: 0
        }}
      />
    );
    const container = component.container;
    rerender = component.rerender;

    const pageNumbers = container.querySelectorAll('.page-number');
    expect(pageNumbers[2]).toBeTruthy();

    fireEvent.click(pageNumbers[2]);
    const offset = 2 * pageSize;
    jest.advanceTimersByTime(10);

    expect(onPageLoad).toHaveBeenLastCalledWith(offset, pageSize);
    shortLinks.slice(offset, pageSize).forEach(link => {
      expect(container.textContent).toContain(link.originalUrl);
      expect(container.textContent).toContain(link.originalUrl);
    });
    jest.clearAllTimers();
  });
});
