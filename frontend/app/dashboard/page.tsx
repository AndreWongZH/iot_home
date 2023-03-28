"use client"

import { AddButton } from '@/components/button';
import { LinkHeader } from './linkheader';
import { useEffect, useState } from 'react';
import Loading from './loading';
import { deleteRoomEP, roomsEP } from '@/data/endpoints';
import { Toggle } from '@/components/toggle';
import { useRouter } from 'next/navigation';
import { Confirm } from 'notiflix';

const RoomTile = ({name, count, onClick} :{ name: string, count: number, onClick: Function }) => {

  return (
    <button
      className="relative h-28 text-center flex flex-col items-center justify-center rounded-lg hover:bg-roomtile-highlight"
      onClick={() => onClick(name)}
    >
      {/* <Link className="z-20 absolute h-28 w-full" href={`/dashboard/room/${name}`}></Link> */}
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
  const router = useRouter()

  const [data, setData] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState("")
  const [deleteMode, setDeleteMode] = useState(false)

  const tileOnClick = (roomName:string) => {
    if (deleteMode) {
      // delete room
      Confirm.show("Delete room",
      `Do you really want to delete ${roomName}?`,
      "Yes", "No",
      () => { 
        fetch(deleteRoomEP(roomName), {
          method: "DELETE",
          headers: {
            "Content-Type": "application/json",
          },
          credentials: "include",
        })
          .then((resp) => resp.json())
          .then(({ success, error }) => {
            if (success) {
              getRoomData()
            }
          })
      },
      () => { })

      return
    }

    // navigate to devices page
    router.push(`/dashboard/room/${roomName}`)
  }

  useEffect(() => {
    getRoomData()
  }, [])

  const getRoomData = () => {
    fetch(roomsEP,{
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
      {
        loading ? <Loading /> :
      
        (<>
          <LinkHeader headerText={"IOT Home"} href={`/dashboard/addroom`} showHome={false} disableMargin>
            <AddButton onClick={null}/>
          </LinkHeader>
          <Toggle toggleOffText={''} toggleOnText={"delete"} setMode={deleteMode} setSetMode={setDeleteMode}/>
          <div className="flex flex-col gap-5 px-4">
            {
            data.map(({name, count}) => { 
              return <RoomTile key={name} name={name} count={count} onClick={tileOnClick}/>
              })
            }
          </div>
          <div className="h-8"></div>
        </>)
      }
    </>
  )
}