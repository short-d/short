import {UIFactory} from './component/UIFactory';
import {CaptchaService, ReCaptcha} from './service/Captcha.service';
import {StaticConfigDecisionService} from './service/StaticConfigDecision.service';
import {UrlService} from './service/Url.service';
import {AuthService} from './service/Auth.service';
import {QrCodeService} from './service/QrCode.service';
import {VersionService} from './service/Version.service';
import {EnvService} from './service/Env.service';
import {CookieService} from './service/Cookie.service';

export function initEnvService(): EnvService {
  return new EnvService();
}

export function initCaptchaService(
  envService: EnvService
): CaptchaService {
  return new CaptchaService(envService);
}

export function initUIFactory(
  reCaptcha: ReCaptcha,
  envService: EnvService
): UIFactory {
  const cookieService = new CookieService();
  const qrCodeService = new QrCodeService();

  const authService = new AuthService(cookieService, envService);
  const urlService = new UrlService(authService, envService);
  const versionService = new VersionService(envService);

  return new UIFactory(
    authService,
    urlService,
    qrCodeService,
    versionService,
    reCaptcha
  );
}