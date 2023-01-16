import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'

// 自动引入element plus所需的工作库
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'

// https://vitejs.dev/config/
export default defineConfig({
    build: {
      outDir: path.join("..", "web"),
      rollupOptions: {
          input: {
              main: path.resolve(__dirname, 'index.html'),
              mobile: path.resolve(__dirname, 'mobile/index.html'),
              login: path.resolve(__dirname, 'login/index.html')
          }
      }
    },
    plugins: [
        vue(),
        AutoImport({
            resolvers: [ElementPlusResolver()],
        }),
        Components({
            resolvers: [ElementPlusResolver()],
        }),
    ],
    resolve: {
        alias: {
            '@': path.join(__dirname, 'src')
        }
    },
    css: {
        preprocessorOptions: {
            scss: {
                additionalData: `@import "@/assets/mixin.scss";`
            }
        }
    }
})
