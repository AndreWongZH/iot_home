import { LinkHeader } from '@/components/header';
import Link from 'next/link';

import { AddButton } from '../../../components/button'

async function getRoomData() {
  const res = await fetch('http://127.0.0.1:3001/rooms')

  if (!res.ok) {
    throw new Error("failed to fetch rooms")
  }

  return res.json()
}

const RoomTile = ({name, devices} :{ name: string, devices: int }) => {
  return (
    <button className="relative h-28 text-center flex flex-col items-center justify-center rounded-lg hover:bg-roomtile-highlight">
      <Link className="z-20 absolute h-28 w-full" href={`/app/dashboard/room/${name}`}></Link>
      <div className="absolute flex flex-col items-center justify-center z-10">
        <h1 className="text-white font-bold capitalize text-2xl">{name}</h1>
        <p className="text-white font-bold capitalize text-xs">{devices} devices connected</p>
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
  devices: RegisteredDevice[];
}

interface RegisteredDevice {
  hostname: string;
  ipaddr: string;
  nickname: string;
  type: string;
}

export default async function Page() {
  const rooms: Room[] = await getRoomData();
  console.log(rooms)

  return (
    <>
      <LinkHeader headerText={"Header"} href={`/app/dashboard/addroom`}>
        <AddButton />
      </LinkHeader>
      <div className="flex flex-col gap-5 px-4">
        {
        rooms.map(({name, devices}) => { 
          return <RoomTile key={name} name={name} devices={devices.length} />
          })
        }
      </div>
      <div className="h-8"></div>
    </>
  )
}