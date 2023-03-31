import { DeleteDevice } from "@/components/button";
import { DeviceInfo } from "@/components/devInfo";
import { BackHeader } from "@/components/header";

export default function Page({ params } : { params: {roomname: string, ip: string} }) {
  return (
    <div className="w-full">
      <BackHeader headerText="switch config"/>
      <DeviceInfo roomName={params.roomname} ip={params.ip} />
      <DeleteDevice roomName={params.roomname} ip={params.ip} />
    </div>
  )
}