import type { Config } from "tailwindcss";

export default {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        background: "var(--background)",
        foreground: "var(--foreground)",
        customGray: "#6f6f6f",
        headerHoverGray: "#181818",
        hoverGray: "#151515",
        headerGray: "#131313",
        sidebarBackgroundGray: "#191919",
        sidebarHoverGray:"#191919",
        sidebarTextGray:"#191919",
        sidebarButtonYellow: "#fea613",
        sidebarButtonHover: "#feb43c"
      },
    },
  },
  plugins: [],
} satisfies Config;
