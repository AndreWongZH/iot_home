"use client"

import { BackHeader } from '@/components/header';
import { usePathname } from 'next/navigation';
import { useState } from 'react';

export default function Page() {
  const pathname = usePathname()
  const roomname = pathname?.split("/")[4]

  const [nickname, setNickname] = useState("");
  const [hostname, setHostname] = useState("");
  const [ipaddr, setIpaddr] = useState("");
  const [type, setType] = useState("");

  async function addDevice(event) {
    event.preventDefault(); 
    console.log("adding device")

    const jsonData = {
      nickname: nickname,
      hostname: hostname,
      ipaddr: ipaddr,
      type: type
    }

    console.log(jsonData)

    await fetch(`http://127.0.0.1:3001/${roomname}/add_device`,
      {
        body: JSON.stringify(jsonData),
        headers: {
          'Content-Type': 'application/json'
        },
        method: 'POST'
      }
    )

    router.replace(`/app/dashboard/room/${roomname}`)
  }

  return (
    <>
      <BackHeader />
      <div className="flex flex-col items-center">
        <form>
          <div className="flex flex-col justify-center items-center gap-y-2.5">
            <input
              type="text"
              name="nickname"
              id="nickname"
              placeholder="nickname"
              onChange={(e) => setNickname(e.target.value)}
              className="rounded-md border border-[#e0e0e0] bg-white py-3 px-3 w-full text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"
            />

            <input
              type="text"
              name="hostname"
              id="hostname"
              placeholder="hostname"
              onChange={(e) => setHostname(e.target.value)}
              className="rounded-md border border-[#e0e0e0] bg-white py-3 px-3 w-full text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"
            />

            <input
              type="text"
              name="ipaddr"
              id="ipaddr"
              placeholder="IP Address"
              onChange={(e) => setIpaddr(e.target.value)}
              className="rounded-md border border-[#e0e0e0] bg-white py-3 w-full px-3 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"
            />

            <input
              type="text"
              name="type"
              id="type"
              placeholder="device type"
              onChange={(e) => setType(e.target.value)}
              className="rounded-md border border-[#e0e0e0] bg-white py-3 w-full px-3 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"
            />

            <button
              className="w-full hover:shadow-lg rounded-md bg-[#6A64F1] py-3 mt-3 text-center text-base font-semibold text-white outline-none"
              onClick={(event) => { addDevice(event) }}
            >
              Add device
            </button>
          </div>
        </form>
      </div>
    </>
  )
}