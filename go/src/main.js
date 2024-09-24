import { $, ready, onClick } from "/utils.js";

ready(() => {
    let buttons = $('.main.content')
    if (buttons instanceof NodeList) {
        buttons.forEach(button => {
            onClick(button, () => {
                console.log('clicked')
            });
        });
    } else {
        onClick(buttons, () => {
        console.log('clicked')
        });
    }
});
