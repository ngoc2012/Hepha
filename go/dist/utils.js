export const $ = s => {
    const e = document.querySelector(s)
    if (!e) throw new Error(`Element not found: ${s}`)
}
export const ready = c => document.addEventListener('DOMContentLoaded', c)
export const onClick = (e, c) => e.addEventListener('click', c)