const colors = require('tailwindcss/colors')

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'internal/app/src/*.templ',
    'internal/app/src/components/*.templ',
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
    require("daisyui"),
  ],
  daisyui: {
    themes: false,
    themes: ["cupcake"],
  },

}
