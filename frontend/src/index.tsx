import React from 'react';
import ReactDOM from 'react-dom';
import './index.scss';
import { App } from './component/App';
import * as serviceWorker from './serviceWorker';

import { CaptchaService } from './service/Captcha.service';

CaptchaService.InitRecaptchaV3().then(reCaptcha => {
  ReactDOM.render(
    <App reCaptcha={reCaptcha} />,
    document.getElementById('root')
  );
  // If you want your app to work offline and load faster, you can change
  // unregister() to register() below. Note this comes with some pitfalls.
  // Learn more about service workers: https://bit.ly/CRA-PWA
  serviceWorker.unregister();
});
