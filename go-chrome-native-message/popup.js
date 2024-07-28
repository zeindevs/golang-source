(() => {
  "use strict";

  document.addEventListener("DOMContentLoaded", () => {
    document.getElementById("send").addEventListener("click", (e) => {
      let start = performance.now();
      e.preventDefault();
      const message = document.getElementById("message").value;
      chrome.runtime.sendMessage({ message: message }, (response) => {
        let end = performance.now();
        document.getElementById("result").innerText =
          `${JSON.stringify(response)} took: ${end - start}ms`;
        console.log(`Response from native app:`, response);
      });
    });
  });
})();
