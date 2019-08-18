export class EnvService {
  static getVal(name: string): string {
    return process.env[`REACT_APP_${name}`] || '';
  }
}
