import { $, ready, onClick } from "/utils.js";

ready(() => {
    let buttons = $('.main.content')
    console.log(buttons)
    buttons.forEach(button => {
        onClick(button, () => {
            console.log('clicked')
        });
    });
});
