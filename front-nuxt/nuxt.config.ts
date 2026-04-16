// https://nuxt.com/docs/api/configuration/nuxt-config


export default defineNuxtConfig({

  compatibilityDate: '2025-07-15',
  devtools: { enabled: false },

  // SSR habilitado para SEO óptimo
  ssr: true,

  // Módulos
  modules: [
    '@nuxtjs/sitemap',
  ],

  // Tailwind V4 via PostCSS
  postcss: {
    plugins: {
      '@tailwindcss/postcss': {},
    },
  },



  // CSS global
  css: [
    '~/assets/css/main.css',
    '@fortawesome/fontawesome-svg-core/styles.css'
  ],



  // Variables de entorno — accesibles en el servidor y cliente
  runtimeConfig: {
    // Solo servidor (privadas)
    adminPin: process.env.NUXT_ADMIN_PIN || '',
    // Públicas (cliente + servidor) — Nuxt las reemplaza en runtime con NUXT_PUBLIC_*
    public: {
      adminPin: process.env.NUXT_PUBLIC_ADMIN_PIN || '',
      siteUrl: process.env.NUXT_PUBLIC_SITE_URL || 'https://personalbarber.vip',
    }
  },

  // Vite define — inyecta el PIN en build-time para compatibilidad con <ClientOnly>
  vite: {
    define: {
      __ADMIN_PIN__: JSON.stringify(process.env.NUXT_PUBLIC_ADMIN_PIN || ''),
    }
  },

  // App head global — metadatos base SEO
  app: {
    head: {
      charset: 'utf-8',
      viewport: 'width=device-width, initial-scale=1',
      htmlAttrs: { lang: 'es' },
      title: 'PersonalBarber — Barbería Premium en Medellín',
      meta: [
        { name: 'description', content: 'Barbería premium en Medellín con el mejor estilo. Cortes exclusivos, barba profesional y productos especializados. Reserva tu cita online al instante.' },
        { name: 'theme-color', content: '#0A0A0A' },
        { property: 'og:type', content: 'website' },
        { property: 'og:site_name', content: 'PersonalBarber' },
        { property: 'og:image', content: '/og-image.webp' },
        { name: 'twitter:card', content: 'summary_large_image' },
      ],
      link: [
        { rel: 'icon', type: 'image/svg+xml', href: '/favicon.svg' },
        { rel: 'preconnect', href: 'https://fonts.googleapis.com' },
        { rel: 'preconnect', href: 'https://fonts.gstatic.com', crossorigin: '' },
        { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Oswald:wght@400;500;600;700&display=swap' },
      ]
    }
  },

  // Sitemap automático (@nuxtjs/sitemap v8 — auto-discovers rutas)
  sitemap: {},

  // Nitro — preset Netlify para deploy
  nitro: {
    preset: 'netlify',
  },

  // Dev server proxy para API de Go — excluye endpoints propios de Nitro
  routeRules: {
    '/api/get_products': { proxy: process.env.NODE_ENV === 'development' ? 'https://personalbarber.vip/api/get_products' : undefined },
    '/api/get_categories': { proxy: process.env.NODE_ENV === 'development' ? 'https://personalbarber.vip/api/get_categories' : undefined },
    '/api/get_cuts': { proxy: process.env.NODE_ENV === 'development' ? 'https://personalbarber.vip/api/get_cuts' : undefined },
    '/api/add_product': { proxy: process.env.NODE_ENV === 'development' ? 'https://personalbarber.vip/api/add_product' : undefined },
    '/api/update_product': { proxy: process.env.NODE_ENV === 'development' ? 'https://personalbarber.vip/api/update_product' : undefined },
    '/api/delete_product': { proxy: process.env.NODE_ENV === 'development' ? 'https://personalbarber.vip/api/delete_product' : undefined },
    '/api/manage_categories': { proxy: process.env.NODE_ENV === 'development' ? 'https://personalbarber.vip/api/manage_categories' : undefined },
    '/api/manage_cuts': { proxy: process.env.NODE_ENV === 'development' ? 'https://personalbarber.vip/api/manage_cuts' : undefined },
    '/api/reservations': { proxy: process.env.NODE_ENV === 'development' ? 'https://personalbarber.vip/api/reservations' : undefined },
    // /api/admin-auth NO se proxea — es handler nativo de Nitro
  },
})

