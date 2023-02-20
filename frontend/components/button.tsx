"use client"

import { GrFormAdd } from 'react-icons/gr'
import { IoMdArrowRoundBack } from 'react-icons/io'
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
      <Icon size={35} color="white"/>
    </button>
  )
}


export const AddButton = ({onClick} : {onClick: any}) => {
  return <Button onClick={onClick} Icon={GrFormAdd} />
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

    console.log(on)
  
    let reply = await fetch(`http://127.0.0.1:3001/${roomname}/${ip}/${on ? "off" : "on"}`,
      {
        method: 'POST'
      }
    )

    setOn(!on)
  }

  return (
    <button
      onClick={(e) => {toggleSwitch(e)}}
      className="h-48 w-48 bg-white flex flex-col p-3 rounded-md drop-shadow shadow hover:bg-highlight focus:shadow-outline focus:outline-none hover:text-gray-500"
    >
      <div className="flex-1 pl-1 pt-2.5">
        {icon}
      </div>
      <div className='pl-3'>
        <h3 className="font-bold text-lg">{nickname}</h3>
        {/* <p className="font-light text-left">{status ? on ? "on" : "off" : "disconnected"}</p> */}
        <p className="font-light text-left">{on ? "on" : "off"}</p>
      </div>
    </button>
  )
}