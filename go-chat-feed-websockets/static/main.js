(() => {
  "use strict"

  let socket = new WebSocket("ws://localhost:3000/ws/orderbookfeed")

  socket.onopen = () => {
    console.info("WebSocket connected")

    socket.send("Hello! ASL?")
  }

  socket.onmessage = (msg) => {
    console.info("got message", msg)
  }

  socket.onclose = () => {
    console.info("WebSocket closed")
  }
})()
