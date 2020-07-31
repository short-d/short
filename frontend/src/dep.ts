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
import { SearchService } from './service/Search.service';
import { BrowserExtensionFactory } from './service/extensionService/BrowserExtension.factory';
import { ChangeLogService } from './service/ChangeLog.service';
import { ClipboardServiceFactory } from './service/clipboardService/Clipboard.service.factory';
import { GraphQLService } from './service/GraphQL.service';
import { FetchHTTPService } from './service/HTTP.service';
import { ShortHTTPApi } from './service/ShortHTTP.api';
import { DynamicDecisionService } from './service/feature-decision/DynamicDecision.service';
import { ShortLinkService } from './service/ShortLink.service';
import { AnalyticsService } from './service/Analytics.service';
import { ChangeLogGraphQLApi } from './service/shortGraphQL/ChangeLogGraphQL.api';
import { ShortLinkGraphQLApi } from './service/shortGraphQL/ShortLinkGraphQL.api';

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
  const shortHTTPApi = new ShortHTTPApi(authService, httpService, envService);
  const dynamicDecisionService = new DynamicDecisionService(shortHTTPApi);

  const graphQLService = new GraphQLService(httpService);
  const versionService = new VersionService(envService);
  const store = initStore();
  const searchService = new SearchService();
  const changeLogGraphQLApi = new ChangeLogGraphQLApi(
    authService,
    envService,
    captchaService,
    graphQLService
  );
  const changeLogService = new ChangeLogService(
    changeLogGraphQLApi,
    errorService
  );
  const extensionService = new BrowserExtensionFactory().makeBrowserExtensionService(
    envService
  );
  const clipboardService = new ClipboardServiceFactory().makeClipboardService();
  const shortLinkGraphQLApi = new ShortLinkGraphQLApi(
    authService,
    envService,
    graphQLService,
    captchaService
  );
  const shortLinkService = new ShortLinkService(
    authService,
    envService,
    captchaService,
    graphQLService,
    shortLinkGraphQLApi,
    errorService
  );
  const analyticsService = new AnalyticsService(shortHTTPApi);

  return new UIFactory(
    authService,
    clipboardService,
    extensionService,
    qrCodeService,
    versionService,
    errorService,
    searchService,
    changeLogService,
    store,
    dynamicDecisionService,
    shortLinkService,
    analyticsService
  );
}

export function initStore(): Store<IAppState> {
  return createStore(reducers, initialAppState);
}
