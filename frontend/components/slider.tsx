"use client"

import { useState } from 'react'

export const Slider = ({ defaultValue, onBriChange }: { defaultValue: number, onBriChange: Function }) => {
  const [value, setValue] = useState(defaultValue)

  return (
    <div className='flex items-center my-10 gap-3'>
      <h1 className='font-bold'>Brightness</h1>
      <input
        className="rounded-lg overflow-hidden appearance-none bg-gray-400 h-3 w-128" type="range" min="1" max="255" step="1"
        value={value}
        onChange={(e) => {
          setValue(parseInt(e.target.value))
          onBriChange(e.target.value)
        }}
      />
      <h1 className=''>{ value }</h1>
    </div>
  )
}