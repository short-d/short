import {createStore} from 'redux';
import {initialAppState, reducers} from './reducers';
import {updateAlias, updateLongLink} from './actions';

describe('createStore', () => {
  test('initializes app state', () => {
    const store = createStore(reducers);
    expect(store.getState()).toBe(initialAppState);
  });
});

describe('UPDATE_LONG_URL', () => {
  test('updates originalUrl', () => {
    const store = createStore(reducers);
    const appState = store.getState();
    expect(appState.editingUrl.originalUrl).toBe('');

    store.subscribe(() => {
      const newState = store.getState();
      expect(newState.editingUrl.originalUrl).toBe('http://www.example.com');
    });
    store.dispatch(updateLongLink('http://www.example.com'));
  });
});

describe('UPDATE_ALIAS', () => {
  test('updates alias', () => {
    const store = createStore(reducers);
    const appState = store.getState();
    expect(appState.editingUrl.alias).toBe('');

    store.subscribe(() => {
      const newState = store.getState();
      expect(newState.editingUrl.alias).toBe('eg');
    });
    store.dispatch(updateAlias('eg'));
  });
});
