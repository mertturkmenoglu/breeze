const { fontFamily } = require("tailwindcss/defaultTheme");

/** @type {import('tailwindcss').Config} */

module.exports = {
  content: [
    "../views/**/*.{templ,go}",
    "../partials/**/*.{templ,go}",
    "../layouts/**/*.{templ,go}",
  ],
  theme: {
    extend: {},
  },
};
