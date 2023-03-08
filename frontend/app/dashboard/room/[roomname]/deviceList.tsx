"use client"

import { Device } from "@/components/button"
import { Toggle } from "@/components/toggle"
import { useState } from "react"

interface DeviceInfo {
  hostname: string;
  ipaddr: string;
  name: string;
  type: string;
}

interface DeviceStatus {
  connected: boolean;
  on_state: boolean;
}

interface RoomDevicesData {
  devList: Array<DeviceInfo>;
  devStatus: Record<string, DeviceStatus>;
}

export const DeviceList = ({ data, roomName }: {data: RoomDevicesData, roomName: string}) => {
  const [setMode, setSetMode] = useState(false)

  return (
    <>
      <Toggle setMode={setMode} setSetMode={setSetMode} />


      <div className="flex flex-wrap gap-5 justify-center">
      {
        data.devList.map((dev) => {
          return <Device setMode={setMode} key={dev.name} name={dev.name} devStatus={data.devStatus[dev.ipaddr]} type={dev.type} roomName={roomName} ip={dev.ipaddr} />
        })
      }
      </div>
    </>
  )
}