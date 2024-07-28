chrome.runtime.onInstalled.addListener(() => {
  console.log("Extension Instaled");
});

function sendNativeMessage(message) {
  return new Promise((resolve, reject) => {
    const port = chrome.runtime.connectNative("com.example.native_messaging");
    console.log("Native messaging connected");
    port.onMessage.addListener((response) => {
      console.log("Reply", response);
      resolve(response);
      port.disconnect();
    });
    port.onDisconnect.addListener(() => {
      console.log("Native messaging disconnect");

      if (chrome.runtime.lastError) {
        console.error(chrome.runtime.lastError);
        reject(chrome.runtime.lastError);
      }
    });
    port.postMessage(message);
  });
}

chrome.runtime.onMessage.addListener((request, sender, sendResponse) => {
  if (request.message) {
    sendNativeMessage({ text: request.message })
      .then((response) => {
        sendResponse(response);
      })
      .catch((error) => {
        sendResponse({ error: error.message });
      });
  }
  return true;
});
