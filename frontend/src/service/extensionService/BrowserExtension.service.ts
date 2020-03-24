export interface IBrowserExtensionService {
  isSupported(): boolean;
  isInstalled(): Promise<boolean>;
}
