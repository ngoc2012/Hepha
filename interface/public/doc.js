function toggleMode() {
  const currentTheme = document.documentElement.getAttribute('data-theme');
  const newTheme = currentTheme === 'dark' ? 'light' : 'dark';
  document.documentElement.setAttribute('data-theme', newTheme);

  const toggleIcon = document.getElementById('toggle-icon');
    if (newTheme === 'dark') {
        toggleIcon.data = '/sun.svg';
    } else {
        toggleIcon.data = '/moon.svg';
    }
}

// Apply dark mode class if user prefers dark mode
if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
  document.documentElement.setAttribute('data-theme', 'dark');
  document.getElementById('toggle-icon').data = '/sun.svg';
} else {
  document.documentElement.setAttribute('data-theme', 'light');
  document.getElementById('toggle-icon').data = '/moon.svg';
}