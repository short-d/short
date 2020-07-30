import { ShortLink } from '../entity/ShortLink';
import { Err } from './Error.service';

export class SearchService {
  getAutoCompleteSuggestions(alias: String): Promise<Array<ShortLink>> {
    return new Promise(async (resolve, reject) => {
      resolve(await this.invokeSearchShortLinkApi(alias));
    });
  }

  private async invokeSearchShortLinkApi(alias: String): Promise<Array<ShortLink>> {
    return new Promise<Array<ShortLink>>(
      (resolve, reject: (errCodes: Err[]) => any) => {
        if (alias === '') {
          resolve([]);
        }
        resolve([
          {
            originalUrl: 'https://www.google.com/',
            alias: 'google'
          },
          {
            originalUrl: 'https://github.com/short-d/short/',
            alias: 'short'
          },
          {
            originalUrl: 'https://developer.mozilla.org/en-US/',
            alias: 'mozilla'
          }
        ]);
      }
    );
  }
}
