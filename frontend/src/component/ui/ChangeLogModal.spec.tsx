import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import { ChangeLogModal } from './ChangeLogModal';

describe('ChangeLogModal', () => {
  beforeAll(() => {
    jest.useFakeTimers();
  });

  test('should render without crash', () => {
    render(<ChangeLogModal />);
  });

  test('should expand changelog when clicked on "View All Updates"', () => {
    const changeLogModalRef = React.createRef<ChangeLogModal>();
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
        ref={changeLogModalRef}
        changeLog={changeLog}
        defaultVisibleLogs={defaultVisibleLogs}
      />
    );

    expect(queryAllByText('View All Updates').length).toBe(0);
    expect(container.getElementsByTagName('li').length).toBe(0);

    changeLogModalRef.current.open();
    jest.runAllTimers();

    expect(queryAllByText('View All Updates').length).toBe(1);
    expect(container.getElementsByTagName('li').length).toBe(
      defaultVisibleLogs
    );

    fireEvent.click(getByText('View All Updates'));

    expect(queryAllByText('View All Updates').length).toBe(0);
    expect(container.getElementsByTagName('li').length).toBe(changeLog.length);
  });

  test('should show content correctly when open', () => {
    const changeLogModalRef = React.createRef<ChangeLogModal>();
    const { container } = render(<ChangeLogModal ref={changeLogModalRef} />);

    expect(container.textContent).not.toContain("Since You've Been Gone");

    changeLogModalRef.current.open();
    jest.runAllTimers();

    expect(container.textContent).toContain("Since You've Been Gone");
  });

  test('should hide content correctly when explicitly closed', () => {
    const changeLogModalRef = React.createRef<ChangeLogModal>();
    const { container } = render(<ChangeLogModal ref={changeLogModalRef} />);

    expect(container.textContent).not.toContain("Since You've Been Gone");

    changeLogModalRef.current.open();
    jest.runAllTimers();

    expect(container.textContent).toContain("Since You've Been Gone");

    changeLogModalRef.current!.close();
    jest.runAllTimers();

    expect(container.textContent).not.toContain("Since You've Been Gone");
  });
});
