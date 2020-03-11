import { Update } from '../entity/Update';
import { CookieService } from './Cookie.service';

export class UpdatesService {
  constructor(private cookieService: CookieService) {}

  updateLastSeenChangeLog = () => {
    this.cookieService.set(
      'whatsNewSettings.lastSeenDate',
      Date.now().toString()
    );
  };

  getLastSeenChangeLog = () => {
    return parseInt(
      this.cookieService.get('whatsNewSettings.lastSeenDate') || '0'
    );
  };

  getChangeLog(): Promise<Array<Update>> {
    return new Promise(async (resolve, reject) => {
      resolve(await this.invokeUpdatesApi());
    });
  }

  private async invokeUpdatesApi(): Promise<Array<Update>> {
    return new Promise<Array<Update>>((resolve, reject) => {
      resolve([
        {
          title: 'Added public url toggle',
          releasedAt: 1583587586801,
          summary: 'It is now possible to make links public or private'
        },
        {
          title: 'Added search bar',
          releasedAt: 1583587586800,
          summary: 'Added search bar to the header'
        },
        {
          title: 'Added sign out button',
          releasedAt: 1583562845043,
          summary: 'Added sign out button to the header'
        },
        {
          title: 'Added sign out button',
          releasedAt: 1583562845042,
          summary: 'Added sign out button to the header'
        }
      ]);
    });
  }
}
