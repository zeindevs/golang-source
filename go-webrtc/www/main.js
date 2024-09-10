(() => {
  'use strict'
  let ws
  let video
  let partner
  let peer
  let userStream

  let player = (id, col) => `
    <div class="video-player border col-span-${col} rounded border-zinc-800 bg-zinc-900 aspect-video">
      <video id="player-${id}" class="w-full h-full object-fit" autoplay controls></video>
    </div>`

  async function createRoom() {
    const resp = await fetch('/create')
    const { room_id } = await resp.json()
    return room_id
  }

  async function openCamera() {
    const allDevices = await navigator.mediaDevices.enumerateDevices()
    const cameras = allDevices.filter((d) => d.kind === 'videoinput')

    console.log(cameras)

    const constraints = {
      audio: true,
      video: {
        deviceId: cameras[0].deviceId
      }
    }

    try {
      return await navigator.mediaDevices.getUserMedia(constraints)
    } catch (err) {
      console.error('Open Camera error:', err)
    }
  }

  async function connectWS(roomId) {
    const main = document.getElementById('main')
    const connect = document.getElementById('connect')

    connect.classList.add('hidden')
    let html = player('me', 1)
    html += player('partner', 1)

    main.innerHTML = html

    openCamera().then((stream) => {
      video = document.getElementById(`player-me`)
      partner = document.getElementById(`player-partner`)

      video.srcObject = stream
      userStream = stream

      ws = new WebSocket(`/join?roomID=${roomId}`)

      ws.addEventListener('open', () => {
        ws.send(JSON.stringify({ join: 'true' }))
      })

      ws.addEventListener('message', async (msg) => {
        console.log('WS message:', msg)
        try {
          const message = JSON.parse(msg.data)

          if (message.join) {
            callUser()
          }

          if (message.iceCandidate) {
            console.log('receiving and adding ice candidate')

            if (peer) {
              try {
                await peer.addIceCandidate(message.iceCandidate)
              } catch (err) {
                console.error(`error receiving ice candidate:`, err)
              }
            }
          }

          if (message.offer) {
            handleOffer(message.offer)
          }

          if (message.answer) {
            console.log('receiving answer')

            if (peer) {
              peer.setRemoteDescription(new RTCSessionDescription(message.answer))
            }
          }
        } catch (err) {
          console.log('handle WS message err:', err)
        }
      })
    })


  }

  async function handleOffer(offer) {
    console.log('received offer, creating answer')

    peer = createPeer()

    await peer.setRemoteDescription(new RTCSessionDescription(offer))

    userStream.getTracks().forEach((track) => {
      peer.addTrack(track, userStream)
    })

    const answer = await peer.createAnswer()
    await peer.setLocalDescription(answer)

    ws.send(JSON.stringify({ answer: peer.localDescription }))
  }

  function callUser() {
    console.log('calling other user')

    peer = createPeer()

    userStream.getTracks().forEach((track) => {
      peer.addTrack(track, userStream)
    })
  }

  function createPeer() {
    console.log('creating peer connection')

    const peer = new RTCPeerConnection({
      iceServers: [{ urls: 'stun:stun.l.google.com:19302' }]
    })

    peer.onnegotiationneeded = handleNegotiationNeeded
    peer.onicecandidate = handleIceCandidateEvent
    peer.ontrack = handleTrackEvent

    return peer
  }

  async function handleNegotiationNeeded() {
    console.log('creating offer')

    try {
      const myOffer = await peer.createOffer()
      await peer.setLocalDescription(myOffer)

      ws.send(JSON.stringify({ offer: peer.localDescription }))
    } catch (err) {
      console.error(err)
    }
  }

  function handleIceCandidateEvent(e) {
    console.log('found ice candidate')

    if (e.candidate) {
      console.log(e.candidate)

      ws.send(JSON.stringify({ iceCandidate: e.candidate }))
    }
  }

  function handleTrackEvent(e) {
    console.log('received tracks', e.streams)

    partner.srcObject = e.streams[0]
  }

  document.addEventListener('DOMContentLoaded', () => {
    const create = document.getElementById('create')
    const form = document.getElementById('form')
    const room = document.getElementById('roomId')
    const connect = document.getElementById('connect')

    create.addEventListener('click', async () => {
      let roomId = await createRoom()

      console.log(`roomId:`, roomId)

      action.classList.add('hidden')
      connect.classList.remove('hidden')

      connectWS(roomId)
    })

    form.addEventListener('submit', (e) => {
      e.preventDefault()

      console.log(`join roomId:`, room.value)

      action.classList.add('hidden')
      connect.classList.remove('hidden')

      connectWS(room.value)
    })
  })
})()
