import React from 'react';
import { render } from '@testing-library/react';
import { Toast } from './Toast';

describe('Toast component', () => {
  beforeAll(() => {
    jest.useFakeTimers();
  });

  test('should render without crash', () => {
    render(<Toast />);
  });

  test('should show content correctly when triggered to show', () => {
    const toastRef = React.createRef<Toast>();
    const toastMessage = 'Toast Message';
    const { container } = render(<Toast ref={toastRef} />);

    expect(container.textContent).not.toContain(toastMessage);
    toastRef.current!.notify(toastMessage, 1000);
    expect(container.textContent).toContain(toastMessage);
  });

  test('should automatically hide content after delay', () => {
    const toastRef = React.createRef<Toast>();
    const toastMessage = 'Toast Message';
    const { container } = render(<Toast ref={toastRef} />);
    const TOTAL_DURATION = 2000;
    const HALF_TIME = 1000;

    expect(container.textContent).not.toContain(toastMessage);
    toastRef.current!.notify(toastMessage, TOTAL_DURATION);

    jest.advanceTimersByTime(HALF_TIME);
    expect(container.textContent).toContain(toastMessage);

    jest.advanceTimersByTime(HALF_TIME);
    expect(container.textContent).not.toContain(toastMessage);

    jest.clearAllTimers();
  });

  test('second notify call should replace first toast', () => {
    const toastRef = React.createRef<Toast>();
    const firstToastMessage = 'First Toast Message';
    const secondToastMessage = 'Second Toast Message';
    const { container } = render(<Toast ref={toastRef} />);
    const TOTAL_DURATION = 2000;
    const HALF_TIME = 1000;

    expect(container.textContent).not.toContain(firstToastMessage);
    toastRef.current!.notify(firstToastMessage, TOTAL_DURATION);

    jest.advanceTimersByTime(HALF_TIME);
    // second notify before the first one closes
    toastRef.current!.notify(secondToastMessage, TOTAL_DURATION);
    expect(container.textContent).not.toContain(firstToastMessage);
    expect(container.textContent).toContain(secondToastMessage);

    jest.advanceTimersByTime(HALF_TIME);
    expect(container.textContent).toContain(secondToastMessage);

    jest.advanceTimersByTime(HALF_TIME);
    expect(container.textContent).not.toContain(secondToastMessage);

    jest.clearAllTimers();
  });
});
