import React from 'react';

import { UserShortLinksSection } from './UserShortLinksSection';
import { fireEvent, render } from '@testing-library/react';
import { IPagedShortLinks } from '../../../service/ShortLink.service';

const sampleShortLinks = [
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

describe('UserShortLinksSection component', () => {
  test('should render without crash', () => {
    render(<UserShortLinksSection onPageLoad={jest.fn} />);
  });

  test('should render nothing when no data is sent', () => {
    const { container } = render(
      <UserShortLinksSection onPageLoad={jest.fn} />
    );

    expect(container.textContent).not.toContain('Created Short Links');
    expect(container.textContent).not.toContain('Long URL');
    expect(container.textContent).not.toContain('Alias');
  });

  test('should render nothing when "total" attribute in pagedShortLinks is zero', () => {
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

  test('should render when "total" attribute in pagedShortLinks is non zero with empty shortLinks', () => {
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

  test('should render passed short links correctly', () => {
    const pagedShortLinks: IPagedShortLinks = {
      shortLinks: sampleShortLinks.slice(0, 3),
      totalCount: sampleShortLinks.length
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

  test('should render correctly when total short links passed is less than page size', () => {
    const pagedShortLinks: IPagedShortLinks = {
      shortLinks: sampleShortLinks.slice(0, 3),
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

  test('should render multiple pages when total short link count exceeds page size', () => {
    const pageSize = 5;
    const pagedShortLinks: IPagedShortLinks = {
      shortLinks: sampleShortLinks.slice(0, pageSize),
      totalCount: sampleShortLinks.length
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

  test('should request short links for initial page correctly', () => {
    const onPageLoad = jest.fn();

    render(<UserShortLinksSection onPageLoad={onPageLoad} />);

    expect(onPageLoad).toBeCalledTimes(1);
  });

  test('should request short links with non default page size when passed as prop', () => {
    const onPageLoad = jest.fn();
    const pageSize = 5;

    render(
      <UserShortLinksSection onPageLoad={onPageLoad} pageSize={pageSize} />
    );

    expect(onPageLoad).toBeCalledTimes(1);
    expect(onPageLoad).toBeCalledWith(0, pageSize);
  });

  test('should request and render short links with correct offset when page changed', () => {
    const pageSize = 3;
    let pagedShortLinks: IPagedShortLinks = {
      shortLinks: sampleShortLinks.slice(0, pageSize),
      totalCount: sampleShortLinks.length
    };
    const onPageLoad = jest.fn((offset: number, pageSize: number) => {
      pagedShortLinks = {
        shortLinks: sampleShortLinks.slice(offset, offset + pageSize),
        totalCount: sampleShortLinks.length
      };
    });

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
    const expectedShortLinks = sampleShortLinks.slice(
      offset,
      offset + pageSize
    );
    expect(onPageLoad).toBeCalledTimes(2);
    expect(onPageLoad).toHaveBeenLastCalledWith(offset, pageSize);
    expect(pagedShortLinks.shortLinks).toEqual(expectedShortLinks);
  });
});
