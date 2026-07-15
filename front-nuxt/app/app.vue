<template>
  <main id="app-container">
    <!-- Navbar global persistente en todas las páginas, excepto checkout y admin -->
    <AppNavbar v-if="showNavbar" />

    <NuxtPage />

    <ClientOnly>
      <!-- Carrito: solo cuando no estamos en Home ni en páginas que no corresponde -->
      <template v-if="showCart">
        <CartDrawer :is-open="cartOpen" @close="cartOpen = false" />
        <CartIcon v-if="!cartOpen" @open="cartOpen = true" />
      </template>

      <!-- Botón flotante de WhatsApp — oculto en checkout -->
      <WhatsAppFab v-if="!cartOpen && route.name !== 'checkout'" />

      <!-- Botón Scroll Top -->
      <button
        v-if="showScrollTop && !cartOpen"
        @click="scrollToTop"
        :aria-label="t('common.scrollTop')"
        class="fixed bottom-10 sm:bottom-6 right-6 z-50 bg-neon-green text-black px-4 py-2 rounded-full shadow-[0_0_15px_rgba(57,255,20,0.4)] hover:bg-neon-green-dark transition-all duration-300 focus:outline-none hover:scale-110"
      >
        ↑ Top
      </button>
    </ClientOnly>
  </main>
</template>

<script setup lang="ts">
import { useLanguage } from '~/composables/useLanguage'

const { initLang, t } = useLanguage()
const route = useRoute()
const config = useRuntimeConfig()

// Configurar URL Canónica Dinámica para evitar advertencias de contenido duplicado en Google Search Console
useHead(() => {
  const siteUrl = config.public.siteUrl || 'https://personalbarber.vip'
  const cleanBase = siteUrl.endsWith('/') ? siteUrl.slice(0, -1) : siteUrl
  const canonicalUrl = `${cleanBase}${route.path}`
  
  return {
    link: [
      {
        rel: 'canonical',
        href: canonicalUrl
      }
    ]
  }
})

const cartOpen = ref(false)
const showScrollTop = ref(false)

// Mostrar navbar global excepto en checkout y admin
const showNavbar = computed(() =>
  route.name !== 'checkout' && route.name !== 'admin'
)

// Mostrar carrito en todas las páginas de tienda (home, producto, etc.) excepto admin y checkout
const showCart = computed(() =>
  route.name !== 'admin' && route.name !== 'checkout'
)

function scrollToTop() {
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

function handleScroll() {
  const shouldShow = window.scrollY > 200
  if (showScrollTop.value !== shouldShow) {
    showScrollTop.value = shouldShow
  }
}

onMounted(() => {
  initLang()
  window.addEventListener('scroll', handleScroll, { passive: true })
})
onUnmounted(() => window.removeEventListener('scroll', handleScroll))
</script>
