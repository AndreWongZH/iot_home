// "use client"


import { Socket } from '@/components/socket';
import { AddButton, Device } from '@/components/button'
import Link from 'next/link'
import { LinkHeader } from '../../linkheader';
import { DeviceList } from './deviceList';

async function getDeviceData(roomName: string) {
  const res = await fetch(`http://127.0.0.1:3001/${roomName}/devices`, {
    cache: 'no-store'
  })

  if (!res.ok) {
    throw new Error("failed to fetch devices")
  }

  return res.json()
}


export default async function Page({ params }: { params: {roomname: string;}}) {
  const {success, data} = await getDeviceData(params.roomname)
  console.log(data)

  return (
    <>
      <LinkHeader disableMargin={true} showHome={true} headerText={"Welcome home, XXX"} href={`/dashboard/room/${params.roomname}/adddevice`}>
        <AddButton onClick={null}/>
      </LinkHeader>

      <DeviceList data={data} roomName={params.roomname}/>
      {
        data.devList.length > 0
          ? <></>
          :
          <div className='flex flex-col items-center justify-center'>
            <h1 className='font-bold text-xl mb-3'>You have no devices</h1>
            <h2 className='mb-5'>How about adding one now?</h2>
            
            <Link className='w-full px-10' href={`/dashboard/room/${params.roomname}/adddevice`}>
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