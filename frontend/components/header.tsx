"use client"

import { AppRouterInstance } from "next/dist/shared/lib/app-router-context"
import Link from "next/link"
import { BackButton } from "./button"


export const BackHeader = ({ router } : { router: AppRouterInstance }) => {
  return (
    <div className="mb-12 h-10 px-3 py-3 bg-white h-16 flex items-center justify-between">
      <Link href={`/app/dashboard/`}>
        <BackButton onClick={() => { router.back() }} />
      </Link>
    </div>
  )
}