import {Component} from "react";

import "./Gallery.css";
import Loader from "./Loader";
import Card from "./Card";
import axios from "axios";
import {useSearchParams} from "react-router-dom";
import CardModal from "./CardModal";
import {EmptyFolder} from "./Folder";

function Gallery(props) {
    const [searchParams] = useSearchParams();
    return <GalleryMeta selectedFolder={props.selectedFolder} modalCardID={searchParams.get("card")}/>;
}

class GalleryMeta extends Component {
    state = {
        // Might seem redundant to store both an array and a map for cards
        // but need to retain the order and also have quick lookup by card ID.
        cards: [],
        loaded: false,
        cardsIndex: {}
    };

    async refresh() {
        const resp = await axios.get(`/api/rohit/cards/${this.props.selectedFolder}`);
        if (resp.status === 200) {
            const cards = resp.data;

            for (const card of cards) {
                this.state.cardsIndex[card.id] = card;
            }

            this.setState({cards, loaded: true});
        }
    }

    async componentDidMount() {
        await this.refresh();
    }

    async componentDidUpdate(prevProps, prevState, snapshot) {
        // Don't re-render if user clicks on one of the cards
        for (const card of this.state.cards) {
            if (prevProps.selectedFolder === card.id) {
                this.setState(prevState);
            }
        }

        if (this.props.selectedFolder !== prevProps.selectedFolder) {
            this.setState({cards: [], loaded: false});
            await this.refresh()
        }
    }

    render() {
        const modalRequested = this.props.modalCardID !== null;
        const modalCard = this.state.cardsIndex[this.props.modalCardID];

        return <div id="gallery">
            {
                this.state.cards.length > 0 ?
                    this.state.cards.map(card => <Card selectedFolder={this.props.selectedFolder} key={card.id}
                                                       card={card}/>)
                    : this.state.loaded ? <EmptyFolder/> : <Loader/>
            }
            {
                modalRequested ? <CardModal card={modalCard}/> : <></>
            }
        </div>;
    }
}

export default Gallery;