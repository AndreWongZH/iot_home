"use client"

import { useState } from 'react'
import styles from './slider.css'

export const Slider = ({ defaultValue, onBriChange }) => {
  const [value, setValue] = useState(defaultValue)

  return (
    <div className='flex items-center'>
      <h1 className='mx-10 font-bold'>Brightness</h1>
      <input
        className="rounded-lg overflow-hidden appearance-none bg-gray-400 h-3 w-128" type="range" min="1" max="255" step="1"
        value={value}
        onChange={(e) => {
          setValue(e.target.value)
          onBriChange(e.target.value)
        }}
      />
      <h1 className='ml-10'>{ value }</h1>
    </div>
  )
}