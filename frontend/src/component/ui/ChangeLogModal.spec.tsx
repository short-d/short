import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { ChangeLogModal } from './ChangeLogModal';

it('renders without crashing', () => {
  render(<ChangeLogModal />);
});

it('expands changelog when clicked on "View All Updates"', () => {
  let changeLogModalRef = React.createRef<ChangeLogModal>();
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
      ref={changeLogModalRef}
      defaultVisibleLogs={defaultVisibleLogs}
    />
  );
  expect(changeLogModalRef).toBeTruthy();
  expect(changeLogModalRef.current).toBeTruthy();
  if (changeLogModalRef && changeLogModalRef.current) {
    changeLogModalRef.current.open();
    changeLogModalRef.current.close();
  }

  expect(queryAllByText('View All Updates').length).toStrictEqual(1);
  expect(container.getElementsByTagName('li').length).toStrictEqual(
    defaultVisibleLogs
  );

  fireEvent.click(getByText('View All Updates'));

  expect(queryAllByText('View All Updates').length).toStrictEqual(0);
  expect(container.getElementsByTagName('li').length).toStrictEqual(
    changeLog.length
  );
});
