import React from 'react';

export default function(
  WrappedComponent: React.ComponentType<any>,
  makeFeatureDecision: () => Promise<boolean>
): React.ComponentType<any> {
  interface IState {
    isFeatureEnabled: boolean;
  }

  return class extends React.Component<any, IState> {
    private isComponentMounted: boolean;

    constructor(props: any) {
      super(props);
      this.state = {
        isFeatureEnabled: false
      };
      this.isComponentMounted = false;
    }

    componentDidMount(): void {
      this.isComponentMounted = true;

      makeFeatureDecision().then(decision => {
        if (!this.isComponentMounted) {
          return;
        }
        this.setState({ isFeatureEnabled: decision });
      });
    }

    componentWillUnmount(): void {
      this.isComponentMounted = false;
    }

    render() {
      const { isFeatureEnabled } = this.state;
      if (!isFeatureEnabled) {
        return <div />;
      }
      return <WrappedComponent {...this.props} />;
    }
  };
}
