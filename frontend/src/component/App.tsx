import React, { Component } from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';

import { ShortLinkService } from '../service/ShortLink.service';
import { UIFactory } from './UIFactory';
import { NotFoundPage } from './pages/NotFoundPage';

interface IProps {
  shortLinkService: ShortLinkService;
  uiFactory: UIFactory;
}

export class App extends Component<IProps> {
  render = () => {
    return (
      <Router>
        <Switch>
          <Route
            path={'/'}
            exact
            render={({ location, history }) =>
              this.props.uiFactory.createHomePage(location, history)
            }
          />
          <Route
            path={'/admin'}
            exact
            render={() => {
              return this.props.uiFactory.createAdminPage();
            }}
          />
          <Route
            path={'/r/:alias'}
            render={({ match }) => {
              let alias = match.params['alias'];
              window.location.href = this.props.shortLinkService.aliasToBackendLink(
                alias
              );
              return <div />;
            }}
          />
          <Route component={NotFoundPage} />
        </Switch>
      </Router>
    );
  };
}
