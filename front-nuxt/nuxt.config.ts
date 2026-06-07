// https://nuxt.com/docs/api/configuration/nuxt-config


export default defineNuxtConfig({

  compatibilityDate: '2025-07-15',
  devtools: { enabled: false },

  devServer: {
    host: '0.0.0.0'
  },

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
    pageTransition: { name: 'page', mode: 'out-in' },
    layoutTransition: { name: 'layout', mode: 'out-in' },
    head: {
      charset: 'utf-8',
      viewport: 'width=device-width, initial-scale=1',
      htmlAttrs: { lang: 'es' },
      title: 'PersonalBarber — Barbería Premium en Medellín',
      meta: [
        { name: 'description', content: 'Barbería premium en Medellín con el mejor estilo. Cortes exclusivos, barba profesional y productos especializados. Reserva tu cita online al instante.' },
        { name: 'theme-color', content: '#0A0A0A' },
        { property: 'og:title', content: 'PersonalBarber — Barbería Premium en Medellín' },
        { property: 'og:description', content: 'Barbería premium en Medellín con el mejor estilo. Cortes exclusivos, barba profesional y productos especializados. Reserva tu cita online al instante.' },
        { property: 'og:type', content: 'website' },
        { property: 'og:site_name', content: 'PersonalBarber' },
        { property: 'og:url', content: 'https://personalbarber.vip' },
        { property: 'og:image', content: 'https://personalbarber.vip/og-image.webp' },
        { property: 'og:image:secure_url', content: 'https://personalbarber.vip/og-image.webp' },
        { property: 'og:image:type', content: 'image/webp' },
        { property: 'og:image:width', content: '1200' },
        { property: 'og:image:height', content: '630' },
        { property: 'og:image:alt', content: 'PersonalBarber — Barbería Premium en Medellín' },
        { name: 'twitter:card', content: 'summary_large_image' },
        { name: 'twitter:title', content: 'PersonalBarber — Barbería Premium en Medellín' },
        { name: 'twitter:description', content: 'Barbería premium en Medellín con el mejor estilo. Cortes exclusivos, barba profesional y productos especializados. Reserva tu cita online al instante.' },
        { name: 'twitter:image', content: 'https://personalbarber.vip/og-image.webp' },
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
        { rel: 'icon', type: 'image/svg+xml', href: '/favicon.svg' },
        { rel: 'apple-touch-icon', sizes: '180x180', href: '/apple-touch-icon.png' },
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

  // Proxy de API: en dev apunta a la API de producción para visualizar los
  // productos reales en local sin depender de la base de datos Go/MongoDB local.
  // En producción (SSR build de Netlify), apunta a las funciones directas.
  routeRules: {
    '/api/**': {
      proxy: process.env.NODE_ENV === 'development'
        ? 'https://personalbarber.vip/api/**'
        : 'https://personalbarber.vip/.netlify/functions/**'
    }
  },
})

