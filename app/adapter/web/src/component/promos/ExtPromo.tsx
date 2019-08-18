import React, {Component} from "react";
import {Notice} from "../ui/Notice";

import './ExtPromo.scss';

export class ExtPromo extends Component {
    render() {
        return (
            <Notice>
                <div className={'ext-promo'}>
                    Type less with&nbsp;
                    <a target={'_blank'}
                       title={'Get s/'}
                       href={'https://s.time4hacks.com/r/shortext'}>s/</a>
                    &nbsp;.
                </div>
            </Notice>
        )
    }
}