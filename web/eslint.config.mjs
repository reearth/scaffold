import config from "eslint-config-reearth";

/** @type { import("eslint").Linter.Config[] } */
export default [
  ...config("", { reactRecommended: true }),
  {
    ignores: ["**/__gen__", "**/dist", "**/coverage"]
  }
];
