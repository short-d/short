import React from 'react';
import { Icon, IconID } from '../../frontend/src/component/ui/Icon';
import { action } from '@storybook/addon-actions';
import { withInfo } from '@storybook/addon-info';

export default {
  title: 'UI/Icon',
  component: <Icon iconID={IconID.Menu} />,
  decorators: [withInfo({ header: false, inline: true })]
}

const iconStyle = {
  width: '24px',
  height: '24px'
};

export const menu = () => {
  return (
    <div style={iconStyle}>
      <Icon iconID={IconID.Menu} onClick={action('click')}/>
    </div>
  )
}
export const menuOpen = () => {
  return (
    <div style={iconStyle}>
      <Icon iconID={IconID.MenuOpen} onClick={action('click')}/>
    </div>
  )
}
export const close = () => {
  return (
    <div style={iconStyle}>
      <Icon iconID={IconID.Close} onClick={action('click')}/>
    </div>
  )
}
export const search = () => {
  return (
    <div style={iconStyle}>
      <Icon iconID={IconID.Search} onClick={action('click')}/>
    </div>
  )
}
export const edit = () => {
  return (
    <div style={iconStyle}>
      <Icon iconID={IconID.Edit} onClick={action('click')}/>
    </div>
  )
}
export const check = () => {
  return (
    <div style={iconStyle}>
      <Icon iconID={IconID.Check} onClick={action('click')}/>
    </div>
  )
}
export const deleteIcon = () => {
  return (
    <div style={iconStyle}>
      <Icon iconID={IconID.Delete} onClick={action('click')}/>
    </div>
  )
}
export const rightArrow = () => {
  return (
    <div style={iconStyle}>
      <Icon iconID={IconID.RightArrow} onClick={action('click')}/>
    </div>
  )
}
