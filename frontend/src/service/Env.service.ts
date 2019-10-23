export class EnvService {
  getVal(name: string): string {
    return process.env[`REACT_APP_${name}`] || '';
  }
}
