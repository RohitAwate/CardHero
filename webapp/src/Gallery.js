import {Component} from "react";

import "./Gallery.css";
import Card from "./Card";
import Loader from "./Loader";

class Gallery extends Component {
    render() {
        return <div id="gallery">
            {
                this.props.cards.length > 0 ?
                    this.props.cards.map(card => <Card key={card.id} card={card}/>)
                    : <Loader/>
            }
        </div>;
    }
}

export default Gallery;