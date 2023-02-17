"use client"

import { ColorPicker, useColor } from 'react-color-palette'
import "react-color-palette/lib/css/styles.css";

export const ColorChanger = () => {
  const [color, setColor] = useColor("hex", "#121212")

  return (
    <>
      <ColorPicker
        width={456}
        height={228}
        color={color}
        onChange={setColor}
        hideHSV
        dark
      />
      <button onClick={() => { console.log(color) }}>Click</button>
    </>
  )
}