import {Component} from "react";

import "./CardGallery.css";

class Gallery extends Component {
    render() {
        return <div id="card-gallery">
            {
                this.props.cards.map(card => <div>{card.contents}</div>)
            }
        </div>;
    }
}

export default Gallery;