import { CgSearchLoading } from 'react-icons/cg'

export default function Loading() {
  return (
    <div className="w-full h-full flex flex-row items-center justify-center">
      <div className='flex flex-row items-center justify-center gap-3'>
        <CgSearchLoading size={70} color={"#64748B"} />
        <h1 className="text-4xl font-bold text-slate-500">Loading data</h1>
      </div>
    </div>
  )
}