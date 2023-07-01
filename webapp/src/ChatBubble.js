import React, {Component} from "react";

import "./ChatBubble.css";
import "./Common.css";
import axios from "axios";
import {useNavigate} from "react-router-dom";

function ChatBubble(props) {
    const navigate = useNavigate();
    const openCardHandler = async (e) => {
        const resp = await axios.get(`/api/rohit/card/goto/${props.card.id}`);
        if (resp.status === 200) {
            return navigate(resp.data);
        }
    };

    return <div className="chat-bubble-container">
        <ChatBubbleNavigation openCardHandler={openCardHandler}/>
        <ChatBubbleMeta card={props.card}/>
    </div>;
}

function ChatBubbleNavigation(props) {
    return <div onClick={props.openCardHandler} className="chat-bubble-nav">
        <button className="transparent-btn">
            <img src="/icons/external-link-60.png" alt="close-cross-icon"/>
        </button>
    </div>;
}

class ChatBubbleMeta extends Component {
    render() {
        const card = this.props.card;
        const formattedTime = ChatBubbleMeta.renderTimestamp(card.timestamp);

        return <div key={card.id} className="chat-bubble">
            <p className="chat-bubble-contents">{card.contents}</p>
            <div className="timestamp">{formattedTime}</div>
        </div>;
    }

    static renderTimestamp(timestamp) {
        const date = new Date(timestamp);
        return date.toLocaleTimeString([], {hour: "numeric", minute: "2-digit"});
    }
}

export default ChatBubble;