<template>
  <main 
    id="app-container"
    :style="{
      '--dept-color': accentColor,
      '--dept-glow': accentGlow,
      '--dept-color-10': accentColor + '1a',
      '--dept-color-30': accentColor + '4d',
    }"
    :class="activeDepartment === 'all' ? 'dept-mode-all' : ''"
    class="min-h-screen text-white relative transition-colors duration-500"
  >
    <!-- Navbar global persistente en todas las páginas, excepto checkout y admin -->
    <AppNavbar v-if="showNavbar" />

    <NuxtPage />

    <ClientOnly>
      <!-- Carrito: solo cuando no estamos en Home ni en páginas que no corresponde -->
      <template v-if="showCart">
        <CartDrawer :is-open="cartOpen" @close="cartOpen = false" />
        <CartIcon v-if="!cartOpen" @open="cartOpen = true" />
      </template>

      <!-- Botón Scroll Top (oculto en checkout y admin para UI totalmente limpia) -->
      <button
        v-if="showScrollTop && !cartOpen && route.name !== 'checkout' && route.name !== 'admin'"
        @click="scrollToTop"
        :aria-label="t('common.scrollTop')"
        class="fixed z-50 dept-bg text-black font-black px-3.5 py-1.5 rounded-full shadow-[0_0_15px_var(--dept-glow)] transition-all duration-300 focus:outline-none hover:scale-110 text-xs"
        :class="route.name === 'tienda-producto-slug' ? 'bottom-[4.5rem] left-4 sm:bottom-6 sm:left-auto sm:right-24' : 'bottom-6 right-20 sm:right-24'"
      >
        ↑ Top
      </button>
    </ClientOnly>
  </main>
</template>

<script setup lang="ts">
import { useLanguage } from '~/composables/useLanguage'
import { useDepartment } from '~/composables/useDepartment'

const { initLang, t } = useLanguage()
const { accentColor, accentGlow, activeDepartment } = useDepartment()
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

// Mostrar carrito en todas las páginas de tienda (home, producto, etc.) excepto admin, checkout y agendar
const showCart = computed(() =>
  route.name !== 'admin' && route.name !== 'checkout' && route.name !== 'agendar'
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
