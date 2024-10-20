import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import path from "path";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [svelte()],
    resolve: {
        alias: {
            $lib: path.resolve("./src/lib"),
        },
    },
    server: {
        port: 3000,
        proxy: {
            "/api": {
                target: "http://localhost:8080/api",
                changeOrigin: true,
                secure: false,
                rewrite: (p) => p.replace(/^\/api/, '')
            },
            "/new-load-added": {
                target: "ws://localhost:8080/new-load-added",
                changeOrigin: true,
                secure: false,
                rewrite: (p) => p.replace(/^\/new-load-added/, '')
            }

        }
    }
})
