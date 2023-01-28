import {Component} from "react";
import axios from "axios";
import LogItem from "./LogItem";

import "./Log.css";

class Log extends Component {
    render() {
        if (this.state) {
            return <div id="log-pane">
                {
                    this.state.logs.map(card => {
                        return <LogItem card={card}/>
                    })
                }
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