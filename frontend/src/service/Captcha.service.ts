import {load} from 'recaptcha-v3';
import {EnvService} from './Env.service';

export interface ReCaptcha {
  execute: (action: string) => Promise<string>;
}

export class CaptchaService {
  constructor(private envService: EnvService) {
  }

  public initRecaptchaV3(): Promise<ReCaptcha> {
    return load(
      this
        .envService
        .getVal('RECAPTCHA_SITE_KEY')
    );
  }
}
