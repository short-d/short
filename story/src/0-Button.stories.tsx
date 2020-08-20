import React from 'react';
import { Button } from '../../frontend/src/component/ui/Button';
import { action } from '@storybook/addon-actions';

import styles from './0-Button.stories.module.scss';

export default {
  title: 'UI/Button',
  component: Button,
  argTypes: {
    label: {
      type: 'text',
      description: 'Button text',
      defaultValue: 'Button',
      table: {
        type: { summary: 'string' },
        defaultValue: { summary: 'Button' }
      },
      control: {
        type: 'text'
      }
    },
    color: {
      defaultValue: 'blue',
      description: 'Button color',
      control: {
        type: 'select',
        options: ['pink', 'blue', 'black', 'red', 'green']
      }
    },
    shadow: {
      defaultValue: false,
      description: 'Button shadow',
      control: {
        type: 'boolean'
      }
    }
  }
};

export const Primary = ({ label, color, shadow, ...args }) => {
  let shadowStyle: string;
  if (shadow) {
    shadowStyle = 'shadow';
  } else {
    shadowStyle = '';
  }
  return (
    <Button {...args} styles={[color, shadowStyle]} onClick={action('click')}>
      {label}
    </Button>
  );
};

export const FullWidth = ({ label, color, shadow, ...args }) => {
  let shadowStyle: string;
  if (shadow) {
    shadowStyle = 'shadow';
  } else {
    shadowStyle = '';
  }
  return (
    <Button
      {...args}
      styles={[color, shadowStyle, 'full-width']}
      onClick={action('click')}
    >
      <div className={styles.content}>{label}</div>
    </Button>
  );
};
