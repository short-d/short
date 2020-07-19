import React from 'react';
import { Toggle } from '../../frontend/src/component/ui/Toggle';
import { action } from '@storybook/addon-actions';
import { withInfo } from '@storybook/addon-info';

export default {
  title: 'UI/Toggle',
  component: <Toggle />,
  decorators: [withInfo({ header: false, inline: true })]
};

export const pink = () => {
  return <Toggle defaultIsEnabled={false} onClick={action('click')}></Toggle>;
};
