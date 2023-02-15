import {Component} from "react";
import LeftPane from "./LeftPane";
import Gallery from "./Gallery";

import "./Explorer.css"
import TopBar from "./TopBar";

class Explorer extends Component {
    render() {
        return <div id="card-explorer">
            <TopBar/>
            <div id="card-explorer-main">
                <LeftPane folders={this.props.folders} />
                <Gallery/>
            </div>
        </div>;
    }
}

export default Explorer;