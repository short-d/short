export interface IHTTPService {
  postJSON<Data>(URL: string, body: any): Promise<Data>;
  get<Data>(URL: string): Promise<Data>;
}

export class FetchHTTPService implements IHTTPService {
  postJSON<Data>(url: string, body: any): Promise<Data> {
    return fetch(url, {
      method: 'post',
      headers: {
        'Content-type': 'application/json; charset=UTF-8'
      },
      body: JSON.stringify(body)
    }).then(response => response.json());
  }

  get<Data>(url: string): Promise<Data> {
    return fetch(url, { method: 'get' }).then(response => response.json());
  }
}
