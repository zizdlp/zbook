import type { Config } from "tailwindcss";

const config: Config = {
  darkMode: "class",
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/providers/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  safelist: [
    {
      pattern:
        /(text|bg|from|to|boder)-(gray|violet|green|red|lime|yellow|teal|sky|cyan|pink|rose|indigo)-(200|300|400|500|600|900)\/(10|20|30|40|50|60|70|75|80|90|100)/,
      variants: ["hover", "dark", "hover:dark", "group-hover"],
    },
    {
      pattern:
        /(text|bg|from|to|border)-(violet|green|red|yellow|teal|sky|cyan|pink|rose|indigo)-(200|300|400|500|600|700|800|900)/,
      variants: [
        "hover",
        "dark",
        "hover:dark",
        "group-hover",
        "focus",
        "dark:focus",
      ],
    },
  ],
  theme: {
    extend: {
      fontFamily: {
        jetbrains: ["var(--font-jetbrains-mono)"],
      },
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
    },
  },
  plugins: [
    require("tailwind-scrollbar")({
      nocompatible: true,
      preferredStrategy: "pseudoelements",
    }),
    require("@tailwindcss/typography"),
  ],
};
export default config;
