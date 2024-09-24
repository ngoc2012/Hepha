import { $, ready, onClick } from "/utils.js";

ready(() => {
    let button = $('.main.contenst')
    console.log(button)
    onClick(button, () => {
        console.log('clicked')
    });
    // let buttons = $('.main.content')
    // buttons.forEach(button => {
    //     onClick(button, () => {
    //         console.log('clicked')
    //     });
    // });
});
