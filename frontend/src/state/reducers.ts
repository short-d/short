import { Url } from '../entity/Url';
import { IErr } from '../entity/Err';
import {
  CLEAR_ERROR,
  IPayloadAction,
  RAISE_CREATE_SHORT_LINK_ERROR,
  RAISE_INPUT_ERROR,
  UPDATE_ALIAS,
  UPDATE_CREATED_URL,
  UPDATE_LONG_LINK
} from './actions';
import { Reducer } from 'redux';

export interface IAppState {
  editingUrl: Url;
  createdUrl?: Url;
  qrCodeUrl?: string;
  err?: IErr;
  inputErr?: string;
}

export const initialAppState = {
  editingUrl: {
    originalUrl: '',
    alias: ''
  }
};

export const reducers: Reducer<IAppState> = (
  state: IAppState = initialAppState,
  action: IPayloadAction
): IAppState => {
  switch (action.type) {
    case UPDATE_LONG_LINK:
      return {
        ...state,
        editingUrl: {
          ...state.editingUrl,
          originalUrl: action.payload
        }
      };
    case UPDATE_ALIAS:
      return {
        ...state,
        editingUrl: {
          ...state.editingUrl,
          alias: action.payload
        }
      };
    case RAISE_INPUT_ERROR:
      return {
        ...state,
        inputErr: action.payload
      };
    case RAISE_CREATE_SHORT_LINK_ERROR:
      return {
        ...state,
        err: action.payload
      };
    case UPDATE_CREATED_URL:
      return {
        ...state,
        createdUrl: action.payload
      };
    case CLEAR_ERROR:
      return {
        ...state,
        err: undefined
      };
    default:
      return state;
  }
};
