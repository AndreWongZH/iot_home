"use client"

import { Socket } from '@/components/socket';
import { AddButton } from '@/components/button'
import Link from 'next/link'
import { LinkHeader } from '../../linkheader';
import { DeviceList } from './deviceList';
import instance from '@/components/axiosInst';
import { useEffect, useState } from 'react';
import Loading from '../../loading';

export default function Page({ params }: { params: {roomname: string;}}) {
  const [data, setData] = useState({devList: [], devStatus: {}})
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    getDeviceData()
  }, [])

  const getDeviceData = () => {
    instance.get(`${params.roomname}/devices`)
    .then(function (resp) {
      const {success, data} = resp.data
      setData(data)
      setLoading(false)
    })
    .catch(function (err) {

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
      {/* <Socket /> */}
    </>
  )
}