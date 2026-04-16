// app.config.ts — configuración accesible en cliente y servidor sin hydration issues
// Los valores aquí se pueden reemplazar por proceso.env al momento de inicio del servidor

export default defineAppConfig({
  // PIN de admin: se pone en el servidor desde .env.local / Netlify
  // Accesible en el cliente vía useAppConfig()
  adminPin: process.env.NUXT_PUBLIC_ADMIN_PIN || '',
})
