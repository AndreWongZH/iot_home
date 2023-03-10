"use client"

import { useEffect, useState, useRef } from 'react'
import { ColorPicker, useColor } from 'react-color-palette'
import "react-color-palette/lib/css/styles.css";

export const ColorChanger = ({ defaultColor, onColorChange }: { defaultColor: Array<number>, onColorChange: Function }) => {
  const [dimension, setDims] = useState({ height: 0, width: 0 })
  const pickerRef = useRef<HTMLDivElement>(null);

  const [color, setColor] = useColor("rgb", {
    r: defaultColor[0],
    g: defaultColor[1],
    b: defaultColor[2],
  })

  useEffect(() => {
    if (pickerRef.current) {
      const { height, width } = pickerRef.current.getBoundingClientRect();

      setDims({
        height: height,
        width: width
      })
    }

    
  }, []);

  return (
    <div className='w-3/4 mx-auto' ref={pickerRef}>
      <ColorPicker
        width={dimension.width}
        height={dimension.height}
        color={color}
        onChange={(e) => { setColor(e); onColorChange(e) }}
        hideHSV
        hideHEX
        dark
      />
    </div>
  )
}