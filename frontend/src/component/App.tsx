import React, { Component } from 'react';
import {
  BrowserRouter as Router,
  Redirect,
  Route,
  Switch
} from 'react-router-dom';

import { ReCaptcha } from '../service/Captcha.service';
import { Home } from './pages/Home';
import { Page404 } from './pages/Page404';
import { UrlService } from '../service/Url.service';

interface Props {
  reCaptcha: ReCaptcha;
}

export class App extends Component<Props> {
  urlService = new UrlService();
  render = () => {
    return (
      <Router>
        <Switch>
          <Route
            path={'/'}
            exact
            render={({ location }) => (
              <Home
                urlService={this.urlService}
                location={location}
                reCaptcha={this.props.reCaptcha}
              />
            )}
          />
          <Route
            path={'/r/:alias'}
            render={({ match }) => {
              let alias = match.params['alias'];
              window.location.href = this.urlService.aliasToLink(alias);
              return <div />;
            }}
          />
          <Route component={Page404} />
        </Switch>
      </Router>
    );
  };
}
