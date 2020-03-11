import { EnvService } from '../Env.service';

export interface IBrowserExtensionService {
  isSupported(): boolean;
  isInstalled(): Promise<boolean>;
}
