export interface Styling {
  styles: string[];
}

export function withCSSModule(styles: string[], cssModuleStyles: any): string {
  return styles.map(style => cssModuleStyles[style]).join(' ');
}
