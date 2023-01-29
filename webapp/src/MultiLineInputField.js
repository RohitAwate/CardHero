import React, {Component} from "react";

class MultiLineInputField extends Component {
    constructor(props) {
        super(props);
        this.state = {
            value: "",
        };
    }

    reset = () => {
        this.setState({value: "", offset: 0});
    }

    onInput = (event) => {
        const value = event.target.innerText;
        this.setState({value: value});
    }

    value = () => {
        return this.state.value;
    }

    render() {
        return <div
            contentEditable
            id={this.props.id}
            role="textbox"
            suppressContentEditableWarning={true}
            placeholder={this.props.placeholder}
            onInput={this.onInput}
        >{this.state.value}</div>;
    }
}

export default MultiLineInputField;