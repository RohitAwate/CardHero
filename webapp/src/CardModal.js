import React, {Component} from "react";

import "./CardModal.css";

class CardModal extends Component {
    render() {
        const folders = ["Folder", "Sub-folder", "Another sub-folder"];

        return <div className="card-modal-container" onClick={this.props.onExit}>
            <div className="card-modal" onClick={event => event.stopPropagation()}>
                <div className="card-modal-nav-bar">
                    <div>
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

                    <button onClick={this.props.onExit} className="transparent-btn">
                        <img src="/icons/close-cross-48.png" alt="close-cross-icon"/>
                    </button>
                </div>
                <div className="card-modal-body">
                    <p>{this.props.card.contents}</p>
                </div>
            </div>
        </div>;
    }
}

export default CardModal;