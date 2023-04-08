import {Component} from "react";

import "./Folder.css";
import {Link} from "react-router-dom";

const CHEVRON_DOWN = "/icons/chevron-down-48.png";
const CHEVRON_RIGHT = "/icons/chevron-right-48.png";

class Folder extends Component {
    state = {expanded: false}

    updateExpansion = () => {
        const thisFolderSelected = "/" + this.props.path === this.props.selectedFolder;
        if (thisFolderSelected) {
            this.setState({expanded: true});
        }
    };

    componentDidMount = this.updateExpansion;

    componentDidUpdate(prevProps, prevState, snapshot) {
        if (prevProps.selectedFolder !== this.props.selectedFolder) {
            this.updateExpansion();
        }
    }

    toggleExpansion = () => {
        const expanded = !this.state.expanded;
        this.setState({expanded});
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

        const thisFolderSelected = "/" + this.props.path === this.props.selectedFolder;
        const folderSelectedClass = thisFolderSelected ? " folder-selected" : "";

        return <>
            <Link to={`/folders/${this.props.path}`}>
                <div id="folder-container" onClick={this.toggleExpansion} style={style}>
                    <span className={"folder-name" + folderSelectedClass}>
                        {
                            hasChildren ?
                                <>
                                    <p>{this.props.folder.name}</p>
                                    <img className="folder-chevron" alt="chevron" src={icon}/>
                                </>
                                :
                                <p>{this.props.folder.name}</p>
                        }
                    </span>
                </div>
            </Link>
            {
                this.state.expanded ?
                    this.props.folder.children.map(
                        child => <Folder selected={child.id === this.props.selectedFolder}
                                         selectedFolder={this.props.selectedFolder}
                                         key={child.id} onFolderSelect={this.props.onFolderSelect}
                                         offset={childOffset} indent={this.props.indent} folder={child}
                                         path={`${this.props.path}/${child.name}`}
                        />
                    )
                    : ""
            }
        </>;
    }
}

export default Folder;