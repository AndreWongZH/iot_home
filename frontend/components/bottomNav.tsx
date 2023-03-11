"use client"

import Link from 'next/link'

import { TbLogout } from 'react-icons/tb'
import { AiFillHome } from 'react-icons/ai'
import { useRouter } from 'next/navigation'
import { Notify } from 'notiflix/build/notiflix-notify-aio'
import { logoutEP } from '@/data/endpoints'

export const BottomNav = () => {
  const router = useRouter()

  const logout = (e: React.SyntheticEvent) => {
    e.preventDefault()

    fetch(logoutEP, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    })
      .then((resp) => resp.json())
      .then(({ success, error }) => {
        if (success) {
          router.replace(`/`)
        } else {
          Notify.failure(error, {
            position: 'center-bottom',
            timeout: 1500,
            showOnlyTheLastOne: true,
            clickToClose: true
          })
        }
      })
  }

  return (
    <nav className="fixed bottom-0 w-full max-w-md z-50">
      <div className="px-7 bg-slate-700 shadow-lg rounded-t-2xl">
        <div className="flex">
          <div className="flex-auto hover:w-full group">
            <Link href="/dashboard" className="flex items-center justify-center text-center mx-auto px-4 py-2 group-hover:w-full text-indigo-500">
              <span className="block px-1 py-1 group-hover:bg-indigo-100 rounded-full group-hover:flex-grow">
                <div className='flex items-center justify-center'>
                  <AiFillHome size={45} />
                  <span className="hidden ml-5 group-hover:inline">Home</span>
                </div>
              </span>
            </Link>
          </div>

          <div className="flex-auto hover:w-full group">
            <button onClick={(e) => logout(e)} className="flex items-center justify-center text-center mx-auto px-4 py-2 group-hover:w-full text-indigo-500">
              <span className="block px-1 py-1 group-hover:bg-indigo-100 rounded-full group-hover:flex-grow">
                <div className='flex items-center justify-center'>
                  <TbLogout size={45} />
                  <span className="hidden ml-5 group-hover:inline-block">Logout</span>
                </div>
              </span>
            </button>
          </div>
        </div>
      </div>
    </nav>
  )
}