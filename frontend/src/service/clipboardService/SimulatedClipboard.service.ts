import { IClipboardService } from './Clipboard.service';

export class SimulatedClipboardService implements IClipboardService {
  private setupTextAreaWithText(text: string): HTMLTextAreaElement {
    let textArea = this.createInvisibleTextArea();
    textArea.value = text;
    document.body.appendChild(textArea);

    return textArea;
  }

  private createInvisibleTextArea(): HTMLTextAreaElement {
    let textArea = document.createElement('textarea');

    textArea.style.position = 'fixed';
    textArea.style.top = '0';
    textArea.style.left = '0';

    textArea.style.width = '2em';
    textArea.style.height = '2em';

    textArea.style.padding = '0';
    textArea.style.border = 'none';
    textArea.style.outline = 'none';
    textArea.style.boxShadow = 'none';
    textArea.style.background = 'transparent';

    return textArea;
  }

  private selectTextInTextArea(textArea: HTMLTextAreaElement): void {
    textArea.focus();
    textArea.select();
  }

  private tearDownTextArea(textArea: HTMLTextAreaElement): void {
    document.body.removeChild(textArea);
  }

  private copySelectedText(): boolean {
    try {
      return document.execCommand('copy');
    } catch (err) {
      return false;
    }
  }

  copyTextToClipboard(text: string): Promise<void> {
    let textArea = this.setupTextAreaWithText(text);

    this.selectTextInTextArea(textArea);
    let isSuccessful = this.copySelectedText();

    this.tearDownTextArea(textArea);

    if (isSuccessful) {
      return Promise.resolve();
    }
    return Promise.reject();
  }
}
