"use client"

import instance from "@/components/axiosInst";
import { BackHeader } from "@/components/header";
import { useEffect, useState } from "react";
import { InputsHandler } from "./inputshandler";

export default function Page({ params } : { params: {roomname: string, ip: string} }) {
  // const [wledInfo, setWledInfo] = useState(emptyWledConfig)
  // const [success, setSuccess] = useState(false)

  // useEffect(() => {
  //   getWledInfo()
  // }, [])

  // const getWledInfo = () => {
  //   instance.get(`${params.roomname}/wled_config/${params.ip}`)
  //   .then(function (resp) {
  //     const {success, data} = resp.data
  //     console.log(data)
  //     setWledInfo(data)
  //     setSuccess(success)
  //   })
  //   .catch(function (err) {

  //   })
  // }

  return (
    <div className="w-full">
      <BackHeader headerText="wled config"/>
      <InputsHandler roomName={params.roomname} ip={params.ip} />
      
    </div>
  )
}