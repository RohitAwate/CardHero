import {Component} from "react";

import "./FolderTreeView.css";
import Folder from "./Folder";

class FolderTreeView extends Component {
    render() {
        return <div id="folder-tree-view">
            {
                this.props.folders.map(folder => {
                    return <Folder offset={0} indent={this.props.indent} folder={folder}/>
                })
            }
        </div>;
    }
}

export default FolderTreeView;