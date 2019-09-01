import React, { Component } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';

import { ReCaptcha } from '../service/Captcha.service';
import { Home } from './pages/Home';

interface Props {
  reCaptcha: ReCaptcha;
}

export class App extends Component<Props> {
  render = () => {
    return (
      <Router>
        <Route
          path={'/'}
          exact
          render={() => <Home reCaptcha={this.props.reCaptcha} />}
        />
      </Router>
    );
  };
}
