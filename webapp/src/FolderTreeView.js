import {Component} from "react";

import "./FolderTreeView.css";
import {Folder} from "./Folder";
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
        if (this.props.lastUpdated !== prevProps.lastUpdated) {
            // This is usually called when a new card is added.
            // It might take about a second for the card to be fully ingested.
            // Thus, delaying the update.
            setTimeout(async () => {
                await this.refresh()
            }, 500);
        } else if (this.props.selectedFolder !== prevProps.selectedFolder) {
            await this.refresh();

        }
    }

    render() {
        const indent = this.props.indent ? this.props.indent : 10;

        return <div id="folder-tree-view">
            <p id="left-pane-section-label">F O L D E R S</p>
            {
                this.state.folders.map(folder => {
                    return <Folder
                        path={`${folder.name}`} selectedFolder={this.props.selectedFolder}
                        key={folder.id} offset={0} indent={indent}
                        folder={folder}/>
                })
            }
        </div>;
    }
}

export default FolderTreeView;