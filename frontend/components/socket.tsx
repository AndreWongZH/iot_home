"use client"

import io from 'socket.io-client';
import { useEffect } from 'react';

let socket = io("http://localhost:3001", {
  withCredentials: true,
  transports: ['polling']
})

export function Socket() {
  // useEffect(() => initSocket(), [])

  useEffect(() => {
    if (socket) {
      // socket.on("device_status", (msg) => {
      //   const obj = JSON.parse(msg)
      //   console.log(msg)
      // })

      socket.on("recvStatus", (msg) => {
        console.log(msg)
      })

      return () => {
        socket.off();
      };
    }
  }, [])

  // const initSocket = () => {
  //   socket = io("http://localhost:3001", {
  //     withCredentials: true,
  //   })

  //   socket.on("connect", (msg) => {
  //     console.log("connected")
  //   })

  //   socket.on("recvStatus", (msg) => {
  //     console.log(msg)
  //   })

  //   // return () => {
  //   //   socket.off();
  //   // };
  // }

  const emitMsg = () => {
    console.log("emitting message")
    socket.emit("testping", "pong")
    socket.emit("devStatus", JSON.stringify({
      "roomname": "bedroom",
      "ipaddr": "192.128.0.0"
    }))
  }

  return (
    <div onClick={() => emitMsg()}>
      <p>
        socket component
      </p>
    </div>
  )
}