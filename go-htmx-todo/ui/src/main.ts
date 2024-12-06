import htmx from "htmx.org";
import Sortable from "sortablejs";

import "./style.css";

window.htmx = htmx;
window.htmx.onLoad((_ctx: any) => {
  var el = document.getElementById("items")!;
  Sortable.create(el, { handle: ".handle" });
});
