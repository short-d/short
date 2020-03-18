import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { Modal } from './Modal';

jest.useFakeTimers();

it('renders without crashing', () => {
  render(<Modal />);
});

it('opens and closes correctly', () => {
  const modalRef = React.createRef<Modal>();
  render(<Modal ref={modalRef} canClose={true} />);

  expect(modalRef).toBeTruthy();
  expect(modalRef.current).toBeTruthy();

  modalRef.current!.open();

  expect(modalRef.current!.state.isOpen).toBeTruthy();
  jest.runAllTimers();
  expect(modalRef.current!.state.isShowing).toBeTruthy();

  modalRef.current!.close();

  expect(modalRef.current!.state.isShowing).toBeFalsy();
  jest.runAllTimers();
  expect(modalRef.current!.state.isOpen).toBeFalsy();
});

it('closes when click on backdrop', () => {
  const modalRef = React.createRef<Modal>();
  const { container } = render(<Modal ref={modalRef} canClose={true} />);

  modalRef.current!.open();
  jest.runAllTimers();

  fireEvent.click(container.getElementsByClassName('mask')[0]);

  expect(modalRef.current!.state.isShowing).toBeFalsy();
  jest.runAllTimers();
  expect(modalRef.current!.state.isOpen).toBeFalsy();
});

it('will not close when canClose prop is falsy', () => {
  const modalRef = React.createRef<Modal>();
  const { container } = render(<Modal ref={modalRef} canClose={false} />);

  modalRef.current!.open();
  jest.runAllTimers();

  fireEvent.click(container.getElementsByClassName('mask')[0]);

  expect(modalRef.current!.state.isOpen).toBeTruthy();
  expect(modalRef.current!.state.isShowing).toBeTruthy();
});
