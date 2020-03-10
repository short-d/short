import { IBrowserExtensionService } from './BrowserExtension.service';
import { EnvService } from '../Env.service';

export class ChromeExtensionService implements IBrowserExtensionService {
  private EXTENSION_ID_ENV_KEY: string = 'CHROME_EXTENSION_ID';
  private PING_MESSAGE_TYPE: string = 'PING';

  constructor(private envService: EnvService) {}

  isSupported(): boolean {
    return true;
  }

  isInstalled(): Promise<boolean> {
    return new Promise(resolve => {
      chrome.runtime.sendMessage(
        this.envService.getVal(this.EXTENSION_ID_ENV_KEY),
        { message: this.PING_MESSAGE_TYPE },
        response => {
          return resolve(response !== undefined && response !== null);
        }
      );
    });
  }
}
