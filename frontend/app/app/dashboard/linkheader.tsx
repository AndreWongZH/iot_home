import Link from "next/link"

import { AiFillHome } from 'react-icons/ai'

export const LinkHeader = ({ href, headerText, children, showHome}) => {
  return (
    <div className="mb-12 h-10 px-3 py-3 bg-white h-16 flex items-center justify-between">
      {
        showHome
        ? <Link href={`/app/dashboard`}>
            <AiFillHome size={35} color={"black"}/>
          </Link>
        : <></>
      }
      
      <h1 className="font-bold text-xl text-slate-600">{headerText}</h1>
      <Link href={href}>{children}</Link>
    </div>
  )
}