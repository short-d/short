import { EnvService } from '../Env.service';
import { IBrowserExtensionService } from './BrowserExtension.service';
import { ChromeExtensionService } from './ChromeExtension.service';
import { DefaultExtensionService } from './DefaultExtension.service';

import { detect } from 'detect-browser';

export enum SupportedBrowsers {
  CHROME = 'chrome'
}

export class BrowserExtensionFactory {
  static createBrowserExtensionService(
    envService: EnvService
  ): IBrowserExtensionService {
    const browser = detect();

    switch (browser && browser.name) {
      case SupportedBrowsers.CHROME:
        return new ChromeExtensionService(envService);
      default:
        return new DefaultExtensionService();
    }
  }
}
