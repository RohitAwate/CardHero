import React, {Component} from "react";
import axios from "axios";
import LogItem from "./LogItem";

import "./Log.css";
import MultiLineInputField from "./MultiLineInputField";

class Log extends Component {
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

        const payload = new URLSearchParams();
        payload.append('contents', contents);
        payload.append('timestamp', new Date().toISOString());

        const resp = await axios.post("/api/rohit/logs", payload);
        if (resp.status === 201) {
            const newCard = resp.data;
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
        return <div id="log-container">
            <div id="log-pane">
                {
                    this.state.logs.length ?
                        this.state.logs.map(card => {
                            return <LogItem key={card.id} card={card}/>
                        })
                        :
                        <LogItem card={this.DEFAULT_CARD}/>
                }
                <div ref={this.bottomFocusRef}/>
            </div>
            <div id="log-input-container">
                <form id="log-input-form" onSubmit={this.onSubmit}>
                    <MultiLineInputField
                        id="log-input-text"
                        ref={this.inputRef}
                        placeholder="Type a message"
                    />
                    <button id="log-send-btn"><img src="/icons/send-plane-48.png" alt="send-icon"/></button>
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

export default Log;