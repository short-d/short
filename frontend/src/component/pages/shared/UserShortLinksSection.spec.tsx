import React from 'react';

import { UserShortLinksSection } from './UserShortLinksSection';
import { fireEvent, render } from '@testing-library/react';
import { IQueryUrlData } from '../../../service/ShortLink.service';

describe('UserShortLinksSection component', () => {
  test('should render without crash', () => {
    render(<UserShortLinksSection updateUrlData={jest.fn} />);
  });

  test('should render nothing when no data is sent', () => {
    const { container } = render(
      <UserShortLinksSection updateUrlData={jest.fn} />
    );

    expect(container.textContent).not.toContain('Created Short Links');
    expect(container.textContent).not.toContain('Long URL');
    expect(container.textContent).not.toContain('Alias');
  });

  test('should render nothing when "total" attribute in urlData is zero', () => {
    const urlData: IQueryUrlData = {
      urls: [],
      total: 0
    };

    const { container } = render(
      <UserShortLinksSection updateUrlData={jest.fn} urlData={urlData} />
    );

    expect(container.textContent).not.toContain('Created Short Links');
    expect(container.textContent).not.toContain('Long URL');
    expect(container.textContent).not.toContain('Alias');
  });

  test('should render when "total" attribute in urlData is non zero with empty urls', () => {
    const urlData: IQueryUrlData = {
      urls: [],
      total: 1
    };

    const { container } = render(
      <UserShortLinksSection updateUrlData={jest.fn} urlData={urlData} />
    );

    expect(container.textContent).toContain('Created Short Links');
    expect(container.textContent).toContain('Long URL');
    expect(container.textContent).toContain('Alias');
  });

  test('should render when url data is sent', () => {
    const urlData: IQueryUrlData = {
      urls: [
        { originalUrl: 'https://longurl.com/1', alias: 'alias1' },
        { originalUrl: 'https://longurl.com/2', alias: 'alias2' },
        { originalUrl: 'https://longurl.com/3', alias: 'alias3' }
      ],
      total: 3
    };

    const { container } = render(
      <UserShortLinksSection updateUrlData={jest.fn} urlData={urlData} />
    );

    expect(container.textContent).toContain('Created Short Links');
    expect(container.textContent).toContain('Long URL');
    expect(container.textContent).toContain('Alias');

    for (let urlIdx = 0; urlIdx < urlData.urls.length; urlIdx++) {
      expect(container.textContent).toContain(urlData.urls[urlIdx].originalUrl);
      expect(container.textContent).toContain(urlData.urls[urlIdx].alias);
    }
  });

  test('should call update data function to fetch initial data when component renders ', () => {
    const updateUrlData = jest.fn();

    render(<UserShortLinksSection updateUrlData={updateUrlData} />);

    expect(updateUrlData).toBeCalledTimes(1);
  });
});
