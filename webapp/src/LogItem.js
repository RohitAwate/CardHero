import {Component} from "react";

import "./LogItem.css";

class LogItem extends Component {
    render() {
        const card = this.props.card;
        const formattedTime = LogItem.renderTimestamp(card.timestamp);

        return <div key={card.id} className="log-item">
            <p className="log-item-contents">{card.contents}</p>
            <div className="log-item-timestamp">{formattedTime}</div>
        </div>;
    }

    static renderTimestamp(timestamp) {
        const date = new Date(timestamp);
        return date.toLocaleTimeString([], {hour: "2-digit", minute: "2-digit"});
    }
}

export default LogItem;