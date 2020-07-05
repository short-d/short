import React, { Component } from 'react';
import './CreateChangeSection.scss';
import { TextField } from '../../form/TextField';
import { MarkdownViewer } from '../../ui/MarkdownViewer';
import { TextAreaField } from '../../form/TextAreaField';
import { Section } from '../../ui/Section';
import { Button } from '../../ui/Button';
import { ChangeLogService } from '../../../service/ChangeLog.service';
import { raiseCreateChangeError } from '../../../state/actions';
import { Store } from 'redux';
import { IAppState } from '../../../state/reducers';

interface IProps {
  changeLogService: ChangeLogService;
  onChangeCreated?: () => void;
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
      <div className="create-change-section">
        <Section title="Create New Change">
          <div className="create-change-form">
            <TextField onChange={this.handleTitleChange} placeHolder="Title" />
            <div className="change-summary">
              <TextAreaField
                rows={12}
                onChange={this.handleSummaryChange}
                placeholder="Summary"
              />
              <MarkdownViewer markdown={this.state.summary} />
            </div>
            <Button onClick={this.handleOnPublishClick}>Publish</Button>
          </div>
        </Section>
      </div>
    );
  }

  private handleTitleChange = (title: string) => {
    this.setState({ title });
  };

  private handleSummaryChange = (summary: string) => {
    this.setState({ summary });
  };

  private handleOnPublishClick = () => {
    const { title, summary } = this.state;

    this.props.changeLogService
      .createChange(title, summary)
      .then(_ => {
        if (this.props.onChangeCreated) {
          this.props.onChangeCreated();
        }
      })
      .catch(({ authenticationErr, changeErr }) => {
        if (authenticationErr) {
          this.props.onAuthenticationFailed();
        }
        this.props.store.dispatch(raiseCreateChangeError(changeErr));
      });
  };
}
