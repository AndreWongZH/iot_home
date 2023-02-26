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
      {
        data.devList.length > 0
          ? <></>
          :
          <div className='flex flex-col items-center justify-center'>
            <h1 className='font-bold text-xl mb-3'>You have no devices</h1>
            <h2 className='mb-5'>How about adding one now?</h2>
            
            <Link className='w-full px-10' href={`/app/dashboard/room/${params.roomname}/adddevice`}>
              <button className="w-full hover:shadow-lg rounded-md bg-[#6A64F1] py-3 mt-3 text-center text-base font-semibold text-white outline-none">
                + Add A Device
              </button>
            </Link>
          </div>
      }
      {/* <Socket /> */}
    </>
  )
}