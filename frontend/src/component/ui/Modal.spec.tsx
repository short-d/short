import * as React from 'react';
import { Modal } from './Modal';
import { render, fireEvent } from '@testing-library/react';
jest.useFakeTimers();

describe('Modal', () => {
  test('should render without crash', () => {
    render(<Modal>Content</Modal>);
  });

  test('should show content correctly when open', () => {
    const modalRef = React.createRef<Modal>();
    const { container } = render(<Modal ref={modalRef}>Modal Content</Modal>);

    expect(container.textContent).not.toContain('Modal Content');
    modalRef.current.open();
    jest.runAllTimers();
    expect(container.textContent).toContain('Modal Content');
  });

  test('should hide content correctly when explicitly closed', () => {
    const modalRef = React.createRef<Modal>();
    const { container } = render(<Modal ref={modalRef}>Modal Content</Modal>);

    modalRef.current.open();
    jest.runAllTimers();
    expect(container.textContent).toContain('Modal Content');

    modalRef.current.close();
    jest.runAllTimers();
    expect(container.textContent).not.toContain('Modal Content');
  });

  test('should hide content correctly when click on mask', () => {
    const modalRef = React.createRef<Modal>();
    const { container } = render(
      <Modal ref={modalRef} canClose={true}>
        Modal Content
      </Modal>
    );

    modalRef.current.open();
    jest.runAllTimers();
    expect(container.textContent).toContain('Modal Content');

    const mask = container.querySelector('.mask');
    fireEvent.click(mask);
    jest.runAllTimers();
    expect(container.textContent).not.toContain('Modal Content');
  });

  test('should not close when canClose is false', () => {
    const modalRef = React.createRef<Modal>();
    const { container } = render(
      <Modal ref={modalRef} canClose={false}>
        Modal Content
      </Modal>
    );

    modalRef.current.open();
    jest.runAllTimers();
    expect(container.textContent).toContain('Modal Content');

    const mask = container.querySelector('.mask');
    fireEvent.click(mask);
    jest.runAllTimers();
    expect(container.textContent).toContain('Modal Content');
  });
});
