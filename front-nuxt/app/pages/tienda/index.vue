<template>
  <div class="bg-barber-black min-h-screen text-white relative pt-16">
    <!-- Barra de carga -->
    <Transition name="fade">
      <div v-if="isLoading" class="fixed top-0 left-0 w-full h-[2px] z-[100] overflow-hidden">
        <div class="h-full animate-progress-bar dept-bg shadow-[0_0_10px_var(--dept-glow)] transition-colors duration-500">
        </div>
      </div>
    </Transition>

    <!-- Header fijo -->
    <div class="sticky top-16 z-30 bg-barber-black/80 backdrop-blur-md border-b border-white/10">
      <div class="max-w-6xl mx-auto px-4 py-4 flex items-center justify-between gap-2">
        <NuxtLink to="/" class="flex-shrink-0 flex items-center gap-1.5 text-gray-400 hover:dept-text transition-colors duration-300">
          <fa-icon :icon="['fas', 'arrow-left']" class="text-xs" />
          <span class="text-xs font-semibold">{{ t('nav.home') }}</span>
        </NuxtLink>
        <h1 class="text-sm sm:text-lg font-bold tracking-tight sm:tracking-widest uppercase text-white truncate text-center flex-1 transition-colors duration-500">
          <span class="dept-text">Personal</span>{{ t('tienda.title').replace('Personal', '') }}
        </h1>
        <div class="w-10 sm:w-16 flex-shrink-0"></div>
      </div>
    </div>

    <div class="max-w-6xl mx-auto px-4 pb-20 pt-8">

      <!-- Switch de Departamento -->
      <div class="flex justify-center mt-2 mb-8 fade-in">
        <div class="inline-flex bg-zinc-900 rounded-full p-1 border border-zinc-800 shadow-[inset_0_2px_4px_rgba(0,0,0,0.6)]">
          <button @click="setDepartment('men'); activeFilter = 'all'"
            :class="['px-4 sm:px-6 py-2 rounded-full font-black tracking-widest text-[10px] sm:text-xs uppercase transition-all duration-300 flex items-center gap-2',
              activeDepartment === 'men' ? 'bg-[#39FF14] text-black shadow-[0_0_15px_rgba(57,255,20,0.3)]' : 'text-zinc-500 hover:text-white']">
            <fa-icon :icon="['fas', 'cut']" /> <span class="hidden xs:inline">{{ t('tienda.men') }}</span><span class="xs:hidden">{{ t('tienda.menMobile') }}</span>
          </button>
          <button @click="setDepartment('merch'); activeFilter = 'all'"
            :class="['px-4 sm:px-6 py-2 rounded-full font-black tracking-widest text-[10px] sm:text-xs uppercase transition-all duration-300 flex items-center gap-2',
              activeDepartment === 'merch' ? 'bg-[#22d3ee] text-black shadow-[0_0_15px_rgba(34,211,238,0.3)]' : 'text-zinc-500 hover:text-white']">
            <fa-icon :icon="['fas', 'tshirt']" /> <span class="hidden xs:inline">{{ t('tienda.merch') }}</span><span class="xs:hidden">{{ t('tienda.merchMobile') }}</span>
          </button>
          <button @click="setDepartment('women'); activeFilter = 'all'"
            :class="['px-4 sm:px-6 py-2 rounded-full font-black tracking-widest text-[10px] sm:text-xs uppercase transition-all duration-300 flex items-center gap-2',
              activeDepartment === 'women' ? 'bg-[#ec4899] text-white shadow-[0_0_15px_rgba(236,72,153,0.3)]' : 'text-zinc-500 hover:text-white']">
            <fa-icon :icon="['fas', 'spa']" /> <span class="hidden xs:inline">{{ t('tienda.women') }}</span><span class="xs:hidden">{{ t('tienda.womenMobile') }}</span>
          </button>
        </div>
      </div>

      <!-- Filtros de categoría -->
      <div class="flex flex-wrap gap-3 mb-10 justify-center">
        <button v-for="f in filters" :key="f.id" @click="activeFilter = f.id"
          :class="['px-5 py-2 rounded-full text-sm font-bold tracking-wide border transition-all duration-300',
            activeFilter === f.id
              ? 'dept-active-filter font-black shadow-[0_0_12px_var(--dept-glow)]'
              : 'glass border-white/20 text-gray-300 hover:dept-border hover:text-white']">
          {{ f.label }}
        </button>
      </div>

      <!-- Contador -->
      <div class="mb-6 text-center">
        <p class="text-gray-500 text-sm">
          {{ t('tienda.showing') }} <span class="font-bold dept-text transition-colors duration-300">{{ filteredProducts.length }}</span> {{ t('tienda.products') }}
          <span v-if="activeFilter !== 'all'"> {{ t('tienda.in') }} <span class="text-white">{{ activeFilterLabel }}</span></span>
        </p>
      </div>

      <!-- Loading -->
      <div v-if="isLoading" class="flex flex-col items-center justify-center py-24">
        <div class="animate-spin rounded-full h-12 w-12 border-b-2 dept-border mb-4 transition-colors duration-300"></div>
        <p class="text-gray-400 font-medium">{{ t('tienda.loading') }}</p>
      </div>

      <!-- Grid de productos -->
      <TransitionGroup v-else name="products-grid" tag="div"
        class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 md:gap-6 transition-opacity duration-500"
        :class="{'opacity-40 pointer-events-none': isLoading}">
        <div v-for="(product, index) in filteredProducts" :key="product.id"
          :style="isFirstVisit ? { '--i': index } : {}"
          class="group flex flex-col bg-white/5 border border-white/10 rounded-2xl overflow-hidden transition-premium dept-hover-card">

          <!-- Imagen -->
          <div class="aspect-square overflow-hidden bg-white relative cursor-pointer flex items-center justify-center p-3" @click="goToDetail(product)">
            <img
              :src="optimizeImage(product.images && product.images.length > 0 ? product.images[0] : product.image, 400)"
              :srcset="optimizeSrcSet(product.images && product.images.length > 0 ? product.images[0] : product.image, [160, 320, 400])"
              sizes="(max-width: 640px) calc(50vw - 1.25rem), (max-width: 1024px) calc(33vw - 1.5rem), 220px"
              :alt="product.name"
              class="w-full h-full object-contain transition-image group-hover:scale-[1.03]"
              :class="{'grayscale opacity-50': product.stock <= 0}"
              width="400"
              height="400"
              :loading="index < 2 ? 'eager' : 'lazy'"
              :fetchpriority="index < 2 ? 'high' : 'auto'"
              decoding="async"
            />
            <div v-if="product.stock <= 0" class="absolute inset-0 flex items-center justify-center bg-black/40 backdrop-blur-[2px]">
              <span class="bg-red-600 text-white text-[10px] font-black px-3 py-1 rounded-full uppercase tracking-tighter shadow-lg">{{ t('tienda.soldOut') }}</span>
            </div>
            <div v-else-if="product.stock <= 3" class="absolute top-2 right-2 z-10">
              <span class="bg-yellow-400 text-black text-[9px] font-black px-2 py-0.5 rounded-full uppercase tracking-tighter shadow-lg animate-pulse">{{ t('tienda.lastItems').replace('{n}', String(product.stock)) }}</span>
            </div>
          </div>

          <!-- Info -->
          <div class="p-4 flex flex-col flex-grow justify-between">
            <div class="cursor-pointer" @click="goToDetail(product)">
              <span class="text-[10px] text-gray-500 uppercase tracking-widest">{{ product.brand }}</span>
              <h3 class="text-sm font-bold text-white transition-colors duration-300 leading-tight mt-0.5 group-hover:dept-text">
                {{ product.name }}
              </h3>
              <p class="text-gray-500 text-xs mt-1.5 leading-relaxed line-clamp-2 italic">{{ product.description }}</p>
            </div>
            <div class="flex items-center justify-between mt-4">
              <span class="font-bold text-sm transition-colors duration-300 dept-text">
                {{ formatPrice(product.price) }}
              </span>
              <button v-if="product.stock > 0" @click.stop="quickAddToCart(product)"
                :disabled="isStockFull(product)"
                :class="['w-8 h-8 rounded-full glass flex items-center justify-center transition-all duration-300 text-sm text-white hover:text-black',
                  isStockFull(product) ? 'opacity-20 cursor-not-allowed border-transparent'
                    : (justAdded === product.id
                      ? 'dept-bg text-black'
                      : 'dept-hover-btn')]"
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
  </div>
</template>

<script setup lang="ts">
import { useLanguage } from '~/composables/useLanguage'
import { useDepartment } from '~/composables/useDepartment'

// Redirección permanente 301 para unificar la tienda en '/'
definePageMeta({
  middleware: [
    function (to) {
      return navigateTo({ path: '/', query: to.query }, { redirectCode: 301, external: false })
    }
  ]
})

const route = useRoute()
const router = useRouter()
const { addToCart, isStockFull } = useCart()
const { products, categories, isLoading, error: errorMessage, fetchCatalog } = useCatalog()
const { t, lang } = useLanguage()
const { activeDepartment, setDepartment } = useDepartment()

useSeoMeta({
  title: t('tienda.seoTitle'),
  ogTitle: t('tienda.seoTitle'),
  description: t('tienda.seoDesc'),
  ogDescription: t('tienda.seoDesc'),
})

const activeFilter = ref('all')
const justAdded = ref<string | number | null>(null)
const isFirstVisit = ref(true)


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
// Flag para evitar reseteo de filtro cuando syncFilter cambia activeDepartment desde URL
let syncingFromRoute = false

function syncFilter() {
  const cat = route.query.cat as string
  const dept = route.query.dept as string

  // Flag para evitar que el watch de activeDepartment resetee el filtro durante sync
  syncingFromRoute = true

  // Primero aplicar dept si viene en la query
  if (dept && ['men', 'women', 'merch'].includes(dept)) {
    setDepartment(dept as 'men' | 'merch' | 'women')
  }

  // Luego refinar según la categoría (tiene prioridad sobre dept)
  const categoryObj = categories.value.find(c => c.id === cat)
  if (categoryObj) {
    if (categoryObj.department === 'unisex' || categoryObj.style === 'premium') {
      setDepartment('merch')
    } else if (categoryObj.department && ['men', 'women'].includes(categoryObj.department)) {
      setDepartment(categoryObj.department as 'men' | 'women')
    }
  }

  // Activar el filtro si la categoría existe en los filtros del departamento activo
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

// Solo resetear el filtro cuando el usuario cambia el departamento MANUALMENTE (no vía URL)
watch(activeDepartment, () => {
  if (!syncingFromRoute) {
    activeFilter.value = 'all'
  }
})


// SEO dinámico por categoría
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

  let list = products.value.filter(p => p.is_active !== false && p.category && activeDeptCats.includes(p.category))
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
// Cargar catálogo con SSR al final para no romper contexto
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

/* Transición ultrasuave y lenta para el zoom de la imagen */
.transition-image {
  transition: transform 2.5s ease-in-out !important;
}
</style>
