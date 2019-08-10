import React, {Component} from 'react';

import './Modal.scss';
import classNames from 'classnames';

interface Props {
}

interface State {
    isOpen: boolean
    isShowing: boolean
}

const transitionDuration = 300;

export class Modal extends Component<Props, State> {
    constructor(props: Props) {
        super(props);

        this.state = {
            isOpen: false,
            isShowing: false
        };
    }

    open() {
        this.setState({
            isOpen: true
        });

        setTimeout(() =>
            this.setState({
                isShowing: true
            }), 10);
    }

    close() {
        this.setState({
            isShowing: false
        });

        setTimeout(() =>
            this.setState({
                isOpen: false
            }), transitionDuration);

    }

    handleOnMaskClick = () => {
        this.close();
    };

    render() {
        return (
            <div className={classNames('modal', {'shown': this.state.isOpen, 'showing': this.state.isShowing})} style={{
                transitionDuration: `${transitionDuration}ms`
            }}>
                <div className={'card'}>
                    {this.props.children}
                </div>
                <div className={'mask'} onClick={this.handleOnMaskClick}/>
            </div>
        );
    }
}