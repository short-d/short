import * as QrCode from 'qrcode';

export class QrCodeService {
  newQrCode(data: string): Promise<string> {
    return new Promise<string>((resolve, reject) => {
      QrCode.toDataURL(data, (err, url) => {
        if (err) {
          reject(err);
          return;
        }
        resolve(url);
      });
    });
  }
}
