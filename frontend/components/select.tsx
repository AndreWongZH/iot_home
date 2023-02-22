"use client"

import { useState } from 'react'

export const Select = ({ optionList, onSelectChange, textInfo, defaultvalue }) => {
  const [option, setOption] = useState(defaultvalue)

  return (
    <div className='my-7 flex items-center gap-3 justify-between'>
      <h1 className='font-bold'>{textInfo}</h1>
      <select
        className='w-full bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500'
        value={option}
        onChange={(e) => {
          setOption(e.target.value)
          onSelectChange(e.target.value)
        }}
      >
          {
            optionList.map((opt, idx) => (
              <option value={idx} key={idx}>{opt}</option>
            ))
          }
      </select>
    </div>
  )
}