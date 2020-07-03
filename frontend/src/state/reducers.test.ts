import { createStore } from 'redux';
import { initialAppState, reducers } from './reducers';

describe('createStore', () => {
  test('initializes app state', () => {
    const store = createStore(reducers);
    expect(store.getState()).toBe(initialAppState);
  });
});
