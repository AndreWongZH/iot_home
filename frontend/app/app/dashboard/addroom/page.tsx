"use client"

import { BackHeader } from '@/components/header';
import { useRouter } from 'next/navigation';
import React, { useState } from 'react';

export default function Page() {
  const router = useRouter();
  const [room, setRoom] = useState("");

  async function addRoom(event: React.SyntheticEvent) {
    event.preventDefault(); 
    console.log("adding room")

    const jsonData = {
      name: room,
    }

    await fetch(`http://127.0.0.1:3001/createrm`,
      {
        body: JSON.stringify(jsonData),
        headers: {
          'Content-Type': 'application/json'
        },
        method: 'POST'
      }
    )
    
    router.push(`/app/dashboard/room/${room}`)
  }

  return (
    <>
      <BackHeader headerText={`Create a room`} />
      <div className="flex flex-col items-center">
        <form>
          <div className="flex flex-col justify-center items-center gap-y-2.5">
            <input
              type="text"
              name="room"
              id="room"
              placeholder="room name"
              onChange={(e) => setRoom(e.target.value)}
              className="rounded-md border border-[#e0e0e0] bg-white py-3 px-3 w-full text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"
            />

            <button
              className="w-full hover:shadow-lg rounded-md bg-[#6A64F1] py-3 mt-3 text-center text-base font-semibold text-white outline-none"
              onClick={(event) => { addRoom(event) }}
            >
              Add Room
            </button>
          </div>
        </form>
      </div>
    </>
  )
}