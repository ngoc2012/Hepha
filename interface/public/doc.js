import { $, ready, onClick } from "/selector.js";

if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
  document.documentElement.setAttribute('data-theme', 'dark');
  setModeIcon('dark');
} else {
  document.documentElement.setAttribute('data-theme', 'light');
  setModeIcon('light');
}

onClick($("#app button[name=mode-toggler]")[0], toggleMode);

function toggleMode() {
  const currentTheme = document.documentElement.getAttribute('data-theme');
  const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
  document.documentElement.setAttribute('data-theme', newTheme);
  setModeIcon(newTheme);
}

function setModeIcon(theme) {
  var pathElement = $("#app button[name=mode-toggler] path")[0];
  if (theme === 'dark') {
    pathElement.setAttribute('d', "M12 3v2.25m6.364.386-1.591 1.591M21 12h-2.25m-.386 6.364-1.591-1.591M12 18.75V21m-4.773-4.227-1.591 1.591M5.25 12H3m4.227-4.773L5.636 5.636M15.75 12a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0Z");
  } else {
    pathElement.setAttribute('d', "M21.752 15.002A9.72 9.72 0 0 1 18 15.75c-5.385 0-9.75-4.365-9.75-9.75 0-1.33.266-2.597.748-3.752A9.753 9.753 0 0 0 3 11.25C3 16.635 7.365 21 12.75 21a9.753 9.753 0 0 0 9.002-5.998Z");
  }
}


console.log($(".page p"));
console.log($(".page *"));

console.log($(".page")[0].getBoundingClientRect());
