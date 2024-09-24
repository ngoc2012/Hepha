import { $a, ready, onClick } from "/utils.js";

ready(() => {
    let buttons = $a('.main.content')
    buttons.forEach(button => {
        onClick(button, () => {
            console.log('clicked')
        });
    });
});
