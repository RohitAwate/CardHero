import {Component} from "react";

import "./Card.css";
import "./Common.css";

class Card extends Component {
    render() {
        const card = this.props.card;
        const formattedTime = Card.renderTimestamp(card.timestamp);

        return <div className="card">
            <p className="card-content">{card.contents}</p>
            <div className="timestamp">{formattedTime}</div>
        </div>;
    }

    static renderTimestamp(timestamp) {
        const ts = new Date(timestamp);
        const now = new Date();

        const options = {hour: "numeric", minute: "2-digit"};

        const currentYear = now.getFullYear() !== ts.getFullYear();
        if (currentYear) {
            options["year"] = "numeric";
        }

        const today = now.toDateString() === ts.toDateString();
        if (!today) {
            options["day"] = "numeric";
            options["month"] = "short";
        }

        let tsString = ts.toLocaleTimeString([], options);

        if (today) {
            tsString = "Today, " + tsString;
        }

        return tsString;
    }
}

export default Card;