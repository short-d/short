import React from 'react';
import { Icon, IconID } from '../../frontend/src/component/ui/Icon';
import { action } from '@storybook/addon-actions';

import styles from './2-Icon.stories.module.scss';

export default {
  title: 'UI/Icon',
  component: Icon
};

export const menu = () => {
  return (
    <div className={styles.icon}>
      <Icon iconID={IconID.Menu} onClick={action('click')} />
    </div>
  );
};

export const menuOpen = () => {
  return (
    <div className={styles.icon}>
      <Icon iconID={IconID.MenuOpen} onClick={action('click')} />
    </div>
  );
};

export const close = () => {
  return (
    <div className={styles.icon}>
      <Icon iconID={IconID.Close} onClick={action('click')} />
    </div>
  );
};

export const search = () => {
  return (
    <div className={styles.icon}>
      <Icon iconID={IconID.Search} onClick={action('click')} />
    </div>
  );
};

export const edit = () => {
  return (
    <div className={styles.icon}>
      <Icon iconID={IconID.Edit} onClick={action('click')} />
    </div>
  );
};
export const check = () => {
  return (
    <div className={styles.icon}>
      <Icon iconID={IconID.Check} onClick={action('click')} />
    </div>
  );
};

export const deleteItem = () => {
  return (
    <div className={styles.icon}>
      <Icon iconID={IconID.Delete} onClick={action('click')} />
    </div>
  );
};

export const rightArrow = () => {
  return (
    <div className={styles.icon}>
      <Icon iconID={IconID.RightArrow} onClick={action('click')} />
    </div>
  );
};
