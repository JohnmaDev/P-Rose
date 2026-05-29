<template>
  <div class="bg-barber-black min-h-screen text-white flex flex-col items-center">
    <!-- Hero section -->
    <HeroSection @reserve="goToReserva" class="animate-on-scroll opacity-0 translate-y-8 transition-all duration-1000 ease-out" />

    <!-- Métricas — prueba social cuantitativa inmediata -->
    <MetricsSection class="w-full" />

    <!-- Diferenciador: Barber a domicilio -->
    <ServiceFeatureSection class="w-full" />

    <div class="w-full max-w-6xl px-4 flex flex-col items-center pb-20 overflow-hidden">
      <ShopCategories class="animate-on-scroll w-full opacity-0 translate-y-12 transition-all duration-1000 ease-out" />
      <MasonryGallery class="animate-on-scroll w-full mt-8 opacity-0 translate-y-12 transition-all duration-1000 ease-out" />
    </div>

    <!-- Reviews — cierre de conversión -->
    <ReviewsSection class="w-full" />

    <!-- Footer -->
    <div class="w-full max-w-6xl px-4">
      <AppFooter class="mt-14 animate-on-scroll opacity-0 translate-y-8 transition-all duration-1000 ease-out" />
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

// Precargar catálogo desde server para que ShopCategories tenga datos
const { fetchCatalog } = useCatalog()
await fetchCatalog()

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
</script>
