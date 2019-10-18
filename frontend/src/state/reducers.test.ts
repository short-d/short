import {createStore} from 'redux';
import {initialAppState, reducers} from './reducers';
import {updateAlias, updateLongLink} from './actions';

describe('createStore', () => {
  test('initializes app state', () => {
    let store = createStore(reducers);
    expect(store.getState()).toBe(initialAppState);
  });
});

describe('UPDATE_LONG_URL', () => {
  test('updates originalUrl', () => {
    let store = createStore(reducers);
    let appState = store.getState();
    expect(appState.editingUrl.originalUrl).toBe('');

    store.dispatch(updateLongLink('http://www.example.com'))
  });
});

describe('UPDATE_ALIAS', () => {
  test('updates alias', () => {
    let store = createStore(reducers);
    let appState = store.getState();
    expect(appState.editingUrl.alias).toBe('');

    store.dispatch(updateAlias('eg'))
  });
});