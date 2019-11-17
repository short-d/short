import { EnvService } from './Env.service';

export class VersionService {
  constructor(private envService: EnvService) {}

  getAppVersion(): string {
    return this.envService.getVal('GIT_SHA');
  }
}
