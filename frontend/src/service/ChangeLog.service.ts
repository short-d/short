import { ChangeLog } from '../entity/ChangeLog';
import { ChangeLogGraphQLApi } from './ChangeLogGraphQL.api';
import { ErrorService, Err } from './Error.service';
import { Change } from '../entity/Change';

export class ChangeLogService {
  constructor(
    private changeLogGraphQLApi: ChangeLogGraphQLApi,
    private errorService: ErrorService
  ) {}

  hasUpdates(): Promise<boolean> {
    return new Promise(async (resolve, reject) => {
      try {
        const changeLog = await this.changeLogGraphQLApi.getChangeLog();
        if (!changeLog.changes || changeLog.changes.length < 1) {
          resolve(false);
          return;
        }

        if (!changeLog.lastViewedAt) {
          resolve(true);
          return;
        }

        changeLog.changes = this.sortChanges(changeLog.changes);
        if (
          new Date(changeLog.lastViewedAt).getTime() <
          new Date(changeLog.changes[0].releasedAt).getTime()
        ) {
          resolve(true);
          return;
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
        changeLog.changes = this.sortChanges(changeLog.changes);
        resolve(changeLog);
      } catch (errCode) {
        reject({
          changeLogErr: this.errorService.getErr(errCode)
        });
      }
    });
  }

  sortChanges(changes: Change[]) {
    changes.sort((a: Change, b: Change) => {
      return (
        new Date(b.releasedAt).getTime() - new Date(a.releasedAt).getTime()
      );
    });

    return changes;
  }
}
