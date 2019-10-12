import { load } from 'recaptcha-v3';
import { EnvService } from './Env.service';

export interface ReCaptcha {
  execute: (action: string) => Promise<string>;
}

export class CaptchaService {
  static InitRecaptchaV3(): Promise<ReCaptcha> {
    return load(EnvService.getVal('RECAPTCHA_SITE_KEY'));
  }
}
