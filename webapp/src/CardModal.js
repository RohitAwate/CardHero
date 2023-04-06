import React, {Component} from "react";

import "./Modal.css";
import "./CardModal.css";
import Card from "./Card";
import {Link, Navigate} from "react-router-dom";

class CardModal extends Component {
    state = {
        mouseInside: false
    }

    onMouseMovement = (inside) => {
        this.setState({mouseInside: inside});
    }

    onKeyDown = (e) => {
        if (e.key === "Escape") {
            return <Navigate to=".."/>
        }
    }

    componentDidMount() {
        document.addEventListener("keydown", this.onKeyDown, false);
    }

    componentWillUnmount() {
        document.removeEventListener("keydown", this.onKeyDown, false);
    }

    render() {
        const card = this.props.card;
        const folderPath = ["See", "My", "Dummy", "Folder", "Path"];

        return <Link to=".." relative="path">
            <div className="modal-container" onKeyDown={this.onKeyDown}>
                <div
                    className="card-modal"
                    onClick={event => event.stopPropagation()}
                    onMouseOver={_ => this.onMouseMovement(true)}
                    onMouseLeave={_ => this.onMouseMovement(false)}
                >
                    <div className="card-modal-nav-bar">
                        {
                            this.state.mouseInside ?
                                <Link to=".." relative="path">
                                    <button className="transparent-btn">
                                        <img src="/icons/close-cross-48.png" alt="close-cross-icon"/>
                                    </button>
                                </Link> : ""
                        }
                    </div>
                    <div className="card-modal-body">
                        <p>{card.contents}</p>
                    </div>
                    <div className="card-modal-bottom-bar">
                        <div className="card-modal-nav-bar-folders-container">
                            {
                                folderPath.map((folder, i, row) => {
                                    const link = <a key={i} className="folder-link" href="#">{folder}</a>;
                                    if (i + 1 === row.length) {
                                        return link;
                                    } else {
                                        const spacerKey = `${folder}-${i}-spacer`;
                                        return [link, <p key={spacerKey} className="folder-link-spacer">⟩</p>];
                                    }
                                })
                            }
                        </div>

                        <p className="timestamp">{Card.renderTimestamp(card.timestamp)}</p>
                    </div>
                </div>
            </div>
        </Link>;
    }
}

export default CardModal;