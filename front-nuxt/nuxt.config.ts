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



  // Variables de entorno
  runtimeConfig: {
    // Solo servidor (privadas) — NUNCA expuestas al cliente
    adminPin: process.env.NUXT_ADMIN_PIN || process.env.NUXT_PUBLIC_ADMIN_PIN || '',
    // Públicas (cliente + servidor)
    public: {
      siteUrl: process.env.NUXT_PUBLIC_SITE_URL || 'https://personalbarber.vip',
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
      title: 'PersonalBarber — Tienda Online de Barbería | Medellín',
      meta: [
        { name: 'description', content: 'Tienda online de productos profesionales de barbería, cuidado personal y moda en Medellín. Envíos a toda Colombia. También agenda tu cita con el barber a domicilio.' },
        { name: 'theme-color', content: '#0A0A0A' },
        { property: 'og:title', content: 'PersonalBarber — Tienda Online de Barbería | Medellín' },
        { property: 'og:description', content: 'Tienda online de barbería premium. Ceras, maquinas, cuidado de barba, skincare y más. Compra online con envío a toda Colombia.' },
        { property: 'og:type', content: 'website' },
        { property: 'og:site_name', content: 'PersonalBarber' },
        { property: 'og:url', content: 'https://personalbarber.vip' },
        { property: 'og:image', content: 'https://personalbarber.vip/og-image.webp' },
        { property: 'og:image:secure_url', content: 'https://personalbarber.vip/og-image.webp' },
        { property: 'og:image:type', content: 'image/webp' },
        { property: 'og:image:width', content: '1200' },
        { property: 'og:image:height', content: '630' },
        { property: 'og:image:alt', content: 'PersonalBarber — Tienda Online de Barbería en Medellín' },
        { name: 'twitter:card', content: 'summary_large_image' },
        { name: 'twitter:title', content: 'PersonalBarber — Tienda Online de Barbería | Medellín' },
        { name: 'twitter:description', content: 'Tienda online de barbería premium en Medellín. Compra online con envío a toda Colombia.' },
        { name: 'twitter:image', content: 'https://personalbarber.vip/og-image.webp' },
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
        { rel: 'icon', type: 'image/svg+xml', href: '/favicon.svg' },
        { rel: 'apple-touch-icon', sizes: '180x180', href: '/apple-touch-icon.png' },
        // Preconnect a Cloudinary — elimina ~300ms de latencia DNS/TLS en imágenes
        { rel: 'preconnect', href: 'https://res.cloudinary.com' },
        // Preload hero image — mejora LCP crítico (mobile vs desktop)
        { rel: 'preload', as: 'image', href: '/bg_vertical_mobile.webp', media: '(max-width: 640px)' },
        { rel: 'preload', as: 'image', href: '/bg_vertical.webp', media: '(min-width: 641px)' },
        // Google Fonts: preconnect first — NON-BLOCKING via preload+onload trick
        { rel: 'preconnect', href: 'https://fonts.googleapis.com' },
        { rel: 'preconnect', href: 'https://fonts.gstatic.com', crossorigin: 'anonymous' },
        // Preload the font CSS non-blocking — eliminates 780ms render-blocking
        {
          rel: 'preload',
          as: 'style',
          href: 'https://fonts.googleapis.com/css2?family=Oswald:wght@400;500;600;700&family=Source+Serif+4:ital,wght@0,300;0,400;1,300&display=swap',
          onload: "this.onload=null;this.rel='stylesheet'"
        },
      ],
      noscript: [
        { innerHTML: '<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Oswald:wght@400;500;600;700&family=Source+Serif+4:ital,wght@0,300;0,400;1,300&display=swap">' }
      ]
    }
  },

  // Sitemap automático (@nuxtjs/sitemap v8 — auto-discovers rutas)
  sitemap: {},

  // Nitro — preset Netlify para deploy
  nitro: {
    preset: 'netlify',
  },

  // Proxy de API + headers de seguridad para mejorar Best Practices score
  routeRules: {
    '/api/**': {
      proxy: process.env.NODE_ENV === 'development'
        ? 'https://personalbarber.vip/api/**'
        : 'https://personalbarber.vip/.netlify/functions/**'
    },
    // Aplicar security headers a todo el sitio
    '/**': {
      headers: {
        'Strict-Transport-Security': 'max-age=63072000; includeSubDomains; preload',
        'Cross-Origin-Opener-Policy': 'same-origin',
        'X-Content-Type-Options': 'nosniff',
        'X-Frame-Options': 'SAMEORIGIN',
        'Referrer-Policy': 'strict-origin-when-cross-origin',
      }
    }
  },
})

