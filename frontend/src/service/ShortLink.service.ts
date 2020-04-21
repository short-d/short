import { Url } from '../entity/Url';
import { Err, ErrorService } from './Error.service';
import { AuthService } from './Auth.service';
import { ApiService } from './Api.service';

export interface IPagedShortLinks {
  shortLinks: Url[];
  totalCount: number;
}

export class ShortLinkService {
  constructor(
    private apiService: ApiService,
    private authService: AuthService,
    private errorService: ErrorService
  ) {}

  getUserCreatedShortLinks(
    offset: number,
    pageSize: number
  ): Promise<IPagedShortLinks> {
    return new Promise((resolve, reject) => {
      // TODO(issue#673): support pagination for user created Short Links in API.
      this.apiService
        .invokeGetUserShortLinksApi(offset, pageSize)
        .then(URLs => {
          resolve(this.getPagedShortLinks(URLs, offset, pageSize));
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

  private getPagedShortLinks(
    urls: Url[],
    offset: number,
    pageSize: number
  ): IPagedShortLinks {
    return {
      shortLinks: urls.slice(offset, offset + pageSize),
      totalCount: urls.length
    };
  }
}
