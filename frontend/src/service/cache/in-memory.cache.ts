import { Cache } from './cache';

export class InMemoryCache implements Cache {
  private readonly buffer: { [key: string]: any };

  constructor() {
    this.buffer = {};
  }

  get<Data>(key: string): Data {
    return this.buffer[key];
  }

  set<Data>(key: string, value: Data): void {
    this.buffer[key] = value;
  }
}
