import React from 'react';
import { NotFoundPage } from '../pages/NotFoundPage';

export default function(
  WrappedComponent: React.ComponentType<any>,
  makeAuthDecision: Promise<boolean>
) {
  enum AuthDecision {
    AUTHORIZED,
    UNAUTHORIZED,
    PENDING
  }

  interface IState {
    decision: AuthDecision;
  }

  return class extends React.Component<any, IState> {
    private isComponentMounted: boolean;

    constructor(props: any) {
      super(props);
      this.state = {
        decision: AuthDecision.PENDING
      };
      this.isComponentMounted = false;
    }

    componentDidMount(): void {
      this.isComponentMounted = true;

      makeAuthDecision.then(decision => {
        if (!this.isComponentMounted) {
          return;
        }
        this.setState({ decision: this.toAuthDecision(decision) });
      });
    }

    componentWillUnmount(): void {
      this.isComponentMounted = false;
    }

    render() {
      const { decision } = this.state;
      if (decision === AuthDecision.AUTHORIZED) {
        return <WrappedComponent {...this.props} />;
      }

      if (decision === AuthDecision.UNAUTHORIZED) {
        return <NotFoundPage />;
      }

      // TODO(issue#816): add loading page component
      return <div />;
    }

    private toAuthDecision(featureDecision: boolean): AuthDecision {
      if (!featureDecision) {
        return AuthDecision.UNAUTHORIZED;
      }
      return AuthDecision.AUTHORIZED;
    }
  };
}
