import React from 'react';
import { SearchBar } from '../../frontend/src/component/ui/SearchBar';
import { action } from '@storybook/addon-actions';
import { withInfo } from '@storybook/addon-info';
import { withKnobs, object } from '@storybook/addon-knobs';

export default {
  title: 'UI/SearchBar',
  component: <SearchBar onChange={(_) => {}} />,
  decorators: [withKnobs, withInfo({ header: false, inline: true })]
};

export const standard = () => {
  return (
    <SearchBar
      onChange={action('change')}
      autoCompleteSuggestions={object('Search Results', [
        {
          longLink: 'https://www.google.com/',
          alias: 'google'
        }
      ])}
    />
  );
};
