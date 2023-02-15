import React, {Component} from "react";
import axios from "axios";
import ChatBubble from "./ChatBubble";

import "./Chat.css";
import MultiLineInput from "./MultiLineInput";

class Chat extends Component {
    DEFAULT_CARD = {
        contents: "Add something to the log!",
        timestamp: new Date().toISOString()
    };

    constructor(props) {
        super(props);
        this.inputRef = React.createRef();
        this.bottomFocusRef = React.createRef();
        this.state = {logs: []}
    }

    onSubmit = async (event) => {
        event.preventDefault();
        const contents = this.inputRef.current.value();
        if (contents.trim() === "") return;

        const newCard = await this.props.onNewCard(contents);
        if (newCard !== null) {
            const logs = this.state.logs;
            logs.push(newCard);
            this.setState({logs: logs});
            this.inputRef.current.reset();
        }
    }

    scrollToBottom = (behavior) => {
        this.bottomFocusRef.current.scrollIntoView({behavior: behavior});
    }

    render() {
        return <div id="chat-container">
            <div id="chat-pane">
                {
                    this.state.logs.length ?
                        this.state.logs.map(card => {
                            return <ChatBubble key={card.id} card={card}/>
                        })
                        :
                        <ChatBubble card={this.DEFAULT_CARD}/>
                }
                <div ref={this.bottomFocusRef}/>
            </div>
            <div id="chat-input-container">
                <form id="chat-input-form" onSubmit={this.onSubmit}>
                    <MultiLineInput
                        id="chat-input-text"
                        ref={this.inputRef}
                        placeholder="Type a message"
                        onSubmit={this.onSubmit}
                    />
                    <button id="chat-send-btn"><img src="/icons/send-plane-48.png" alt="send-icon"/></button>
                </form>
            </div>
        </div>;
    }

    async componentDidMount() {
        const resp = await axios.get("/api/rohit/logs");
        if (resp.status === 200) {
            this.setState({logs: resp.data});
        }

        this.scrollToBottom("auto");
    }

    componentDidUpdate() {
        this.scrollToBottom("smooth");
    }
}

export default Chat;