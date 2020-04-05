import React, { ReactText } from 'react';
import { render, fireEvent } from '@testing-library/react';
import { Pagination } from './Pagination';

const range = (from: number, to: number) => {
  const result = [];
  for (let i = from; i <= to; i++) {
    result.push(i);
  }
  return result;
};

const NAVIGATOR_BUTTONS_COUNT = 2;
const PREV_NAV_BUTTON_TEXT = '< Previous';
const NEXT_NAV_BUTTON_TEXT = 'Next >';

describe('Pagination component', () => {
  test('should render without crash', () => {
    render(<Pagination />);
  });

  test('should render without hidden pages when number of pages less than or equal to 5', () => {
    const pageLimit = 5;
    const totalPages = 5;
    const totalRecords = totalPages * pageLimit;

    const { container } = render(
      <Pagination pageLimit={pageLimit} totalRecords={totalRecords} />
    );

    const paginationBlocks = container.querySelectorAll('button');
    expect(paginationBlocks).toHaveLength(totalPages + NAVIGATOR_BUTTONS_COUNT);

    const expectedBlockContents = range(1, totalPages);
    for (let i = 0; i < expectedBlockContents.length; i++) {
      expect(paginationBlocks[i + 1].textContent).toContain(
        expectedBlockContents[i]
      );
    }
  });

  test('should have previous button disabled when in first page', () => {
    const pageLimit = 5;
    const totalPages = 5;
    const totalRecords = totalPages * pageLimit;

    const { getByText } = render(
      <Pagination pageLimit={pageLimit} totalRecords={totalRecords} />
    );

    expect(
      getByText(PREV_NAV_BUTTON_TEXT).hasAttribute('disabled')
    ).toBeTruthy();
  });

  test('should have next button disabled when in last page', () => {
    const pageLimit = 5;
    const totalPages = 5;
    const totalRecords = totalPages * pageLimit;
    const lastPage = totalPages;

    const { getByText } = render(
      <Pagination pageLimit={pageLimit} totalRecords={totalRecords} />
    );
    getByText(lastPage.toString()).click();

    expect(
      getByText(NEXT_NAV_BUTTON_TEXT).hasAttribute('disabled')
    ).toBeTruthy();
  });

  test('should have previous and next button enabled when not in first or last page', () => {
    const pageLimit = 5;
    const totalPages = 5;
    const totalRecords = totalPages * pageLimit;
    const secondPage = 2;

    const { getByText } = render(
      <Pagination pageLimit={pageLimit} totalRecords={totalRecords} />
    );
    getByText(secondPage.toString()).click();

    expect(
      getByText(PREV_NAV_BUTTON_TEXT).hasAttribute('disabled')
    ).toBeFalsy();
    expect(
      getByText(NEXT_NAV_BUTTON_TEXT).hasAttribute('disabled')
    ).toBeFalsy();
  });

  test('should navigate to selected page when clicked on page button', () => {
    const paginationRef = React.createRef<Pagination>();
    const pageLimit = 5;
    const totalPages = 5;
    const totalRecords = totalPages * pageLimit;
    const secondPage = 2;

    const { getByText } = render(
      <Pagination
        ref={paginationRef}
        pageLimit={pageLimit}
        totalRecords={totalRecords}
      />
    );
    getByText(secondPage.toString()).click();

    expect(paginationRef.current!.state.currentPage).toBe(secondPage);
  });

  test('should navigate to next page when clicked on next button', () => {
    const paginationRef = React.createRef<Pagination>();
    const pageLimit = 5;
    const totalPages = 5;
    const totalRecords = totalPages * pageLimit;
    const firstPage = 1;
    const secondPage = 2;

    const { getByText } = render(
      <Pagination
        ref={paginationRef}
        pageLimit={pageLimit}
        totalRecords={totalRecords}
      />
    );

    expect(paginationRef.current!.state.currentPage).toBe(firstPage);

    getByText(NEXT_NAV_BUTTON_TEXT).click();
    expect(paginationRef.current!.state.currentPage).toBe(secondPage);
  });

  test('should navigate to previous page when clicked on previous button', () => {
    const paginationRef = React.createRef<Pagination>();
    const pageLimit = 5;
    const totalPages = 7;
    const totalRecords = totalPages * pageLimit;
    const secondPage = 2;
    const thirdPage = 3;

    const { getByText } = render(
      <Pagination
        ref={paginationRef}
        pageLimit={pageLimit}
        totalRecords={totalRecords}
      />
    );

    getByText(thirdPage.toString()).click();
    expect(paginationRef.current!.state.currentPage).toBe(thirdPage);

    getByText(PREV_NAV_BUTTON_TEXT).click();
    expect(paginationRef.current!.state.currentPage).toBe(secondPage);
  });

  test('should call onPageChanged when navigating to a different page', () => {
    const pageLimit = 5;
    const totalPages = 7;
    const totalRecords = totalPages * pageLimit;
    const thirdPage = 3;
    const onPageChanged = jest.fn();

    const { getByText } = render(
      <Pagination
        pageLimit={pageLimit}
        totalRecords={totalRecords}
        onPageChanged={onPageChanged}
      />
    );
    getByText(thirdPage.toString()).click();
    getByText(PREV_NAV_BUTTON_TEXT).click();
    getByText(NEXT_NAV_BUTTON_TEXT).click();

    expect(onPageChanged).toHaveBeenCalledTimes(3);
  });
});
