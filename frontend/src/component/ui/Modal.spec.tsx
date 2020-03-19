import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { Modal } from './Modal';

jest.useFakeTimers();

it('renders without crashing', () => {
  render(<Modal />);
});

it('opens correctly', () => {
  const modalRef = React.createRef<Modal>();
  render(<Modal ref={modalRef} canClose={true} />);

  expect(modalRef).toBeTruthy();
  expect(modalRef.current).toBeTruthy();

  modalRef.current!.open();

  expect(modalRef.current!.state.isOpen).toBe(true);
  jest.runAllTimers();
  expect(modalRef.current!.state.isShowing).toBe(true);
});

it('closes modal when click on close button', () => {
  const modalRef = React.createRef<Modal>();
  render(<Modal ref={modalRef} canClose={true} />);

  modalRef.current!.open();
  jest.runAllTimers();

  modalRef.current!.close();

  expect(modalRef.current!.state.isShowing).toBe(false);
  jest.runAllTimers();
  expect(modalRef.current!.state.isOpen).toBe(false);
});

it('closes when click on backdrop', () => {
  const modalRef = React.createRef<Modal>();
  const { container } = render(<Modal ref={modalRef} canClose={true} />);

  modalRef.current!.open();
  jest.runAllTimers();

  fireEvent.click(container.getElementsByClassName('mask')[0]);

  expect(modalRef.current!.state.isShowing).toBe(false);
  jest.runAllTimers();
  expect(modalRef.current!.state.isOpen).toBe(false);
});

it('will not close when canClose prop is false', () => {
  const modalRef = React.createRef<Modal>();
  const { container } = render(<Modal ref={modalRef} canClose={false} />);

  modalRef.current!.open();
  jest.runAllTimers();

  fireEvent.click(container.getElementsByClassName('mask')[0]);

  expect(modalRef.current!.state.isOpen).toBe(true);
  expect(modalRef.current!.state.isShowing).toBe(true);
});
