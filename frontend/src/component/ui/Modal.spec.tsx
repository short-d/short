import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { Modal } from './Modal';

jest.useFakeTimers();

it('renders content correctly', () => {
  const { container } = render(<Modal>Modal Content</Modal>);
  expect(container.textContent).toMatch('Modal Content');
});

it('opens correctly', () => {
  const modalRef = React.createRef<Modal>();
  render(<Modal ref={modalRef} canClose={true} />);

  expect(modalRef).toBeTruthy();
  expect(modalRef.current).toBeTruthy();

  expect(modalRef.current!.state.isOpen).toBe(false);
  expect(modalRef.current!.state.isShowing).toBe(false);

  modalRef.current!.open();

  expect(modalRef.current!.state.isOpen).toBe(true);
  jest.runAllTimers();
  expect(modalRef.current!.state.isShowing).toBe(true);
});

it('closes modal when click on close button', () => {
  const modalRef = React.createRef<Modal>();
  render(<Modal ref={modalRef} canClose={true} />);

  expect(modalRef.current!.state.isOpen).toBe(false);
  expect(modalRef.current!.state.isShowing).toBe(false);

  modalRef.current!.open();

  expect(modalRef.current!.state.isOpen).toBe(true);
  jest.runAllTimers();
  expect(modalRef.current!.state.isShowing).toBe(true);

  modalRef.current!.close();

  expect(modalRef.current!.state.isShowing).toBe(false);
  jest.runAllTimers();
  expect(modalRef.current!.state.isOpen).toBe(false);
});

it('closes when click on backdrop', () => {
  const modalRef = React.createRef<Modal>();
  const { container } = render(<Modal ref={modalRef} canClose={true} />);

  expect(modalRef.current!.state.isOpen).toBe(false);
  expect(modalRef.current!.state.isShowing).toBe(false);

  modalRef.current!.open();
  
  expect(modalRef.current!.state.isOpen).toBe(true);
  jest.runAllTimers();
  expect(modalRef.current!.state.isShowing).toBe(true);

  fireEvent.click(container.getElementsByClassName('mask')[0]);

  expect(modalRef.current!.state.isShowing).toBe(false);
  jest.runAllTimers();
  expect(modalRef.current!.state.isOpen).toBe(false);
});

it('will not close when canClose prop is false', () => {
  const modalRef = React.createRef<Modal>();
  const { container } = render(<Modal ref={modalRef} canClose={false} />);

  expect(modalRef.current!.state.isOpen).toBe(false);
  expect(modalRef.current!.state.isShowing).toBe(false);

  modalRef.current!.open();
  
  expect(modalRef.current!.state.isOpen).toBe(true);
  jest.runAllTimers();
  expect(modalRef.current!.state.isShowing).toBe(true);

  fireEvent.click(container.getElementsByClassName('mask')[0]);

  expect(modalRef.current!.state.isOpen).toBe(true);
  expect(modalRef.current!.state.isShowing).toBe(true);
});
