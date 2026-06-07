<template>
  <div class="flex flex-col items-center w-full">
    <!-- Logo SVG con enlace de regreso al inicio (accesibilidad + SEO) -->
    <NuxtLink to="/" aria-label="PersonalBarber — Ir al inicio">
      <img
        src="/PersonalBarber.svg"
        class="w-24 h-24 sm:w-32 sm:h-32 opacity-75 hover:opacity-100 transition-all duration-700 hover:scale-110"
        alt="PersonalBarber — Barbería Premium Medellín"
        width="128"
        height="128"
        loading="lazy"
      />
    </NuxtLink>

    <!-- Footer con NAP (Name + Address + Phone) — señal de SEO local critical -->
    <footer class="text-center text-[10px] sm:text-xs text-gray-500 py-10 tracking-widest uppercase">
      <!-- Datos barbero — Schema NAP para SEO local de Google -->
      <address class="not-italic flex flex-col items-center gap-1 mb-4 text-gray-600 text-[10px]">
        <span>📍 La 4 Sur · Medellín, Antioquia, Colombia</span>
        <a
          href="tel:+573045840264"
          class="hover:text-neon-green transition duration-300 cursor-pointer"
          aria-label="Llamar a PersonalBarber"
        >
          📞 +57 304 584 0264
        </a>
      </address>

      <!-- Redes sociales del barbero -->
      <nav class="flex items-center justify-center gap-4 mb-4" :aria-label="t('footer.socialLabel')">
        <a
          href="https://www.instagram.com/pipehp_/"
          target="_blank"
          rel="noopener noreferrer"
          class="text-gray-600 hover:text-pink-500 transition duration-300"
          aria-label="Instagram de PersonalBarber"
        >
          <fa-icon :icon="['fab', 'instagram']" />
        </a>
        <a
          href="https://www.tiktok.com/@pipehpbarber"
          target="_blank"
          rel="noopener noreferrer"
          class="text-gray-600 hover:text-white transition duration-300"
          aria-label="TikTok de PersonalBarber"
        >
          <fa-icon :icon="['fab', 'tiktok']" />
        </a>
        <a
          href="https://wa.me/573045840264"
          target="_blank"
          rel="noopener noreferrer"
          class="text-gray-600 hover:text-neon-green transition duration-300"
          aria-label="WhatsApp de PersonalBarber"
        >
          <fa-icon :icon="['fab', 'whatsapp']" />
        </a>
      </nav>

      <!-- Créditos del desarrollador -->
      <p class="flex flex-col sm:flex-row items-center justify-center gap-2">
        <span class="opacity-50">
          &lt;/&gt; 2026 — Built with ❤️ by
        </span>
        <a
          href="https://www.linkedin.com/in/john-mario-echavarria-bermudez/"
          target="_blank"
          rel="noopener noreferrer"
          class="text-neon-green font-bold hover:text-white transition duration-500 cursor-pointer"
          aria-label="Perfil de LinkedIn de JohnmaDev"
        >
          JohnmaDev
        </a>
        <span class="hidden sm:inline opacity-30">·</span>
        <span class="opacity-50">{{ t('footer.rights') }}</span>
      </p>
    </footer>
  </div>
</template>

<script setup lang="ts">
import { useLanguage } from '~/composables/useLanguage'
const { t } = useLanguage()
// Componente footer con logo, datos de contacto, redes sociales y créditos del desarrollador
// NAP (Name + Address + Phone) incluido para SEO local de Google
</script>
