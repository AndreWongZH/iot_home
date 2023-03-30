"use client"

import { websocketEP } from "@/data/endpoints";

var W3CWebSocket = require('websocket').w3cwebsocket;

interface SocketService {
  client: any;
}

class SocketService {
  constructor() {
    // check for cookie first

    this.client = new W3CWebSocket(websocketEP, "", {
      headers: {
        Cookie: document.cookie
      }
    })

    this.client.onerror = function() {
      console.log('Connection Error');
    };
  
    this.client.onopen = function() {
      console.log('WebSocket Client Connected');
    };
  
    this.client.onclose = function() {
      console.log('echo-protocol Client Closed');
    };
  }

  close() {
    // call this during logout
    this.client.close()
  }
}

let socket: any = null;

export function getSocketInstance() {
  if (!socket) {
    socket = new SocketService();
  }
  return socket;
}


// export function SocketConn() {
//   const pathname = usePathname();
//   const [data, setData] = useState([])
//   const [cli, setCli] = useState<any>(null)

//   // useEffect(() => {
//   //   if (cli) {
//   //     cli.close()
//   //   }
//   // }, [pathname])
  

//   useEffect(() => {
    
//     const client = new W3CWebSocket('ws://localhost:3001/ws', "", {
//       headers: {
//         Cookie: document.cookie
//       }
//     })

//     setCli(client)

//     client.onerror = function() {
//       console.log('Connection Error');
//     };
  
//     client.onopen = function() {
//         console.log('WebSocket Client Connected');
    
//         // function sendNumber() {
//         //     if (client.readyState === client.OPEN) {
//         //         var number = Math.round(Math.random() * 0xFFFFFF);
//         //         client.send(number.toString());
//         //         setTimeout(sendNumber, 1000);
//         //     }
//         // }
//         // sendNumber();
//     };
  
//     client.onclose = function() {
//         console.log('echo-protocol Client Closed');
//     };
    
//     client.onmessage = function(e: any) {
//         // if (typeof e.data === 'string') {
//         //     console.log("Received: '" + e.data + "'");
//         // }
//         console.log(e.data)
//     };

//     return () => {
//       if (cli) {
//         cli.close()
//       }
//     }

//   }, [])

//   const senddata = () => {
//     var data = {
//       success: "hello world!"
//     }

//     if (cli) {
//       cli.send(JSON.stringify(data))
//     }
//   }

//   return (
//     <>
//       <div onClick={() => {senddata()}}>click me</div>
//     </>
//   )
// }

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