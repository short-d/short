import * as Cookies from 'es-cookie';

export class CookieService {
  get(key: string): string {
    return Cookies.get(key) || '';
  }

  set(key: string, value?: string) {
    return Cookies.set(key, value || '');
  }
}
