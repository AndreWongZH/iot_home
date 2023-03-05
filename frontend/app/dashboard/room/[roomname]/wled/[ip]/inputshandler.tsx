"use client"

import instance from '@/components/axiosInst'
import { ColorChanger } from '@/components/colorChanger'
import { Select } from '@/components/select'
import { Slider } from '@/components/slider'
import { effects, palettes } from '@/data/wled'
import { useEffect, useState } from 'react'
import { Color } from 'react-color-palette'

const emptyWledConfig = {
  bri: 0,
  seg: [
    {
      col: [[0,0,0]],
      fx: 0,
      pal: 0
    }
  ]
}

interface Seg {
  col: Array<Array<number>>
  fx: number;
  pal: number;
}

interface WledConfig {
  bri: number;
  seg: Array<Seg>;
}

export const InputsHandler = ({ roomName, ip }: { roomName: string, ip: string }) => {
  const [wledInfo, setWledInfo] = useState<WledConfig>(emptyWledConfig)
  const [success, setSuccess] = useState(false)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    getWledInfo()
  }, [])

  const getWledInfo = () => {
    instance.get(`${roomName}/wled_config/${ip}`)
    .then(function (resp) {
      const {success, data} = resp.data
      console.log(data)
      setWledInfo(data)
      setSuccess(success)
      setLoading(false)
    })
    .catch(function (err) {

    })
  }

  const setWledConfig = async () => {

    console.log(wledInfo)

    instance.post(`${roomName}/wled_config/set/${ip}`, wledInfo)
    .then(function (resp) {
      const {success, data} = resp.data
    })
    .catch(function (err) {

    })

  }

  const onColorChange = (color: Color) => {
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

  const onBriChange = (bri: string) => {
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
    <>
      {
        success 
        ? <></>
        : 
        <div className="mx-auto text-center mb-4 block w-3/4 rounded-lg bg-orange-500 p-4 text-base leading-5 text-white opacity-100">
          <h1 className="font-bold mb-2">Error fetching wled data</h1>
          <h1 className="font-bold">Wled device is offline</h1>
        </div>
      }
      {
        loading ? <div>loading...</div> :
      
        <div className='p-4'>
          <ColorChanger defaultColor={wledInfo.seg[0].col[0]} onColorChange={onColorChange} />
          <Slider defaultValue={wledInfo.bri} onBriChange={onBriChange} />
          <Select onSelectChange={onEffectChange} defaultValue={wledInfo.seg[0].fx} textInfo={"Effect"} optionList={effects} />
          <Select onSelectChange={onPaletteChange} defaultValue={wledInfo.seg[0].pal} textInfo={"Palette"} optionList={palettes} />

          <button
            className="w-full hover:shadow-lg rounded-md bg-[#6A64F1] py-3 mt-3 text-center text-base font-semibold text-white outline-none"
            onClick={() => { setWledConfig() }}
          >
            Change
          </button>
        </div>
      }
    </>
  )
}