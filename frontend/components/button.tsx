"use client"

import { GrFormAdd } from 'react-icons/gr'
import { IoMdArrowRoundBack, IoMdAdd } from 'react-icons/io'
import { IconType } from 'react-icons/lib'

import { GoLightBulb } from 'react-icons/go'
import { ImSwitch } from 'react-icons/im'
import { TbDeviceSpeaker } from 'react-icons/tb'
import { useState } from 'react'

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


export const AddButton = ({onClick} : {onClick: any}) => {
  return <Button onClick={onClick} Icon={IoMdAdd} />
}

export const BackButton = ({onClick} : {onClick: any}) => {
  return <Button onClick={onClick} Icon={IoMdArrowRoundBack} />
}

const getDeviceIcon = (type) => {
  if (type === "wled") {
    return <GoLightBulb size={80} />
  } else if (type === "switch") {
    return <ImSwitch size={70} />
  }

  return <TbDeviceSpeaker size={70} />
}

export const Device = ({ nickname, devStatus, type, roomname, ip }) => {
  let icon = getDeviceIcon(type)
  const [status, setStatus] = useState(devStatus.status)
  const [on, setOn] = useState(devStatus.on)

  async function toggleSwitch(event) {
    event.preventDefault();

    if (!status) {
      return
    }
  
    let reply = await fetch(`http://127.0.0.1:3001/${roomname}/${ip}/${on ? "off" : "on"}`,
      {
        method: 'POST'
      }
    )

    setOn(!on)
  }

  return (
    <button
      onClick={(e) => { toggleSwitch(e) }}
      className={`group relative h-48 w-48 bg-white flex flex-col p-3 rounded-md drop-shadow shadow focus:shadow-outline focus:outline-none hover:text-gray-500`}
    >
      <div className={`absolute inset-0 w-3  ${status ? "bg-highlight" : "bg-amber-400"} transition-all duration-[250ms] ease-out group-hover:w-full rounded-l-md group-hover:rounded-md`}></div>
      <div className="relative flex-1 pl-1 pt-2.5 group-hover:text-white">
        {icon}
      </div>
      <div className='relative pl-3 group-hover:text-white'>
        <h3 className="font-bold text-lg">{nickname}</h3>
        {/* <p className="font-light text-left">{status ? on ? "on" : "off" : "disconnected"}</p> */}
        <p className="font-light text-left">{on ? "on" : "off"}</p>
      </div>
    </button>
  )
}