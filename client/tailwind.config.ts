import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
        colors: {
            yellow: {
                950: '#F7BC1C',
            },
            green: {
                950: '#90843C',
                955: '#95893e',
            },
            blue : {
                950: '#283D5F'
            },
        },
        fontFamily: {
            pacifico: ["Pacifico", "sans-serif"],
            grotesque: ["Grotesque", "sans-serif"],
        },
    },
  },
  plugins: [],
};
export default config;
