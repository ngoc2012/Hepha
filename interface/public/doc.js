function toggleMode() {
  document.documentElement.setAttribute('data-theme', document.documentElement.getAttribute('data-theme') === 'dark' ? 'light' : 'dark');

  const toggleIcon = document.getElementById('toggle-icon');
    if (newTheme === 'dark') {
        toggleIcon.src = 'dark-mode-icon.svg';
    } else {
        toggleIcon.src = 'light-mode-icon.svg';
    }
}

// Apply dark mode class if user prefers dark mode
if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
  document.documentElement.setAttribute('data-theme', 'dark');
  document.getElementById('toggle-icon').src = 'dark-mode-icon.svg';
} else {
  document.documentElement.setAttribute('data-theme', 'light');
  document.getElementById('toggle-icon').src = 'light-mode-icon.svg';
}