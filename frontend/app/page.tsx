"use client"

import { Login } from "./login";

export default function Page() {
  console.log(document.cookie)
  return (
    <>
      {
        <Login />
      }
    </>
  )
}
