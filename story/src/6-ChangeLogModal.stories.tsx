import React from 'react';
import { ChangeLogModal } from '../../frontend/src/component/ui/ChangeLogModal';

export default {
  title: 'UI/ChangeLogModal',
  component: <ChangeLogModal />,
}


export const pink = () => {
  const changeLog = [
    {
      id: '12345',
      title: 'Lorem ipsum',
      releasedAt: new Date('01 Jan 2010 00:00:00'),
      summaryMarkdown:
        'Lorem ipsum dolor sit amet, consectetuer adipiscing elit'
    },
    {
      id: '12346',
      title: 'Lorem ipsum',
      releasedAt: new Date('01 Jan 2010 00:00:00'),
      summaryMarkdown:
        'Lorem ipsum dolor sit amet, consectetuer adipiscing elit'
    },
    {
      id: '12347',
      title: 'Lorem ipsum',
      releasedAt: new Date('01 Jan 2010 00:00:00'),
      summaryMarkdown:
        'Lorem ipsum dolor sit amet, consectetuer adipiscing elit'
    }
  ];
  return <ChangeLogModal changeLog={changeLog}></ChangeLogModal>
}
