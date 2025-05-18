/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.html"
  ],
  theme: {
    extend: {
      colors: {
        cs_color: 'hsl(197, 71%, 43%)',
        cs_white: 'hsl(197, 1%, 100%)',
        cs_black: 'hsl(197, 1%, 07%)',
        ord_plc: 'hsl(217, 91%, 51%)', //blue
        ord_pdg: 'hsl(267, 81%, 53%)', //purple
        ord_cld: 'hsl(217, 21%, 63%)', //gray
        ord_rdy: 'hsl(33, 93%, 47%)', //orange
        ord_trs: 'hsl(51, 93%, 53%)', //yellow
        ord_dlv: 'hsl(147, 81%, 37%)', //green
        ord_dnd: 'hsl(1, 91%, 53%)', //red
      }
    },
  },
  plugins: [],
}

