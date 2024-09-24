import { $, ready, onClick } from "/utils.js";

class Button {
    constructor(element) {
        this.element = element;
        this.state = 'default'; // Initial state
    }

    setState(newState) {
        this.state = newState;
    }

    getState() {
        return this.state;
    }
}

ready(() => {
    let buttons = $('.main.content')
    buttons.forEach(button => {
        onClick(button, () => {
            console.log('clicked')
        });
    });
});
