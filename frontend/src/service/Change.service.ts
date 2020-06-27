import { Err, ErrorService } from './Error.service';
import { ChangeGraphQLApi } from './shortGraphQL/ChangeGraphQL.api';
import { Change } from '../entity/Change';

export class ChangeService {
  constructor(
    private changeGraphQLApi: ChangeGraphQLApi,
    private errorService: ErrorService
  ) {}

  createChange(title: string, summaryMarkdown: string): Promise<Change> {
    return new Promise<Change>((resolve, reject) => {
      this.changeGraphQLApi
        .createChange(title, summaryMarkdown)
        .then(resolve)
        .catch(errCode => {
          if (errCode === Err.Unauthenticated) {
            reject({ authenticationErr: 'User is not authenticated' });
            return;
          }
          reject({ changeErr: this.errorService.getErr(errCode) });
        });
    });
  }
}
