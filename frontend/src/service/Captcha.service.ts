import {load} from 'recaptcha-v3';
import {EnvService} from './Env.service';

export interface ReCaptcha {
  execute: (action: string) => Promise<string>;
}

export const CREATE_SHORT_LINK = 'createShortLink';

export class CaptchaService {
  private reCaptcha?: ReCaptcha;
  constructor(private envService: EnvService) {
  }

  public initRecaptchaV3(): Promise<void> {
    return load(
      this
        .envService
        .getVal('RECAPTCHA_SITE_KEY')
    ).then(reCaptcha => {
      this.reCaptcha = reCaptcha;
    });
  }

  public execute(action: string): Promise<string> {
    if(!this.reCaptcha) {
      return Promise.reject('ReCaptcha is not ready yet');
    }
    return this.reCaptcha.execute(action);
  }
}
