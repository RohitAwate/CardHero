import {Component} from "react";

import "./FolderTreeView.css";
import Folder from "./Folder";
import axios from "axios";

class FolderTreeView extends Component {
    state = {
        folders: []
    }

    async refresh() {
        const resp = await axios.get("/api/rohit/folders/");
        if (resp.status === 200) {
            this.setState({folders: resp.data.children});
        }
    }

    async componentDidMount() {
        await this.refresh();
    }

    async componentDidUpdate(prevProps, prevState, snapshot) {
        // TODO: Refresh only if folders have changed
        // await this.refresh();
    }

    render() {
        return <div id="folder-tree-view">
            <p id="left-pane-section-label">F O L D E R S</p>
            {
                this.state.folders.map(folder => {
                    return <Folder
                        path={`${folder.name}`} selectedFolder={this.props.selectedFolder}
                        key={folder.id} offset={0} indent={this.props.indent}
                        folder={folder}/>
                })
            }
        </div>;
    }
}

export default FolderTreeView;