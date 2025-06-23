document.body.addEventListener("htmx:wsOpen", () => {
  const pulseEl = document.getElementById("pulse");
  pulseEl.classList.add("active");
});
document.body.addEventListener("htmx:wsClose", () => {
  const pulseEl = document.getElementById("pulse");
  pulseEl.classList.remove("active");
});
