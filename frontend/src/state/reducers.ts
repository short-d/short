import {Url} from '../entity/Url';
import {Err} from '../entity/Err';
import {
  CLEAR_ERROR,
  IPayloadAction, RAISE_CREATE_SHORT_LINK_ERROR,
  RAISE_INPUT_ERROR,
  UPDATE_ALIAS, UPDATE_CREATED_URL,
  UPDATE_LONG_LINK
} from './actions';
import {Reducer} from 'redux';

export interface IAppState {
  editingUrl: Url;
  createdUrl?: Url;
  qrCodeUrl?: string;
  err?: Err;
  inputErr?: string;
}

export const initialAppState = {
  editingUrl: {
    originalUrl: '',
    alias: ''
  },
};

export const reducers: Reducer<IAppState> =
  (state: IAppState = initialAppState, action: IPayloadAction): IAppState => {
    switch (action.type) {
      case UPDATE_LONG_LINK:
        return Object.assign({}, state, {
          editingUrl: Object.assign({}, state.editingUrl, {
            originalUrl: action.payload
          })
        });
      case UPDATE_ALIAS:
        return Object.assign({}, state, {
          editingUrl: Object.assign({}, state.editingUrl, {
            alias: action.payload
          })
        });
      case RAISE_INPUT_ERROR:
        return Object.assign({}, state, {
          inputErr: action.payload
        });
      case RAISE_CREATE_SHORT_LINK_ERROR:
        return Object.assign({}, state, {
          err: action.payload
        });
      case UPDATE_CREATED_URL:
        return Object.assign({}, state, {
          createdUrl: action.payload
        });
      case CLEAR_ERROR:
        return Object.assign({}, state, {
          err: null
        });
      default:
        return state;
    }
  };