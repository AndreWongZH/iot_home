import { BackHeader } from "@/components/header";
import { InputsHandler } from "./inputshandler";

export default function Page({ params } : { params: {roomname: string, ip: string} }) {
  return (
    <div className="w-full">
      <BackHeader headerText="wled config"/>
      <InputsHandler roomName={params.roomname} ip={params.ip} />
    </div>
  )
}