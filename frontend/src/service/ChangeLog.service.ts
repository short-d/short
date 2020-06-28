import { ChangeLog } from '../entity/ChangeLog';
import { ChangeLogGraphQLApi } from './shortGraphQL/ChangeLogGraphQL.api';
import { ErrorService, Err } from './Error.service';
import { Change } from '../entity/Change';

export class ChangeLogService {
  constructor(
    private changeLogGraphQLApi: ChangeLogGraphQLApi,
    private errorService: ErrorService
  ) {}

  hasUpdates(): Promise<boolean> {
    return new Promise(async (resolve, reject) => {
      let changeLog;
      try {
        changeLog = await this.changeLogGraphQLApi.getChangeLog();
      } catch (errCode) {
        if (errCode === Err.Unauthenticated) {
          resolve(false);
          return;
        }

        reject({
          changeLogErr: this.errorService.getErr(errCode)
        });
      }

      if (!changeLog || !changeLog.changes || changeLog.changes.length < 1) {
        resolve(false);
        return;
      }

      if (!changeLog.lastViewedAt) {
        resolve(true);
        return;
      }

      changeLog.changes = this.sortChanges(changeLog.changes);
      if (
        changeLog.lastViewedAt.getTime() <
        changeLog.changes[0].releasedAt.getTime()
      ) {
        resolve(true);
        return;
      }

      resolve(false);
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

  viewChangeLog(): Promise<Date> {
    return this.changeLogGraphQLApi.viewChangeLog();
  }

  sortChanges(changes: Change[]) {
    changes.sort((a: Change, b: Change) => {
      return b.releasedAt.getTime() - a.releasedAt.getTime();
    });

    return changes;
  }

  createChange(title: string, summaryMarkdown: string): Promise<Change> {
    return new Promise<Change>((resolve, reject) => {
      this.changeLogGraphQLApi
        .createChange(title, summaryMarkdown)
        .then(resolve)
        .catch(errCode => {
          // TODO(issue#904): impose definite error handling mechanism in client classes.
          if (errCode === Err.Unauthenticated) {
            reject({ authenticationErr: 'User is not authenticated' });
            return;
          }
          reject({ changeErr: this.errorService.getErr(errCode) });
        });
    });
  }

  deleteChange(changeId: string): Promise<string> {
    return new Promise<string>((resolve, reject) => {
      this.changeLogGraphQLApi
        .deleteChange(changeId)
        .then(resolve)
        .catch(errCode => {
          // TODO(issue#904): impose definite error handling mechanism in client classes.
          if (errCode === Err.Unauthenticated) {
            reject({ authenticationErr: 'User is not authenticated' });
            return;
          }
          reject({ changeErr: this.errorService.getErr(errCode) });
        });
    });
  }
}
