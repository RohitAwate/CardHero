import React, {Component} from "react";

import "./CardModal.css";
import Card from "./Card";

class CardModal extends Component {
    state = {
        mouseInside: false
    }

    onMouseMovement = (inside) => {
        console.log("inside:", inside)
        this.setState({mouseInside: inside});
    }

    render() {
        const card = this.props.card;
        const folders = ["Folder", "Sub-folder", "Another sub-folder"];

        return <div className="card-modal-container" onClick={this.props.onExit}>
            <div
                className="card-modal"
                onClick={event => event.stopPropagation()}
                onMouseOver={_ => this.onMouseMovement(true)}
                onMouseLeave={_ => this.onMouseMovement(false)}
            >
                <div className="card-modal-nav-bar">
                    {
                        this.state.mouseInside ?
                            <button onClick={this.props.onExit} className="transparent-btn">
                                <img src="/icons/close-cross-48.png" alt="close-cross-icon"/>
                            </button> : ""
                    }
                </div>
                <div className="card-modal-body">
                    <p>{card.contents}</p>
                </div>
                <div className="card-modal-bottom-bar">
                    <div className="card-modal-nav-bar-folders-container">
                        {
                            folders.map((folder, i, row) => {
                                const link = <a className="folder-link" href="#">{folder}</a>;
                                if (i + 1 === row.length) {
                                    // Last one.
                                    return link;
                                } else {
                                    // Not last one.
                                    return [link, "  >  "];
                                }
                            })
                        }
                    </div>

                    <p className="timestamp">{Card.renderTimestamp(card.timestamp)}</p>
                </div>
            </div>
        </div>;
    }
}

export default CardModal;