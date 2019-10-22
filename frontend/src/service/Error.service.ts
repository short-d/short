
import {IErr} from '../entity/Err';

export enum ErrUrl {
  AliasAlreadyExist = 'aliasAlreadyExist',
  UserNotHuman = 'requesterNotHuman',
  Unauthorized = 'invalidAuthToken'
}

const unknownErr = {
  name: 'Unknown error',
  description: `
                I am not aware of this error. 
                Please email byliuyang11@gmail.com the screenshots and detailed 
                steps to reproduce it so that I can investigate.
                `
};

const aliasNotAvailableErr = {
  name: 'Alias not available',
  description: `
                The alias you choose is not available, please choose a 
                different one. Leaving custom alias field empty will automatically 
                generate a available alias.
                `
};

const userNotHumanErr = {
  name: 'User not human',
  description: `
                The algorithm believes you are an automated script instead of a
                human user. Please email byliuyang11@gmail.com if this is a
                mistake.
                `
};

export class ErrorService {
  getErr(errCode: ErrUrl): IErr {
    switch (errCode) {
      case ErrUrl.AliasAlreadyExist:
        return aliasNotAvailableErr;
      case ErrUrl.UserNotHuman:
        return userNotHumanErr;
      default:
        return unknownErr;
    }
  }
}
