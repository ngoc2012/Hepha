import { $, ready, onClick } from "/utils.js";

ready(() => {
    let buttons = $('.main.content')
    buttons.forEach(button => {
        onClick(button, () => {
            console.log('clicked')
        });
    });
});
