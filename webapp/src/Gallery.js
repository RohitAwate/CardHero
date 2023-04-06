import {Component} from "react";

import "./Gallery.css";
import Loader from "./Loader";
import Card from "./Card";
import axios from "axios";
import {Outlet, Route, Routes} from "react-router-dom";
import CardModal from "./CardModal";

class Gallery extends Component {
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
        if (this.props.selectedFolder !== prevProps.selectedFolder) {
            this.setState({cards: []});
            await this.refresh()
        }
    }

    render() {
        return <div id="gallery">
            <CardsMosaic cards={this.state.cards}/>
            <Routes>
                <Route path={"/cards/:id"} element={<CardModal/>}/>
            </Routes>
        </div>;
    }
}

function CardsMosaic(props) {
    return <>
        {
            props.cards.length > 0 ?
                props.cards.map(card => <Card key={card.id} card={card}/>)
                : <Loader/>
        }
        <Outlet/>
    </>;
}

export default Gallery;