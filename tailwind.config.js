/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./views/**/*.html"
  ],
  theme: {
    extend: {
      colors: {
        posblue100: 'oklch(0.72, 0.11, 173)',
      }
    },
  },
  plugins: [],
}

