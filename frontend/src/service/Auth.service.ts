import { EnvService } from './Env.service';
import { CookieService } from './Cookie.service';

export class AuthService {
  static getAuthToken(): string {
    return CookieService.get('token');
  }

  static signOut() {
    CookieService.set('token', '');
  }

  static isSignedIn(): boolean {
    let token = AuthService.getAuthToken();
    return token.length > 0;
  }

  static githubSignInLink(): string {
    return `${EnvService.getVal('HTTP_API_BASE_URL')}/oauth/github/sign-in`;
  }
}
