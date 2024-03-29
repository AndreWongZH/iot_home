"use client"

import { BackHeader } from '@/components/header';
import { useRouter } from 'next/navigation';
import React, { useState } from 'react';
import { Notify } from 'notiflix/build/notiflix-notify-aio';
import { addDeviceEP } from '@/data/endpoints';
import { Discover } from './discover';

export default function Page({ params }: { params: { roomname: string } }) {
  const router = useRouter();

  const [name, setName] = useState("");
  const [ipAddr, setIpAddr] = useState("");
  const [type, setType] = useState("");

  async function addDevice(event: React.SyntheticEvent) {
    event.preventDefault();

    const jsonData = {
      name: name,
      ipaddr: ipAddr,
      type: type
    }

    fetch(addDeviceEP(params.roomname), {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(jsonData)
    })
      .then((resp) => resp.json())
      .then(({ success, error }) => {
        if (success) {
          router.replace(`/dashboard/room/${params.roomname}`)
        } else {
          Notify.failure(error, {
            position: 'center-bottom',
            timeout: 1500,
            showOnlyTheLastOne: true,
            clickToClose: true
          })
        }
      })
  }

  return (
    <>
      <BackHeader headerText={`Add a device to ${params.roomname}`} />
      <div className="flex flex-col items-center">
        <div className="flex flex-col justify-center items-center gap-y-2.5">
          <input
            type="text"
            name="nickname"
            id="nickname"
            placeholder="nickname"
            onChange={(e) => setName(e.target.value)}
            className="rounded-md border border-[#e0e0e0] bg-white py-3 px-3 w-full text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"
          />

          <input
            type="text"
            name="ipaddr"
            id="ipaddr"
            placeholder="IP Address"
            value={ipAddr || ""}
            onChange={(e) => setIpAddr(e.target.value)}
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
            onClick={(event: React.SyntheticEvent) => { addDevice(event) }}
          >
            Add device
          </button>

          <Discover setIpAddr={setIpAddr} />
        </div>
      </div>
    </>
  )
}