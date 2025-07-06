/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.html"
  ],
  theme: {
    extend: {
      colors: {
        cs_color: 'hsl(81, 21%, 35%))',
        cs_white: 'hsl(205, 1%, 100%)',
        cs_light: 'hsl(205, 1%, 63%)',
        cs_black: 'hsl(205, 1%, 07%)',

        ord_plc: 'hsl(217, 99%, 56%)',
        ord_cld: 'hsl(357, 95%, 50%)',
        ord_trs: 'hsl(81, 97%, 45%)',
        ord_dlv: 'hsl(141, 89%, 43%)',
        ord_pdg: 'hsl(267, 81%, 57%)',
        ord_rdy: 'hsl(51, 93%, 45%)',
      }
    },
  },
  plugins: [],
}
