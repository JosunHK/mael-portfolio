const {
    fontFamily
} = require("tailwindcss/defaultTheme")
/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: "class",
    content: [
        './web/static/global.css',
        './web/templates/**/*.{go,js,html,templ}',
    ],
    theme: {
        container: {
            center: true,
            padding: "2rem",
            screens: {
                "sm": "350px",
                "desktop": "1400px",
            },
        },
        backgroundSize: {
            'full':"100% 100%"
        },
        extend: {
            fontFamily: {
                'sans': ['Inter', 'sans-serif', "ui-sans-serif", "system-ui", "sans-serif", "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"],
                'marker': ["Poiret One", "sans-serif"],
                'chill': ["LXGW WenKai Mono TC", "monospace"],
                'barlow': ["Barlow", "sans-serif"],
            },
            colors: {
                border: "hsl(var(--border))",
                input: "hsl(var(--input))",
                ring: "hsl(var(--ring))",
                background: "hsl(var(--background))",
                foreground: "hsl(var(--foreground))",
                eRed: "hsl(var(--eRed))",
                eOrange: "hsl(var(--eOrange))",
                eYellow: "hsl(var(--eYellow))",
                eGreen: "hsl(var(--eGreen))",
                eBlue: "hsl(var(--eBlue))",
                eLavender: "hsl(var(--eLavender))",
                ePurple: "hsl(var(--ePurple))",
                eBrown: "hsl(var(--eBrown))",
                primary: {
                    DEFAULT: "hsl(var(--primary))",
                    foreground: "hsl(var(--primary-foreground))",
                },
                secondary: {
                    DEFAULT: "hsl(var(--secondary))",
                    foreground: "hsl(var(--secondary-foreground))",
                },
                destructive: {
                    DEFAULT: "hsl(var(--destructive))",
                    foreground: "hsl(var(--destructive-foreground))",
                },
                muted: {
                    DEFAULT: "hsl(var(--muted))",
                    foreground: "hsl(var(--muted-foreground))",
                },
                accent: {
                    DEFAULT: "hsl(var(--accent))",
                    foreground: "hsl(var(--accent-foreground))",
                },
                popover: {
                    DEFAULT: "hsl(var(--popover))",
                    foreground: "hsl(var(--popover-foreground))",
                },
                card: {
                    DEFAULT: "hsl(var(--card))",
                    foreground: "hsl(var(--card-foreground))",
                },
            },
            borderRadius: {
                lg: `var(--radius)`,
                md: `calc(var(--radius) - 2px)`,
                sm: "calc(var(--radius) - 4px)",
            },
            keyframes: {
                "fade-in": {
                    from: {
                        opacity: "0"
                    },
                    to: {
                        opacity: "1"
                    },
                },
                "slide-in": {
                    from: {
                        width: "0"
                    },
                    to: {
                        width: "55dvw"
                    },
                },
                "slide-out": {
                    from: {
                        width: "55dvw"
                    },
                    to: {
                        width: "0"
                    },
                },
                "accordion-down": {
                    from: {
                        height: "0"
                    },
                    to: {
                        height: "var(--radix-accordion-content-height)"
                    },
                },
                "accordion-up": {
                    from: {
                        height: "var(--radix-accordion-content-height)"
                    },
                    to: {
                        height: "0"
                    },
                },
            },
            animation: {
                "fade-in": "fade-in 0.2s ease-in",
                "slide-in": "slide-in 0.1s ease-in forwards",
                "slide-out": "slide-out 0.1s ease-out forwards",
                "accordion-down": "accordion-down 0.2s ease-out",
                "accordion-up": "accordion-up 0.2s ease-out",
            },
        },
        plugins: [
            require('@tailwindcss/forms'),
            require('@tailwindcss/typography'),
        ]
    },
}
