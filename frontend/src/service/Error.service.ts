import { IErr } from '../entity/Err';

export enum ErrUrl {
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
              Could not get a response from the server. Please ensure you are
              connected to the Internet and try again. If the problem persists,
              please email byliuyang11@gmail.com with screenshots or logs.
              `
};

export class ErrorService {
  getErr(errCode: ErrUrl): IErr {
    switch (errCode) {
      case ErrUrl.AliasAlreadyExist:
        return aliasNotAvailableErr;
      case ErrUrl.UserNotHuman:
        return userNotHumanErr;
      case ErrUrl.NetworkError:
        return networkErr;
      case ErrUrl.Unknown:
      default:
        return unknownErr;
    }
  }
}
