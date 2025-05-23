export default {
  locales: ["en", "ja"],
  output: "src/i18n/locales/$LOCALE.json",
  input: ["src/**/*.{ts,tsx}"],
  // allow keys to be phrases having `:`, `.`
  namespaceSeparator: false,
  keySeparator: false,
  createOldCatalogs: false,
};
