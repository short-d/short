import React from 'react';
import { render } from '@testing-library/react';
import { Toast } from './Toast';

describe('toast', () => {
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
    toastRef.current!.show();
    jest.runAllTimers();
    expect(container.textContent).toContain('Toast message.');
  });

  test('should hide content correctly when triggered to hide', () => {
    const toastRef = React.createRef<Toast>();
    const { container } = render(
      <Toast ref={toastRef} toastMessage={'Toast message.'} />
    );

    toastRef.current!.show();
    jest.runAllTimers();
    expect(container.textContent).toContain('Toast message.');

    toastRef.current!.hide();
    jest.runAllTimers();
    expect(container.textContent).not.toContain('Toast message.');
  });

  test('should automatically hide content after delay', () => {
    const toastRef = React.createRef<Toast>();
    const { container } = render(
      <Toast ref={toastRef} toastMessage={'Toast message.'} />
    );

    expect(container.textContent).not.toContain('Toast message.');
    toastRef.current!.showAndHide(2000);

    jest.advanceTimersByTime(1000);
    expect(container.textContent).toContain('Toast message.');

    jest.advanceTimersByTime(1000);
    expect(container.textContent).not.toContain('Toast message.');

    jest.clearAllTimers();
  });

  test('second showAndHide call should not affect first toast', () => {
    const toastRef = React.createRef<Toast>();
    const { container } = render(
      <Toast ref={toastRef} toastMessage={'Toast message.'} />
    );

    expect(container.textContent).not.toContain('Toast message.');
    toastRef.current!.showAndHide(2000);

    jest.advanceTimersByTime(1000);
    // second showAndHide trigger before the first one closes
    toastRef.current!.showAndHide(3000);
    expect(container.textContent).toContain('Toast message.');

    jest.advanceTimersByTime(1000);
    expect(container.textContent).not.toContain('Toast message.');

    jest.clearAllTimers();
  });
});
