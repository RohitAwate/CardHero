import "./LeftPane.css"
import FolderTreeView from "./FolderTreeView";


function LeftPane(props) {
    return <div id="left-pane">
        <FolderTreeView lastUpdated={props.lastUpdated} selectedFolder={props.selectedFolder}/>
    </div>;
}

export default LeftPane;