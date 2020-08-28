import React from 'react';
import { Notice } from '../../frontend/src/component/ui/Notice';
import { text } from '@storybook/addon-knobs';

import styles from './2-Notice.stories.module.scss';

export default {
  title: 'UI/Notice',
  component: Notice
};

export const notice = () => {
  return (
    <Notice>
      <div className={styles.content}>
        {text('Notice Message', 'Welcome to Short')}
      </div>
    </Notice>
  );
};
