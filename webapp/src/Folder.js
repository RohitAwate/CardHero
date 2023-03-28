import {Component} from "react";

import "./Folder.css";

const CHEVRON_DOWN = "/icons/chevron-down-48.png";
const CHEVRON_RIGHT = "/icons/chevron-right-48.png";

class Folder extends Component {
    state = {
        expanded: false
    }

    onFolderClick = (e) => {
        this.setState({expanded: !this.state.expanded});
        e.folderID = this.props.folder.id;
        this.props.onFolderSelect(e);
    }

    render() {
        const style = {
            marginLeft: this.props.offset
        };

        const childOffset = this.props.offset + this.props.indent;
        const hasChildren = this.props.folder.children.length > 0;

        let icon;
        if (this.state.expanded) {
            icon = CHEVRON_DOWN;
        } else {
            icon = CHEVRON_RIGHT;
        }

        const folderSelectedClass = this.props.folder.id === this.props.selectedFolder ? " folder-selected" : "";

        return <div id="folder-container" style={style}>
            {
                hasChildren ?
                    <span className={"folder-name" + folderSelectedClass} onClick={this.onFolderClick}>
                        <p>{this.props.folder.name}</p>
                        <img className="folder-chevron" alt="chevron" src={icon}/>
                    </span>
                    :
                    <span className={"folder-name" + folderSelectedClass} onClick={this.onFolderClick}>
                        <p>{this.props.folder.name}</p>
                    </span>
            }
            {
                this.state.expanded ?
                    this.props.folder.children.map(
                        child => <Folder selected={child.id === this.props.selectedFolder}
                                         selectedFolder={this.props.selectedFolder}
                                         key={child.id} onFolderSelect={this.props.onFolderSelect}
                                         offset={childOffset} indent={this.props.indent} folder={child}/>
                    )
                    : ""
            }
        </div>;
    }
}

export default Folder;