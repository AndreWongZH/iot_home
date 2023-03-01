"use client"

export const Toggle = ({setMode, setSetMode} : {setMode: boolean, setSetMode: Function}) => {

  return (
    <div className="mb-12 flex justify-end items-center m-2 cursor-pointer cm-toggle-wrapper"
      onClick={() => {setSetMode((prev: boolean) => {
        return !prev
      })}}
    >
      <span className="font-semibold text-xs mr-1">
          Switch
      </span>
      <div className={`rounded-full w-12 h-7 p-0.5 ${setMode ? 'bg-amber-200' : 'bg-highlight'}`}>
          <div className={`rounded-full w-6 h-6 bg-white transform mx-auto duration-300 ease-in-out ${setMode ? "translate-x-2" : "-translate-x-2"}`}></div>
      </div>
      <span className="font-semibold text-xs ml-1">
          Settings
      </span>
    </div>
  )
}