# webrtc

**What ths end goal?**

to equip you will all the knowledge required to bulild and deploy a very simple video chat service.

**What are we going to cover?**

- Multithreading in Golang
- An Introduction to webRTC
- A brief into WebSockets
- Putting the pieces together - The Final Demo

**Prerequistes?**

- Basic JavaScript and Go
- Basic Backend Development

## Goroutine

Goroutine are light weight threads of executioin managed by the go runtime.

- In simpler terms, they are just functions you can run concurrenyly.
- Just append to "go" keyword before the function to run it as a goroutine.
- This simplicity of starting a goroutine is one of the reasons go is widely used to build backend apps that need concurrenyly.

## Channeles

`ch := make(chan int)`

- A channel is a communication mechanism that lets one goroutine send values to another goroutine.
- Each channel is a conduit for values of a particular types, called the channel's elements type.
- Channels are how goroutines communicate

## WebRTC Introduction and Demo

1. What is WebRTC?
2. WebRTC Connection Lifecycle
3. SDP and Signalling using WebSockets
4. NAT, STUN & TURN Servers, ICE Candidates
5. System Design for Final Demo

## What is WebRTC?

An open framework for the we that enables Real-Time-Communications in the browser.

WebRTC (Web Real-Time-Communication) is a technology which enables Web applicaitons and sites to capture and optioanally stream audio and/or video media, as well as to exchange arbitrary data between browsers without requiring an intermediary.

It's the tech that enables use to build something like Google Meet, by using standard Web APIs

## WebRTC Connection Lifecycle: Overview

- WebRTC works by exchanging **Offers** and **Answers** between 2 peers.
- This negotiation is called **Signalling**
- After Signalling Peers exchange p2p connectivity information using something called **ICE Candidates**
- Offers, Answers are expressed using the **Session Description Protocol** (SDP)
- Signalling Server (mostly WebSockets)

## WebRTC Connection Lifecycle: Detail

1. Create Offer
2. Local Description  = Offer
3. Send Offer
4. Receive Offer
5. Remote Description = Offer
6. Create Answer
7. Local Description = Answer
8. Send Answer
9. Receive Answer
10. Remote Description = Answer
11. Generate and exchange ICE
12. Peers Connected

## Signalling, SDP and WebSockets

The process of signalling can be done with any medium. Heck you can even tweet the offer.

But in serious terms, we usually use a WebSockets server because of uts duplex communication capabilities

We can use regular HTTP, but we'll have to use long polling

**WebSockets**

- The WebSockets API makes it possible to open a two-way interactive communication session bwetween the user's browser and a server.
- With this API, we can send messages to a server and receive event-drivent responses without having to pol the server for a reply.
- <https://tools.ietf.org/html/rfc6455>
- We'll be using github.com/gorilla/websocket to setup a WS server in GO

**Session Description Protocol (SDP)**

- The configuration of an endpoint of a WebRTC connection is called a session description. It is expressed using the SDP.
- The description includses information about the kind of media being sent, its format, the transfer protocol being used, the endpoint's IP address and port, and other information needed to describe a mediaa transfer endpoint.
- Offers and Answers are special descriptions.
- Each peer keeps two descriptions on hand: **the local description**, describing itself, and **the remote description**, describing the other end of the call.

## Okay, What exactly happens when the "Peers Connect"

- Once the negotiation is done, we try to establish a p2p connection
- An external service is usually used for discovering the possible candidates for connecting to a peer. (Trying to find a path in the graph)
- This service is called **ICE** (interactive Connectivity Establishment) and is using either a STUN or a TURN server.
- STUN stands for Session Traversal Utilities for NAT, and is usually used indirectly in most WebRTC applications.

## ICE Candidates

- ICE candidates detail the available methods the peer is able to communicate (directly or through a TURN server)
- Typically, each peer will propose its best candidates first, making their way down the line toward their worse candidates.

> Ideally, candidates are UDP (since it's faster, and media streams are able to recover from interruptions relatively easily) but the ICE standard does allow TCP candidates as well.

**Some STUN servers you can use**

- stun.l.google.com:19302
- stun1.l.google.com:19302
- stun2.l.google.com:19302
- stun3.l.google.com:19302
- stun4.l.google.com:19302
- stun.ekiga.net
- stun.ideasip.com
- stun.rixtelecom.se
- stun.schlund.de
- stun.stunprotocol.org:3478
- stun.voiparound.com
- stun.voipbuster.com
- stun.voipstunt.com
- stun.voxgratia.org

Full list Available at bit.ly/webrtc-stun-list

## The Entire Exchange in a VERY Complicated Diagram

https://developer.mozilla.org/en-US/docs/Web/API/WebRTC_API/Connectivity/#ice_candidates
