export interface IHTTPService {
  postJSON(URL: string, body: any): Promise<any>;
}

export class FetchHTTPService implements IHTTPService {
  postJSON(url: string, body: any): Promise<any> {
    return fetch(url, {
      method: 'post',
      headers: {
        'Content-type': 'application/json; charset=UTF-8'
      },
      body: JSON.stringify(body)
    }).then(response => response.json());
  }
}
