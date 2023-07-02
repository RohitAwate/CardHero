import {Component} from "react";
import LeftPane from "./LeftPane";
import Gallery from "./Gallery";

import "./Explorer.css"
import TopBar from "./TopBar";
import {Navigate, Route, Routes, useLocation} from "react-router-dom";

class Explorer extends Component {
    DEFAULT_FOLDER = "/folders/Default";

    setPageTitle = () => {

    }

    render() {
        this.setPageTitle();

        return <div id="card-explorer">
            <TopBar/>
            <div id="card-explorer-main">
                <Routes>
                    <Route path={"/"} element={<Navigate to={this.DEFAULT_FOLDER}/>}/>
                    <Route path={"/folders/*"} element={<ExplorerMeta lastUpdated={this.props.lastUpdated}/>}/>
                    <Route path={"/cards/:id"} element={<ExplorerMeta lastUpdated={this.props.lastUpdated}/>}/>
                </Routes>
            </div>
        </div>;
    }
}

function ExplorerMeta(props) {
    const location = useLocation();

    // Grab the path and strip away the "/folders" part
    const selectedFolder = location.pathname.substring("/folders".length);

    // Update page title
    const folders = selectedFolder.split("/");
    const currentFolder = folders[folders.length - 1];
    document.title = currentFolder + " | CardHero";

    return <>
        <LeftPane lastUpdated={props.lastUpdated} selectedFolder={selectedFolder}/>
        <Gallery lastUpdated={props.lastUpdated} selectedFolder={selectedFolder}/>
    </>;
}

export default Explorer;