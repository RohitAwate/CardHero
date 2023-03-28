import {Component} from "react";

import "./Gallery.css";
import CardModal from "./CardModal";
import Loader from "./Loader";
import Card from "./Card";

class Gallery extends Component {
    state = {cardModal: null};

    onCardClick = (card) => {
        this.setState({cardModal: card})
    }

    onModalExit = () => {
        this.setState({cardModal: null})
    }

    render() {
        return <div id="gallery">
            {
                this.state.cardModal ? <CardModal card={this.state.cardModal} folderPath={this.props.folderPath} onExit={this.onModalExit} /> : ""
            }
            {
                this.props.cards.length > 0 ?
                    this.props.cards.map(card => <Card onCardClick={this.onCardClick} key={card.id} card={card}/>)
                    : <Loader/>
            }
        </div>;
    }
}

export default Gallery;