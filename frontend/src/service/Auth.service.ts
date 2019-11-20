import { EnvService } from './Env.service';
import { CookieService } from './Cookie.service';
import { RoutingService } from './Routing.service';

export class AuthService {
  constructor(
    private cookieService: CookieService,
    private envService: EnvService,
    private routingService: RoutingService
  ) {}

  cacheAuthToken(pageUrl: string) {
    const params = new URLSearchParams(pageUrl);
    const token = params.get('token');
    if (!token || token.length < 1) {
      return;
    }
    this.saveAuthToken(token);
    this.routingService.navigateTo('/');
  }

  saveAuthToken(token: string | null) {
    if (token == null) {
      return;
    }
    this.cookieService.set('token', token);
  }

  getAuthToken(): string {
    return this.cookieService.get('token');
  }

  signOut() {
    this.cookieService.set('token', '');
  }

  isSignedIn(): boolean {
    let token = this.getAuthToken();
    return token.length > 0;
  }

  githubSignInLink(): string {
    return `${this.envService.getVal(
      'HTTP_API_BASE_URL'
    )}/oauth/github/sign-in`;
  }

  googleSignInLink(): string {
    return `${this.envService.getVal(
      'HTTP_API_BASE_URL'
    )}/oauth/google/sign-in`;
  }

  facebookSignInLink(): string {
    return `${this.envService.getVal(
      'HTTP_API_BASE_URL'
    )}/oauth/facebook/sign-in`;
  }
}
