"use client"

import { ColorChanger } from '@/components/colorChanger'
import { Slider } from '@/components/slider'
import { effects } from '@/data/wled'
import { useState } from 'react'

export const InputsHandler = ({ default_wled_info, roomname, ip }) => {
  const [wledInfo, setWledInfo] = useState(default_wled_info)
  console.log(wledInfo)

  const setWledConfig = async () => {
    console.log(wledInfo)

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


  return (
    <>
      <ColorChanger defaultColor={wledInfo.seg[0].col[0]} onColorChange={onColorChange} />
      {/* <h1 className="p-10 mr-10 mx-10">hello world</h1> */}
      <Slider defaultValue={wledInfo.bri} onBriChange={onBriChange} />
      <select data-te-select-init>
        {
          effects.map((eff, idx) => (
            <option value={eff} key={idx}>{eff}</option>
          ))
        }
      </select>
      <button onClick={() => { setWledConfig() }}>Change</button>
    </>
  )
}