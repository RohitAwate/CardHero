import {Component} from "react";
import Chat from "./Chat";

import "./App.css";
import Explorer from "./Explorer";
import axios from "axios";

class App extends Component {
    state = {
        folders: []
    }

    async componentDidMount() {
        await this.fetchCards();
    }

    onNewCard = async (contents) => {
        const newCard = await this.createNewCard(contents);
        this.fetchCards();
        return newCard;
    }

    fetchCards = async () => {
        const resp = await axios.get("/api/rohit/folders/");
        if (resp.status === 200) {
            this.setState({folders: resp.data.children});
        }
    }

    createNewCard = async (contents) => {
        const payload = new URLSearchParams();
        payload.append('contents', contents);
        payload.append('timestamp', new Date().toISOString());

        const resp = await axios.post("/api/rohit/logs", payload);
        if (resp.status === 201) {
            return resp.data;
        }
    }

    getFolderPath = (folderID) => {
        if (!this.state.folders) {
            return [];
        }

        const dfs = (children, parentPath) => {
            for (const child of children) {
                const childPath = parentPath.concat(child.name);

                if (child.id === folderID) {
                    return childPath;
                } else {
                    const result = dfs(child.children, childPath);
                    if (result !== childPath) {
                        return result;
                    }
                }
            }

            return parentPath;
        };

        return dfs(this.state.folders, []);
    }

    render() {
        return <div id="app-container">
            <Explorer folders={this.state.folders} getFolderPath={this.getFolderPath}/>
            <Chat onNewCard={this.onNewCard}/>
        </div>;
    }
}

export default App;