import {Component} from "react";

import "./LeftPane.css"
import FolderTreeView from "./FolderTreeView";


class LeftPane extends Component {
    render() {
        return <div id="left-pane">
            <FolderTreeView
                selectedFolder={this.props.selectedFolder}
                onFolderSelect={this.props.onFolderSelect}
                folders={this.props.folders} indent={15}
            />
        </div>;
    }
}

export default LeftPane;