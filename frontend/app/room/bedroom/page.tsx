import { GoLightBulb } from 'react-icons/go'
import { ImSwitch } from 'react-icons/im'
import { TbDeviceSpeaker } from 'react-icons/tb'

const obj = [
  { nickname: "Monitor", state: "on", type: "wled" },
  { nickname: "Desk lamp", state: "off", type: "wled" },
  { nickname: "head lamp", state: "on", type: "switch" },
  { nickname: "Fan", state: "on", type: "IR" },
  { nickname: "Air purifier", state: "on", type: "purifier" },
]

const getDeviceIcon = (type) => {
  if (type === "wled") {
    return <GoLightBulb size={80} />
  } else if (type === "switch") {
    return <ImSwitch size={70} />
  }

  return <TbDeviceSpeaker size={70} />
}

const Device = ({ nickname, state, type }) => {
  let icon = getDeviceIcon(type)

  return (
    <div className="h-48 w-48 bg-white flex flex-col p-3 rounded-md drop-shadow">
      <div className="flex-1 pl-1 pt-2.5">
        {icon}
      </div>
      <div className='pl-3'>
        <h3 className="font-bold text-lg">{nickname}</h3>
        <p className="font-light">{state}</p>
      </div>
    </div>
  )
}


export default function Page() {
  return (
    <div className='max-w-md mx-auto  bg-background-default'>
      <div className="mb-8 h-10 p-3">
        <h1 className="font-bold text-xl text-slate-400">Welcome home, XXX</h1>
      </div>

      <div className="flex flex-wrap gap-2 justify-center">
      {
        obj.map((dev) => {
          return <Device key={dev.nickname} nickname={dev.nickname} state={dev.state} type={dev.type} />
        })
      }
      </div>

      <div className="py-8 h-10" />

    </div>
  )
}