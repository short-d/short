export interface Cache {
  get<Data>(key: string): Data;
  set<Data>(key: string, value: Data): void;
}
