import { BackHeader } from "@/components/header";
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
      <BackHeader headerText="wled config"/>
      {
        wledInfo.success 
        ? <></>
        : 
        <div className="mx-auto text-center mb-4 block w-3/4 rounded-lg bg-orange-500 p-4 text-base leading-5 text-white opacity-100">
          <h1 className="font-bold mb-2">Error fetching wled data</h1>
          <h1 className="font-bold">Wled device is offline</h1>
        </div>
      }
      <InputsHandler roomName={params.roomname} ip={params.ip} default_wled_info={wledInfo.data} />
      
    </div>
  )
}