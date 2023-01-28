import {Component} from "react";
import axios from "axios";

class Log extends Component {
    render() {
        if (this.state) {
            return <div>
                {
                    this.state.logs.map(card => {
                        return <h2 key={card.id}>{card.contents}</h2>
                    })
                }
            </div>;
        } else {
            return <h1>Add something to the log!</h1>;
        }
    }

    async componentDidMount() {
        const resp = await axios.get("/api/rohit/logs");
        this.setState({logs: resp.data});
    }
}

export default Log;