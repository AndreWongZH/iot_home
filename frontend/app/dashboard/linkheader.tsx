"use client"

import Link from "next/link"
import { useRouter } from "next/navigation"

import { AiFillHome } from 'react-icons/ai'

interface LinkHeaderArgs {
  href: string;
  headerText: string;
  children: JSX.Element;
  showHome: boolean;
  disableMargin?: boolean;
}

export const LinkHeader = ({ href, headerText, children, showHome, disableMargin = false }: LinkHeaderArgs) => {
  const router = useRouter()

  return (
    <div className={`${disableMargin ? "" : "mb-12"} h-10 px-3 py-3 bg-white h-16 flex items-center justify-between`}>
      {
        showHome?
        //  <Link href={`/dashboard`}>
          <span className="cursor-pointer">
            <AiFillHome size={35} color={"#475569"} className="hover:scale-110" onClick={() => router.push("/dashboard")} />
          </span>
          // </Link>
        : <></>
      }
      
      <h1 className="font-bold text-xl text-slate-600">{headerText}</h1>
      <Link href={href}>{children}</Link>
    </div>
  )
}