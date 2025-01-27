
import React, { Component } from 'react'

import "./ChatHistory.css"

import Message from '../Message/Message'
class ChatHistory extends Component {

  render() {
    console.log("Chat History", this.props.chatHistory);
    const messages = this.props.chatHistory.map(msg => <Message key={msg.timeStamp} message={msg.data} />);

    return (
      <div className='ChatHistory'>
        <h2>Chat History</h2>
        {messages.length > 0 ? messages : <p>No messages yet...</p>}
      </div>
    );
  }
}

export default ChatHistory
