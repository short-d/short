import { IErr } from '../entity/Err';

export enum Err {
  AliasAlreadyExist = 'aliasAlreadyExist',
  UserNotHuman = 'requesterNotHuman',
  Unauthorized = 'invalidAuthToken',
  NetworkError = 'networkError',
  Unknown = 'unknownError',
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

const networkErr = {
  name: 'Network error',
  description: `
              Unable to reach the server. Please double check your Internet connection and try again.
              If this happens consistently, please email byliuyang11@gmail.com screenshots and the necessary steps to reproduce the error.
              `
};

export class ErrorService {
  getErr(errCode: Err): IErr {
    switch (errCode) {
      case Err.AliasAlreadyExist:
        return aliasNotAvailableErr;
      case Err.UserNotHuman:
        return userNotHumanErr;
      case Err.NetworkError:
        return networkErr;
      case Err.Unknown:
      default:
        return unknownErr;
    }
  }
}
