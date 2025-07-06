/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.html"
  ],
  theme: {
    extend: {
      colors: {
        pos_col: 'hsl(75, 29%, 39%)',
        pos_whi: 'hsl(205, 1%, 100%)',
        pos_lgh: 'hsl(205, 1%, 63%)',
        pos_bla: 'hsl(205, 1%, 07%)',

        pos_plc: 'hsl(217, 99%, 56%)',
        pos_cld: 'hsl(357, 95%, 50%)',
        pos_trs: 'hsl(81, 97%, 45%)',
        pos_dlv: 'hsl(141, 89%, 43%)',
        pos_pdg: 'hsl(267, 81%, 57%)',
        pos_rdy: 'hsl(51, 93%, 45%)',
      }
    },
  },
  plugins: [],
}
