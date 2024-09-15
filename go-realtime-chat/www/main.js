(() => {
  'use strict'
  let btnJoin
  let ws

  async function login(data) {
    return await fetch('/login', {
      method: 'post',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    }).then((res) => res.json())
  }

  async function signup(data) {
    return await fetch('/signup', {
      method: 'post',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    }).then((res) => res.json())
  }

  async function createRoom(data) {
    return await fetch('/ws/createRoom', {
      method: 'post',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    }).then((res) => res.json())
  }

  async function getRooms() {
    return await fetch('/ws/getRooms', {
      method: 'get',
      headers: {
        'Accept': 'application/json'
      },
    }).then((res) => res.json())
  }

  async function joinRoom(id) {
    //return await fetch(`/ws/joinRoom/${id}`, {
    //  method: 'get',
    //  headers: {
    //    'Accept': 'application/json'
    //  },
    //}).then((res) => res.json())
    //
    ws = new WebSocket(`/ws/joinRoom/${id}`)

    ws.addEventListener('open', (e) => {
      console.log('ws connected', e)
    })

    ws.addEventListener('message', (msg) => {
      console.log('ws message', msg)
    })

    ws.addEventListener('error', (e) => {
      console.error('ws error', e)
    })

    ws.addEventListener('close', (e) => {
      console.log('ws close', e)
    })
  }

  async function sendMessage(data) {

  }

  async function getClients(id) {
    return await fetch(`/ws/getClients/${id}`, {
      method: 'get',
      headers: {
        'Accept': 'application/json'
      },
    }).then((res) => res.json())
  }

  let roomHtml = (id) => `<li class="border rounded border-zinc-300 p-5 flex items-center gap-3 hover:border-blue-600">
        <div class="flex-1">
          <p class="text-sm text-zinc-400">Room</p>
          <h3 class="text-lg text-blue-600 font-medium">${id}</h3>
        </div>
        <button data-id="${id}"
          class="btn-join py-1 px-3 rounded border border-blue-600 bg-blue-600 text-white hover:bg-blue-700 flex items-center text-sm font-medium justify-center">join</button>
      </li>`

  function getAuth() {
    try {
      return JSON.parse(localStorage.getItem("auth"))
    } catch (err) {
      throw new Error('auth not found')
    }
  }

  async function updateRooms() {
    const listRooms = document.getElementById('list-rooms')
    await getRooms().then((res) => {
      if (res.length > 0) {
        let html = ''
        res.map((room) => {
          html += roomHtml(room.id)
        })
        listRooms.innerHTML = html
        btnJoin = document.querySelectorAll(".btn-join")
        btnJoin.forEach((btn) => {
          btn.addEventListener('click', (e) => {
            let id = e.target.getAttribute('data-id')
            joinRoom(id)
          })
        })
      }
    })

  }

  document.addEventListener('DOMContentLoaded', async () => {
    const formLogin = document.getElementById('login')
    const formSignup = document.getElementById('signup')
    const formCreateRoom = document.getElementById('create-room')
    const formMessage = document.getElementById('send-message')
    const btnSignup = document.getElementById('btn-signup')
    const btnLogin = document.getElementById('btn-login')

    btnLogin.addEventListener('click', () => {
      formSignup.classList.remove('hidden')
      formLogin.classList.add('hidden')
    })

    btnSignup.addEventListener('click', () => {
      formLogin.classList.remove('hidden')
      formSignup.classList.add('hidden')
    })

    const auth = getAuth()

    if (auth) {
      updateRooms()
    }

    if (!auth) {
      formLogin.classList.remove('hidden')
      formLogin.addEventListener('submit', async (e) => {
        e.preventDefault()
        let data = new FormData(formLogin)
        await login(Object.fromEntries(data)).then((res) => {
          auth = res
          localStorage.setItem("auth", JSON.stringify(res))
          formLogin.classList.add('hidden')
        }).catch((err) => {
          console.error(err)
        })
      })
    }

    formSignup.addEventListener('submit', async (e) => {
      e.preventDefault()
      let data = new FormData(formLogin)
      await signup(Object.fromEntries(data)).then((res) => {
        console.log(res)
        formSignup.classList.add('hidden')
        formLogin.classList.remove('hidden')
      }).catch((err) => {
        console.error(err)
      })
    })

    formCreateRoom.addEventListener('submit', async (e) => {
      e.preventDefault()
      let data = new FormData(formCreateRoom)
      data.set("name", auth.username)
      await createRoom(Object.fromEntries(data)).then(async (res) => {
        console.log(res)
        updateRooms()
      }).catch((err) => {
        console.error(err)
      })
    })

    formMessage.addEventListener('submit', async (e) => {
      e.preventDefault()
      let data = new FormData(formCreateRoom);
      data.set("name", auth.username)
    })
  })
})()
