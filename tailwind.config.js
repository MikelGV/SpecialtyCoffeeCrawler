module.exports = {
  purge: [],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {},
  },
    content: [
        "./cmd/web/*.templ",
        "./cmd/web/templates/*.templ",
        "./cmd/web/**/*.html",
        "./cmd/web/**/*.go",
    ],
  variants: {
    extend: {},
  },
  plugins: [],
}
