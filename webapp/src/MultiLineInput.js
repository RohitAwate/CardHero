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
        if (!this.allowedWhitespace && event.target.innerText.trim() === '') {
            event.target.innerText = "";
        }

        this.inputValue = event.target.innerText;
    }

    onSubmit = (event) => {
        this.allowedWhitespace = false;

        if (event.shiftKey && event.which === 13) {
            this.allowedWhitespace = true;
        } else if ((event.metaKey || event.ctrlKey) && event.which === 13) {
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