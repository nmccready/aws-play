const globals = require("globals");
const js = require("@eslint/js");
const { defineConfig } = require("eslint/config");

module.exports = defineConfig([
  {
    files: ["**/*.js"],
    languageOptions: {
      globals: {
        ...globals.commonjs,
        ...globals.node,
      },
    },
    plugins: {
      js,
    },
    extends: ["js/recommended"],
    rules: {
      "no-unused-vars": "warn",
      "comma-dangle": ["error", "always-multiline"],
    },
  },
]);
