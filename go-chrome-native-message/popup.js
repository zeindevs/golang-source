(() => {
  "use strict";

  document.addEventListener("DOMContentLoaded", () => {
    document.getElementById("send").addEventListener("click", () => {
      console.log("send message");
      const message = document.getElementById("message").value;
      chrome.runtime.sendMessage({ message: message }, (response) => {
        alert(JSON.stringify(response));
        console.log(`Response from native app:`, response);
      });
    });
  });
})();
