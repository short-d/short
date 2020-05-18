import { Err } from './Error.service';
import { IGraphQLError, IGraphQLRequestError } from './GraphQL.service';

export function getErrorCodes(errors: IGraphQLRequestError): string[] {
  const { networkError, graphQLErrors } = errors;
  if (networkError) {
    return [Err.NetworkError];
  }
  if (!graphQLErrors || graphQLErrors.length === 0) {
    return [Err.Unknown];
  }
  return graphQLErrors.map(gqlErrorToCode);
}

function gqlErrorToCode(graphQLError: IGraphQLError): string {
  switch (graphQLError.extensions.code) {
    case Err.Unauthenticated:
      return Err.Unauthenticated;
    default:
      return Err.Unknown;
  }
}
