(() => {
  "use strict";

  document.addEventListener("DOMContentLoaded", () => {
    const go = new Go(); // Defined in wasm_exec.js
    const WASM_URL = "wasm.wasm";

    // Providing the environment object, used in WebAssembly.instantiateStreaming.
    // This part goes after "const go = new Go();" declaration.
    go.importObject.env = {
      add: function (x, y) {
        return x + y;
      },
      // ... other functions
    };

    let wasm;

    function updateResult() {
      wasm.exports.update();
    }

    document.querySelector("#a").oninput = updateResult;
    document.querySelector("#b").oninput = updateResult;

    if ("instantiateStreaming" in WebAssembly) {
      WebAssembly.instantiateStreaming(fetch(WASM_URL), go.importObject).then(
        function (obj) {
          wasm = obj.instance;
          go.run(wasm);
          updateResult();

          // Calling the multiply function:
          console.log("multiplied two numbers:", wasm.exports.multiply(5, 3));
        },
      );
    } else {
      fetch(WASM_URL)
        .then((resp) => resp.arrayBuffer())
        .then((bytes) =>
          WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
            wasm = obj.instance;
            go.run(wasm);
            updateResult();

            // Calling the multiply function:
            console.log("multiplied two numbers:", wasm.exports.multiply(5, 3));
          }),
        );
    }
  });
})();
