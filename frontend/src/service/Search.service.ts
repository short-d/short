import { Url } from '../entity/Url';
import { Err } from './Error.service';

export class SearchService {
  getAutoCompleteSuggestions(alias: String): Promise<Array<Url>> {
    return new Promise(async (resolve, reject) => {
      resolve(await this.invokeSearchShortLinkApi(alias));
    });
  }

  private async invokeSearchShortLinkApi(alias: String): Promise<Array<Url>> {
    return new Promise<Array<Url>>(
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
