// "use client"


import { Socket } from '@/components/socket';
import { AddButton, Device } from '@/components/button'
import Link from 'next/link'
import { LinkHeader } from '../../linkheader';

const obj = [
  { nickname: "Monitor", state: "on", type: "wled" },
  { nickname: "Desk lamp", state: "off", type: "wled" },
  { nickname: "head lamp", state: "on", type: "switch" },
  { nickname: "Fan", state: "on", type: "IR" },
  { nickname: "Air purifier", state: "on", type: "purifier" },
]

async function getDeviceData(roomName) {
  const res = await fetch(`http://127.0.0.1:3001/${roomName}/devices`, {
    cache: 'no-store'
  })

  if (!res.ok) {
    throw new Error("failed to fetch devices")
  }

  return res.json()
}


export default async function Page({ params }) {
  const {success, data} = await getDeviceData(params.roomname)
  console.log(data)
  
  return (
    <>
      <LinkHeader showHome={true} headerText={"Welcome home, andre"} href={`/app/dashboard/room/${params.roomname}/adddevice`}>
        <AddButton />
      </LinkHeader>
      <div className="flex flex-wrap gap-5 justify-center">
      {
        data.devList.map((dev) => {
          return <Device key={dev.nickname} nickname={dev.nickname} devStatus={data.devStatus[dev.ipaddr]} type={dev.type} roomname={params.roomname} ip={dev.ipaddr} />
        })
      }
      </div>
      {/* <Socket /> */}
    </>
  )
}