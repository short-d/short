import { Url } from '../entity/Url';
import { ErrorService, Err } from './Error.service';
import { EnvService } from './Env.service';
import { AuthService } from './Auth.service';
import { CaptchaService, SEARCH_SHORT_LINK } from './Captcha.service';

export class SearchService {
  constructor(
    private authService: AuthService,
    private envService: EnvService,
    private errorService: ErrorService,
    private captchaService: CaptchaService
  ) {}
  searchShortLink(alias: String): Promise<Array<Url>> {
    return new Promise(async (resolve, reject) => {
      try {
        resolve(await this.invokeSearchShortLinkApi(alias));
      } catch (errCodes) {
        const errCode = errCodes[0];
        if (errCode === Err.Unauthorized) {
          reject({
            authorizationErr: 'Unauthorized to create short link'
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

  private async invokeSearchShortLinkApi(alias: String): Promise<Array<Url>> {
    const captchaResponse = await this.captchaService.execute(
      SEARCH_SHORT_LINK
    );
    return new Promise<Array<Url>>(
      (resolve, reject: (errCodes: Err[]) => any) => {
        if (!alias) {
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
