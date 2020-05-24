import { IHTTPService } from './HTTP.service';
import { EnvService } from './Env.service';

export class ShortHTTPApi {
  private readonly baseURL: string;

  constructor(
    private httpService: IHTTPService,
    private envService: EnvService
  ) {
    this.baseURL = this.envService.getVal('HTTP_API_BASE_URL');
  }

  getFeatureToggle(featureID: string): Promise<boolean> {
    const url = `${this.baseURL}/features/${featureID}`;
    return this.httpService.getJSON<boolean>(url);
  }

  trackEvent(event: string): Promise<Body> {
    const url = `${this.baseURL}/analytics/track/${event}`;
    return this.httpService.get(url);
  }
}
