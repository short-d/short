import React, {Component} from 'react';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom';

import {UrlService} from '../service/Url.service';
import {UIFactory} from './UIFactory';
import {NotFoundPage} from './pages/NotFoundPage';

interface IProps {
  urlService: UrlService;
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
            render={({location}) =>
              this.props.uiFactory.createHomePage(location)
            }
          />
          <Route
            path={'/r/:alias'}
            render={({match}) => {
              let alias = match.params['alias'];
              window.location.href = this.props.urlService.aliasToLink(alias);
              return <div/>;
            }}
          />
          <Route component={NotFoundPage}/>
        </Switch>
      </Router>
    );
  };
}
