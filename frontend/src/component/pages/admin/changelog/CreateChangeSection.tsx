import React, { Component } from 'react';
import './CreateChangeSection.scss';
import { ChangeLogService } from '../../../../service/ChangeLog.service';
import { Store } from 'redux';
import { IAppState } from '../../../../state/reducers';
import { TextField } from '../../../form/TextField';
import { TextAreaField } from '../../../form/TextAreaField';
import { MarkdownViewer } from '../../../ui/MarkdownViewer';
import { Button } from '../../../ui/Button';
import { raiseCreateChangeError } from '../../../../state/actions';

interface IProps {
  changeLogService: ChangeLogService;
  onChangeCreated: () => void;
  onAuthenticationFailed: () => void;
  store: Store<IAppState>;
}

interface IState {
  title: string;
  summary: string;
}

export class CreateChangeSection extends Component<IProps, IState> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      title: '',
      summary: ''
    };
  }

  render() {
    return (
      <div className={'create-change-section'}>
        <div className={'title'}>Create New Change</div>
        <div className={'form'}>
          <div>
            <div className={'label'}>Title</div>
            <TextField onChange={this.handleOnTitleChange} />
          </div>
          <div className={'summary-input'}>
            <div className={'label'}>Summary</div>
            <div className={'change-summary'}>
              <div className={'summary-textarea'}>
                <TextAreaField onChange={this.handleOnSummaryChange} />
              </div>
              <div className={'summary-markdown'}>
                <MarkdownViewer markdown={this.state.summary} />
              </div>
            </div>
          </div>
          <div className={'publish'}>
            <div className={'publish-button'}>
              <Button
                styles={['green', 'full-width', 'full-height']}
                onClick={this.handleOnPublishClick}
              >
                Publish
              </Button>
            </div>
          </div>
        </div>
      </div>
    );
  }

  private handleOnTitleChange = (title: string) => {
    this.setState({ title });
  };

  private handleOnSummaryChange = (summary: string) => {
    this.setState({ summary });
  };

  private handleOnPublishClick = () => {
    const { title, summary } = this.state;

    this.props.changeLogService
      .createChange(title, summary)
      .then(_ => this.props.onChangeCreated())
      .catch(({ authenticationErr, changeErr }) => {
        if (authenticationErr) {
          this.props.onAuthenticationFailed();
        }
        this.props.store.dispatch(raiseCreateChangeError(changeErr));
      });
  };
}
