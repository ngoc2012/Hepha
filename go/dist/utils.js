export const $ = s => document.querySelectorAll(s)
export const ready = c => document.addEventListener('DOMContentLoaded', c)
export const onClick = (e, c) => e.addEventListener('click', c)

// If you need to return a single element, you can use the following code:
// export const $ = s => {
//     const e = document.querySelectorAll(s)
//     if (e.length === 0) throw new Error(`Element not found: ${s}`)
//     if (e.length === 1) return e[0]
//     return e
// }