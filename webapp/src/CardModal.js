import React, {Component} from "react";

import "./Modal.css";
import "./CardModal.css";
import Card from "./Card";
import {Link, useNavigate} from "react-router-dom";
import axios from "axios";

function CardModal(props) {
    const navigate = useNavigate();
    return <CardModalMeta navigate={navigate} cardID={props.cardID}/>;
}

class CardModalMeta extends Component {
    state = {
        mouseInside: false,
        card: {},
    };

    modalRef = React.createRef();

    async refresh() {
        const resp = await axios.get(`/api/rohit/card/${this.props.cardID}`);
        if (resp.status === 200) {
            const card = resp.data;
            this.setState({card});
        }
    }

    async componentDidMount() {
        document.addEventListener("keydown", this.onKeyDown, false);
        document.addEventListener("mousedown", this.onClickOutside);

        await this.refresh();
    }

    componentWillUnmount() {
        document.removeEventListener("keydown", this.onKeyDown, false);
        document.removeEventListener("mousedown", this.onClickOutside);
    }

    onMouseMovement = (inside) => {
        this.setState({mouseInside: inside});
    }

    closeModal = () => {
        return this.props.navigate(-1);
    }

    onClickOutside = (e) => {
        if (this.modalRef && !this.modalRef.current.contains(e.target)) {
            return this.closeModal();
        }
    }

    onKeyDown = (e) => {
        if (e.key === "Escape") {
            return this.closeModal();
        }
    }

    render() {
        const card = this.state.card;
        const folderPath = ["See", "My", "Dummy", "Folder", "Path"];

        return <div className="modal-container" onKeyDown={this.onKeyDown}>
            <div
                ref={this.modalRef}
                className="card-modal"
                onClick={event => event.stopPropagation()}
                onMouseOver={_ => this.onMouseMovement(true)}
                onMouseLeave={_ => this.onMouseMovement(false)}
            >
                <div className="card-modal-nav-bar">
                    {
                        this.state.mouseInside ?
                            <Link to=".." relative="path" className="transparent-btn">
                                <img src="/icons/close-cross-48.png" alt="close-cross-icon"/>
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
        </div>;
    }
}

export default CardModal;