<template>
  <nav
    ref="navRef"
    class="fixed top-0 left-0 right-0 z-50 transition-all duration-500"
    :class="scrolled ? 'bg-barber-black/90 backdrop-blur-xl shadow-[0_4px_30px_rgba(0,0,0,0.5)]' : 'bg-transparent'"
  >
    <div class="max-w-7xl mx-auto px-4 sm:px-6 h-16 flex items-center justify-between gap-4">

      <!-- Logo -->
      <div class="flex-1 flex justify-start">
        <NuxtLink to="/" aria-label="PersonalBarber — Inicio" class="flex-shrink-0 group">
          <img
            src="/PersonalBarber.svg"
            alt="PersonalBarber"
            fetchpriority="high"
            class="h-9 w-9 transition-all duration-300 group-hover:scale-110 group-hover:drop-shadow-[0_0_8px_rgba(57,255,20,0.6)]"
          />
        </NuxtLink>
      </div>

      <!-- Desktop Navigation -->
      <div class="hidden md:flex flex-1 items-center justify-center gap-2">
        <NuxtLink
          v-for="link in navLinks"
          :key="link.to"
          :to="link.to"
          class="px-4 py-2 text-xs font-black tracking-[0.15em] uppercase text-gray-400 hover:dept-text transition-colors duration-300"
          active-class="dept-text"
        >
          {{ link.label }}
        </NuxtLink>
      </div>

      <!-- Right side: Language toggle + Book CTA -->
      <div class="flex-1 flex items-center justify-end gap-3 flex-shrink-0">
        <!-- Language Toggle -->
        <ClientOnly>
          <button
            @click="toggleLang"
            class="hidden sm:flex items-center gap-1.5 px-3 py-1.5 rounded-full border border-white/15 text-[10px] font-black tracking-widest uppercase transition-all duration-300 hover:dept-border hover:dept-text group"
            :title="lang === 'es' ? 'Switch to English' : 'Cambiar a Español'"
          >
            <span class="text-sm leading-none">{{ lang === 'es' ? '🇨🇴' : '🇺🇸' }}</span>
            <span class="transition-colors duration-300 text-gray-400 group-hover:dept-text">{{ lang === 'es' ? 'ES' : 'EN' }}</span>
          </button>
        </ClientOnly>

        <!-- Mobile Hamburger -->
        <button
          @click="mobileOpen = !mobileOpen"
          class="md:hidden w-9 h-9 flex flex-col items-center justify-center gap-1.5 group"
          :aria-label="mobileOpen ? t('common.closeMenu') : t('common.openMenu')"
        >
          <span
            v-for="i in 3"
            :key="i"
            class="block h-[2px] bg-white transition-all duration-300 rounded-full"
            :class="[
              i === 1 ? (mobileOpen ? 'w-5 translate-y-[7px] rotate-45' : 'w-5') : '',
              i === 2 ? (mobileOpen ? 'opacity-0 w-0' : 'w-3.5') : '',
              i === 3 ? (mobileOpen ? 'w-5 -translate-y-[7px] -rotate-45' : 'w-5') : '',
            ]"
          ></span>
        </button>
      </div>
    </div>

    <!-- Mobile Menu -->
    <Transition name="mobile-menu">
      <div
        v-if="mobileOpen"
        class="md:hidden bg-barber-black/95 backdrop-blur-xl border-t border-white/10 px-4 py-6 flex flex-col gap-4"
      >
        <NuxtLink
          v-for="link in navLinks"
          :key="link.to"
          :to="link.to"
          @click="mobileOpen = false"
          class="flex items-center gap-3 px-4 py-3 rounded-xl text-sm font-black tracking-widest uppercase text-gray-400 hover:text-white hover:bg-white/5 transition-all duration-200"
        >
          <fa-icon :icon="link.icon" class="dept-text text-xs w-4" />
          {{ link.label }}
        </NuxtLink>

        <!-- Mobile Language + Book -->
        <div class="flex items-center gap-3 mt-2 pt-4 border-t border-white/10">
          <ClientOnly>
            <button
              @click="toggleLang"
              class="flex-1 flex items-center justify-center gap-2 px-4 py-2.5 rounded-xl border border-white/15 text-xs font-black tracking-widest uppercase text-gray-400 hover:dept-border hover:dept-text transition-all duration-300"
            >
              <span>{{ lang === 'es' ? '🇨🇴 ES' : '🇺🇸 EN' }}</span>
            </button>
          </ClientOnly>
        </div>
      </div>
    </Transition>
  </nav>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useLanguage } from '~/composables/useLanguage'

const { lang, t, toggleLang } = useLanguage()

const scrolled = ref(false)
const mobileOpen = ref(false)
const navRef = ref<HTMLElement | null>(null)

const navLinks = computed(() => [
  { to: '/', label: t('nav.home'), icon: ['fas', 'store'] as [string, string] },
  { to: '/agendar', label: t('nav.book'), icon: ['fas', 'calendar-check'] as [string, string] },
])

function handleScroll() {
  scrolled.value = window.scrollY > 30
}

function handleClickOutside(event: MouseEvent) {
  if (mobileOpen.value && navRef.value && !navRef.value.contains(event.target as Node)) {
    mobileOpen.value = false
  }
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll, { passive: true })
  document.addEventListener('click', handleClickOutside)
  handleScroll()
})

onUnmounted(() => {
  window.removeEventListener('scroll', handleScroll)
  document.removeEventListener('click', handleClickOutside)
})
</script>

<style scoped>
.mobile-menu-enter-active,
.mobile-menu-leave-active {
  transition: opacity 0.25s ease, transform 0.25s ease;
}
.mobile-menu-enter-from,
.mobile-menu-leave-to {
  opacity: 0;
  transform: translateY(-8px);
}
</style>
