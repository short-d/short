export class ClipboardService {
  private fallbackCopyTextToClipboard(text: string): boolean {
    let textArea = document.createElement('textarea');

    // set minimal visual effects in case textarea flashes
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

    textArea.value = text;

    document.body.appendChild(textArea);
    textArea.focus();
    textArea.select();

    let successful = false;
    try {
      successful = document.execCommand('copy');
    } catch (err) {
      console.log(`Failed to copy ${text} into Clipboard`);
    }

    document.body.removeChild(textArea);
    return successful;
  }

  public copyTextToClipboard(text: string): Promise<void> {
    if (navigator.clipboard) {
      return navigator.clipboard.writeText(text);
    }
    const isSuccessful = this.fallbackCopyTextToClipboard(text);
    if (isSuccessful) {
      return Promise.resolve();
    }
    return Promise.reject();
  }
}
