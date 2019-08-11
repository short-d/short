import React, {Component} from "react";
import {Notice} from "../ui/Notice";

import './ExtPromo.scss';

export class ExtPromo  extends Component{
    render() {
        return (
            <Notice>
                <div className={'ext-promo'}>
                    Type less with <a target={'_blank'}
                                      title={'Get s/'}
                                      href={'https://github.com/byliuyang/short-ext'}>s/</a>.
                </div>
            </Notice>
        )
    }
}