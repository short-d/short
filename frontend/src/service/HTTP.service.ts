export interface IHTTPService {
  postJSON<Data>(URL: string, body: any): Promise<Data>;
  getJSON<Data>(URL: string): Promise<Data>;
  get(URL: string): Promise<Body>;
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

  getJSON<Data>(url: string): Promise<Data> {
    return this.get(url).then(response => response.json());
  }

  get(url: string): Promise<Body> {
    return fetch(url, { method: 'get' });
  }
}
