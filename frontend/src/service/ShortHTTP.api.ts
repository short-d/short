import { IHTTPService } from './HTTP.service';
import { EnvService } from './Env.service';
import { AuthService } from './Auth.service';

export class ShortHTTPApi {
  private readonly baseURL: string;

  constructor(
    private authService: AuthService,
    private httpService: IHTTPService,
    private envService: EnvService
  ) {
    this.baseURL = this.envService.getVal('HTTP_API_BASE_URL');
  }

  getFeatureToggle(featureID: string): Promise<boolean> {
    const url = `${this.baseURL}/features/${featureID}`;
    const headers = { Authorization: this.authService.getBearerToken() };

    return this.httpService.getJSON<boolean>(url, headers);
  }

  trackEvent(event: string): Promise<Body> {
    const url = `${this.baseURL}/analytics/track/${event}`;
    return this.httpService.get(url);
  }

  isUserAdmin(): Promise<boolean> {
    const url = `${this.baseURL}/features/admin-panel`;
    const headers = { Authorization: this.authService.getBearerToken() };

    return this.httpService.getJSON<boolean>(url, headers);
  }
}
