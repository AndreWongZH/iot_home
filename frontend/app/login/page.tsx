"use client"

import { loginEP } from "@/data/endpoints"
import { useRouter } from "next/navigation"
import { Notify } from "notiflix/build/notiflix-notify-aio"
import { useState } from "react"

export default function Page() {
  const router = useRouter()
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")

  const tryLogin = async (e: React.SyntheticEvent) => {
    e.preventDefault()

    const jsonData = {
      username,
      password
    }

    fetch(loginEP, {
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
          router.push(`/dashboard/`)
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
    <div className="flex flex-col justify-start items-center w-full h-full">
      <h1 className="font-bold text-2xl mb-20 pt-24">Welcome to IOT home</h1>
      <div className="max-w-xs h-64 mx-auto bg-highlight mt-16 p-7 rounded-lg">
        <div className="flex flex-col justify-center items-center gap-y-2.5">
          <input
            type="text"
            name="username"
            id="username"
            placeholder="Username"
            onChange={(e) => { setUsername(e.target.value) }}
            className="rounded-md border border-[#e0e0e0] bg-white py-3 px-3 w-full text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"
          />

          <input
            type="password"
            name="password"
            id="password"
            placeholder="Password"
            onChange={(e) => { setPassword(e.target.value) }}
            className="rounded-md border border-[#e0e0e0] bg-white py-3 w-full px-3 text-base font-medium text-[#6B7280] outline-none focus:border-[#6A64F1] focus:shadow-md"
          />

          <button
            className="w-full hover:shadow-lg rounded-md bg-[#6A64F1] py-3 mt-3 text-center text-base font-semibold text-white outline-none"
            onClick={(e) => { tryLogin(e) }}
          >Login</button>
        </div>
      </div>
    </div>
  )
}