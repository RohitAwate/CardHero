import {Component} from "react";

import "./Gallery.css";
import Card from "./Card";

class Gallery extends Component {
    render() {
        return <div id="gallery">
            {
                this.props.cards.length > 0 ?
                    this.props.cards.map(card => <Card key={card.id} card={card}/>)
                    : <img
                        className="loader"
                        src="/loaders/ellipsis-green.svg"
                        alt="Loading..."
                        style={{alignSelf: "center"}}
                    />
            }
        </div>;
    }
}

export default Gallery;