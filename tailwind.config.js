/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/*.html", "./templates/fragments/*.html"],
  theme: {
    extend: {
      fontFamily: {
        sans: ["Poppins", "sans-serif"]
      }
    }
  },
  plugins: []
}
 