export class CookieService {
  static read(key: string): string {
    const pattern = new RegExp(`${key}\\s*=\\s*([^;]*)`);
    const matches = document.cookie.match(pattern);
    if (!matches) return '';
    if (matches && matches.length < 2) {
      return '';
    }
    return matches[1];
  }
}
