import {Component} from "react";

import "./Gallery.css";
import Loader from "./Loader";
import Card from "./Card";
import axios from "axios";
import {useSearchParams} from "react-router-dom";
import CardModal from "./CardModal";

function Gallery(props) {
    const [searchParams, setSearchParams] = useSearchParams();
    return <GalleryMeta selectedFolder={props.selectedFolder} modalCardID={searchParams.get("card")}/>;
}

class GalleryMeta extends Component {
    state = {
        cards: []
    };

    async refresh() {
        const resp = await axios.get(`/api/rohit/cards/${this.props.selectedFolder}`);
        if (resp.status === 200) {
            const cards = resp.data;
            this.setState({cards});
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
            this.setState({cards: []});
            await this.refresh()
        }
    }

    render() {
        return <div id="gallery">
            {
                this.state.cards.length > 0 ?
                    this.state.cards.map(card => <Card selectedFolder={this.props.selectedFolder} key={card.id}
                                                       card={card}/>)
                    : <Loader/>
            }
            {
                this.props.modalCardID !== null ? <CardModal cardID={this.props.modalCardID}/> : ""
            }
        </div>;
    }
}

export default Gallery;