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
    } catch (err) {}

    document.body.removeChild(textArea);
    return successful;
  }

  public copyTextToClipboard(text: string): Promise<void> {
    if (!navigator.clipboard) {
      console.log('other');
      return this.fallbackCopyTextToClipboard(text)
        ? Promise.resolve()
        : Promise.reject();
    }

    return navigator.clipboard.writeText(text);
  }
}
