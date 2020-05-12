import { ChangeLog } from '../entity/ChangeLog';
import { ChangeLogGraphQLApi } from './ChangeLogGraphQL.api';
import { ErrorService, Err } from './Error.service';

export class ChangeLogService {
  constructor(
    private changeLogGraphQLApi: ChangeLogGraphQLApi,
    private errorService: ErrorService
  ) {}

  hasUpdates(): Promise<boolean> {
    return new Promise(async (resolve, reject) => {
      try {
        const changeLog = await this.changeLogGraphQLApi.getChangeLog();
        if (!changeLog.changes) {
          resolve(false);
          return;
        }

        if (!changeLog.changes[0]) {
          resolve(false);
          return;
        }

        if (!changeLog.lastViewedAt) {
          resolve(true);
          return;
        }

        for (let i = 0; i < changeLog.changes.length; i++) {
          if (
            new Date(changeLog.lastViewedAt).getTime() <
            new Date(changeLog.changes[i].releasedAt).getTime()
          ) {
            resolve(true);
            return;
          }
        }

        resolve(false);
      } catch (errCode) {
        if (errCode === Err.Unauthenticated) {
          resolve(false);
          return;
        }

        reject({
          changeLogErr: this.errorService.getErr(errCode)
        });
      }
    });
  }

  getChangeLog(): Promise<ChangeLog> {
    return new Promise(async (resolve, reject) => {
      try {
        const changeLog = await this.changeLogGraphQLApi.getChangeLog();
        resolve(changeLog);
      } catch (errCode) {
        reject({
          changeLogErr: this.errorService.getErr(errCode)
        });
      }
    });
  }
}
