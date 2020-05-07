import React from 'react';
import { fireEvent, render } from '@testing-library/react';
import { PageControl } from './PageControl';

describe('PageControl component', () => {
  test('should render without crash', () => {
    render(<PageControl totalPages={0} />);
  });

  test('should render nothing when there is no page', () => {
    const { container } = render(<PageControl totalPages={0} />);

    expect(container.textContent).not.toContain('Previous');
    expect(container.textContent).not.toContain('Next');
    expect(container.textContent).not.toContain('0');
  });

  test('should render previous, next and page number 1 when this is 1 page', () => {
    const { container } = render(<PageControl totalPages={1} />);

    expect(container.textContent).toContain('Previous');
    expect(container.textContent).toContain('Next');
    expect(container.textContent).toContain('1');
  });

  test('should render previous, next and page numbers when there are 10 pages', () => {
    const { container } = render(<PageControl totalPages={10} />);

    expect(container.textContent).toContain('Previous');
    expect(container.textContent).toContain('Next');

    for (let pageIdx = 0; pageIdx < 10; pageIdx++) {
      expect(container.textContent).toContain(`${pageIdx + 1}`);
    }
  });

  test('should go to correct page when clicking on page numbers', () => {
    let currentPageIdx = 0;
    const handlePageChanged = (pageIdx: number) => {
      currentPageIdx = pageIdx;
    };

    const { container } = render(
      <PageControl totalPages={7} onPageChanged={handlePageChanged} />
    );

    const pageNumbers = container.querySelectorAll('.page-number');

    for (let pageIdx = 0; pageIdx < 7; pageIdx++) {
      fireEvent.click(pageNumbers[pageIdx]);
      expect(currentPageIdx).toBe(pageIdx);
    }
  });

  test('should go to previous page when clicking Previous', () => {
    let currentPageIdx = 0;
    const handlePageChanged = (pageIdx: number) => {
      currentPageIdx = pageIdx;
    };

    const { container } = render(
      <PageControl totalPages={7} onPageChanged={handlePageChanged} />
    );

    const pageNumbers = container.querySelectorAll('.page-number');
    fireEvent.click(pageNumbers[3]);
    expect(currentPageIdx).toBe(3);

    const prevButton = container.querySelector('.previous');
    expect(prevButton).toBeTruthy();
    expect(prevButton!.textContent).toContain('Previous');
    fireEvent.click(prevButton!);
    expect(currentPageIdx).toBe(2);
  });

  test('should do nothing when on first page and when clicking Previous', () => {
    let currentPageIdx = 0;
    const handlePageChanged = (pageIdx: number) => {
      currentPageIdx = pageIdx;
    };

    const { container } = render(
      <PageControl totalPages={7} onPageChanged={handlePageChanged} />
    );

    const pageNumbers = container.querySelectorAll('.page-number');
    fireEvent.click(pageNumbers[0]);
    expect(currentPageIdx).toBe(0);

    const prevButton = container.querySelector('.previous');
    expect(prevButton).toBeTruthy();
    expect(prevButton!.textContent).toContain('Previous');
    fireEvent.click(prevButton!);
    expect(currentPageIdx).toBe(0);
  });

  test('should go to next page when clicking Next', () => {
    let currentPageIdx = 0;
    const handlePageChanged = (pageIdx: number) => {
      currentPageIdx = pageIdx;
    };

    const { container } = render(
      <PageControl totalPages={7} onPageChanged={handlePageChanged} />
    );

    const pageNumbers = container.querySelectorAll('.page-number');
    fireEvent.click(pageNumbers[3]);
    expect(currentPageIdx).toBe(3);

    const nextButton = container.querySelector('.next');
    expect(nextButton).toBeTruthy();
    expect(nextButton!.textContent).toContain('Next');
    fireEvent.click(nextButton!);
    expect(currentPageIdx).toBe(4);
  });

  test('should do nothing when on last page and when clicking Next', () => {
    let currentPageIdx = 0;
    const handlePageChanged = (pageIdx: number) => {
      currentPageIdx = pageIdx;
    };

    const { container } = render(
      <PageControl totalPages={7} onPageChanged={handlePageChanged} />
    );

    const pageNumbers = container.querySelectorAll('.page-number');
    fireEvent.click(pageNumbers[6]);
    expect(currentPageIdx).toBe(6);

    const nextButton = container.querySelector('.next');
    expect(nextButton).toBeTruthy();
    expect(nextButton!.textContent).toContain('Next');
    fireEvent.click(nextButton!);
    expect(currentPageIdx).toBe(6);
  });

  test('should fallback to last page if changing to page greater than total number of pages', () => {
    fail('Not implemented');
  });

  test('should go to correct page after total number of pages changes', () => {
    fail('Not implemented');
  });
});
