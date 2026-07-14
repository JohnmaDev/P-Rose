<template>
  <div class="bg-barber-black min-h-screen text-white relative">

    <!-- ─── Barra de carga top ─── -->
    <Transition name="fade">
      <div v-if="isLoading" class="fixed top-0 left-0 w-full h-[2px] z-[100] overflow-hidden">
        <div class="h-full animate-progress-bar transition-colors duration-500"
          :class="{
            'bg-neon-green shadow-[0_0_10px_#39FF14]': activeDepartment === 'men',
            'bg-cyan-400 shadow-[0_0_10px_#22d3ee]': activeDepartment === 'merch',
            'bg-pink-500 shadow-[0_0_10px_#ec4899]': activeDepartment === 'women'
          }">
        </div>
      </div>
    </Transition>

    <!-- ─── STORE HERO — fondo dramático + identidad de tienda ─── -->
    <div class="relative w-full min-h-[88svh] flex flex-col justify-center overflow-hidden">

      <!-- Imagen de fondo con overlay multicapa -->
      <div class="absolute inset-0 z-0">
        <picture>
          <source media="(orientation: landscape)" srcset="/bg_horizontal.webp">
          <img
            src="/bg_vertical.webp"
            alt="PersonalBarber Store"
            class="w-full h-full object-cover object-top"
            style="filter: brightness(0.28) saturate(0.85)"
            fetchpriority="high"
          />
        </picture>
        <!-- Gradiente profundo hacia abajo -->
        <div class="absolute inset-0 bg-gradient-to-b from-barber-black/40 via-transparent to-barber-black"></div>
        <!-- Vignette lateral -->
        <div class="absolute inset-0 bg-gradient-to-r from-barber-black/70 via-transparent to-barber-black/50"></div>
      </div>

      <!-- Glow ambiental neón -->
      <div class="absolute top-1/3 left-1/4 w-[45vw] h-[45vw] max-w-[500px] bg-neon-green/7 rounded-full blur-[130px] pointer-events-none z-0"></div>

      <!-- Contenido hero -->
      <div class="relative z-10 max-w-7xl mx-auto px-6 pt-28 pb-16 w-full flex flex-col lg:flex-row items-center lg:items-end gap-10 lg:gap-20">

        <!-- Columna izquierda: tipografía masiva -->
        <div class="flex-1 flex flex-col gap-5">

          <!-- Badge live -->
          <div class="inline-flex items-center gap-2 self-start px-4 py-1.5 rounded-full border border-neon-green/30 bg-neon-green/5 backdrop-blur-sm">
            <span class="w-2 h-2 rounded-full bg-neon-green animate-pulse shadow-[0_0_8px_rgba(57,255,20,0.9)]"></span>
            <span class="text-neon-green text-[10px] font-black tracking-[0.22em] uppercase">Medellín · Premium Store</span>
          </div>

          <!-- Título masivo estilo tienda -->
          <h1 class="text-[3.2rem] sm:text-[5rem] lg:text-[7.5rem] xl:text-[9rem] font-black tracking-tighter italic leading-[0.88] text-shadow-premium">
            <span class="text-neon-green block drop-shadow-[0_0_25px_rgba(57,255,20,0.35)]">{{ t('store.heroTitle1') }}</span>
            <span class="text-white block pt-2">{{ t('store.heroTitle2') }}</span>
          </h1>

          <!-- Subtítulo -->
          <p class="text-gray-300 text-base sm:text-lg max-w-lg leading-relaxed font-medium">
            {{ t('store.heroSub') }}
          </p>

          <!-- Trust badges -->
          <div class="flex flex-wrap gap-4 mt-1">
            <span v-for="badge in trustBadges" :key="badge"
              class="flex items-center gap-2 text-[11px] text-gray-300 font-semibold">
              <fa-icon :icon="['fas', 'circle-check']" class="text-neon-green" />
              {{ badge }}
            </span>
          </div>
        </div>

        <!-- Columna derecha: números de credibilidad (compactos, sin protagonismo) -->
        <div class="flex lg:flex-col gap-3 flex-shrink-0">
          <div class="flex flex-col items-center px-5 py-3 rounded-2xl bg-black/40 backdrop-blur-md border border-white/10">
            <span class="text-2xl font-black text-neon-green">+50</span>
            <span class="text-[9px] text-gray-500 uppercase tracking-widest font-bold mt-0.5">{{ t('hero.trustClients') }}</span>
          </div>
          <div class="flex flex-col items-center px-5 py-3 rounded-2xl bg-black/40 backdrop-blur-md border border-white/10">
            <span class="text-2xl font-black text-yellow-400">5.0★</span>
            <span class="text-[9px] text-gray-500 uppercase tracking-widest font-bold mt-0.5">Rating</span>
          </div>
        </div>
      </div>

      <!-- Scroll indicator -->
      <div class="absolute bottom-5 left-1/2 -translate-x-1/2 z-10 flex flex-col items-center gap-1.5 animate-bounce pointer-events-none">
        <span class="text-[9px] text-white/20 font-black tracking-[0.3em] uppercase">Explorar</span>
        <div class="w-[1px] h-7 bg-gradient-to-b from-white/20 to-transparent rounded-full"></div>
      </div>
    </div>

    <!-- ─── CONTENIDO TIENDA ─── -->
    <div class="max-w-7xl mx-auto px-4 sm:px-6 pb-20 pt-8">

      <!-- Switch de Departamento -->
      <div class="flex justify-center mt-2 mb-8">
        <div class="inline-flex bg-zinc-900 rounded-full p-1 border border-zinc-800 shadow-[inset_0_2px_4px_rgba(0,0,0,0.6)]">
          <button @click="activeDepartment = 'men'; activeFilter = 'all'"
            :class="['px-4 sm:px-6 py-2 rounded-full font-black tracking-widest text-[10px] sm:text-xs uppercase transition-all duration-300 flex items-center gap-2',
              activeDepartment === 'men' ? 'bg-neon-green text-black shadow-[0_0_15px_rgba(57,255,20,0.3)]' : 'text-zinc-500 hover:text-white']">
            <fa-icon :icon="['fas', 'cut']" /> <span class="hidden xs:inline">{{ t('tienda.men') }}</span><span class="xs:hidden">{{ t('tienda.menMobile') }}</span>
          </button>
          <button @click="activeDepartment = 'merch'; activeFilter = 'all'"
            :class="['px-4 sm:px-6 py-2 rounded-full font-black tracking-widest text-[10px] sm:text-xs uppercase transition-all duration-300 flex items-center gap-2',
              activeDepartment === 'merch' ? 'bg-cyan-400 text-black shadow-[0_0_15px_rgba(34,211,238,0.3)]' : 'text-zinc-500 hover:text-white']">
            <fa-icon :icon="['fas', 'tshirt']" /> <span class="hidden xs:inline">{{ t('tienda.merch') }}</span><span class="xs:hidden">{{ t('tienda.merchMobile') }}</span>
          </button>
          <button @click="activeDepartment = 'women'; activeFilter = 'all'"
            :class="['px-4 sm:px-6 py-2 rounded-full font-black tracking-widest text-[10px] sm:text-xs uppercase transition-all duration-300 flex items-center gap-2',
              activeDepartment === 'women' ? 'bg-pink-500 text-white shadow-[0_0_15px_rgba(236,72,153,0.3)]' : 'text-zinc-500 hover:text-white']">
            <fa-icon :icon="['fas', 'spa']" /> <span class="hidden xs:inline">{{ t('tienda.women') }}</span><span class="xs:hidden">{{ t('tienda.womenMobile') }}</span>
          </button>
        </div>
      </div>

      <!-- Filtros de categoría -->
      <div class="flex flex-wrap gap-3 mb-8 justify-center">
        <button v-for="f in filters" :key="f.id" @click="activeFilter = f.id"
          :class="['px-5 py-2 rounded-full text-sm font-bold tracking-wide border transition-all duration-300',
            activeFilter === f.id
              ? (activeDepartment === 'men' ? 'bg-neon-green text-black border-neon-green' : (activeDepartment === 'merch' ? 'bg-cyan-400 text-black border-cyan-400' : 'bg-pink-500 text-white border-pink-500'))
              : (activeDepartment === 'men' ? 'glass border-white/20 text-gray-300 hover:border-neon-green/50 hover:text-white' : (activeDepartment === 'merch' ? 'glass border-white/20 text-gray-300 hover:border-cyan-400/50 hover:text-white' : 'glass border-white/20 text-gray-300 hover:border-pink-500/50 hover:text-white'))]">
          {{ f.label }}
        </button>
      </div>

      <!-- Contador -->
      <div class="mb-6 text-center">
        <p class="text-gray-500 text-sm">
          {{ t('tienda.showing') }} <span class="font-bold transition-colors duration-300" :class="{'text-neon-green': activeDepartment === 'men','text-cyan-400': activeDepartment === 'merch','text-pink-500': activeDepartment === 'women'}">{{ filteredProducts.length }}</span> {{ t('tienda.products') }}
          <span v-if="activeFilter !== 'all'"> {{ t('tienda.in') }} <span class="text-white">{{ activeFilterLabel }}</span></span>
        </p>
      </div>

      <!-- Loading -->
      <div v-if="isLoading" class="flex flex-col items-center justify-center py-24">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 mb-4 transition-colors duration-300" :class="{'border-neon-green': activeDepartment === 'men','border-cyan-400': activeDepartment === 'merch','border-pink-500': activeDepartment === 'women'}"></div>
        <p class="text-gray-400 font-medium">{{ t('tienda.loading') }}</p>
      </div>

      <!-- Grid de productos -->
      <TransitionGroup v-else name="products-grid" tag="div"
        class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-4 md:gap-5"
        :class="{'opacity-40 pointer-events-none': isLoading}">
        <div v-for="(product, index) in filteredProducts" :key="product.id"
          :style="isFirstVisit ? { '--i': index } : {}"
          :class="['group flex flex-col bg-white/5 border border-white/10 rounded-2xl overflow-hidden transition-premium',
            activeDepartment === 'men' ? 'hover:border-neon-green/50' : (activeDepartment === 'merch' ? 'hover:border-cyan-400/50' : 'hover:border-pink-500/50')]">

          <!-- Imagen -->
          <div class="aspect-square overflow-hidden bg-white/5 relative cursor-pointer" @click="goToDetail(product)">
            <img
              :src="optimizeImage(product.images && product.images.length > 0 ? product.images[0] : product.image, 400)"
              :alt="product.name"
              class="w-full h-full object-cover transition-premium group-hover:scale-110"
              :class="{'grayscale opacity-50': product.stock <= 0}"
              loading="lazy"
            />
            <div v-if="product.stock <= 0" class="absolute inset-0 flex items-center justify-center bg-black/40 backdrop-blur-[2px]">
              <span class="bg-red-500 text-white text-[10px] font-black px-3 py-1 rounded-full uppercase tracking-tighter shadow-lg">{{ t('tienda.soldOut') }}</span>
            </div>
            <div v-else-if="product.stock <= 3" class="absolute top-2 right-2">
              <span class="bg-yellow-400 text-black text-[9px] font-black px-2 py-0.5 rounded-full uppercase tracking-tighter shadow-lg animate-pulse">{{ t('tienda.lastItems').replace('{n}', String(product.stock)) }}</span>
            </div>
          </div>

          <!-- Info -->
          <div class="p-4 flex flex-col flex-grow justify-between">
            <div class="cursor-pointer" @click="goToDetail(product)">
              <span class="text-[10px] text-gray-500 uppercase tracking-widest">{{ product.brand }}</span>
              <h3 class="text-sm font-bold text-white transition-colors duration-300 leading-tight mt-0.5"
                :class="{'group-hover:text-neon-green': activeDepartment === 'men','group-hover:text-cyan-400': activeDepartment === 'merch','group-hover:text-pink-500': activeDepartment === 'women'}">
                {{ product.name }}
              </h3>
            </div>
            <div class="flex items-center justify-between mt-4">
              <span class="font-bold text-sm transition-colors duration-300"
                :class="{'text-neon-green': activeDepartment === 'men','text-cyan-400': activeDepartment === 'merch','text-pink-500': activeDepartment === 'women'}">
                {{ formatPrice(product.price) }}
              </span>
              <button v-if="product.stock > 0" @click.stop="quickAddToCart(product)"
                :disabled="isStockFull(product)"
                :class="['w-8 h-8 rounded-full glass flex items-center justify-center transition-all duration-300 text-sm text-white hover:text-black',
                  isStockFull(product) ? 'opacity-20 cursor-not-allowed border-transparent'
                    : (justAdded === product.id
                      ? (activeDepartment === 'men' ? 'bg-neon-green text-black' : (activeDepartment === 'merch' ? 'bg-cyan-400 text-black' : 'bg-pink-500 text-white'))
                      : (activeDepartment === 'men' ? 'hover:bg-neon-green' : (activeDepartment === 'merch' ? 'hover:bg-cyan-400' : 'hover:bg-pink-500')))]"
                :aria-label="isStockFull(product) ? t('tienda.maxStock') : t('tienda.addToCart')">
                <fa-icon :icon="['fas', isStockFull(product) ? 'lock' : (justAdded === product.id ? 'check' : 'plus')]" class="text-[10px]" />
              </button>
              <button v-else disabled class="w-8 h-8 rounded-full bg-white/5 border border-white/10 flex items-center justify-center text-zinc-700 cursor-not-allowed">
                <fa-icon :icon="['fas', 'times']" />
              </button>
            </div>
          </div>
        </div>
      </TransitionGroup>

      <!-- Error State -->
      <div v-if="errorMessage && !isLoading" class="flex flex-col items-center justify-center py-20 text-center px-4">
        <div class="bg-red-500/10 border border-red-500/20 p-6 rounded-3xl max-w-md">
          <fa-icon :icon="['fas', 'exclamation-triangle']" class="text-red-500 text-3xl mb-4" />
          <h3 class="text-white font-bold mb-2">{{ t('tienda.connError') }}</h3>
          <p class="text-zinc-400 text-sm mb-6">{{ errorMessage }}</p>
          <button @click="fetchData" class="px-6 py-2 bg-zinc-800 hover:bg-zinc-700 text-white rounded-xl text-xs font-bold uppercase transition-all">{{ t('tienda.retry') }}</button>
        </div>
      </div>

      <!-- Empty State -->
      <div v-if="!isLoading && filteredProducts.length === 0 && !errorMessage" class="text-center py-24">
        <fa-icon :icon="['fas', 'box-open']" class="text-4xl text-gray-600 mb-4" />
        <p class="text-gray-500">{{ t('tienda.emptySub') }}</p>
      </div>
    </div>

    <!-- Footer -->
    <div class="w-full max-w-7xl mx-auto px-4 sm:px-6">
      <AppFooter class="mt-10" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useLanguage } from '~/composables/useLanguage'

const route = useRoute()
const router = useRouter()
const { addToCart, isStockFull } = useCart()
const { products, categories, isLoading, error: errorMessage, fetchCatalog } = useCatalog()
const { t } = useLanguage()

// SEO — tienda como home
useSeoMeta({
  title: 'PersonalBarber — Tienda Online de Barbería | Medellín',
  ogTitle: 'PersonalBarber — Tienda Online de Barbería | Medellín',
  description: 'Compra productos profesionales de barbería, cuidado personal y moda en PersonalBarber. Envíos a Medellín y toda Colombia. También agenda tu cita con el barber a domicilio.',
  ogDescription: 'Tienda online de barbería premium en Medellín. Ceras, maquinas, cuidado de barba, skincare y más. Compra online con envío.',
  ogUrl: 'https://personalbarber.vip',
})

// JSON-LD — Store + BarberShop (dual schema)
useHead({
  script: [{
    type: 'application/ld+json',
    children: JSON.stringify({
      '@context': 'https://schema.org',
      '@type': ['Store', 'BarberShop'],
      name: 'PersonalBarber',
      description: 'Tienda online de productos de barbería premium y servicio de barbero a domicilio en Medellín.',
      url: 'https://personalbarber.vip',
      telephone: '+573045840264',
      image: 'https://personalbarber.vip/og-image.webp',
      priceRange: '$$',
      currenciesAccepted: 'COP',
      paymentAccepted: 'Cash, Transferencia, Nequi, PSE',
      address: {
        '@type': 'PostalAddress',
        addressLocality: 'Medellín',
        addressRegion: 'Antioquia',
        addressCountry: 'CO',
      },
      sameAs: [
        'https://www.instagram.com/pipehp_/',
        'https://www.tiktok.com/@pipehpbarber',
      ],
    })
  }]
})

const activeDepartment = ref<'men' | 'merch' | 'women'>('men')
const activeFilter = ref('all')
const justAdded = ref<string | number | null>(null)
const isFirstVisit = ref(true)

const trustBadges = computed(() => [
  t('tienda.badgeOriginal'),
  t('tienda.badgeShipping'),
  t('tienda.badgeSupport'),
])


function getCategoryLabel(catId: string) {
  const cat = categories.value.find(c => c.id === catId)
  const label = cat ? cat.label : catId
  const key = `tienda.categorias.${catId}`
  const translated = t(key)
  return translated === key ? label : translated
}

const filters = computed(() => [
  { id: 'all', label: t('tienda.all') },
  ...categories.value
    .filter(c => {
      if (c.comingSoon) return false
      if (activeDepartment.value === 'merch') return c.department === 'unisex' || c.style === 'premium'
      return c.department === activeDepartment.value
    })
    .map(c => ({ id: c.id, label: getCategoryLabel(c.id) }))
])

let syncingFromRoute = false

function syncFilter() {
  const cat = route.query.cat as string
  const dept = route.query.dept as string
  syncingFromRoute = true
  if (dept && ['men', 'women', 'merch'].includes(dept)) {
    activeDepartment.value = dept as 'men' | 'merch' | 'women'
  }
  const categoryObj = categories.value.find(c => c.id === cat)
  if (categoryObj) {
    if (categoryObj.department === 'unisex' || categoryObj.style === 'premium') {
      activeDepartment.value = 'merch'
    } else if (categoryObj.department && ['men', 'women'].includes(categoryObj.department)) {
      activeDepartment.value = categoryObj.department as 'men' | 'women'
    }
  }
  activeFilter.value = (cat && filters.value.find(f => f.id === cat)) ? cat : 'all'
  nextTick(() => { syncingFromRoute = false })
}

function shuffleProducts() {
  if (products.value.length > 0) {
    const shuffled = [...products.value]
    for (let i = shuffled.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1))
      ;[shuffled[i], shuffled[j]] = [shuffled[j], shuffled[i]]
    }
    products.value = shuffled
  }
}

async function fetchData() {
  await fetchCatalog(true)
  shuffleProducts()
  syncFilter()
  setTimeout(() => { isFirstVisit.value = false }, 1000)
}

onMounted(() => {
  shuffleProducts()
  syncFilter()
  setTimeout(() => { isFirstVisit.value = false }, 1000)
})

watch(() => route.query.cat, syncFilter)
watch(() => route.query.dept, syncFilter)
watch(activeDepartment, () => {
  if (!syncingFromRoute) activeFilter.value = 'all'
})
watch(activeFilter, (newFilter) => {
  const label = filters.value.find(f => f.id === newFilter)?.label || 'Tienda'
  useSeoMeta({ title: `${label} | PersonalBarber Medellín` })
}, { immediate: false })

const filteredProducts = computed(() => {
  const activeDeptCats = categories.value
    .filter(c => {
      if (activeDepartment.value === 'merch') return c.department === 'unisex' || c.style === 'premium'
      return c.department === activeDepartment.value
    })
    .map(c => c.id)
  let list = products.value.filter(p => p.category && activeDeptCats.includes(p.category))
  if (activeFilter.value !== 'all') {
    list = list.filter(p => p.category === activeFilter.value)
  }
  return list
})

const activeFilterLabel = computed(() => filters.value.find(f => f.id === activeFilter.value)?.label ?? '')

function goToDetail(product: { id: number; name: string }) {
  router.push({ name: 'tienda-producto-slug', params: { slug: generateProductSlug(product.id, product.name) } })
}

function quickAddToCart(product: Parameters<typeof addToCart>[0]) {
  if (isStockFull(product)) return
  const res = addToCart(product)
  if (res.success) {
    justAdded.value = product.id
    setTimeout(() => { justAdded.value = null }, 2000)
  }
}

if (import.meta.server) {
  await fetchCatalog()
} else {
  fetchCatalog()
}
</script>

<style scoped>
.products-grid-move,
.products-grid-enter-active,
.products-grid-leave-active {
  transition: opacity 0.5s ease-out, transform 0.5s cubic-bezier(0.2, 0.8, 0.2, 1);
  will-change: transform, opacity;
}
.products-grid-enter-active { transition-delay: calc(var(--i, 0) * 0.04s); z-index: 10; }
.products-grid-enter-from { opacity: 0; transform: translateY(20px) scale(0.95); }
.products-grid-leave-to { opacity: 0; transform: scale(0.9); }
.products-grid-leave-active { position: absolute; width: calc(50% - 1rem); z-index: 0; pointer-events: none; }
@media (min-width: 768px) { .products-grid-leave-active { width: calc(33.333% - 1.5rem); } }
@media (min-width: 1024px) { .products-grid-leave-active { width: calc(25% - 2rem); } }
</style>
