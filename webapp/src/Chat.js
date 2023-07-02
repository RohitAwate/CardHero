import React, {Component} from "react";
import axios from "axios";
import ChatBubble from "./ChatBubble";

import "./Chat.css";
import "./Common.css";
import MultiLineInput from "./MultiLineInput";

class Chat extends Component {
    DEFAULT_CARD = {
        contents: "Add something to the log!",
        timestamp: new Date().toISOString()
    };

    state = {
        logs: []
    }

    inputRef = React.createRef();
    bottomFocusRef = React.createRef();

    createNewCard = async (contents) => {
        const payload = new URLSearchParams();
        payload.append('contents', contents);
        payload.append('timestamp', new Date().toISOString());

        const resp = await axios.post("/api/rohit/logs", payload);
        if (resp.status === 201) {
            return resp.data;
        }
    }

    onSubmit = async (event) => {
        event.preventDefault();
        const contents = this.inputRef.current.value();
        if (contents.trim() === "") return;

        const newCard = await this.createNewCard(contents);
        if (newCard !== null) {
            const logs = this.state.logs;
            logs.push(newCard);
            this.setState({logs: logs});

            this.inputRef.current.reset();
            this.props.updateApp();
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
                    <button className="transparent-btn"><img src="/icons/send-plane-48.png" alt="send-icon"/>
                    </button>
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