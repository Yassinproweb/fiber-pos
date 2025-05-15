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
        ord_plc: 'hsl(217, 81%, 57%)', //blue
        ord_pdg: 'hsl(267, 81%, 53%)', //purple
        ord_cld: 'hsl(217, 7%, 71%)', //gray
        ord_rdy: 'hsl(37, 87%, 53%)', //orange
        ord_trs: 'hsl(53, 87%, 53%)', //yellow
        ord_dlv: 'hsl(147, 81%, 37%)', //green
        ord_dnd: 'hsl(1, 81%, 47%)', //red
      }
    },
  },
  plugins: [],
}

