import { Change } from './Change';

export interface ChangeLog {
  changes: Change[];
  lastViewedAt: string;
}
