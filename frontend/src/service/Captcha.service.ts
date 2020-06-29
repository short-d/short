import { load } from 'recaptcha-v3';
import { EnvService } from './Env.service';
import { Err } from './Error.service';

export interface ReCaptcha {
  execute: (action: string) => Promise<string>;
}

export const CREATE_SHORT_LINK = 'createShortLink';
export const UPDATE_SHORT_LINK = 'updateShortLink';
export const SEARCH_SHORT_LINK = 'searchShortLink';
export const VIEW_CHANGE_LOG = 'viewChangeLog';

export const CREATE_CHANGE = 'createChange';
export const UPDATE_CHANGE = 'updateChange';
export const DELETE_CHANGE = 'deleteChange';

const INVALID_SITE_KEY_ERR_MSG = 'Invalid site key or not loaded in api.js';

export class CaptchaService {
  private reCaptcha?: ReCaptcha;

  constructor(private envService: EnvService) {}

  public initRecaptchaV3(): Promise<void> {
    return load(this.envService.getVal('RECAPTCHA_SITE_KEY')).then(
      reCaptcha => {
        this.reCaptcha = reCaptcha;
      }
    );
  }

  public execute(action: string): Promise<string> {
    return new Promise<string>((resolve, reject) => {
      if (!this.reCaptcha) {
        reject(Err.ReCaptchaNotReady);
        return;
      }

      this.reCaptcha.execute(action).then(resolve, err => {
        const errMsg = err.message;
        if (CaptchaService.contains(errMsg, INVALID_SITE_KEY_ERR_MSG)) {
          reject(Err.InvalidReCaptchaSiteKey);
          return;
        }
        reject(Err.Unknown);
      });
    });
  }

  private static contains(text: string, substr: string): boolean {
    return text.indexOf(substr) > -1;
  }
}
