const rooms = [
  'bedroom',
  'living room',
  'store room'
]

const RoomTile = ({name}) => {
  return (
    <button className="relative h-28 text-center flex flex-col items-center justify-center rounded-lg hover:bg-roomtile-highlight">
      <div className="absolute flex flex-col items-center justify-center z-10">
        <h1 className="text-white font-bold capitalize text-2xl">{name}</h1>
        <p className="text-white font-bold capitalize text-xs">3 devices connected</p>
      </div>
      <div className="absolute h-28 rounded-lg w-full bg-roomtile z-5"></div>
      <div
        className="bg-cover bg-center absolute h-28 rounded-lg w-full opacity-70 z-0"
        style={{
          backgroundImage: `url('/images/room.jpg')`,
        }}
      ></div>
    </button>
  )
}


export default function Page() {
  

  return (
    <>
      <div className="h-8"></div>
      <div className="flex flex-col gap-5 px-4">
        {
        rooms.map((rm) => { 
          return <RoomTile key={rm} name={rm} />
          })
        }
      </div>
    </>
  )
}