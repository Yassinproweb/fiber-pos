/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.html"
  ],
  theme: {
    extend: {
      colors: {
        cs_color: 'hsl(207, 89%, 45%)',
        cs_white: 'hsl(207, 1%, 100%)',
        cs_light: 'hsl(207, 1%, 63%)',
        cs_black: 'hsl(207, 1%, 07%)',

        ord_plc: 'hsl(213, 91%, 57%)',
        ord_cld: 'hsl(1, 95%, 53%)',
        ord_trs: 'hsl(81, 93%, 47%)',
        ord_dlv: 'hsl(141, 81%, 37%)',
        ord_pdg: 'hsl(267, 81%, 57%)',
        ord_rdy: 'hsl(51, 93%, 45%)',
      }
    },
  },
  plugins: [],
}
