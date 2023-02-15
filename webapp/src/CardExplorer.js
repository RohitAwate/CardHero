import {Component} from "react";
import LeftPane from "./LeftPane";
import CardGallery from "./CardGallery";

import "./CardExplorer.css"
import TopBar from "./TopBar";

class CardExplorer extends Component {
    render() {
        return <div id="card-explorer">
            <TopBar/>
            <div id="card-explorer-main">
                <LeftPane/>
                <CardGallery/>
            </div>
        </div>;
    }
}

export default CardExplorer;