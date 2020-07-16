import { Url } from '../entity/Url';
import { IErr } from '../entity/Err';
import {
  CLEAR_ERROR,
  IPayloadAction,
  RAISE_CREATE_SHORT_LINK_ERROR,
  RAISE_GET_USER_SHORT_LINKS_ERROR,
  RAISE_GET_CHANGELOG_ERROR,
  RAISE_CREATE_CHANGE_ERROR
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
    case RAISE_CREATE_SHORT_LINK_ERROR:
      return {
        ...state,
        err: action.payload
      };
    case RAISE_GET_USER_SHORT_LINKS_ERROR:
      return {
        ...state,
        err: action.payload
      };
    case RAISE_GET_CHANGELOG_ERROR:
      return {
        ...state,
        err: action.payload
      };
    case RAISE_CREATE_CHANGE_ERROR:
      return {
        ...state,
        err: action.payload
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
