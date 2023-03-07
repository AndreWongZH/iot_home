"use client"

import { useEffect, useState } from 'react'
// import WebSocket from 'websocket'
var W3CWebSocket = require('websocket').w3cwebsocket;

export function SocketConn() {
  const [data, setData] = useState([])
  const [cli, setCli] = useState(null)

  useEffect(() => {
    console.log(document.cookie)
    const client = new W3CWebSocket('ws://localhost:3001/ws', "", {
      headers: {
        Cookie: document.cookie
      }
    })

    setCli(client)

    client.onerror = function() {
      console.log('Connection Error');
    };
  
    client.onopen = function() {
        console.log('WebSocket Client Connected');
    
        // function sendNumber() {
        //     if (client.readyState === client.OPEN) {
        //         var number = Math.round(Math.random() * 0xFFFFFF);
        //         client.send(number.toString());
        //         setTimeout(sendNumber, 1000);
        //     }
        // }
        // sendNumber();
    };
  
    client.onclose = function() {
        console.log('echo-protocol Client Closed');
    };
    
    client.onmessage = function(e: any) {
        // if (typeof e.data === 'string') {
        //     console.log("Received: '" + e.data + "'");
        // }
        console.log(e.data)
    };

  }, [data])

  const senddata = () => {
    var data = {
      success: "hello world!"
    }

    if (cli) {
      cli.send(JSON.stringify(data))
    }
  }

  return (
    <>
      <div onClick={() => {senddata()}}>click me</div>
    </>
  )
}

// import io from 'socket.io-client';
// import { useEffect } from 'react';

// let socket = io("http://localhost:3001", {
//   withCredentials: true,
//   transports: ['polling']
// })

// export function Socket() {
//   // useEffect(() => initSocket(), [])

//   useEffect(() => {
//     if (socket) {
//       // socket.on("device_status", (msg) => {
//       //   const obj = JSON.parse(msg)
//       //   console.log(msg)
//       // })

//       socket.on("recvStatus", (msg) => {
//         console.log(msg)
//       })

//       return () => {
//         socket.off();
//       };
//     }
//   }, [])

//   // const initSocket = () => {
//   //   socket = io("http://localhost:3001", {
//   //     withCredentials: true,
//   //   })

//   //   socket.on("connect", (msg) => {
//   //     console.log("connected")
//   //   })

//   //   socket.on("recvStatus", (msg) => {
//   //     console.log(msg)
//   //   })

//   //   // return () => {
//   //   //   socket.off();
//   //   // };
//   // }

//   const emitMsg = () => {
//     console.log("emitting message")
//     socket.emit("testping", "pong")
//     socket.emit("devStatus", JSON.stringify({
//       "roomname": "bedroom",
//       "ipaddr": "192.128.0.0"
//     }))
//   }

//   return (
//     <div onClick={() => emitMsg()}>
//       <p>
//         socket component
//       </p>
//     </div>
//   )
// }