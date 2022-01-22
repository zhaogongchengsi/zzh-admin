import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";
import { resolve } from "path";
// import analyze from "rollup-plugin-analyzer";
import visualizer from "rollup-plugin-visualizer";

export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
    visualizer(),
  ],
  resolve: {
    alias: {
      "@": resolve(__dirname, "./src/"),
    },
  },
  server: {
    host: "0.0.0.0",
    port: 3000,
    poen: true,
  },
});
