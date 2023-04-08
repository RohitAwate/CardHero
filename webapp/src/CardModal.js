import React, {Component} from "react";

import "./Modal.css";
import "./CardModal.css";
import Card from "./Card";
import {Link, useSearchParams} from "react-router-dom";
import Loader from "./Loader";

function CardModal(props) {
    const [searchParams, setSearchParams] = useSearchParams();
    const closeModal = () => {
        if (searchParams.has("card")) {
            searchParams.delete("card");
            setSearchParams(searchParams);
        }
    };

    return <CardModalMeta
        card={props.card}
        folderPath={props.folderPath}
        closeModal={closeModal}
    />;
}

class CardModalMeta extends Component {
    state = {
        mouseInside: false,
    };

    modalRef = React.createRef();

    async componentDidMount() {
        document.addEventListener("keydown", this.onKeyDown, false);
        document.addEventListener("mousedown", this.onClickOutside);
    }

    componentWillUnmount() {
        document.removeEventListener("keydown", this.onKeyDown, false);
        document.removeEventListener("mousedown", this.onClickOutside);
    }

    onMouseMovement = (inside) => {
        this.setState({mouseInside: inside});
    }

    onClickOutside = (e) => {
        if (this.modalRef && !this.modalRef.current.contains(e.target)) {
            return this.props.closeModal();
        }
    }

    onKeyDown = (e) => {
        if (e.key === "Escape") {
            return this.props.closeModal();
        }
    }

    render() {
        const card = this.props.card;

        let linkPath = "/folders";
        const folderPath = this.props.folderPath.map((folder, i, row) => {
            linkPath += "/" + folder;
            const link = <Link key={i} className="folder-link" to={linkPath}>{folder}</Link>;

            if (i + 1 === row.length) {
                return link;
            } else {
                const spacerKey = `${folder}-${i}-spacer`;
                return [link, <p key={spacerKey} className="folder-link-spacer">⟩</p>];
            }
        });

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
                    {card ? <p>{card.contents}</p> : <Loader/>}
                </div>
                <div className="card-modal-bottom-bar">
                    <div className="card-modal-nav-bar-folders-container">
                        {folderPath}
                    </div>
                    <p className="timestamp">
                        {
                            card ?
                                Card.renderTimestamp(card.timestamp)
                                : <Loader size={30}/>
                        }
                    </p>
                </div>
            </div>
        </div>;
    }
}

export default CardModal;