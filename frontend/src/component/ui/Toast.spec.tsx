import React from 'react';
import { render } from '@testing-library/react';
import { Toast } from './Toast';

describe('Toast component', () => {
  beforeAll(() => {
    jest.useFakeTimers();
  });

  test('should render without crash', () => {
    render(<Toast toastMessage={'Toast message.'} />);
  });

  test('should show content correctly when triggered to show', () => {
    const toastRef = React.createRef<Toast>();
    const { container } = render(
      <Toast ref={toastRef} toastMessage={'Toast message.'} />
    );

    expect(container.textContent).not.toContain('Toast message.');
    toastRef.current!.notify(1000);
    expect(container.textContent).toContain('Toast message.');
  });

  test('should automatically hide content after delay', () => {
    const toastRef = React.createRef<Toast>();
    const { container } = render(
      <Toast ref={toastRef} toastMessage={'Toast message.'} />
    );

    expect(container.textContent).not.toContain('Toast message.');
    toastRef.current!.notify(2000);

    jest.advanceTimersByTime(1000);
    expect(container.textContent).toContain('Toast message.');

    jest.advanceTimersByTime(1000);
    expect(container.textContent).not.toContain('Toast message.');

    jest.clearAllTimers();
  });

  test('second notify call should replace first toast', () => {
    const toastRef = React.createRef<Toast>();
    const { container } = render(
      <Toast ref={toastRef} toastMessage={'Toast message.'} />
    );

    expect(container.textContent).not.toContain('Toast message.');
    toastRef.current!.notify(2000);

    jest.advanceTimersByTime(1000);
    // second notify before the first one closes(at 1000ms)
    toastRef.current!.notify(2000);
    expect(container.textContent).toContain('Toast message.');

    jest.advanceTimersByTime(1500);
    // time is currently 2500ms
    expect(container.textContent).toContain('Toast message.');

    jest.advanceTimersByTime(1000);
    // time is currently 3500ms
    expect(container.textContent).not.toContain('Toast message.');

    jest.clearAllTimers();
  });
});
