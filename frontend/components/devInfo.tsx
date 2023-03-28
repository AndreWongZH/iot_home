"use client"

import { deviceInfoEP } from "@/data/endpoints"
import { useEffect, useState } from "react"

interface DeviceInfo {
  name: string;
  ipaddr: string;
  type: string;
}

const DefaultDeviceInfo: DeviceInfo = {
  name: "",
  ipaddr: "",
  type: "",
}

export const DeviceInfo = ({ roomName, ip }: { roomName: string, ip: string }) => {
  const [deviceInfo, setDeviceInfo] = useState<DeviceInfo>(DefaultDeviceInfo)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState("")

  useEffect(() => {
    getDeviceInfo()
  }, [])

  const getDeviceInfo = () => {
    fetch(deviceInfoEP(roomName, ip), {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    })
    .then((resp) => resp.json())
    .then(({ success, data, error }) => {
      if (success) {
        setDeviceInfo(data)
      } else {
        setError(error)
      }
      setLoading(false)
    })
  }

  return (
    <>
      {
        loading ? <></> :

        (
          <div className="w-4/5 mx-auto py-8">
            <h1 className="font-bold text-xl text-slate-600">name: {deviceInfo.name}</h1>
            <h1 className="font-bold text-xl text-slate-600">ip address: {deviceInfo.ipaddr}</h1>
            <h1 className="font-bold text-xl text-slate-600">device type: {deviceInfo.type}</h1>
          </div>
        )
      }
    </>
  )

}