export const $ = s => {
    const e = document.querySelectorAll(s)
    if (!e) throw new Error(`Element not found: ${s}`)
    console.log(e)
    if (e.length === 1) return e[0]
    return e
}
export const ready = c => document.addEventListener('DOMContentLoaded', c)
export const onClick = (e, c) => e.addEventListener('click', c)