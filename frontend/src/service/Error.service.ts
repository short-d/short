
import {Err} from '../entity/Err';

export enum ErrUrl {
  AliasAlreadyExist = 'aliasAlreadyExist',
  UserNotHuman = 'requesterNotHuman',
  Unauthorized = 'invalidAuthToken'
}

export class ErrorService {
  getErr(errCode: ErrUrl): Err {
    switch (errCode) {
      case ErrUrl.AliasAlreadyExist:
        return {
          name: 'Alias not available',
          description: `
                The alias you choose is not available, please choose a different one. 
                Leaving custom alias field empty will automatically generate a available alias.
                `
        };
      case ErrUrl.UserNotHuman:
        return {
          name: 'User not human',
          description: `
                The algorithm thinks you are an automated script instead of human user.
                Please contact byliuyang11@gmail.com if this is a mistake.
                `
        };
      default:
        return {
          name: 'Unknown error',
          description: `
                I am not aware of this error. 
                Please email byliuyang11@gmail.com the screenshots and detailed steps to reproduce it so that I can investigate.
                `
        };
    }
  }
}