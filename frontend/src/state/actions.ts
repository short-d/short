import { Action } from 'redux';
import { IErr } from '../entity/Err';

export interface IPayloadAction extends Action {
  payload?: any;
}

export const RAISE_CREATE_SHORT_LINK_ERROR = 'RAISE_CREATE_SHORT_LINK_ERROR';
export const CLEAR_ERROR = 'CLEAR_ERROR';
export const RAISE_GET_USER_SHORT_LINKS_ERROR =
  'RAISE_GET_USER_SHORT_LINKS_ERROR';
export const RAISE_GET_CHANGELOG_ERROR = 'RAISE_GET_CHANGELOG_ERROR';

export const raiseCreateShortLinkError = (err: IErr): IPayloadAction => ({
  type: RAISE_CREATE_SHORT_LINK_ERROR,
  payload: err
});

export const raiseGetUserShortLinksError = (err: IErr): IPayloadAction => ({
  type: RAISE_GET_USER_SHORT_LINKS_ERROR,
  payload: err
});

export const raiseGetChangeLogError = (err: IErr): IPayloadAction => ({
  type: RAISE_GET_CHANGELOG_ERROR,
  payload: err
});

export const clearError = (): IPayloadAction => ({
  type: CLEAR_ERROR
});
