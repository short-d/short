export interface IHTTPService {
  postJSON<Data>(URL: string, body: any): Promise<Data>;
  getJSON<Data>(URL: string, headers?: any): Promise<Data>;
  get(URL: string, headers?: any): Promise<Body>;
}

export class FetchHTTPService implements IHTTPService {
  postJSON<Data>(url: string, body: any): Promise<Data> {
    return fetch(url, {
      method: 'post',
      headers: {
        'Content-type': 'application/json; charset=UTF-8'
      },
      body: JSON.stringify(body)
    }).then((response) => response.json());
  }

  getJSON<Data>(url: string, headers?: any): Promise<Data> {
    return this.get(url, headers).then((response) => response.json());
  }

  get(url: string, headers?: any): Promise<Body> {
    return fetch(url, { method: 'get', headers: headers });
  }
}
