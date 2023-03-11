const baseUrl = "http://localhost:"
const port = "3001"

export const loginEP = `${baseUrl}${port}/login`
export const logoutEP = `${baseUrl}${port}/logout`
export const userEP = `${baseUrl}${port}/user`

export const roomsEP = `${baseUrl}${port}/rooms`
export const createRoomEP = `${baseUrl}${port}/create-room`

export const addDeviceEP = (roomName: string) => {
    return `${baseUrl}${port}/${roomName}/add-device`
}
export const devicesEP = (roomName: string) => {
    return `${baseUrl}${port}/${roomName}/devices`
}

export const toggleEP = (roomName: string, ip: string, toggle: string) => {
    return `${baseUrl}${port}/${roomName}/${ip}/${toggle}`
}

export const wledConfigEP = (roomName: string, ip: string) => {
    return `${baseUrl}${port}/${roomName}/${ip}/wled-config`
}

export const websocketEP = `ws://localhost:${port}/ws`