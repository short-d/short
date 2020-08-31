import React from 'react';
import { Notice } from '../../frontend/src/component/ui/Notice';
import { text } from '@storybook/addon-knobs';

import styles from './3-Notice.stories.module.scss';

export default {
  title: 'UI/Notice',
  component: Notice
};

export const notice = () => {
  return (
    <Notice>
      <div className={styles.content}>
        {text('Message', 'Welcome to Short')}
      </div>
    </Notice>
  );
};
