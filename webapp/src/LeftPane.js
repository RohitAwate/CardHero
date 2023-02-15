import {Component} from "react";

import "./LeftPane.css"
import FolderTreeView from "./FolderTreeView";


class LeftPane extends Component {
    state = {
        folders: [
            {name: "Default", children: []},
            {
                name: "Code",
                children: [
                    {name: "Golang", children: []}, {name: "Java", children: []}
                ]
            },
            {
                name: "Recipes",
                children: [
                    {name: "Veg", children: []}, {name: "NonVeg", children: []}
                ]
            }
        ]
    }

    render() {
        return <div id="left-pane">
            <FolderTreeView folders={this.state.folders} indent={15}/>
        </div>;
    }
}

export default LeftPane;