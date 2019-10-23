export class RoutingService {
  navigateTo(pageUrl: string) {
    window.history.replaceState({}, document.title, pageUrl);
  }
}
