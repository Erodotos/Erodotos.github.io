// Theme toggle: follows the system setting until the visitor picks one.
(function () {
  var root = document.documentElement;
  var KEY = "theme";

  function current() {
    var set = root.getAttribute("data-theme");
    if (set) return set;
    return window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
  }

  function render(btn) {
    var dark = current() === "dark";
    btn.textContent = dark ? "☀" : "☾";
    btn.setAttribute("aria-label", dark ? "Switch to light theme" : "Switch to dark theme");
  }

  document.addEventListener("DOMContentLoaded", function () {
    var btn = document.querySelector(".theme-toggle");
    if (!btn) return;
    render(btn);
    btn.addEventListener("click", function () {
      var next = current() === "dark" ? "light" : "dark";
      root.setAttribute("data-theme", next);
      try { localStorage.setItem(KEY, next); } catch (e) {}
      render(btn);
    });
    window.matchMedia("(prefers-color-scheme: dark)").addEventListener("change", function () {
      render(btn);
    });
  });
})();
