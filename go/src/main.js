import { $, ready, onClick } from "/utils.js";

class Button {
    constructor(e) {
        this.e = e
        this.number = 1
        onClick(this.e, () => {
            this.number++
            this.render()
        })
    }
    render() {this.e.innerHTML = "Count " + this.number}
}

ready(() => {
    let buttons = $('.main.content')
    buttons.forEach(button => {new Button(button)});
});
