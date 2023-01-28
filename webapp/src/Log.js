import {Component} from "react";
import axios from "axios";
import LogItem from "./LogItem";

import "./Log.css";

class Log extends Component {
    render() {
        if (this.state) {
            return <div id="log-container">
                <div id="log-pane">
                    {
                        this.state.logs.map(card => {
                            return <LogItem card={card}/>
                        })
                    }
                </div>
                <div id="log-input-container">
                    <form id="log-input-form">
                        <div contentEditable="true" id="log-input-text" placeholder="Type a message"></div>
                        <button id="log-send-btn"><img src="/icons/send-plane-48.png"  alt="send-icon"/></button>
                    </form>
                </div>
            </div>;
        } else {
            return <h1>Add something to the log!</h1>;
        }
    }

    async componentDidMount() {
        const resp = await axios.get("/api/rohit/logs");
        this.setState({logs: resp.data});
    }
}

export default Log;