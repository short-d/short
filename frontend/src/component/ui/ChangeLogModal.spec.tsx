import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { ChangeLogModal } from './ChangeLogModal';

it('renders without crashing', () => {
  render(<ChangeLogModal />);
});

it('expands changelog when clicked on "View All Updates"', () => {
  const changeLog = [
    {
      title: 'Lorem ipsum',
      releasedAt: 1500000000003,
      summary: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit'
    },
    {
      title: 'Lorem ipsum',
      releasedAt: 1500000000002,
      summary: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit'
    },
    {
      title: 'Lorem ipsum',
      releasedAt: 1500000000001,
      summary: 'Lorem ipsum dolor sit amet, consectetuer adipiscing elit'
    }
  ];
  const defaultVisibleLogs = 2;
  const { getByText, queryAllByText, container } = render(
    <ChangeLogModal
      changeLog={changeLog}
      defaultVisibleLogs={defaultVisibleLogs}
    />
  );

  expect(queryAllByText('View All Updates').length).toBe(1);
  expect(container.getElementsByTagName('li').length).toBe(defaultVisibleLogs);

  fireEvent.click(getByText('View All Updates'));

  expect(queryAllByText('View All Updates').length).toBe(0);
  expect(container.getElementsByTagName('li').length).toBe(changeLog.length);
});

it('opens correctly', () => {
  const changeLogModalRef = React.createRef<ChangeLogModal>();
  render(<ChangeLogModal ref={changeLogModalRef} />);
  expect(changeLogModalRef).toBeTruthy();
  expect(changeLogModalRef.current).toBeTruthy();
  changeLogModalRef.current!.open();
});

it('closes correctly', () => {
  const changeLogModalRef = React.createRef<ChangeLogModal>();
  render(<ChangeLogModal ref={changeLogModalRef} />);
  expect(changeLogModalRef).toBeTruthy();
  expect(changeLogModalRef.current).toBeTruthy();
  changeLogModalRef.current!.open();
  changeLogModalRef.current!.close();
});