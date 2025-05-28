import config from "eslint-config-reearth";

/** @type { import("eslint").Linter.Config[] } */
export default [
  ...config("", { reactRecommended: true }),
  {
    ignores: ["/dist", "/coverage", "**/__gen__"],
  },
];
