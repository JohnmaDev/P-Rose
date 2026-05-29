<template>
  <ClientOnly>
    <a
      :href="`https://api.whatsapp.com/send?phone=573045840264&text=${encodeURIComponent(waMessage)}`"
      target="_blank"
      rel="noopener noreferrer"
      :aria-label="t('whatsapp.label')"
      class="fixed bottom-6 left-6 z-50 group"
    >
      <!-- Pulse rings -->
      <span class="absolute inset-0 rounded-full bg-[#25D366] opacity-30 animate-ping"></span>
      <span class="absolute inset-0 rounded-full bg-[#25D366] opacity-20 animate-ping [animation-delay:0.4s]"></span>

      <!-- Main button -->
      <div class="relative w-14 h-14 rounded-full bg-[#25D366] flex items-center justify-center shadow-[0_0_25px_rgba(37,211,102,0.5)] group-hover:shadow-[0_0_40px_rgba(37,211,102,0.8)] group-hover:scale-110 transition-all duration-300">
        <fa-icon :icon="['fab', 'whatsapp']" class="text-white text-2xl" />
      </div>

      <!-- Tooltip -->
      <div class="absolute left-16 top-1/2 -translate-y-1/2 bg-zinc-900 border border-white/10 text-white text-[11px] font-bold tracking-wide px-3 py-2 rounded-xl whitespace-nowrap opacity-0 group-hover:opacity-100 transition-all duration-300 pointer-events-none shadow-lg">
        {{ t('whatsapp.tooltip') }}
        <span class="absolute top-1/2 -translate-y-1/2 -left-1.5 w-3 h-3 bg-zinc-900 border-l border-b border-white/10 rotate-45"></span>
      </div>
    </a>
  </ClientOnly>
</template>

<script setup lang="ts">
import { useLanguage } from '~/composables/useLanguage'

const { t, lang } = useLanguage()

const waMessage = computed(() =>
  lang.value === 'en'
    ? 'Hello! I found you on personalbarber.vip and I would like more information.'
    : '¡Hola! Te encontré en personalbarber.vip y me gustaría más información.'
)
</script>
