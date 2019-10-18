import {UIFactory} from './component/UIFactory';
import {CaptchaService} from './service/Captcha.service';
import {StaticConfigDecisionService} from './service/StaticConfigDecision.service';
import {UrlService} from './service/Url.service';
import {AuthService} from './service/Auth.service';
import {QrCodeService} from './service/QrCode.service';
import {VersionService} from './service/Version.service';
import {EnvService} from './service/Env.service';
import {CookieService} from './service/Cookie.service';
import {createStore, Store} from 'redux';
import {reducers, IAppState} from './state/reducers';
import {ErrorService} from './service/Error.service';
import {RoutingService} from './service/Routing.service';

export function initEnvService(): EnvService {
  return new EnvService();
}

export function initCaptchaService(
  envService: EnvService
): CaptchaService {
  return new CaptchaService(envService);
}

export function initUIFactory(
  envService: EnvService,
  captchaService: CaptchaService
): UIFactory {
  const cookieService = new CookieService();
  const qrCodeService = new QrCodeService();
  const staticConfigDecision = new StaticConfigDecisionService();

  const routingService = new RoutingService();
  const authService = new AuthService(cookieService, envService, routingService);
  const urlService = new UrlService(authService, envService, captchaService);
  const versionService = new VersionService(envService);
  const errorService = new ErrorService();
  const store = initStore();

  return new UIFactory(
    authService,
    urlService,
    qrCodeService,
    versionService,
    captchaService,
    errorService,
    store,
    staticConfigDecision
  );
}

export function initStore(): Store<IAppState> {
  return createStore(reducers);
}