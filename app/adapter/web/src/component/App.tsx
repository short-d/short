import React, { Component } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';

import { ReCaptcha } from '../service/Captcha.service';
import { Home } from './pages/Home';
import { EnvService } from '../service/Env.service';
import { Playground } from '@apollographql/graphql-playground-react';

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
        <Route
          path={'/api/graphql'}
          render={() => (
            <Playground
              endpoint={`${EnvService.getVal('GRAPHQL_API_BASE_URL')}/graphql`}
            />
          )}
        />
        }/>
      </Router>
    );
  };
}
