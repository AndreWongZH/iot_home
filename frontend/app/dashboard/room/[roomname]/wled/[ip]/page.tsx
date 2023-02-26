import { ColorChanger } from "@/components/colorChanger";
import { BackHeader } from "@/components/header";
import { Slider } from "@/components/slider";
import { InputsHandler } from "./inputshandler";

async function getWledInfo(roomName: string, ip: string) {
  const res = await fetch(`http://127.0.0.1:3001/${roomName}/wled_config/${ip}`)

  if (!res.ok) {
    throw new Error("failed to fetch wled config")
  }

  return res.json()
}

export default async function Page({ params } : { params: {roomname: string, ip: string} }) {
  const wledInfo = await getWledInfo(params.roomname, params.ip)
  console.log(wledInfo)

  return (
    <div className="w-full">
      <BackHeader />
      {wledInfo.success ? <InputsHandler roomname={params.roomname} ip={params.ip} default_wled_info={wledInfo.data} /> : <></>}
      {wledInfo.success ? <></> : <h1>Error fetching wled data. Wled device is offline</h1>}
    </div>
  )
}