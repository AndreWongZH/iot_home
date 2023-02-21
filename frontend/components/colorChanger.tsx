"use client"

import { ColorPicker, useColor } from 'react-color-palette'
import "react-color-palette/lib/css/styles.css";

export const ColorChanger = ({ defaultColor, onColorChange }) => {
  const [color, setColor] = useColor("rgb", {
    r: defaultColor[0],
    g: defaultColor[1],
    b: defaultColor[2],
  })

  // const changeColor = async () => {
  //   console.log(color)

  //   const wledJson = {
  //     seg: [
  //       {
  //         col: [[color.rgb.r, color.rgb.g, color.rgb.b]]
  //       }
  //     ]
  //   }

  //   await fetch(`http://127.0.0.1:3001/${roomname}/wled_config/set/${ip}`,
  //     {
  //       body: JSON.stringify(wledJson),
  //       headers: {
  //         'Content-Type': 'application/json'
  //       },
  //       method: 'POST'
  //     }
  //   )
  // }

  return (
    <div>
      <ColorPicker
        width={456}
        height={228}
        color={color}
        onChange={(e) => { setColor(e); onColorChange(e) }}
        hideHSV
        dark
      />
      
    </div>
  )
}