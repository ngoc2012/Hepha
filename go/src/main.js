import { qsa, ready, onClick } from "/utils.js";

ready(() => {
    let buttons = qsa('.main.content')
    buttons.forEach(button => {
        onClick(button, () => {
            console.log('clicked')
        });
    });
});
