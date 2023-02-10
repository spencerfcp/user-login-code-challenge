/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx}",
    "./src/components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      animation: {
        fadeIn: "fadeIn .5s ease-in-out",
      },
      keyframes: () => ({
        fadeIn: {
          "0%": { opacity: 0 },
          "100%": { opacity: 1 },
        },
      }),
      colors: {
        blue: {
          DEFAULT: "#15769A",
          50: "#7ACEED",
          100: "#68C7EA",
          200: "#44BAE5",
          300: "#20ACE1",
          400: "#1A92BE",
          500: "#15769A",
          600: "#0E5069",
          700: "#082A37",
          800: "#010506",
          900: "#000000",
        },
      },
    },
  },
  plugins: [],
};
