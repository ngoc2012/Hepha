export const $ = (s: string): NodeListOf<Element> => document.querySelectorAll(s);
export const ready = (c: (event: Event) => void): void => {document.addEventListener('DOMContentLoaded', c);};
export const onClick = (e: EventTarget, c: (event: MouseEvent) => void): void => {
    e.addEventListener('click', c as EventListener);
};