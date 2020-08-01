import { ShortLink } from '../entity/ShortLink';
import { Err } from './Error.service';

export class SearchService {
  getAutoCompleteSuggestions(alias: String): Promise<Array<ShortLink>> {
    return new Promise(async (resolve, reject) => {
      resolve(await this.invokeSearchShortLinkApi(alias));
    });
  }

  private async invokeSearchShortLinkApi(
    alias: String
  ): Promise<Array<ShortLink>> {
    return new Promise<Array<ShortLink>>(
      (resolve, reject: (errCodes: Err[]) => any) => {
        if (alias === '') {
          resolve([]);
        }
        resolve([
          {
            longLink: 'https://www.google.com/',
            alias: 'google'
          },
          {
            longLink: 'https://github.com/short-d/short/',
            alias: 'short'
          },
          {
            longLink: 'https://developer.mozilla.org/en-US/',
            alias: 'mozilla'
          }
        ]);
      }
    );
  }
}
