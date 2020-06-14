import { Url } from '../entity/Url';
import { Err, ErrorService } from './Error.service';
import { ShortLinkGraphQLApi } from './ShortGraphQLService/ShortLinkGraphQL.api';

export interface IPagedShortLinks {
  shortLinks: Url[];
  offset: number;
  pageSize: number;
  totalCount: number;
}

export class ShortLinkService {
  constructor(
    private shortLinkGraphQLApi: ShortLinkGraphQLApi,
    private errorService: ErrorService
  ) {}

  getUserCreatedShortLinks(
    offset: number,
    pageSize: number
  ): Promise<IPagedShortLinks> {
    return new Promise((resolve, reject) => {
      // TODO(issue#673): support pagination for user created Short Links in API.
      this.shortLinkGraphQLApi
        .getUserShortLinks(offset, pageSize)
        .then((URLs: Url[]) => {
          resolve(this.getPagedShortLinks(URLs, offset, pageSize));
        })
        .catch((errCode: Err) => {
          if (errCode === Err.Unauthenticated) {
            reject({ authenticationErr: 'User is not authenticated' });
            return;
          }
          reject({
            getUserShortLinksErr: this.errorService.getErr(errCode)
          });
        });
    });
  }

  private getPagedShortLinks(
    urls: Url[],
    offset: number,
    pageSize: number
  ): IPagedShortLinks {
    return {
      shortLinks: urls.slice(offset, offset + pageSize),
      offset: offset,
      pageSize: pageSize,
      totalCount: urls.length
    };
  }
}
