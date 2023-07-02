import {Component} from "react";
import Chat from "./Chat";

import "./App.css";
import Explorer from "./Explorer";
import SearchModal from "./SearchModal";

class App extends Component {
    state = {lastUpdated: Date.now()}

    update = () => {
        this.setState({lastUpdated: Date.now()})
    }

    render() {
        return <div id="app-container">
            <Explorer lastUpdated={this.state.lastUpdated} />
            <Chat updateApp={this.update}/>
            <SearchModal/>
        </div>;
    }
}

export default App;