import { qsa, ready } from "/utils.js";

ready(() => {
    let buttons = qsa('.main.content')
    buttons.forEach(button => {
        button.addEventListener('click', () => {
            console.log('clicked');
        });
    });
});
