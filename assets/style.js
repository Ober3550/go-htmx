window.addEventListener("DOMContentLoaded", (event) => {
  document.body.addEventListener("htmx:load", function (event) {
    document.querySelectorAll("i").forEach((icon) => {
      icon.classList.add("material-symbols-rounded");
    });
    document.querySelectorAll("select").forEach((select) => {
      select.style =
        "appearance: none; padding: 16px; border-radius: 20px; font-size: 20px; margin: 4px";
    });
  });
});
