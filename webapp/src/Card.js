import {Component} from "react";

import "./Card.css";

class Card extends Component {
    render() {
        const card = this.props.card;

        return <div className="card">
            <p>{card.contents}</p>
        </div>;
    }
}

export default Card;