import React, {Component} from 'react';

import './Section.scss';

interface Props {
    title: string
}

export class Section extends Component<Props> {
    render() {
        return (
            <div className={'section'}>
                <div className={'center'}>
                    <div className={'title'}>{this.props.title}</div>
                    {this.props.children}
                </div>
            </div>
        );
    }
}