import {Component} from "react";

import "./Folder.css";

class Folder extends Component {
    state = {
        expanded: true
    }

    toggleChildren = (e) => {
        this.setState({expanded: !this.state.expanded});
    }

    render() {
        const style = {
            marginLeft: this.props.offset
        };

        const childOffset = this.props.offset + this.props.indent;
        const hasChildren = this.props.folder.children.length > 0;

        let icon;
        if (this.state.expanded) {
            icon = "/icons/chevron-down-48.png";
        } else {
            icon = "/icons/chevron-right-48.png";
        }

        return <div id="folder-container" style={style}>
            {
                hasChildren ?
                    <span className="folder-name" onClick={this.toggleChildren}>
                        <img className="folder-chevron" alt="chevron" src={icon}/>
                        <p>{this.props.folder.name}</p>
                    </span>
                    :
                    <span className="folder-name" onClick={this.toggleChildren}>
                        <p>{this.props.folder.name}</p>
                    </span>
            }
            {
                this.state.expanded ?
                    this.props.folder.children.map(
                        child => <Folder offset={childOffset} indent={this.props.indent} folder={child}/>
                    )
                    : ""
            }
        </div>;
    }
}

export default Folder;