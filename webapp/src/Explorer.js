import {Component} from "react";
import LeftPane from "./LeftPane";
import Gallery from "./Gallery";

import "./Explorer.css"
import TopBar from "./TopBar";
import axios from "axios";

class Explorer extends Component {
    state = {
        selectedFolderID: "",
        cards: []
    }

    DEFAULT_FOLDER_NAME = "Default";

    getDefaultFolderID = async () => {
        const resp = await axios.get(`/api/rohit/folders/${this.DEFAULT_FOLDER_NAME}`);
        if (resp.status === 200) {
            return resp.data.id;
        }
    }

    async componentDidMount() {
        const defaultFolderID = await this.getDefaultFolderID();
        await this.onFolderSelect({folderID: defaultFolderID});
    }


    onFolderSelect = async (e) => {
        this.setState({selectedFolderID: e.folderID, cards: []});

        const resp = await axios.get(`/api/rohit/folder/${e.folderID}`);
        if (resp.status === 200) {
            this.setState({selectedFolderID: e.folderID, cards: resp.data});
        }
    }

    render() {
        return <div id="card-explorer">
            <TopBar/>
            <div id="card-explorer-main">
                <LeftPane selectedFolder={this.state.selectedFolderID} onFolderSelect={this.onFolderSelect}
                          folders={this.props.folders}/>
                <Gallery cards={this.state.cards} folderPath={this.props.getFolderPath(this.state.selectedFolderID)}/>
            </div>
        </div>;
    }
}

export default Explorer;