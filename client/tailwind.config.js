/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/*.{html,js,ts,vue}",
    "./src/views/**/*.{html,js,ts,vue}",
    "./src/components/**/*.{html,js,ts,vue}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
};
