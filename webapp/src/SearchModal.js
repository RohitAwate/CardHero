import React, {Component} from "react";

import "./Modal.css";
import "./SearchModal.css";
import "./Common.css";
import axios from "axios";
import Card from "./Card";
import {useNavigate} from "react-router-dom";

class SearchModal extends Component {
    static SEARCH_TYPING_DELAY = 200;

    state = {
        show: false,
        mouseInside: false,
        results: [],
        typingTimer: null,
    }

    modalRef = React.createRef();

    onMouseMovement = (inside) => {
        this.setState({mouseInside: inside});
    }

    hideSearch = () => {
        const show = false;
        const results = [];
        this.setState({show, results});
    }

    onKeyDown = (e) => {
        if (e.key === "Escape") {
            this.hideSearch();
        } else if (e.ctrlKey && e.key === "k") {
            e.preventDefault();
            const show = !this.state.show;
            const results = [];
            this.setState({show, results});
        }
    }

    componentDidMount() {
        document.addEventListener("keydown", this.onKeyDown, false);
        document.addEventListener("mousedown", this.onClickOutside);
    }

    componentWillUnmount() {
        document.removeEventListener("keydown", this.onKeyDown, false);
        document.removeEventListener("mousedown", this.onClickOutside);
    }

    onClickOutside = (e) => {
        // First check is for when the search modal is not showing.
        if (this.modalRef.current && !this.modalRef.current.contains(e.target)) {
            return this.hideSearch();
        }
    }

    onQueryTyped = (e) => {
        const newQuery = e.target.value;

        if (this.state.typingTimer) {
            clearTimeout(this.state.typingTimer);
        }

        const typingTimer = setTimeout(async () => {
            const resp = await axios.get("/api/rohit/search", {params: {query: newQuery}});
            if (resp.status === 200) {
                const results = resp.data;
                this.setState({results})
            }
        }, SearchModal.SEARCH_TYPING_DELAY);

        this.setState({typingTimer});
    }

    render() {
        if (this.state.show) {
            return <div className="modal-container">
                <div
                    ref={this.modalRef}
                    className="search-modal"
                    onClick={event => event.stopPropagation()}
                    onMouseOver={_ => this.onMouseMovement(true)}
                    onMouseLeave={_ => this.onMouseMovement(false)}
                >
                    <div id="search-input-container">
                        <img src="/icons/search-50.png" alt="search-icon" width={20}/>
                        <input autoComplete={"off"} autoFocus onChange={this.onQueryTyped}
                               placeholder="Search or jump to"
                               id="search-input"/>
                    </div>
                    <div className="separator"></div>
                    <div id="search-results-container">
                        {
                            this.state.results.map(result => <SearchResult hideSearch={this.hideSearch} key={result.id}
                                                                           result={result}/>)
                        }
                    </div>
                </div>
            </div>;
        }
    }
}

function SearchResult(props) {
    const result = props.result;
    const formattedTime = Card.renderTimestamp(result.timestamp);
    const navigate = useNavigate();

    const onClickHandler = async (e) => {
        const resp = await axios.get(`/api/rohit/card/goto/${props.result.id}`);
        if (resp.status === 200) {
            props.hideSearch();
            return navigate(resp.data);
        }
    };

    return <div className="search-result" onClick={onClickHandler}>
        <p>{result.contents}</p>
        <div className="timestamp">{formattedTime}</div>
    </div>;
}

export default SearchModal;