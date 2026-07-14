<template>
  <!-- Ícono flotante del carrito con badge de cantidad -->
  <ClientOnly>
    <Transition name="fade">
      <button
        @click="$emit('open')"
        :aria-label="t('cart.openCart')"
        :class="['fixed right-5 z-50 w-14 h-14 rounded-full bg-neon-green text-black flex items-center justify-center shadow-lg shadow-neon-green/30 hover:bg-neon-green-dark transition-all duration-500 ease-out hover:scale-110 active:scale-95',
          y > 50 ? 'bottom-24' : 'bottom-6 right-6']"
      >
        <fa-icon :icon="['fas', 'shopping-bag']" class="text-lg" />
        <!-- Badge de cantidad -->
        <Transition name="badge-pop">
          <span
            v-if="cartCount > 0"
            class="absolute -top-1 -right-1 bg-black text-neon-green text-[10px] font-black w-5 h-5 rounded-full flex items-center justify-center border border-neon-green/50"
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
import { useWindowScroll } from '@vueuse/core'

const { t } = useLanguage()
defineEmits<{ open: [] }>()
const { cartCount } = useCart()
const { y } = useWindowScroll()
</script>

<style scoped>
.badge-pop-enter-active { transition: all 0.25s cubic-bezier(0.34, 1.56, 0.64, 1); }
.badge-pop-leave-active { transition: all 0.15s ease-in; }
.badge-pop-enter-from, .badge-pop-leave-to { opacity: 0; transform: scale(0); }
</style>
