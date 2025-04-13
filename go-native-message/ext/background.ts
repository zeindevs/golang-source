chrome.runtime
  .sendNativeMessage("com.example.native_messaging", {
    text: "Hello from Chrome",
  })
  .then((response) => {
    console.log("Received: " + response.text);
  });
