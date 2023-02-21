import { ColorChanger } from "@/components/colorChanger";
import { Slider } from "@/components/slider";
import { InputsHandler } from "./inputshandler";

async function getWledInfo(roomname, ip) {
  const res = await fetch(`http://127.0.0.1:3001/${roomname}/wled_config/${ip}`)

  if (!res.ok) {
    throw new Error("failed to fetch wled config")
  }

  return res.json()
}

export default async function Page({ params }) {
  const wledInfo = await getWledInfo(params.roomname, params.ip)
  console.log(wledInfo)

  return (
    <div className="w-full">
      {/* <ColorChanger defaultColor={wledInfo.seg[0].col[0]} roomname={params.roomname} ip={params.ip} /> */}
      {/* <h1 className="p-10 mr-10 ml-10">hello world</h1> */}
      {/* <Slider defaultValue={wledInfo.bri}/> */}
      <InputsHandler roomname={params.roomname} ip={params.ip} default_wled_info={wledInfo} />
    </div>
  )
}