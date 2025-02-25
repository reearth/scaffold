import react from "@vitejs/plugin-react-swc";
import { defineConfig } from "vite";
import gitRevision from "vite-plugin-git-revision";
import tsconfigPaths from "vite-tsconfig-paths";

export default defineConfig({
  plugins: [react(), tsconfigPaths(), gitRevision({})],
});
