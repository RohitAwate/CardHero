import React, {Component} from "react";

import "./Modal.css";
import "./SearchModal.css";
import axios from "axios";
import Card from "./Card";

class SearchModal extends Component {
    static SEARCH_TYPING_DELAY = 200;

    state = {
        mouseInside: false,
        results: [],
        typingTimeoutID: null,
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

    onQueryTyped = (e) => {
        const newQuery = e.target.value;

        if (this.state.typingTimeoutID) {
            clearTimeout(this.state.typingTimeoutID);
        }

        const typingTimeoutID = setTimeout(async () => {
            const resp = await axios.get("/api/rohit/search", {params: {query: newQuery}});
            if (resp.status === 200) {
                const results = resp.data;
                this.setState({results})
            }
        }, SearchModal.SEARCH_TYPING_DELAY);

        this.setState({typingTimeoutID});
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
                    <img src="/icons/search-50.png" alt="search-icon" width={20}/>
                    <input autoComplete={"off"} autoFocus onChange={this.onQueryTyped} placeholder="Search or jump to" id="search-input"/>
                </div>
                <div>
                    {
                        this.state.results.map(result => <SearchResult key={result.id} result={result}/>)
                    }
                </div>
            </div>
        </div>;
    }
}

function SearchResult(props) {
    const result = props.result;
    const formattedTime = Card.renderTimestamp(result.timestamp);

    return <a href={`/cards/${result.id}`} className="search-result-anchor">
        <div className="search-result">
            <p>{result.contents}</p>
            <div className="timestamp">{formattedTime}</div>
        </div>
    </a>;
}

export default SearchModal;