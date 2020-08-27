import React, { Component } from 'react';

import styles from './Spinner.module.scss';

export class Spinner extends Component {
  render() {
    return (
      <div className={styles['spinner']}>
        <div className={styles['solar-system']}>
          <div className={`${styles['earth-orbit']} ${styles['orbit']}`}>
            <div className={`${styles['planet']} ${styles['earth']}`} />
            <div className={`${styles['venus-orbit']} ${styles['orbit']}`}>
              <div className={`${styles['planet']} ${styles['venus']}`} />
              <div className={`${styles['mercury-orbit']} ${styles['orbit']}`}>
                <div className={`${styles['planet']} ${styles['mercury']}`} />
                <div className={styles['sun']} />
              </div>
            </div>
          </div>
        </div>
      </div>
    );
  }
}
