import { Url } from '../entity/Url';

export interface IPagedShortLinks {
  shortLinks: Url[];
  totalCount: number;
}
