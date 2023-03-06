"use client"

import { BackButton } from "./button"
import { useRouter } from 'next/navigation';


export const BackHeader = ({ headerText }: { headerText: string}) => {
  const router = useRouter();

  return (
    <div className="mb-12 h-10 px-3 py-3 bg-white h-16 flex items-center justify-start">
      <BackButton onClick={() => { router.back() }} />
      <h1 className="font-bold text-xl text-slate-600 ml-10">{headerText}</h1>
    </div>
  )
}