import { Action } from 'redux';
import { IErr } from '../entity/Err';
import { Url } from '../entity/Url';

export interface IPayloadAction extends Action {
  payload?: any;
}

export const UPDATE_LONG_LINK = 'UPDATE_LONG_URL';
export const UPDATE_ALIAS = 'UPDATE_ALIAS';
export const RAISE_INPUT_ERROR = 'RAISE_INPUT_ERROR';
export const RAISE_CREATE_SHORT_LINK_ERROR = 'RAISE_CREATE_SHORT_LINK_ERROR';
export const UPDATE_CREATED_URL = 'UPDATE_CREATED_URL';
export const CLEAR_ERROR = 'CLEAR_ERROR';
export const RAISE_GET_USER_SHORT_LINKS_ERROR =
  'RAISE_GET_USER_SHORT_LINKS_ERROR';

export const updateLongLink = (longLink: string): IPayloadAction => ({
  type: UPDATE_LONG_LINK,
  payload: longLink
});

export const updateAlias = (alias: string): IPayloadAction => ({
  type: UPDATE_ALIAS,
  payload: alias
});

export const raiseInputError = (inputError: string | null): IPayloadAction => ({
  type: RAISE_INPUT_ERROR,
  payload: inputError || ''
});

export const raiseCreateShortLinkError = (err: IErr): IPayloadAction => ({
  type: RAISE_CREATE_SHORT_LINK_ERROR,
  payload: err
});

export const raiseGetUserShortLinksError = (err: IErr): IPayloadAction => ({
  type: RAISE_GET_USER_SHORT_LINKS_ERROR,
  payload: err
});

export const updateCreatedUrl = (url: Url): IPayloadAction => ({
  type: UPDATE_CREATED_URL,
  payload: url
});

export const clearError = (): IPayloadAction => ({
  type: CLEAR_ERROR
});
