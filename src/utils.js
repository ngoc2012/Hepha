export const $ = s => document.querySelectorAll(s)
export const ready = c => document.addEventListener('DOMContentLoaded', c)
export const onClick = (e, c) => e.addEventListener('click', c)