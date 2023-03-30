
let baseUrl: string
let wsBaseURL: string
let LOCALIP: string

if (process.env.NEXT_PUBLIC_LOCALIP) {
    LOCALIP = process.env.NEXT_PUBLIC_LOCALIP
    baseUrl = "http://"+LOCALIP+":"
    wsBaseURL = "ws://"+LOCALIP+":"
    console.log("running backend ip at: ")
    console.log(LOCALIP)
} else {
    console.log("using backend localhost")
    baseUrl = "http://localhost:"
    wsBaseURL = "ws://localhost:"
}

const port = "3001"

export const discoverEP = `${baseUrl}${port}/discover`

export const loginEP = `${baseUrl}${port}/login`
export const registerEP = `${baseUrl}${port}/register`
export const logoutEP = `${baseUrl}${port}/logout`
export const userEP = `${baseUrl}${port}/user`

export const roomsEP = `${baseUrl}${port}/rooms`
export const createRoomEP = `${baseUrl}${port}/create-room`
export const deleteRoomEP = (roomName: string) => {
    return `${baseUrl}${port}/delete-room/${roomName}`
}

export const addDeviceEP = (roomName: string) => {
    return `${baseUrl}${port}/${roomName}/add-device`
}
export const devicesEP = (roomName: string) => {
    return `${baseUrl}${port}/${roomName}/devices`
}

export const deviceInfoEP = (roomName: string, ip: string) => {
    return `${baseUrl}${port}/${roomName}/${ip}`
}
export const delDeviceEP = (roomName: string, ip: string) => {
    return `${baseUrl}${port}/${roomName}/${ip}/delete-device`
}
export const toggleEP = (roomName: string, ip: string, toggle: string) => {
    return `${baseUrl}${port}/${roomName}/${ip}/${toggle}`
}

export const wledConfigEP = (roomName: string, ip: string) => {
    return `${baseUrl}${port}/${roomName}/${ip}/wled-config`
}

export const websocketEP = `${wsBaseURL}${port}/ws`
