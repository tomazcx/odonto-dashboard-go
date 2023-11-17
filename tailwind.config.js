/** @type {import('tailwindcss').Config} */
export default {
  content: ["**/*.templ"],
  theme: {
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["winter"]
  }
}

