"use client"

import { userEP } from "@/data/endpoints";
import Link from "next/link"
import { useRouter } from "next/navigation"
import { useEffect, useState } from "react";

import { AiFillHome } from 'react-icons/ai'

interface LinkHeaderArgs {
  href: string;
  headerText?: string;
  children: JSX.Element;
  showHome: boolean;
  disableMargin?: boolean;
}

export const LinkHeader = ({ href, headerText, children, showHome, disableMargin = false }: LinkHeaderArgs) => {
  const router = useRouter()
  const [username, setUsername] = useState("")

  useEffect(() => {
    getUsername()
  }, [])

  const getUsername = () => {
    fetch(userEP, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    })
      .then((resp) => resp.json())
      .then(({ success, data }) => {
        if (success) {
          setUsername(data)
        }
      })
  }

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
      
      <h1 className="font-bold text-xl text-slate-600">{headerText ? headerText : "Welcome home, " + username}</h1>
      <Link href={href}>{children}</Link>
    </div>
  )
}