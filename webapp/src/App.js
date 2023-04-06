import {Component} from "react";
import Chat from "./Chat";

import "./App.css";
import Explorer from "./Explorer";

class App extends Component {
    render() {
        return <div id="app-container">
            <Explorer/>
            <Chat/>
        </div>;
    }
}

export default App;