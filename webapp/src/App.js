import {Component} from "react";
import Chat from "./Chat";

import "./App.css";
import CardExplorer from "./CardExplorer";

class App extends Component {
    render() {
        return <div id="app-container">
            <CardExplorer/>
            <Chat/>
        </div>;
    }
}

export default App;