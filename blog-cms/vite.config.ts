import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite';
import { ArcoResolver } from 'unplugin-vue-components/resolvers';
import { resolve } from "path";

export default defineConfig({
  plugins: [
    vue(),
    Components({
      resolvers: [
        ArcoResolver()
      ]
    }),
  ],
    resolve: {
    alias: {
      "@": resolve(__dirname, "./src/"),
    },
  },
  server: {
    port: 3000,
    host: '0.0.0.0',
  }
})
