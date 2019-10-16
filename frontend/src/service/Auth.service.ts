import { EnvService } from './Env.service';
import { CookieService } from './Cookie.service';

export class AuthService {
  constructor(
    private cookieService: CookieService,
    private envService: EnvService,
  ) {}
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
    return `${this.envService.getVal('HTTP_API_BASE_URL')}/oauth/github/sign-in`;
  }
}
