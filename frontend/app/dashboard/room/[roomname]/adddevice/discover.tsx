"use client"

import { CgSearchLoading } from 'react-icons/cg'
import { discoverEP } from "@/data/endpoints"
import { useState } from "react"

const DiscoveredAddress = ({ ipAddr, setIpAddr }: { ipAddr: string, setIpAddr: Function }) => {
  return (
    <button className="flex items-center justify-between h-14 px-4 py-2 hover:bg-highlight hover:rounded-md"
      onClick={() => setIpAddr(ipAddr)}
    >
      <h1 className="font-bold uppercase">{ipAddr}</h1>
    </button>
  )
}

export const Discover = ({ setIpAddr }: { setIpAddr: Function }) => {
  const [data, setData] = useState<Array<string>>([])
  const [loading, setLoading] = useState(true)
  const [wait, setWaiting] = useState(true)

  const getIpList = () => {
    setWaiting(false)

    fetch(discoverEP, {
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
        console.log(data)
        setData(data)
        setLoading(false)
      })
  }

  return (
    <>
      {
        wait ?
          <button
            className="w-full hover:shadow-lg rounded-md bg-[#6A64F1] py-3 mt-5 text-center text-base font-semibold text-white outline-none"
            onClick={() => { getIpList() }}
          >
            Find devices
          </button>
          :
          loading ? (
            <div className='flex flex-row items-center justify-center gap-3 mt-5'>
              <CgSearchLoading size={40} color={"#64748B"} />
              <h1 className="text-2xl font-bold text-slate-500">Loading data</h1>
            </div>
          )
            :
            <div className="flex flex-col items-center">
              <h1 className="font-bold text-2xl py-2 mt-5">{data.length} devices found</h1>
              <div className="w-full px-5">
                {data.map((ipAddr, i) => {
                  return <DiscoveredAddress key={i} ipAddr={ipAddr} setIpAddr={setIpAddr} />
                })}
              </div>
            </div>
      }
    </>
  )
}