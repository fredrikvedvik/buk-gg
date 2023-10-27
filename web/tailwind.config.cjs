/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./index.html", "./src/**/*.{vue,js,ts,jsx,tsx}"],
    theme: {
        extend: {
            fontFamily: {
                barlow: ["'Barlow'", "sans-serif"],
            },
        },
    },
    plugins: [require("@tailwindcss/typography")],
};
