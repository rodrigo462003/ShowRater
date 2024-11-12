/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors')
module.exports = {
  content: ["views/*.html"],
  theme: {
    extend: {
     colors: {
      primary: colors.blue,
      "primary-accent": colors.blue,
     },
    }
  },
  plugins: [
   require('@tailwindcss/forms'),
  ],
}

