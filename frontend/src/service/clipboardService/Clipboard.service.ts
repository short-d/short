export interface IClipboardService {
  copyTextToClipboard(text: string): Promise<void>;
}
