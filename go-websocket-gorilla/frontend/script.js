(() => {
  "use strict";
  let conn;
  let selectedChat = "general";

  class Event {
    constructor(type, payload) {
      this.type = type;
      this.payload = payload;
    }
  }

  class SendMessageEvent {
    constructor(message, from) {
      this.message = message;
      this.from = from;
    }
  }

  class NewMessageEvent {
    constructor(message, from, sent) {
      this.message = message;
      this.from = from;
      this.sent = sent;
    }
  }

  class ChangeChatRoomEvent {
    constructor(name) {
      this.name = name;
    }
  }

  function routeEvent(event) {
    if (event.type === undefined) {
      alert("no 'type' field in event");
    }
    switch (event.type) {
      case "new_message":
        console.log("new message");
        const messageEvent = Object.assign(
          new NewMessageEvent(),
          event.payload,
        );
        appendChatMessage(messageEvent);
        break;
      default:
        alert("unsupported message type");
        break;
    }
  }

  function appendChatMessage(messageEvent) {
    let date = new Date(messageEvent.sent);

    const formattedMsg = `${date.toLocaleString()}: ${messageEvent.message}`;

    let textarea = document.getElementById("chatmessages");
    textarea.innerHTML = textarea.innerHTML + "\n" + formattedMsg;
    textarea.scrollTop = textarea.scrollHeight;
  }

  function changeChatRoom(e) {
    e.preventDefault();
    let newChat = document.getElementById("chatroom");
    if (newChat != null && newChat.value != selectedChat) {
      selectedChat = newChat.value;
      document.getElementById("chat-header").innerHTML =
        "Currently in chat: " + selectedChat;

      let changeEvent = new ChangeChatRoomEvent(selectedChat);
      sendEvent("change_room", changeEvent);
      let textarea = document.getElementById("chatmessages");
      textarea.innerHTML = `You changed room into: ${selectedChat}`;
      console.log(newChat.value);
    }
    return false;
  }

  function sendMessage(e) {
    e.preventDefault();
    let newMessage = document.getElementById("message");
    if (newMessage != null) {
      console.log(newMessage.value);
      let outgointEvent = new SendMessageEvent(newMessage.value, "percy");
      sendEvent("send_message", outgointEvent);
    }
    return false;
  }

  function sendEvent(eventName, payload) {
    const event = new Event(eventName, payload);
    conn.send(JSON.stringify(event));
  }

  function login(e) {
    e.preventDefault();
    let formData = {
      username: document.getElementById("username").value,
      password: document.getElementById("password").value,
    };

    fetch("/login", {
      method: "post",
      body: JSON.stringify(formData),
      mode: "cors",
    })
      .then((response) => {
        if (response.ok) {
          return response.json();
        } else {
          throw "unauthorized";
        }
      })
      .then((data) => {
        connectWebsocket(data.otp);
      })
      .catch((e) => {
        alert(e);
      });
    return false;
  }

  function connectWebsocket(otp) {
    if (window["WebSocket"]) {
      console.log("supports WebSockets");

      conn = new WebSocket(
        "wss://" + document.location.host + "/ws?otp=" + otp,
      );

      conn.onopen = function (evt) {
        document.getElementById("connection-header").innerHTML =
          "Connected to Websocket: true";
      };

      conn.onclose = function (evt) {
        // Set disconnected
        document.getElementById("connection-header").innerHTML =
          "Connected to Websocket: false";
      };

      conn.onmessage = function (evt) {
        console.log(evt);

        const eventData = JSON.parse(evt.data);
        const event = Object.assign(new Event(), eventData);

        routeEvent(event);
      };
    } else {
      alert("Not supporting WebSockets");
    }
  }

  window.onload = function () {
    document.getElementById("chatroom-selection").onsubmit = changeChatRoom;
    document.getElementById("chatroom-message").onsubmit = sendMessage;
    document.getElementById("login-form").onsubmit = login;
  };
})();
