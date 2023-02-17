// "use client"


import { Socket } from '@/components/socket';
import { AddButton } from '@/components/button'
import { LinkHeader } from '@/components/header'
import Link from 'next/link'
import { GoLightBulb } from 'react-icons/go'
import { ImSwitch } from 'react-icons/im'
import { TbDeviceSpeaker } from 'react-icons/tb'

const obj = [
  { nickname: "Monitor", state: "on", type: "wled" },
  { nickname: "Desk lamp", state: "off", type: "wled" },
  { nickname: "head lamp", state: "on", type: "switch" },
  { nickname: "Fan", state: "on", type: "IR" },
  { nickname: "Air purifier", state: "on", type: "purifier" },
]

const getDeviceIcon = (type) => {
  if (type === "wled") {
    return <GoLightBulb size={80} />
  } else if (type === "switch") {
    return <ImSwitch size={70} />
  }

  return <TbDeviceSpeaker size={70} />
}

const Device = ({ nickname, currentStatus, type }) => {
  let icon = getDeviceIcon(type)

  return (
    <button className="h-48 w-48 bg-white flex flex-col p-3 rounded-md drop-shadow shadow hover:bg-highlight focus:shadow-outline focus:outline-none hover:text-gray-500">
      <div className="flex-1 pl-1 pt-2.5">
        {icon}
      </div>
      <div className='pl-3'>
        <h3 className="font-bold text-lg">{nickname}</h3>
        <p className="font-light text-left">{currentStatus.state ? currentStatus.on ? "on" : "off" : "disconnected"}</p>
      </div>
    </button>
  )
}

async function getDeviceData(roomname) {
  const res = await fetch(`http://127.0.0.1:3001/${roomname}/devices`)

  if (!res.ok) {
    throw new Error("failed to fetch rooms")
  }

  return res.json()
}


export default async function Page({ params }) {
  console.log(params.roomname)
  const devicesData = await getDeviceData(params.roomname)
  console.log(devicesData)
  
  return (
    <>
      <LinkHeader headerText={"Welcome home, andre"} href={`/app/dashboard/room/${params.roomname}/adddevice`}>
        <AddButton />
      </LinkHeader>
      <div className="flex flex-wrap gap-5 justify-center">
      {
        devicesData.devices.map((dev) => {
          return <Device key={dev.nickname} nickname={dev.nickname} currentStatus={devicesData.DeviceInfo[dev.ipaddr]} type={dev.type} />
        })
      }
      </div>
      <Socket />
    </>
  )
}