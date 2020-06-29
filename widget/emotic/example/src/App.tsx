import React, { Component } from 'react'

import { Emotic, EmotionType, Feedback } from 'emotic'
import 'emotic/dist/index.css'

import './App.scss'

interface IState {
  showDiscount: boolean
  showRefund: boolean
}

export class App extends Component<any, IState> {
  constructor(props: any) {
    super(props)

    this.state = {
      showDiscount: false,
      showRefund: false
    }
  }

  render() {
    return (
      <div className={'App'}>
        {this.state.showDiscount && (
          <div className={'message'}>
            Hey, we saw you are not satisfied with our service.
            <br />
            We highly value you as our customer.
            <br />
            <br />
            We want to keep you here. How about having your next month's
            subscription for <a href={'/next-month-free'}>free</a>?
            <br />
            That's on us!
          </div>
        )}
        {this.state.showRefund && (
          <div className={'message'}>
            We are very sorry to hear that you are frustrated with our service.
            <br />
            <br />
            We will carefully review your feedback and let you know immediately
            when the problem get fixed.
            <br />
            <br />
            Your trust is our top priority.
            <br />
            <br />
            In case you want to <a href={'/refund'}>cancel</a> the service, we
            can provide you the full refund.
          </div>
        )}
        <Emotic onFeedbackFiled={this.handleOnFeedbackFiled} />
      </div>
    )
  }

  handleOnFeedbackFiled = (feedback: Feedback) => {
    this.setState({
      showDiscount: false,
      showRefund: false
    })
    switch (feedback.emotion) {
      case EmotionType.Hate:
        this.showDiscount()
        break
      case EmotionType.Terrible:
        this.showRefund()
        break
    }
  }

  showDiscount = () => {
    this.setState({
      showDiscount: true
    })
  }

  showRefund = () => {
    this.setState({
      showRefund: true
    })
  }
}

export default App
