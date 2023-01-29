import React, {Component} from "react";

class MultiLineInput extends Component {
    constructor(props) {
        super(props);
        this.inputValue = "";
    }

    reset = () => {
        this.inputValue = "";
        this.setState({});
    }

    onInput = (event) => {
        if (event.target.innerHTML.trim() === '<div><br></div>') {
            event.target.innerHTML = "";
        }

        this.inputValue = event.target.innerHTML;
    }

    onSubmit = (event) => {
        if ((event.metaKey || event.ctrlKey) && event.which === 13) {
            this.props.onSubmit(event);
        }
    }

    value = () => {
        return this.inputValue;
    }

    render() {
        return <div
            contentEditable
            id={this.props.id}
            ref={this.inputRef}
            role="textbox"
            suppressContentEditableWarning={true}
            placeholder={this.props.placeholder}
            onKeyDown={this.onSubmit}
            onInput={this.onInput}
        >{this.inputValue}</div>;
    }
}

export default MultiLineInput;