import { IClipboardService } from './Clipboard.service';
import { NativeClipboardService } from './NativeClipboard.service';
import { SimulatedClipboardService } from './SimulatedClipboard.service';

export class ClipboardServiceFactory {
  public makeClipboardService(): IClipboardService {
    if (navigator.clipboard) {
      return new NativeClipboardService();
    }
    return new SimulatedClipboardService();
  }
}
