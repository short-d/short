import ReactDOM from 'react-dom';
import './index.scss';
import * as serviceWorker from './serviceWorker';
import { initCaptchaService, initEnvService, initUIFactory } from './dep';

const envService = initEnvService();
const captchaService = initCaptchaService(envService);

captchaService.initRecaptchaV3().then(() => {
  const uiFactory = initUIFactory(envService, captchaService);
  ReactDOM.render(uiFactory.createApp(), document.getElementById('root'));
  // If you want your app to work offline and load faster, you can change
  // unregister() to register() below. Note this comes with some pitfalls.
  // Learn more about service workers: https://bit.ly/CRA-PWA
  serviceWorker.unregister();
});
