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
          warmWhite: '#F7F3EC',
          cream: '#EFE7DA',
          latte: '#D7C4A8',
          brown: '#7A5638',
          espresso: '#2A1E17',
          charcoal: '#1E1E1E',
          softGray: '#A8A29A',
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
