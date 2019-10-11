import React, { Component } from 'react';
import {
  BrowserRouter as Router,
  Route,
  Switch,
} from 'react-router-dom';

import { ReCaptcha } from '../service/Captcha.service';
import { Home } from './pages/Home';
import { Page404 } from './pages/Page404';

interface Props {
  reCaptcha: ReCaptcha;
}

export class App extends Component<Props> {
  render = () => {
    return (
      <Router>
        <Switch>
          <Route
            path={'/'}
            exact
            render={({ location }) => (
              <Home location={location} reCaptcha={this.props.reCaptcha} />
            )}
          />
          <Route component={Page404} />
        </Switch>
      </Router>
    );
  };
}
