import React, {Component} from "react";

import "./Modal.css";
import "./SearchModal.css";

class SearchModal extends Component {
    state = {
        mouseInside: false
    }

    onMouseMovement = (inside) => {
        this.setState({mouseInside: inside});
    }

    onKeyDown = (e) => {
        if (e.key === "Escape") {
            this.props.onExit();
        }
    }

    componentDidMount() {
        document.addEventListener("keydown", this.onKeyDown, false);
    }

    componentWillUnmount() {
        document.removeEventListener("keydown", this.onKeyDown, false);
    }

    render() {
        return <div className="modal-container" onClick={this.props.onExit} onKeyDown={this.onKeyDown}>
            <div
                className="search-modal"
                onClick={event => event.stopPropagation()}
                onMouseOver={_ => this.onMouseMovement(true)}
                onMouseLeave={_ => this.onMouseMovement(false)}
            >
                <div id="search-input-container">
                    <img src="/icons/search-50.png" alt="search-icon" width={20} />
                    <input autoFocus placeholder="Search or jump to" id="search-input" />
                </div>
            </div>
        </div>;
    }
}

export default SearchModal;