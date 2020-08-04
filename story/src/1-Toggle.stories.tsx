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

export const blue = () => {
  return (
    <Toggle
      styles={['blue']}
      defaultIsEnabled={false}
      onClick={action('click')}
    ></Toggle>
  );
};

export const black = () => {
  return (
    <Toggle
      styles={['black']}
      defaultIsEnabled={false}
      onClick={action('click')}
    ></Toggle>
  );
};

export const red = () => {
  return (
    <Toggle
      styles={['red']}
      defaultIsEnabled={false}
      onClick={action('click')}
    ></Toggle>
  );
};

export const green = () => {
  return (
    <Toggle
      styles={['green']}
      defaultIsEnabled={false}
      onClick={action('click')}
    ></Toggle>
  );
};

export const shadow = () => {
  return (
    <Toggle
      styles={['pink', 'shadow']}
      defaultIsEnabled={false}
      onClick={action('click')}
    ></Toggle>
  );
};
