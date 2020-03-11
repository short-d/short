import { IBrowserExtensionService } from './BrowserExtension.service';

export class DefaultExtensionService implements IBrowserExtensionService {
  isSupported(): boolean {
    return false;
  }

  isInstalled(): Promise<boolean> {
    return Promise.resolve(false);
  }
}
