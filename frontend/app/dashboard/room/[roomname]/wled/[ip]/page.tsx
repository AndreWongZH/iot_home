import { DeviceInfo } from "@/components/devInfo";
import { BackHeader } from "@/components/header";
import { InputsHandler } from "./inputshandler";
import { DeleteDevice } from "@/components/button";

export default function Page({ params } : { params: {roomname: string, ip: string} }) {
  return (
    <div className="w-full">
      <BackHeader headerText="wled config"/>
      <DeviceInfo roomName={params.roomname} ip={params.ip}/>
      <InputsHandler roomName={params.roomname} ip={params.ip} />
      <DeleteDevice roomName={params.roomname} ip={params.ip} />
    </div>
  )
}