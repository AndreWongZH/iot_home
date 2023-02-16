"use client"

import { GrFormAdd } from 'react-icons/gr'
import { IoMdArrowRoundBack } from 'react-icons/io'
import { IconType } from 'react-icons/lib'

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