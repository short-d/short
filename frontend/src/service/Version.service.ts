import { EnvService } from './Env.service';

export class VersionService {
  static getAppVersion(): string {
    return EnvService.getVal('GIT_SHA');
  }
}
