import React from 'react';

import { UserShortLinksSection } from './UserShortLinksSection';
import { fireEvent, render } from '@testing-library/react';
import { IPagedShortLinks } from '../../../service/ShortLink.service';

describe('UserShortLinksSection component', () => {
  test('should render without crash', () => {
    render(<UserShortLinksSection onPageLoad={jest.fn} />);
  });

  test('should render nothing when there is no short link', () => {
    const { container } = render(
      <UserShortLinksSection onPageLoad={jest.fn} />
    );

    expect(container.textContent).not.toContain('Created Short Links');
    expect(container.textContent).not.toContain('Long URL');
    expect(container.textContent).not.toContain('Alias');
  });

  test('should render nothing when there is 0 total pages', () => {
    const pagedShortLinks: IPagedShortLinks = {
      shortLinks: [],
      totalCount: 0
    };

    const { container } = render(
      <UserShortLinksSection
        onPageLoad={jest.fn}
        pagedShortLinks={pagedShortLinks}
      />
    );

    expect(container.textContent).not.toContain('Created Short Links');
    expect(container.textContent).not.toContain('Long URL');
    expect(container.textContent).not.toContain('Alias');
  });

  test('should render short links correctly when given at least 1 page', () => {
    const pagedShortLinks: IPagedShortLinks = {
      shortLinks: [],
      totalCount: 1
    };

    const { container } = render(
      <UserShortLinksSection
        onPageLoad={jest.fn}
        pagedShortLinks={pagedShortLinks}
      />
    );

    expect(container.textContent).toContain('Created Short Links');
    expect(container.textContent).toContain('Long URL');
    expect(container.textContent).toContain('Alias');
  });

  test('should render short links correctly', () => {
    const pagedShortLinks: IPagedShortLinks = {
      shortLinks: [
        { originalUrl: 'https://longurl.com/1', alias: 'alias1' },
        { originalUrl: 'https://longurl.com/2', alias: 'alias2' },
        { originalUrl: 'https://longurl.com/3', alias: 'alias3' }
      ],
      totalCount: 12
    };

    const { container } = render(
      <UserShortLinksSection
        onPageLoad={jest.fn}
        pagedShortLinks={pagedShortLinks}
      />
    );

    expect(container.textContent).toContain('Created Short Links');
    expect(container.textContent).toContain('Long URL');
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
      totalCount: 3
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
      totalCount: 5
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

  test('should load short links for initial page correctly', () => {
    const onPageLoad = jest.fn();

    render(<UserShortLinksSection onPageLoad={onPageLoad} />);

    expect(onPageLoad).toBeCalledTimes(1);
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

  test('should load correct short links when page changed', () => {
    const pageSize = 3;
    let pagedShortLinks: IPagedShortLinks = {
      shortLinks: [
        { originalUrl: 'https://longurl.com/1', alias: 'alias1' },
        { originalUrl: 'https://longurl.com/2', alias: 'alias2' },
        { originalUrl: 'https://longurl.com/3', alias: 'alias3' }
      ],
      totalCount: 12
    };
    const onPageLoad = jest.fn();

    const { container } = render(
      <UserShortLinksSection
        onPageLoad={onPageLoad}
        pageSize={pageSize}
        pagedShortLinks={pagedShortLinks}
      />
    );

    const pageNumbers = container.querySelectorAll('.page-number');
    expect(pageNumbers[2]).toBeTruthy();

    fireEvent.click(pageNumbers[2]);
    const offset = 2 * pageSize;
    expect(onPageLoad).toHaveBeenLastCalledWith(offset, pageSize);
  });
});
