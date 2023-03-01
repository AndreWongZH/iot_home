const data = [
  { macname: "sun microsystems", ipaddr: "192.128.1.2", macaddr: "4D:4E:4F:02:14:B0" },
  { macname: "Ubuntu", ipaddr: "192.128.1.30", macaddr: "4D:4E:4F:02:14:B1" },
  { macname: "Espressif", ipaddr: "192.128.1.222", macaddr: "4D:4E:4F:02:14:B2" },
]


const DiscoveredAddress = ({ macname, ipaddr, macaddr }: { macname: string, ipaddr: string, macaddr: string }) => {
  return (
    <div className="flex items-center justify-between h-14 px-4 py-6 my-5 hover:bg-highlight hover:rounded-md">
      <div className="flex flex-col">
        <h1 className="font-bold uppercase">{macname}</h1>
        <p className="text-sm">already in network</p>
      </div>
      <div className="flex flex-col items-end justify-center">
        <p>{ipaddr}</p>
        <p>{macaddr}</p>
      </div>
    </div>
  )
}


export default function Page() {
  return (
    <div className="flex flex-col items-center">
      <h1 className="font-bold text-2xl py-8">10 devices found</h1>
      <div className="w-full px-5">
        {data.map((d) => {
          return <DiscoveredAddress key={d.ipaddr} macname={d.macname} ipaddr={d.ipaddr} macaddr={d.macaddr} />
        })}
      </div>
    </div>
  )
}