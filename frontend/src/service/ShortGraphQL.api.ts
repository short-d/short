import { Err } from './Error.service';
import { IGraphQLError, IGraphQLRequestError } from './GraphQL.service';

export class ShortGraphQLApi {
  public getErrorCodes(errors: IGraphQLRequestError): string[] {
    const { networkError, graphQLErrors } = errors;
    if (networkError) {
      return [Err.NetworkError];
    }
    if (!graphQLErrors || graphQLErrors.length === 0) {
      return [Err.Unknown];
    }
    return graphQLErrors.map(this.gqlErrorToCode);
  }

  private gqlErrorToCode(graphQLError: IGraphQLError): string {
    switch (graphQLError.extensions.code) {
      case Err.Unauthenticated:
        return Err.Unauthenticated;
      default:
        return Err.Unknown;
    }
  }
}
