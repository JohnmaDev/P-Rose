<template>
  <section class="w-full py-24 relative overflow-hidden">
    <!-- Fondo con patrón sutil -->
    <div class="absolute inset-0 bg-gradient-to-b from-barber-black via-zinc-950 to-barber-black"></div>
    <div class="absolute inset-0 opacity-[0.03]" style="background-image: repeating-linear-gradient(0deg, transparent, transparent 40px, white 40px, white 41px), repeating-linear-gradient(90deg, transparent, transparent 40px, white 40px, white 41px);"></div>

    <div class="relative max-w-5xl mx-auto px-4">
      <!-- Headline -->
      <div class="text-center mb-16 animate-on-scroll opacity-0 translate-y-8 transition-all duration-1000">
        <h2 class="text-[2.5rem] sm:text-[4rem] lg:text-[80px] font-black tracking-tighter italic uppercase text-shadow-premium leading-none">
          {{ t('metrics.headline') }}
          <span class="text-neon-green block sm:inline drop-shadow-[0_0_15px_rgba(57,255,20,0.3)]">{{ t('metrics.headlineHighlight') }}</span>
        </h2>
        <p class="text-gray-400 text-lg sm:text-xl mt-5 max-w-xl mx-auto italic font-bold tracking-wide">
          {{ t('metrics.sub') }}
        </p>
      </div>

      <!-- Grid de métricas -->
      <div class="grid grid-cols-2 lg:grid-cols-4 gap-5 sm:gap-6">
        <div
          v-for="(metric, i) in metrics"
          :key="metric.key"
          class="animate-on-scroll opacity-0 translate-y-8 transition-all duration-700 group"
          :style="`transition-delay: ${i * 100}ms`"
        >
          <div class="relative bg-zinc-900/60 backdrop-blur-sm border border-white/8 rounded-3xl p-6 sm:p-8 flex flex-col items-center text-center gap-3 overflow-hidden transition-all duration-500 hover:border-white/20 hover:bg-zinc-900/80">
            <!-- Glow background on hover -->
            <div class="absolute inset-0 opacity-0 group-hover:opacity-100 transition-opacity duration-500 rounded-3xl" :style="`background: radial-gradient(circle at 50% 0%, ${metric.color}15 0%, transparent 60%)`"></div>

            <!-- Icon -->
            <div class="relative w-14 h-14 rounded-2xl flex items-center justify-center transition-transform duration-300 group-hover:scale-110" :style="`background: ${metric.color}15; border: 1px solid ${metric.color}30`">
              <fa-icon :icon="metric.icon" class="text-xl transition-all duration-300" :style="`color: ${metric.color}`" />
            </div>

            <!-- Number with count-up -->
            <div class="relative">
              <p class="text-4xl sm:text-5xl font-black tracking-tighter leading-none" :style="`color: ${metric.color}`">
                <span ref="counterRefs" :data-target="metric.value" :data-prefix="metric.prefix" :data-suffix="metric.suffix">
                  {{ metric.prefix }}{{ displayed[i] }}{{ metric.suffix }}
                </span>
              </p>
            </div>

            <!-- Label -->
            <p class="text-[11px] sm:text-xs font-black tracking-[0.15em] uppercase text-gray-400 leading-tight">
              {{ t(`metrics.${metric.key}`) }}
            </p>
            <p v-if="metric.subKey" class="text-[10px] text-gray-600 tracking-widest uppercase -mt-1">
              {{ t(`metrics.${metric.subKey}`) }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { useLanguage } from '~/composables/useLanguage'

const { t } = useLanguage()

const metrics = [
  { key: 'years', value: 10, prefix: '+', suffix: '', icon: ['fas', 'gem'], color: '#39FF14' },
  { key: 'clients', value: 50, prefix: '', suffix: '+', icon: ['fas', 'scissors'], color: '#39FF14' },
  { key: 'reviews', value: 100, prefix: '+', suffix: '', icon: ['fab', 'whatsapp'], color: '#25D366' },
  { key: 'delivery', value: 100, prefix: '', suffix: '%', icon: ['fas', 'truck'], color: '#39FF14', subKey: 'deliverySub' },
]

const displayed = ref(metrics.map(() => 0))
let hasAnimated = false

function animateCount(index: number, target: number) {
  const duration = 1800
  const start = performance.now()
  const step = (now: number) => {
    const progress = Math.min((now - start) / duration, 1)
    // Easing out cubic
    const eased = 1 - Math.pow(1 - progress, 3)
    displayed.value[index] = Math.floor(eased * target)
    if (progress < 1) requestAnimationFrame(step)
    else displayed.value[index] = target
  }
  requestAnimationFrame(step)
}

onMounted(() => {
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting && !hasAnimated) {
        hasAnimated = true
        metrics.forEach((m, i) => {
          setTimeout(() => animateCount(i, m.value), i * 150)
        })
        observer.disconnect()
      }
    })
  }, { threshold: 0.4 })

  const el = document.querySelector('.metrics-trigger')
  if (el) observer.observe(el)

  // Fallback: observe the section itself
  const sections = document.querySelectorAll('section')
  sections.forEach(s => {
    if (s.textContent?.includes('+10') || s.querySelector('[data-target]')) {
      observer.observe(s)
    }
  })
})
</script>
