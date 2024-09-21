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
      pattern: /(bg)-(cyan|green|red|yellow|lime)-(600)\/(75)$/,
    },
    {
      pattern: /(bg)-(cyan|green|red|yellow|lime)-(500)\/(25)$/,
      variants: ["dark"],
    },
    {
      pattern:
        /(text)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(500|600|700|900)$/,
    },
    {
      pattern:
        /(text)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(500|600)$/,
      variants: ["hover"],
    },
    {
      pattern:
        /(text)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(400)$/,
      variants: ["dark", "hover:dark"],
    },
    {
      pattern:
        /(border)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(600)$/,
      variants: ["hover:dark"],
    },
    {
      pattern:
        /(border)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(500)$/,
      variants: ["hover"],
    },
    {
      pattern:
        /(border)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(400)$/,
    },
    {
      pattern:
        /(border)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(800)$/,
      variants: ["dark"],
    },
    {
      pattern:
        /(border)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(500)$/,
      variants: ["focus"],
    },
    {
      pattern:
        /(border)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(600)$/,
      variants: ["dark:focus"],
    },
    {
      pattern:
        /(bg)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(400)\/(10)$/,
      variants: ["hover"],
    },
    {
      pattern:
        /(bg)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(400)\/(10)$/, //submenufile item
    },
    {
      pattern:
        /(bg)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(500)\/(10)$/, //submenufile item
      variants: ["dark"],
    },
    {
      pattern:
        /(bg)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(500)$/, //repo sidebar item
      variants: ["group-hover"],
    },
    {
      pattern:
        /(bg)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(600)$/, //repo sidebar item
    },
    {
      pattern:
        /(bg)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(700)$/, //repo sidebar item
      variants: ["hover"],
    },
    {
      pattern:
        /(bg)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(700)\/(50)$/, //repo sidebar item
      variants: ["dark"],
    },
    {
      pattern:
        /(bg)-(yellow|violet|teal|sky|pink|indigo|green|red|cyan)-(800)\/(50)$/, //repo sidebar item
      variants: ["hover:dark"],
    },
  ],
  // safelist: [
  //   {
  //     pattern:
  //       /(text|bg|from|to|boder)-(gray|violet|green|red|lime|yellow|teal|sky|cyan|pink|rose|indigo)-(200|300|400|500|600|900)\/(10|20|30|40|50|60|70|75|80|90|100)/,
  //     variants: ["hover", "dark", "hover:dark", "group-hover"],
  //   },
  //   {
  //     pattern:
  //       /(text|bg|from|to|border)-(violet|green|red|yellow|teal|sky|cyan|pink|rose|indigo)-(200|300|400|500|600|700|800|900)/,
  //     variants: [
  //       "hover",
  //       "dark",
  //       "hover:dark",
  //       "group-hover",
  //       "focus",
  //       "dark:focus",
  //     ],
  //   },
  // ],
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
    }),
    require("@tailwindcss/typography"),
  ],
};
export default config;
