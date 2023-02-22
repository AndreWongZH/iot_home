/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'background-default': '#e0e8ea',
        'outbounds': '#F2F2F2',
        'highlight': '#BDDCBD',
        'roomtile': '#000000B3',
        'roomtile-highlight': '#0000004d',
      },
    }
  },
  plugins: [],
}
