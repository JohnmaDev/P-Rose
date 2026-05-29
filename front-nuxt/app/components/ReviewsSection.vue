<template>
  <section class="w-full py-24 relative overflow-hidden">
    <div class="absolute inset-0 bg-gradient-to-b from-barber-black via-zinc-950 to-barber-black"></div>

    <div class="relative max-w-6xl mx-auto px-4">
      <!-- Header -->
      <div class="text-center mb-16 animate-on-scroll opacity-0 translate-y-8 transition-all duration-1000">
        <h2 class="text-[2.5rem] sm:text-[4rem] lg:text-[80px] font-black tracking-tighter italic uppercase text-shadow-premium leading-tight">
          {{ t('reviews.headline') }}
          <span class="text-neon-green drop-shadow-[0_0_15px_rgba(57,255,20,0.3)] block sm:inline">
            {{ t('reviews.headlineHighlight') }}
          </span>
        </h2>
        <p class="text-gray-400 text-lg sm:text-xl mt-5 max-w-xl mx-auto italic font-bold tracking-wide">
          {{ t('reviews.sub') }}
        </p>
        <!-- WhatsApp source badge -->
        <div class="inline-flex items-center gap-2 mt-4 px-4 py-2 rounded-full border border-[#25D366]/30 bg-[#25D366]/5">
          <fa-icon :icon="['fab', 'whatsapp']" class="text-[#25D366] text-sm" />
          <span class="text-[#25D366] text-[11px] font-black tracking-widest uppercase">Verificado · WhatsApp</span>
        </div>
      </div>

      <!-- Reviews Grid -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-5">
        <div
          v-for="(review, i) in reviews"
          :key="i"
          class="animate-on-scroll opacity-0 translate-y-8 transition-all duration-700 group"
          :style="`transition-delay: ${i * 80}ms`"
        >
          <div class="h-full bg-zinc-900/60 backdrop-blur-sm border border-white/8 rounded-3xl p-6 flex flex-col gap-4 transition-all duration-500 hover:border-white/20 hover:bg-zinc-900 hover:-translate-y-1 hover:shadow-[0_20px_60px_rgba(0,0,0,0.5)]">
            <!-- Stars -->
            <div class="flex items-center gap-1">
              <fa-icon
                v-for="s in 5"
                :key="s"
                :icon="['fas', 'star']"
                class="text-xs text-yellow-400"
              />
              <span class="text-[10px] text-gray-600 ml-2 font-bold tracking-wider">WhatsApp</span>
            </div>

            <!-- Quote text -->
            <p class="text-gray-300 text-sm leading-relaxed flex-1 italic">
              "{{ review.text }}"
            </p>

            <!-- Reviewer -->
            <div class="flex items-center gap-3 pt-3 border-t border-white/8">
              <!-- Avatar with initials -->
              <div
                class="w-10 h-10 rounded-full flex items-center justify-center text-xs font-black flex-shrink-0"
                :style="`background: ${review.color}20; border: 1.5px solid ${review.color}40; color: ${review.color}`"
              >
                {{ review.initials }}
              </div>
              <div>
                <p class="text-white text-sm font-bold leading-tight">{{ review.name }}</p>
                <p class="text-gray-600 text-[11px] tracking-wider">{{ review.location }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- CTA below reviews -->
      <div class="text-center mt-14 animate-on-scroll opacity-0 translate-y-8 transition-all duration-1000">
        <a
          href="https://api.whatsapp.com/send?phone=573045840264"
          target="_blank"
          rel="noopener noreferrer"
          class="inline-flex items-center gap-3 px-10 py-4 border-2 border-[#25D366]/40 hover:border-[#25D366] text-white hover:text-[#25D366] font-black italic tracking-widest uppercase text-sm transition-all duration-300 rounded-2xl hover:bg-[#25D366]/5 hover:shadow-[0_0_30px_rgba(37,211,102,0.2)]"
        >
          <fa-icon :icon="['fab', 'whatsapp']" class="text-[#25D366]" />
          {{ lang === 'es' ? 'Únete a nuestros clientes satisfechos' : 'Join our satisfied clients' }}
          <fa-icon :icon="['fas', 'arrow-left']" class="rotate-180 text-xs" />
        </a>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { useLanguage } from '~/composables/useLanguage'

const { t, lang } = useLanguage()

const reviews = [
  {
    name: 'Carlos M.',
    initials: 'CM',
    location: 'El Poblado, Medellín 🇨🇴',
    color: '#39FF14',
    text: 'El mejor barbero a domicilio de Medellín. Puntual, profesional y el corte quedó impecable. Llevo 2 años yendo con Andrés y no existe otro igual.',
  },
  {
    name: 'James W.',
    initials: 'JW',
    location: 'Boston, USA 🇺🇸',
    color: '#25D366',
    text: "Best haircut I got in all of Medellín! He came right to my Airbnb, super professional with all his equipment. Will book every time I visit Colombia.",
  },
  {
    name: 'Santiago R.',
    initials: 'SR',
    location: 'Laureles, Medellín 🇨🇴',
    color: '#39FF14',
    text: 'El servicio a domicilio es la mejor idea. Andrés llega con todo el equipo y trabaja como si fuera su propio salón. La barba quedó perfecta.',
  },
  {
    name: 'Michael T.',
    initials: 'MT',
    location: 'Digital Nomad, UK 🇬🇧',
    color: '#a855f7',
    text: "Found PersonalBarber through a friend. Home barber service in Medellín — absolutely perfect for expats and travelers. Highly recommend to anyone visiting.",
  },
  {
    name: 'Andrés V.',
    initials: 'AV',
    location: 'Envigado, Antioquia 🇨🇴',
    color: '#39FF14',
    text: 'Ya son 3 años yendo con Andrés. El corte siempre evoluciona con las últimas tendencias y la atención es de primera. Lo recomiendo a todos mis amigos.',
  },
  {
    name: 'Lucas B.',
    initials: 'LB',
    location: 'Buenos Aires 🇦🇷 → Medellín',
    color: '#f59e0b',
    text: 'Viví 6 meses en Medellín y Andrés fue mi barbero en todo ese tiempo. Nivel internacional, sin duda. El corte y la barba siempre impecables.',
  },
]
</script>
