import { EnvService } from '../Env.service';
import { IBrowserExtensionService } from './BrowserExtension.service';
import { ChromeExtensionService } from './ChromeExtension.service';
import { DefaultExtensionService } from './DefaultExtension.service';

import { detect } from 'detect-browser';

export class BrowserExtensionFactory {
  private static CHROME_BROWSER_NAME: string = 'chrome';

  static createBrowserExtensionService(
    envService: EnvService
  ): IBrowserExtensionService {
    const browser = detect();

    switch (browser && browser.name) {
      case BrowserExtensionFactory.CHROME_BROWSER_NAME:
        return new ChromeExtensionService(envService);
      default:
        return new DefaultExtensionService();
    }
  }
}
