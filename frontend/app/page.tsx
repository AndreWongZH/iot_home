"use client"

import { useState } from "react";
import { useRouter } from "next/navigation";
import { loginEP, registerEP } from "@/data/endpoints";
import { Notify } from "notiflix";
import { Auth } from "./auth";

export default function Page() {
  const [isLogin, setLogin] = useState(true)
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const router = useRouter()
  // if (document) {
  //   console.log(document.cookie)
  // }

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
  

  const register = async (e: React.SyntheticEvent) => {
    e.preventDefault()

    const jsonData = {
      username,
      password
    }

    fetch(registerEP, {
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
    <>
      <Auth isLogin={isLogin} setLogin={setLogin} setUsername={setUsername} setPassword={setPassword} tryLogin={tryLogin} register={register} />
    </>
  )
}
