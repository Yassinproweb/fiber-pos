/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.html"
  ],
  theme: {
    extend: {
      colors: {
        cs_color: 'hsl(187, 41%, 43%)',
        cs_white: 'hsl(187, 1%, 100%)',
        cs_black: 'hsl(187, 1%, 07%)',
        ord_plc: 'hsl(217, 71%, 57%)', //blue
        ord_pdg: 'hsl(277, 71%, 57%)', //purple
        ord_cld: 'hsl(217, 11%, 83%)', //gray
        ord_rdy: 'hsl(47, 81%, 53%)', //orange
        ord_trs: 'hsl(77, 81%, 53%)', //yellow
        ord_dlv: 'hsl(147, 71%, 37%)', //green
        ord_dnd: 'hsl(1, 71%, 47%)', //red
      }
    },
  },
  plugins: [],
}

