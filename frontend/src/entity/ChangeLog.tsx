import { Change } from './Change';

export interface ChangeLog {
  changes: Change[];
  lastViewedAt?: Date;
}
