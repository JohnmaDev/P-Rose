<template>
  <!-- Ícono flotante del carrito con badge de cantidad -->
  <ClientOnly>
    <Transition name="fade">
      <button
        @click="$emit('open')"
        :aria-label="t('cart.openCart')"
        :class="['fixed right-5 z-50 w-14 h-14 rounded-full dept-bg text-black flex items-center justify-center shadow-lg shadow-[0_0_15px_var(--dept-glow)] transition-all duration-500 ease-out hover:scale-110 active:scale-95',
          y > 50 ? 'bottom-24' : 'bottom-6 right-6']"
      >
        <fa-icon :icon="['fas', 'shopping-bag']" class="text-lg" />
        <!-- Badge de cantidad -->
        <Transition name="badge-pop">
          <span
            v-if="cartCount > 0"
            class="absolute -top-1 -right-1 bg-black dept-text text-[10px] font-black w-5 h-5 rounded-full flex items-center justify-center border dept-border"
          >
            {{ cartCount > 9 ? '9+' : cartCount }}
          </span>
        </Transition>
      </button>
    </Transition>
  </ClientOnly>
</template>

<script setup lang="ts">
import { useLanguage } from '~/composables/useLanguage'

const { t } = useLanguage()
defineEmits<{ open: [] }>()
const { cartCount } = useCart()

const y = ref(0)

onMounted(() => {
  const onScroll = () => {
    y.value = window.scrollY
  }
  window.addEventListener('scroll', onScroll, { passive: true })
  onScroll()
  
  onUnmounted(() => {
    window.removeEventListener('scroll', onScroll)
  })
})
</script>

<style scoped>
.badge-pop-enter-active { transition: all 0.25s cubic-bezier(0.34, 1.56, 0.64, 1); }
.badge-pop-leave-active { transition: all 0.15s ease-in; }
.badge-pop-enter-from, .badge-pop-leave-to { opacity: 0; transform: scale(0); }
</style>
