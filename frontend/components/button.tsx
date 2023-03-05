"use client"

import { IoMdArrowRoundBack, IoMdAdd } from 'react-icons/io'
import { IconType } from 'react-icons/lib'

import { GoLightBulb } from 'react-icons/go'
import { ImSwitch } from 'react-icons/im'
import { TbDeviceSpeaker } from 'react-icons/tb'
import React, { useState } from 'react'
import { useRouter } from 'next/navigation';
import instance from './axiosInst'

export const Button = ({onClick, Icon} : {onClick: any, Icon: IconType}) => {
  return (
    <button
      className="shadow hover:bg-highlight focus:shadow-outline focus:outline-none text-white text-xs p-1 rounded bg-background-default"
      onClick={onClick}
    >
      <span className='text-slate-500 hover:text-white'>
        <Icon size={35} />
      </span>
    </button>
  )
}


export const AddButton = ({onClick} : {onClick: Function | null}) => {
  return <Button onClick={onClick} Icon={IoMdAdd} />
}

export const BackButton = ({onClick} : {onClick: Function}) => {
  return <Button onClick={onClick} Icon={IoMdArrowRoundBack} />
}

const getDeviceIcon = (type : string) => {
  if (type === "wled") {
    return <GoLightBulb size={80} />
  } else if (type === "switch") {
    return <ImSwitch size={70} />
  }

  return <TbDeviceSpeaker size={70} />
}

interface DeviceArgs {
  name: string;
  devStatus: {
    status: boolean;
    on: boolean;
  }
  type: string;
  roomName: string;
  ip: string;
  setMode: boolean;
}

export const Device = ({ name, devStatus, type, roomName, ip, setMode }: DeviceArgs) => {
  let icon = getDeviceIcon(type)
  const [status, setStatus] = useState(devStatus.status)
  const [on, setOn] = useState(devStatus.on)
  const router = useRouter();

  async function toggleSwitch(event: React.SyntheticEvent) {
    event.preventDefault();

    // if (!status) {
    //   return
    // }

    instance.post(`${roomName}/${ip}/${on ? "off" : "on"}`)
    .then(function (resp) {
      const {success, data} = resp.data
    })
    .catch(function (err) {

    })

    setOn(!on)
  }

  function goToSettings() {
    if (type == "wled") {
      router.push(`/dashboard/room/${roomName}/wled/${ip}`)
    }
    
    return
  }

  return (
    <button
      onClick={(e) => { setMode ? goToSettings() : toggleSwitch(e) }}
      className={`group relative h-48 w-48 bg-white flex flex-col p-3 rounded-md drop-shadow shadow focus:shadow-outline focus:outline-none hover:text-gray-500`}
    >
      <div className={`absolute inset-0 w-3  ${status ? "bg-highlight" : "bg-amber-400"} transition-all duration-[250ms] ease-out group-hover:w-full rounded-l-md group-hover:rounded-md`}></div>
      <div className="relative flex-1 pl-1 pt-2.5 group-hover:text-white">
        {icon}
      </div>
      <div className='relative pl-3 group-hover:text-white'>
        <h3 className="font-bold text-lg">{name}</h3>
        {/* <p className="font-light text-left">{status ? on ? "on" : "off" : "disconnected"}</p> */}
        <p className="font-light text-left">{on ? "on" : "off"}</p>
      </div>
    </button>
  )
}