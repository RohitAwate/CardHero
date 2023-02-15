import {Component} from "react";

import "./LeftPane.css"
import FolderTreeView from "./FolderTreeView";


class LeftPane extends Component {
    render() {
        return <div id="left-pane">
            <FolderTreeView folders={this.props.folders} indent={15}/>
        </div>;
    }
}

export default LeftPane;