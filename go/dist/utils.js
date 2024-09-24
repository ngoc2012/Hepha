export const $ = s => document.querySelector(s)
export const $a = s => document.querySelectorAll(s)
export const ready = c => document.addEventListener('DOMContentLoaded', c)
export const onClick = (e, c) => e.addEventListener('click', c)