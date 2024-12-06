import { defineConfig } from "vite";
import path from "path";

export default defineConfig({
  build: {
    lib: {
      entry: [path.resolve(__dirname, "src/main.ts")],
      formats: ["es"],
      name: "[name]",
    },
    outDir: "../static",
    emptyOutDir: false,
  },
});
