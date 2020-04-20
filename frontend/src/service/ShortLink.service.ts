import { Url } from '../entity/Url';
import { Err, ErrorService } from './Error.service';
import { AuthService } from './Auth.service';
import { IApiService } from './Api.service';

export interface IPagedShortLinks {
  shortLinks: Url[];
  totalCount: number;
}

export class ShortLinkService {
  constructor(
    private apiService: IApiService,
    private authService: AuthService,
    private errorService: ErrorService
  ) {}

  getUserCreatedShortLinks(
    offset: number,
    pageSize: number
  ): Promise<IPagedShortLinks> {
    return new Promise((resolve, reject) => {
      this.apiService
        .invokeGetUserShortLinksApi(offset, pageSize)
        .then(URLs => {
          resolve(this.getPagedShortLinksFromURLs(URLs, offset, pageSize));
        })
        .catch(errCode => {
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

  private getPagedShortLinksFromURLs(
    urls: Url[],
    offset: number,
    pageSize: number
  ): IPagedShortLinks {
    // TODO(issue#673): support pagination for user created Short Links in API.
    return {
      shortLinks: urls.slice(offset, offset + pageSize),
      totalCount: urls.length
    };
  }
}
