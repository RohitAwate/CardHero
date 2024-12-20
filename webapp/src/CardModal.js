import React, {Component} from "react";

import "./Modal.css";
import "./CardModal.css";
import Card from "./Card";
import {Link, useSearchParams} from "react-router-dom";
import Loader from "./Loader";
import AddTextToClipboard from "./services/Clipboard";
import axios from "axios";

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

    componentDidMount() {
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

    setPageTitle = () => {
        if (!this.props.card) {
            return;
        }

        const card = this.props.card;

        // If longer than n characters, shorten them.
        const n = 20;

        let cardContentsForTitle;
        if (card.contents.length > n) {
            cardContentsForTitle = card.contents.substring(0, n) + "...";
        } else {
            cardContentsForTitle = card.contents + "...";
        }

        document.title = cardContentsForTitle + " | CardHero";
    }

    render() {
        const card = this.props.card;
        this.setPageTitle();

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
                            <Link to="." relative="path" className="transparent-btn">
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
                    <div className="card-modal-interactions-container">
                        <button onClick={async _ => AddTextToClipboard(await CardModalMeta.getCardPermalink(card.id))}
                                className="permalink-copy-btn transparent-btn">
                            <img src="/icons/permalink-50.png" alt="permalink-icon"/>
                        </button>
                        <button onClick={_ => AddTextToClipboard(card.contents)}
                                className="contents-copy-btn transparent-btn">
                            <img src="/icons/copy-48.png" alt="copy-icon"/>
                        </button>
                    </div>
                </div>
            </div>
        </div>;
    }

    static async getCardPermalink(uuid) {
        const apiURL = `/api/rohit/card/goto/${uuid}`;
        const resp = await axios.get(apiURL);
        if (resp.status !== 200) {
            alert("Could not get permalink for the card");
            return;
        }

        return `${window.location.origin}${resp.data}`;
    }
}

export default CardModal;