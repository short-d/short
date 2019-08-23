import { CookieService } from './Cookie.service';

export class AuthService {
  static getAuthToken(): string {
    return CookieService.read('token');
  }
}
