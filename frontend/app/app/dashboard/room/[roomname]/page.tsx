import { GoLightBulb } from 'react-icons/go'
import { ImSwitch } from 'react-icons/im'
import { TbDeviceSpeaker } from 'react-icons/tb'
import { GrFormAdd } from 'react-icons/gr'

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
    <button className="h-48 w-48 bg-white flex flex-col p-3 rounded-md drop-shadow shadow hover:bg-highlight focus:shadow-outline focus:outline-none hover:text-gray-500">
      <div className="flex-1 pl-1 pt-2.5">
        {icon}
      </div>
      <div className='pl-3'>
        <h3 className="font-bold text-lg">{nickname}</h3>
        <p className="font-light text-left">{state}</p>
      </div>
    </button>
  )
}

const Button = () => {
  return (
    <button className="shadow hover:bg-highlight focus:shadow-outline focus:outline-none text-white text-xs p-1 rounded bg-background-default">
      <GrFormAdd size={35} color="white"/>
    </button>
  )
}


const Header = () => {
  return (
    <div className="mb-12 h-10 px-3 py-3 bg-white h-16 flex items-center justify-between">
      <h1 className="font-bold text-xl text-slate-600">Welcome home, Andre</h1>
      <Button />
    </div>
  )
}


export default function Page({ params }) {
  console.log(params.roomname)
  return (
    <>
      <Header />
      <div className="flex flex-wrap gap-5 justify-center">
      {
        obj.map((dev) => {
          return <Device key={dev.nickname} nickname={dev.nickname} state={dev.state} type={dev.type} />
        })
      }
        </div>
    </>
  )
}