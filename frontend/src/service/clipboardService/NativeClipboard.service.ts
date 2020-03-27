import { IClipboardService } from './Clipboard.service';

export class NativeClipboardService implements IClipboardService {
  public copyTextToClipboard(text: string): Promise<void> {
    return navigator.clipboard.writeText(text);
  }
}
