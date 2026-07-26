<template>
  <section class="w-full py-24 relative overflow-hidden">

    <div class="relative max-w-7xl mx-auto px-6 lg:px-12">
      <div class="flex flex-col lg:flex-row items-center gap-12 lg:gap-20">

        <!-- Left: Text Content -->
        <div class="w-full lg:w-1/2 text-center lg:text-left animate-on-scroll opacity-0 translate-y-8 transition-all duration-1000">
          <!-- Badge -->
          <div class="inline-flex items-center gap-2 mb-6 px-4 py-2 rounded-full border border-neon-green/30 bg-neon-green/5 mx-auto lg:mx-0">
            <span class="w-2 h-2 rounded-full bg-neon-green animate-pulse shadow-[0_0_8px_rgba(57,255,20,0.8)]"></span>
            <span class="text-neon-green text-[11px] font-black tracking-[0.2em] uppercase">{{ t('feature.badge') }}</span>
          </div>

          <!-- Headline -->
          <h2 class="text-[3rem] sm:text-[4.5rem] lg:text-[80px] font-black tracking-tighter italic uppercase leading-tight text-shadow-premium">
            <span class="text-white block">{{ t('feature.headline1') }}</span>
            <span class="text-neon-green drop-shadow-[0_0_20px_rgba(57,255,20,0.4)] block">{{ t('feature.headline2') }}</span>
          </h2>

          <!-- Sub -->
          <p class="text-gray-300 text-lg sm:text-xl mt-6 mb-8 leading-relaxed max-w-lg mx-auto lg:mx-0">
            {{ t('feature.sub') }}
          </p>

          <!-- Bullets -->
          <ul class="flex flex-col gap-4 mb-10">
            <li
              v-for="(bullet, i) in bullets"
              :key="i"
              class="flex items-center gap-4 text-left"
            >
              <div class="flex-shrink-0 w-8 h-8 rounded-full bg-neon-green/10 border border-neon-green/30 flex items-center justify-center shadow-[0_0_10px_rgba(57,255,20,0.15)]">
                <fa-icon :icon="bullet.icon" class="text-neon-green text-xs" />
              </div>
              <span class="text-gray-300 text-sm sm:text-base font-medium">{{ t(bullet.key) }}</span>
            </li>
          </ul>

          <!-- CTA -->
          <NuxtLink
            to="/agendar"
            class="inline-flex items-center gap-3 px-10 py-5 bg-neon-green text-black text-lg font-black italic tracking-wider uppercase rounded-2xl transition-all duration-300 shadow-[0_0_30px_rgba(57,255,20,0.3)] hover:shadow-[0_0_60px_rgba(57,255,20,0.6)] hover:-translate-y-1 hover:scale-105"
          >
            {{ t('feature.cta') }}
            <fa-icon :icon="['fas', 'map-marker-alt']" />
          </NuxtLink>
        </div>

        <!-- Right: Visual Card -->
        <div class="w-full lg:w-1/2 animate-on-scroll opacity-0 translate-y-8 transition-all duration-1000 [transition-delay:200ms]">
          <div class="relative max-w-md mx-auto">
            <!-- Offset glow -->
            <div class="absolute -inset-4 bg-neon-green/10 rounded-[50px] blur-2xl"></div>

            <!-- Main card -->
            <div class="relative bg-zinc-900/80 backdrop-blur-xl rounded-[40px] border border-white/10 overflow-hidden shadow-2xl p-8 sm:p-10">
              <!-- Top badge -->
              <div class="flex items-center justify-between mb-8">
                <div class="flex items-center gap-2">
                  <img src="/PersonalBarber.svg" alt="PersonalBarber" class="w-8 h-8" width="32" height="32" loading="lazy" decoding="async" />
                  <span class="text-white text-sm font-black tracking-widest uppercase">PersonalBarber</span>
                </div>
                <span class="text-[10px] font-black tracking-widest text-neon-green uppercase bg-neon-green/10 border border-neon-green/20 px-3 py-1 rounded-full">Premium</span>
              </div>

              <!-- Service items -->
              <div class="space-y-4 mb-8">
                <div
                  v-for="(service, i) in featureServices"
                  :key="i"
                  class="flex items-center gap-4 p-4 rounded-2xl border transition-all duration-300"
                  :class="i === 0 ? 'bg-neon-green/8 border-neon-green/30' : 'bg-white/3 border-white/8 hover:border-white/20'"
                >
                  <div class="w-10 h-10 rounded-xl flex items-center justify-center flex-shrink-0" :class="i === 0 ? 'bg-neon-green/15' : 'bg-white/5'">
                    <fa-icon :icon="service.icon" :class="i === 0 ? 'text-neon-green' : 'text-gray-400'" class="text-sm" />
                  </div>
                  <div class="flex-1 min-w-0">
                    <p class="text-white text-sm font-bold truncate">{{ service.name }}</p>
                    <p class="text-gray-500 text-xs">{{ service.detail }}</p>
                  </div>
                  <fa-icon v-if="i === 0" :icon="['fas', 'check']" class="text-neon-green text-xs flex-shrink-0" />
                </div>
              </div>

              <!-- Bottom stat -->
              <div class="border-t border-white/8 pt-6 flex items-center justify-between">
                <div class="text-center">
                  <p class="text-2xl font-black text-neon-green">+100</p>
                  <p class="text-[10px] text-gray-500 uppercase tracking-widest mt-0.5">Reviews WA</p>
                </div>
                <div class="h-8 w-[1px] bg-white/10"></div>
                <div class="text-center">
                  <p class="text-2xl font-black text-white">10+</p>
                  <p class="text-[10px] text-gray-500 uppercase tracking-widest mt-0.5">{{ t('feature.statYears') }}</p>
                </div>
                <div class="h-8 w-[1px] bg-white/10"></div>
                <div class="text-center">
                  <p class="text-2xl font-black text-white">Medellin</p>
                  <p class="text-[10px] text-gray-500 uppercase tracking-widest mt-0.5">{{ t('feature.statCoverage') }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { useLanguage } from '~/composables/useLanguage'

const { t } = useLanguage()

const bullets = [
  { key: 'feature.bullet1', icon: ['fas', 'map-marker-alt'] as [string, string] },
  { key: 'feature.bullet2', icon: ['fas', 'scissors'] as [string, string] },
  { key: 'feature.bullet3', icon: ['fas', 'bolt'] as [string, string] },
]

const featureServices = computed(() => [
  { name: t('feature.service1Name'), detail: t('feature.service1Detail'), icon: ['fas', 'scissors'] as [string, string] },
  { name: t('feature.service2Name'), detail: t('feature.service2Detail'), icon: ['fas', 'spa'] as [string, string] },
  { name: t('feature.service3Name'), detail: t('feature.service3Detail'), icon: ['fas', 'cut'] as [string, string] },
])
</script>
