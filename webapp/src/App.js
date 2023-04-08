import {Component} from "react";
import Chat from "./Chat";

import "./App.css";
import Explorer from "./Explorer";
import SearchModal from "./SearchModal";

class App extends Component {
    render() {
        return <div id="app-container">
            <Explorer/>
            <Chat/>
            <SearchModal/>
        </div>;
    }
}

export default App;