import {Component} from "react";

import "./FolderTreeView.css";
import Folder from "./Folder";

class FolderTreeView extends Component {
    render() {
        return <div id="folder-tree-view">
            <p id="left-pane-section-label">F O L D E R S</p>
            {
                this.props.folders.map(folder => {
                    return <Folder
                        selectedFolder={this.props.selectedFolder}
                        onFolderSelect={this.props.onFolderSelect}
                        key={folder.id} offset={0} indent={this.props.indent}
                        folder={folder}/>
                })
            }
        </div>;
    }
}

export default FolderTreeView;