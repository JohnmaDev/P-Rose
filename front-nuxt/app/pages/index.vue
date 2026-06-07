<template>
  <div class="bg-barber-black min-h-screen text-white flex flex-col items-center relative overflow-hidden">
    <!-- Ambient Glow Orbs that visually connect all sections into a single canvas -->
    <div class="absolute top-[18%] left-[-15%] w-[50vw] h-[50vw] max-w-[600px] bg-neon-green/6 rounded-full blur-[140px] pointer-events-none z-0"></div>
    <div class="absolute top-[45%] right-[-15%] w-[55vw] h-[55vw] max-w-[700px] bg-neon-green/4 rounded-full blur-[160px] pointer-events-none z-0"></div>
    <div class="absolute top-[75%] left-[-10%] w-[45vw] h-[45vw] max-w-[550px] bg-neon-green/5 rounded-full blur-[130px] pointer-events-none z-0"></div>

    <!-- Hero section -->
    <HeroSection @reserve="goToReserva" class="relative z-10" />

    <!-- Métricas — prueba social cuantitativa inmediata -->
    <MetricsSection class="animate-on-scroll w-full opacity-0 translate-y-12 transition-all duration-[1200ms] ease-out relative z-10" />

    <!-- Diferenciador: Barber a domicilio -->
    <ServiceFeatureSection class="animate-on-scroll w-full opacity-0 translate-y-12 transition-all duration-[1200ms] ease-out relative z-10" />

    <div class="w-full max-w-6xl px-4 flex flex-col items-center pb-20 overflow-hidden relative z-10">
      <ShopCategories class="animate-on-scroll w-full opacity-0 translate-y-12 transition-all duration-[1200ms] ease-out" />
      <MasonryGallery class="animate-on-scroll w-full mt-8 opacity-0 translate-y-12 transition-all duration-[1200ms] ease-out" />
    </div>

    <!-- Reviews — cierre de conversión -->
    <ReviewsSection class="animate-on-scroll w-full opacity-0 translate-y-12 transition-all duration-[1200ms] ease-out relative z-10" />

    <!-- Footer -->
    <div class="w-full max-w-6xl px-4 relative z-10">
      <AppFooter class="mt-14 animate-on-scroll opacity-0 translate-y-8 transition-all duration-[1200ms] ease-out" />
    </div>
  </div>
</template>

<script setup lang="ts">
// SEO — renderizado en servidor, Google lo indexa directamente
useSeoMeta({
  title: 'PersonalBarber — Barbería Premium en Medellín | Reserva tu Cita Online',
  ogTitle: 'PersonalBarber — Barbería Premium en Medellín',
  description: 'Barbería premium a domicilio en Medellín. Cortes exclusivos, barba profesional y productos especializados. Más de 10 años de experiencia. Reserva tu cita online al instante.',
  ogDescription: 'Barbería premium a domicilio en Medellín. Cortes exclusivos, barba profesional y +100 clientes satisfechos. Reserva online en segundos.',
  ogUrl: 'https://personalbarber.vip',
})

// JSON-LD BarberShop — para Google Knowledge Panel y Rich Results
useHead({
  script: [{
    type: 'application/ld+json',
    children: JSON.stringify({
      '@context': 'https://schema.org',
      '@type': 'BarberShop',
      name: 'PersonalBarber',
      description: 'Barbería premium a domicilio en Medellín. Cortes exclusivos, barba profesional y productos especializados. Más de 10 años de experiencia.',
      url: 'https://personalbarber.vip',
      telephone: '+573045840264',
      image: 'https://personalbarber.vip/og-image.webp',
      priceRange: '$$',
      currenciesAccepted: 'COP',
      paymentAccepted: 'Cash, Transferencia, Nequi',
      address: {
        '@type': 'PostalAddress',
        streetAddress: 'La 4 Sur',
        addressLocality: 'Medellín',
        addressRegion: 'Antioquia',
        addressCountry: 'CO',
      },
      geo: {
        '@type': 'GeoCoordinates',
        latitude: 6.2518,
        longitude: -75.5636,
      },
      sameAs: [
        'https://www.instagram.com/pipehp_/',
        'https://www.tiktok.com/@pipehpbarber',
      ],
      openingHoursSpecification: [
        {
          '@type': 'OpeningHoursSpecification',
          dayOfWeek: ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday'],
          opens: '07:00',
          closes: '21:00',
        },
        {
          '@type': 'OpeningHoursSpecification',
          dayOfWeek: ['Saturday'],
          opens: '08:00',
          closes: '18:00',
        },
      ],
      hasOfferCatalog: {
        '@type': 'OfferCatalog',
        name: 'Servicios de Barbería',
        itemListElement: [
          { '@type': 'Offer', itemOffered: { '@type': 'Service', name: 'Corte de Cabello a Domicilio' } },
          { '@type': 'Offer', itemOffered: { '@type': 'Service', name: 'Arreglo de Barba' } },
          { '@type': 'Offer', itemOffered: { '@type': 'Service', name: 'Corte + Barba Combo' } },
          { '@type': 'Offer', itemOffered: { '@type': 'Service', name: 'Cejas y Limpieza Facial' } },
        ]
      }
    })
  }]
})

const router = useRouter()
const { fetchCatalog } = useCatalog()

function goToReserva() {
  router.push({ name: 'agendar' })
}

// Scroll reveal animations — solo en cliente
onMounted(() => {
  const observer = new IntersectionObserver((entries) => {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        entry.target.classList.remove('opacity-0', 'translate-y-8', 'translate-y-12')
        entry.target.classList.add('opacity-100', 'translate-y-0')
        observer.unobserve(entry.target)
      }
    })
  }, { threshold: 0.15, rootMargin: '0px 0px -50px 0px' })

  document.querySelectorAll('.animate-on-scroll').forEach(el => observer.observe(el))

  onUnmounted(() => observer.disconnect())
})

// Precargar catálogo al final para no romper el contexto de hooks síncronos
if (import.meta.server) {
  await fetchCatalog()
} else {
  fetchCatalog()
}
</script>
