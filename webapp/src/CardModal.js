import React, {Component} from "react";

import "./CardModal.css";

class CardModal extends Component {
    render() {
        return <div className="card-modal-container" onClick={this.props.onExit}>
            <div className="card-modal" onClick={event => event.stopPropagation()}>
                <div className="card-modal-nav-bar">
                    <p>Folder > SubFolder > Another SubFolder</p>
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