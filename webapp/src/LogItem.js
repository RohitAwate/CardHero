import {Component} from "react";

import "./LogItem.css";

class LogItem extends Component {
    render() {
        const card = this.props.card;
        const formattedTime = LogItem.renderTimestamp(card.timestamp);

        return <div key={card.id} className="chat-item">
            <p className="chat-item-contents">{card.contents}</p>
            <div className="chat-item-timestamp">{formattedTime}</div>
        </div>;
    }

    static renderTimestamp(timestamp) {
        const date = new Date(timestamp);
        return date.toLocaleTimeString([], {hour: "2-digit", minute: "2-digit"});
    }
}

export default LogItem;