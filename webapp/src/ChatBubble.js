import {Component} from "react";

import "./ChatBubble.css";

class ChatBubble extends Component {
    render() {
        const card = this.props.card;
        const formattedTime = ChatBubble.renderTimestamp(card.timestamp);

        return <div key={card.id} className="chat-bubble">
            <p className="chat-bubble-contents">{card.contents}</p>
            <div className="chat-bubble-timestamp">{formattedTime}</div>
        </div>;
    }

    static renderTimestamp(timestamp) {
        const date = new Date(timestamp);
        return date.toLocaleTimeString([], {hour: "2-digit", minute: "2-digit"});
    }
}

export default ChatBubble;