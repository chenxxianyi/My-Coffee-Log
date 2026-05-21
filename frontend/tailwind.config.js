/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        coffee: {
          warmWhite: '#FFF2DB',
          cream: '#FDE8C2',
          latte: '#E76F51',
          brown: '#5C3D2E',
          espresso: '#5C3D2E',
          charcoal: '#2A1A0E',
          softGray: '#C0A07C',
        }
      },
      fontFamily: {
        serif: ['"Cormorant Garamond"', 'serif'],
        sans: ['"Plus Jakarta Sans"', 'sans-serif'],
      }
    },
  },
  plugins: [],
}
