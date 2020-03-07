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
          title: 'Added search bar',
          publishedAt: 1583587586800,
          excerpt: 'Added search bar to the header'
        },
        {
          title: 'Added sign out button',
          publishedAt: 1583562845043,
          excerpt: 'Added sign out button to the header'
        },
        {
          title: 'Added search bar',
          publishedAt: 1583562845042,
          excerpt: 'Added search bar to the header'
        },
        {
          title: 'Added sign out button',
          publishedAt: 1583562845041,
          excerpt: 'Added sign out button to the header'
        },
        {
          title: 'Added search bar',
          publishedAt: 1583562845040,
          excerpt: 'Added search bar to the header'
        },
        {
          title: 'Added sign out button',
          publishedAt: 1583562845039,
          excerpt: 'Added sign out button to the header'
        },
        {
          title: 'Added search bar',
          publishedAt: 1583562845038,
          excerpt: 'Added search bar to the header'
        },
        {
          title: 'Added sign out button',
          publishedAt: 1583562845037,
          excerpt: 'Added sign out button to the header'
        }
      ]);
    });
  }
}
