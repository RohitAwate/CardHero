import {Component} from "react";

import "./Gallery.css";
import Card from "./Card";

class Gallery extends Component {
    render() {
        return <div id="gallery">
            {this.props.cards.map(card => <Card card={card}/>)}
        </div>;
    }
}

export default Gallery;