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
      this.apiService
        .invokeGetUserShortLinksApi(offset, pageSize)
        .then(pagedShortLinks => resolve(pagedShortLinks))
        .catch(errCode => {
          if (errCode === Err.Unauthenticated) {
            reject({ authenticationErr: 'User is not authenticated' });
            return;
          }
          reject({
            getUserShortLinksErr: this.errorService.getErr(errCode)
          });
        });
      return;
    });
  }
}
