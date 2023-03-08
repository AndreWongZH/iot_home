"use client"

import { AddButton } from '@/components/button'
import Link from 'next/link'
import { LinkHeader } from '../../linkheader';
import { DeviceList } from './deviceList';
import { useEffect, useState } from 'react';
import Loading from '../../loading';
import { getSocketInstance } from '@/components/socket';

export default function Page({ params }: { params: {roomname: string;}}) {
  const [data, setData] = useState({devList: [], devStatus: {}})
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    getDeviceData()
  }, [])

  useEffect(() => {
    let soc = getSocketInstance()
    soc.client.onmessage = function(e: any) {
      // if (typeof e.data === 'string') {
      //     console.log("Received: '" + e.data + "'");
      // }
      console.log(e.data)
  };
  }, [])

  const getDeviceData = () => {
    fetch(`http://localhost:3001/${params.roomname}/devices`,{
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    })
    .then((resp) => {
      return resp.json()
    })
    .then(({ success, data }) => {
      setData(data)
      setLoading(false)
    })

  }

  return (
    <>
      <LinkHeader disableMargin={true} showHome={true} headerText={"Welcome home, XXX"} href={`/dashboard/room/${params.roomname}/adddevice`}>
        <AddButton onClick={null}/>
      </LinkHeader>

      {
        loading ? <Loading /> : <DeviceList data={data} roomName={params.roomname}/>
      }
      
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
    </>
  )
}