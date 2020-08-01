import {ShortLink} from '../entity/ShortLink';
import {Err, ErrorService} from './Error.service';
import {ShortLinkGraphQLApi} from './shortGraphQL/ShortLinkGraphQL.api';
import {IErr} from '../entity/Err';
import {EnvService} from './Env.service';
import {validateLongLinkFormat} from '../validators/LongLink.validator';
import {validateCustomAliasFormat} from '../validators/CustomAlias.validator';

export interface IPagedShortLinks {
  shortLinks: ShortLink[];
  offset: number;
  pageSize: number;
  totalCount: number;
}

interface ICreateShortLinkErrs {
  authorizationErr?: string;
  createShortLinkErr?: IErr;
}

export class ShortLinkService {
  constructor(
    private envService: EnvService,
    private shortLinkGraphQLApi: ShortLinkGraphQLApi,
    private errorService: ErrorService
  ) {}

  createShortLink(
    editingShortLink: ShortLink,
    isPublic: boolean = false
  ): Promise<ShortLink> {
    return new Promise(async (resolve, reject) => {
      const longLink = editingShortLink.longLink;
      const customAlias = editingShortLink.alias;

      const err = this.validateInputs(longLink, customAlias);
      if (err) {
        reject(err);
        return;
      }

      try {
        const shortLink = await this.shortLinkGraphQLApi.createShortLink(
          editingShortLink,
          isPublic
        );
        resolve(shortLink);
        return;
      } catch (errCode) {
        if (errCode === Err.Unauthenticated) {
          reject({
            authenticationErr: 'User is not authenticated'
          });
          return;
        }

        const error = this.errorService.getErr(errCode);
        reject({
          createShortLinkErr: error
        });
      }
    });
  }

  aliasToFrontendLink(alias: string): string {
    return `${window.location.protocol}//${window.location.host}/r/${alias}`;
  }

  aliasToBackendLink(alias: string): string {
    return `${this.envService.getVal('HTTP_API_BASE_URL')}/r/${alias}`;
  }

  private validateInputs(
    longLink?: string,
    customAlias?: string
  ): ICreateShortLinkErrs | null {
    let err = validateLongLinkFormat(longLink);
    if (err) {
      return {
        createShortLinkErr: {
          name: 'Invalid Long Link',
          description: err
        }
      };
    }

    err = validateCustomAliasFormat(customAlias);
    if (err) {
      return {
        createShortLinkErr: {
          name: 'Invalid Custom Alias',
          description: err
        }
      };
    }
    return null;
  }

  getUserCreatedShortLinks(
    offset: number,
    pageSize: number
  ): Promise<IPagedShortLinks> {
    return new Promise((resolve, reject) => {
      // TODO(issue#673): support pagination for user created Short Links in API.
      this.shortLinkGraphQLApi
        .getUserShortLinks(offset, pageSize)
        .then((shortLinks: ShortLink[]) => {
          resolve(this.getPagedShortLinks(shortLinks, offset, pageSize));
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

  async updateShortLink(oldAlias: string, shortLink: Partial<ShortLink>) {
    return this.shortLinkGraphQLApi.updateShortLink(oldAlias, shortLink);
  }

  private getPagedShortLinks(
    shortLinks: ShortLink[],
    offset: number,
    pageSize: number
  ): IPagedShortLinks {
    return {
      shortLinks: shortLinks.slice(offset, offset + pageSize),
      offset: offset,
      pageSize: pageSize,
      totalCount: shortLinks.length
    };
  }
}
