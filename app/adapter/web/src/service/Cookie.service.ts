import * as Cookies from "es-cookie";

export class CookieService {
    static get(key: string): string {
        return Cookies.get('token') || ''
    }

    static set(key: string, value?: string) {
        return Cookies.set(key, value || '');
    }
}