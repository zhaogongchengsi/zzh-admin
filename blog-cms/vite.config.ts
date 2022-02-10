import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite';
import { ArcoResolver } from 'unplugin-vue-components/resolvers';
import { resolve, join } from "path";

export default defineConfig({
  base: "./",
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
  build: {
      outDir: join(__dirname, 'app/dist/'),
      assetsDir: '', // 相对路径 加载问题
      rollupOptions: {
        output: {
          format: 'es',
        },
        external: ['electron'], // 告诉 Rollup 不要去打包 electron
      },
  },
  server: {
    port: 3000,
    host: '0.0.0.0',
  }
})
