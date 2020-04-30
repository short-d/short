import { ShortHTTPApi } from './ShortHTTP.api';

export class AnalyticsService {
  constructor(private shortHTTPApi: ShortHTTPApi) {}

  track(event: string) {
    this.shortHTTPApi.trackEvent(event).then();
  }
}
