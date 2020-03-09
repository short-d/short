import { EnvService } from './Env.service';

export class ChromeExtensionService {
  private EXTENSION_ID_ENV_KEY: string = 'CHROME_EXTENSION_ID';
  private PING_MESSAGE_TYPE: string = 'PING';

  constructor(private envService: EnvService) {}

  isExtensionInstalled() {
    return new Promise((resolve, reject) => {
      chrome.runtime.sendMessage(
        this.envService.getVal(this.EXTENSION_ID_ENV_KEY),
        { message: this.PING_MESSAGE_TYPE },
        response => {
          if (response !== undefined && response !== null) resolve();
        }
      );
    });
  }
}
