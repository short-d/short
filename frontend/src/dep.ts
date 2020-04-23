import { UIFactory } from './component/UIFactory';
import { CaptchaService } from './service/Captcha.service';
import { AuthService } from './service/Auth.service';
import { QrCodeService } from './service/QrCode.service';
import { VersionService } from './service/Version.service';
import { EnvService } from './service/Env.service';
import { CookieService } from './service/Cookie.service';
import { createStore, Store } from 'redux';
import { IAppState, initialAppState, reducers } from './state/reducers';
import { ErrorService } from './service/Error.service';
import { RoutingService } from './service/Routing.service';
import { UrlService } from './service/Url.service';
import { SearchService } from './service/Search.service';
import { BrowserExtensionFactory } from './service/extensionService/BrowserExtension.factory';
import { ChangeLogService } from './service/ChangeLog.service';
import { ClipboardServiceFactory } from './service/clipboardService/Clipboard.service.factory';
import { GraphQLService } from './service/GraphQL.service';
import { FetchHTTPService } from './service/HTTP.service';
import { ShortHTTPApi } from './service/ShortHTTP.api';
import { DynamicDecisionService } from './service/feature-decision/DynamicDecision.service';

export function initEnvService(): EnvService {
  return new EnvService();
}

export function initCaptchaService(envService: EnvService): CaptchaService {
  return new CaptchaService(envService);
}

export function initUIFactory(
  envService: EnvService,
  captchaService: CaptchaService
): UIFactory {
  const cookieService = new CookieService();
  const qrCodeService = new QrCodeService();

  const routingService = new RoutingService();
  const authService = new AuthService(
    cookieService,
    envService,
    routingService
  );
  const errorService = new ErrorService();
  const httpService = new FetchHTTPService();
  const shortHTTPApi = new ShortHTTPApi(httpService, envService);
  const dynamicDecisionService = new DynamicDecisionService(shortHTTPApi);

  const graphQLService = new GraphQLService(httpService);
  const urlService = new UrlService(
    authService,
    envService,
    errorService,
    captchaService,
    graphQLService
  );
  const versionService = new VersionService(envService);
  const store = initStore();
  const searchService = new SearchService();
  const changeLogService = new ChangeLogService();
  const extensionService = new BrowserExtensionFactory().makeBrowserExtensionService(
    envService
  );
  const clipboardService = new ClipboardServiceFactory().makeClipboardService();

  return new UIFactory(
    authService,
    clipboardService,
    extensionService,
    urlService,
    qrCodeService,
    versionService,
    errorService,
    searchService,
    changeLogService,
    store,
    dynamicDecisionService
  );
}

export function initStore(): Store<IAppState> {
  return createStore(reducers, initialAppState);
}
