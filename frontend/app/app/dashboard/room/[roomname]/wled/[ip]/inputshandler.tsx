"use client"

import { ColorChanger } from '@/components/colorChanger'
import { Select } from '@/components/select'
import { Slider } from '@/components/slider'
import { effects, palettes } from '@/data/wled'
import { useState } from 'react'

export const InputsHandler = ({ default_wled_info, roomname, ip }) => {
  const [wledInfo, setWledInfo] = useState(default_wled_info)

  const setWledConfig = async () => {

    await fetch(`http://127.0.0.1:3001/${roomname}/wled_config/set/${ip}`,
      {
        body: JSON.stringify(wledInfo),
        headers: {
          'Content-Type': 'application/json'
        },
        method: 'POST'
      }
    )
  }

  const onColorChange = (color) => {
    setWledInfo((prev) => ({
      ...prev,
      seg: [
        {
          ...prev.seg[0],
          col: [
            [color.rgb.r, color.rgb.g, color.rgb.b],
            [0, 0, 0],
            [0, 0, 0],
          ]
        },
      ]
    }))
  }

  const onBriChange = (bri) => {
    setWledInfo((prev) => ({
      ...prev,
      bri: parseInt(bri)
    }))
  }

  const onEffectChange = (idx: string) => {
    setWledInfo((prev) => ({
      ...prev,
      seg: [
        {
          ...prev.seg[0],
          fx: parseInt(idx)
        },
      ]
    }))
  }

  const onPaletteChange = (idx: string) => {
    setWledInfo((prev) => ({
      ...prev,
      seg: [
        {
          ...prev.seg[0],
          pal: parseInt(idx)
        },
      ]
    }))
  }


  return (
    <div className='p-4'>
      <ColorChanger defaultColor={wledInfo.seg[0].col[0]} onColorChange={onColorChange} />
      <Slider defaultValue={wledInfo.bri} onBriChange={onBriChange} />
      <Select onSelectChange={onEffectChange} defaultvalue={wledInfo.seg[0].fx} textInfo={"Effect"} optionList={effects} />
      <Select onSelectChange={onPaletteChange} defaultvalue={wledInfo.seg[0].pal} textInfo={"Palette"} optionList={palettes} />

      <button
        className="w-full hover:shadow-lg rounded-md bg-[#6A64F1] py-3 mt-3 text-center text-base font-semibold text-white outline-none"
        onClick={() => { setWledConfig() }}
      >
        Change
      </button>
    </div>
  )
}