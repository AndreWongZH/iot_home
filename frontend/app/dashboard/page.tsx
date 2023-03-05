"use client"

import Link from 'next/link';
import { AddButton } from '@/components/button';
import { LinkHeader } from './linkheader';
import instance from '@/components/axiosInst';
import { useEffect, useState } from 'react';

const RoomTile = ({name, count} :{ name: string, count: number }) => {
  return (
    <button className="relative h-28 text-center flex flex-col items-center justify-center rounded-lg hover:bg-roomtile-highlight">
      <Link className="z-20 absolute h-28 w-full" href={`/dashboard/room/${name}`}></Link>
      <div className="absolute flex flex-col items-center justify-center z-10">
        <h1 className="text-white font-bold capitalize text-2xl">{name}</h1>
        <p className="text-white font-bold capitalize text-xs">{count} devices connected</p>
      </div>
      <div className="absolute h-28 rounded-lg w-full bg-roomtile z-5"></div>
      <div
        className="bg-cover bg-center absolute h-28 rounded-lg w-full opacity-70 z-0"
        style={{
          backgroundImage: `url('/images/room.jpg')`,
        }}
      >
      </div>
    </button>
  )
}

interface Room {
  name: string;
  count: number;
}

interface RegisteredDevice {
  hostname: string;
  ipaddr: string;
  nickname: string;
  type: string;
}

export default function Page() {
  // const data = await getRoomData()
  // console.log(data)
  // const {success, data} : {success: Boolean, data: Room[]} = await getRoomData();
  // console.log(data)

  const [data, setData] = useState([])
  const [error, setError] = useState("")

  useEffect(() => {
    getRoomData()
  }, [])

  const getRoomData = () => {
    instance.get('rooms')
      .then(function (resp) {
        const {success, data} = resp.data
        console.log(success)
        setData(data)
      })
      .catch(function (err) {

      })
  }


  console.log(data)
  return (
    <>
      <LinkHeader headerText={"IOT Home"} href={`/dashboard/addroom`} showHome={false}>
        <AddButton onClick={null}/>
      </LinkHeader>
      <div className="flex flex-col gap-5 px-4">
        {
        data.map(({name, count}) => { 
          return <RoomTile key={name} name={name} count={count} />
          })
        }
      </div>
      <div className="h-8"></div>
    </>
  )
}