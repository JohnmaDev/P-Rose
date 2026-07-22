<template>
  <div class="bg-[#0D0D0D] min-h-screen text-white relative pt-16 font-sans">
    <!-- Barra de carga dinámica -->
    <Transition name="fade">
      <div v-if="isLoading" class="fixed top-0 left-0 w-full h-[2px] z-[100] overflow-hidden">
        <div class="h-full animate-progress-bar bg-[#00FF00] shadow-[0_0_10px_#00FF00]"></div>
      </div>
    </Transition>

    <div v-if="isLoading && !product" class="flex flex-col items-center justify-center min-h-[70vh]">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-[#00FF00] mb-4"></div>
      <p class="text-[#A1A1AA] font-medium tracking-widest text-xs uppercase">{{ t('tienda.loading') }}</p>
    </div>

    <div v-else-if="product" class="max-w-[1400px] mx-auto px-4 md:px-8 py-6 pb-40 lg:pb-20 transition-opacity duration-500" :class="{'opacity-40 pointer-events-none': isLoading}">
      
      <!-- 1. BREADCRUMB (minimalista) -->
      <nav class="flex items-center gap-2 text-[10px] md:text-[11px] text-[#555] uppercase tracking-[0.18em] font-bold mb-10">
        <NuxtLink to="/tienda" class="hover:text-[#00FF00] transition-colors duration-200">Tienda</NuxtLink>
        <span class="text-[#333]">›</span>
        <NuxtLink :to="`/tienda?cat=${product.category}`" class="hover:text-[#00FF00] transition-colors duration-200">{{ getCategoryLabel(product.category) }}</NuxtLink>
        <span class="text-[#333]">›</span>
        <span class="text-[#A1A1AA] truncate max-w-[200px] sm:max-w-[400px]">{{ product.name }}</span>
      </nav>

      <!-- Main Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8 lg:gap-12 items-start">
        
        <!-- 2. ÁREA DE MEDIOS (Izquierda) -->
        <div class="flex flex-col gap-3 md:sticky md:top-24">
          <div class="relative w-full aspect-square bg-[#161616] rounded-3xl border border-[#262626] overflow-hidden group shadow-[0_4px_40px_rgba(0,0,0,0.5)]">
            <div ref="carouselRef" @scroll="handleScroll" class="w-full h-full overflow-x-auto overflow-y-hidden snap-x snap-mandatory scroll-smooth hide-scrollbar flex">
              <div v-for="(img, idx) in (product.images?.length ? product.images : ['/hero_barber.webp'])" :key="idx" class="w-full h-full flex-shrink-0 snap-center bg-white">
                <img :src="optimizeImage(img, 800)" :alt="`${product.name} ${idx + 1}`" class="w-full h-full object-contain transition-transform duration-700 group-hover:scale-105" loading="lazy" />
              </div>
            </div>
          </div>

          <!-- Carrusel de Miniaturas -->
          <div v-if="product.images && product.images.length > 1" class="flex gap-2 overflow-x-auto hide-scrollbar py-1">
            <button v-for="(img, idx) in product.images" :key="idx" @click="scrollToImage(idx)"
              class="relative w-16 h-16 rounded-xl overflow-hidden border-2 transition-all duration-300 shrink-0 bg-white"
              :class="activeIdx === idx ? 'border-[#00FF00] shadow-[0_0_12px_rgba(0,255,0,0.25)]' : 'border-[#262626] opacity-50 hover:opacity-100 hover:border-gray-500'">
              <img :src="optimizeImage(img, 150)" :alt="`Thumb ${idx + 1}`" class="w-full h-full object-contain p-1" />
            </button>
          </div>
        </div>

        <!-- 3. ÁREA DE COMPRA Y DETALLES (Derecha) -->
        <div class="flex flex-col gap-7">
          <!-- Títulos y Precio -->
          <div>
            <div class="flex items-center gap-3 mb-3">
              <span class="text-[#00FF00] text-[10px] font-black tracking-widest uppercase bg-[#00FF00]/10 px-2 py-1 rounded border border-[#00FF00]/20">{{ product.brand || 'Premium' }}</span>
              <span class="text-[#A1A1AA] text-[10px] font-bold tracking-widest uppercase">{{ getCategoryLabel(product.category) }}</span>
            </div>
            <h1 class="text-3xl lg:text-4xl font-black font-oswald tracking-tight leading-[1.1] mb-4 text-white drop-shadow-md">
              {{ product.name }}
            </h1>
            <div class="flex items-end gap-4 mt-2">
              <span class="text-4xl lg:text-5xl font-black text-[#00FF00] tracking-tighter filter drop-shadow-[0_0_10px_rgba(0,255,0,0.3)]">
                {{ formatPrice(product.price) }}
              </span>
              <span v-if="product.comparePrice" class="text-xl text-[#A1A1AA] line-through font-bold mb-1">
                {{ formatPrice(product.comparePrice) }}
              </span>
            </div>
          </div>

          <!-- SELECTOR DE VARIANTES (Capsule Pills — estilo Sephora) -->
          <div v-if="product.variants && product.variants.length > 0" class="space-y-3">
            <span class="text-[10px] font-bold uppercase tracking-[0.2em] text-[#A1A1AA]">Presentación</span>
            <div class="flex flex-wrap gap-2">
              <button v-for="(variant, vIdx) in product.variants" :key="vIdx"
                @click="selectedVariant = vIdx"
                class="px-5 py-2.5 rounded-full text-xs font-bold uppercase tracking-wider transition-all duration-300"
                :class="selectedVariant === vIdx
                  ? 'border border-[#00FF00] bg-[#00FF00]/10 text-[#00FF00] shadow-[0_0_12px_rgba(0,255,0,0.2)]'
                  : 'border border-[#262626] bg-transparent text-[#A1A1AA] hover:border-[#555] hover:text-white'">
                {{ variant }}
              </button>
            </div>
          </div>

          <!-- SELECTOR DE CANTIDAD + STOCK (inline, limpio) -->
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-4">
              <span class="text-[10px] font-bold uppercase tracking-[0.2em] text-[#A1A1AA]">Cantidad</span>
              <div class="flex items-center border border-[#262626] rounded-full">
                <button @click="qty = Math.max(1, qty - 1)" class="w-9 h-9 rounded-full hover:bg-white/5 flex items-center justify-center transition-all text-[#A1A1AA] hover:text-white">−</button>
                <span class="text-white font-black w-8 text-center text-sm tabular-nums">{{ qty }}</span>
                <button @click="qty < (product.stock - getProductQty(product.id)) ? qty++ : null"
                  :disabled="qty >= (product.stock - getProductQty(product.id))"
                  class="w-9 h-9 rounded-full hover:bg-white/5 flex items-center justify-center transition-all text-[#00FF00] disabled:opacity-30 disabled:cursor-not-allowed">+</button>
              </div>
            </div>
            <span v-if="product.stock > 5" class="text-[10px] text-[#A1A1AA] font-bold uppercase tracking-wider">En Stock</span>
            <span v-else-if="product.stock > 0" class="text-[10px] text-amber-500 font-bold uppercase tracking-wider animate-pulse">¡Últimas {{ product.stock }} uds!</span>
            <span v-else class="text-[10px] text-red-500 font-bold uppercase tracking-wider">Agotado</span>
          </div>

          <!-- CTAs -->
          <div class="flex flex-col gap-3">
            <button @click="handleBuyNow" :disabled="product.stock <= 0 || isStockFull(product)"
              class="w-full py-4 bg-[#00FF00] hover:bg-[#10B981] text-black font-black uppercase tracking-widest text-sm rounded-2xl transition-all duration-300 shadow-[0_0_20px_rgba(0,255,0,0.3)] hover:shadow-[0_0_30px_rgba(0,255,0,0.5)] hover:-translate-y-0.5 flex items-center justify-center gap-2 disabled:bg-zinc-800 disabled:text-zinc-600 disabled:shadow-none disabled:transform-none disabled:cursor-not-allowed">
              <fa-icon :icon="['fas', 'bolt']" class="text-lg" />
              {{ product.stock <= 0 ? 'Agotado' : 'Comprar Ahora' }}
            </button>
            <button @click="handleAddToCart" :disabled="product.stock <= 0 || isStockFull(product)"
              class="w-full py-4 bg-transparent border-2 border-[#262626] hover:border-[#00FF00] text-white hover:text-[#00FF00] font-bold uppercase tracking-widest text-xs rounded-2xl transition-all duration-300 flex items-center justify-center gap-2 disabled:opacity-30 disabled:cursor-not-allowed group">
              <fa-icon :icon="['fas', 'shopping-bag']" class="group-hover:animate-bounce" />
              Agregar al Carrito
            </button>
            <Transition name="fade">
              <div v-if="showConfirm" class="text-center text-[10px] font-bold text-[#00FF00] tracking-widest uppercase mt-1">
                ✓ Producto agregado exitosamente
              </div>
            </Transition>
          </div>

          <!-- WIDGET DE CONFIANZA Y LOGÍSTICA -->
          <div class="grid grid-cols-3 gap-3">
            <div class="bg-[#161616] border border-[#262626] rounded-2xl p-3 flex flex-col items-center justify-center text-center gap-2 group hover:border-[#00FF00]/50 transition-colors">
              <div class="w-8 h-8 rounded-full bg-[#0D0D0D] flex items-center justify-center border border-[#262626] group-hover:border-[#00FF00] transition-colors">
                <fa-icon :icon="['fas', 'shield-alt']" class="text-[#00FF00] text-xs" />
              </div>
              <span class="text-[9px] text-[#A1A1AA] uppercase tracking-wider font-bold leading-tight">Garantía<br>Real</span>
            </div>
            <div class="bg-[#161616] border border-[#262626] rounded-2xl p-3 flex flex-col items-center justify-center text-center gap-2 group hover:border-[#00FF00]/50 transition-colors">
              <div class="w-8 h-8 rounded-full bg-[#0D0D0D] flex items-center justify-center border border-[#262626] group-hover:border-[#00FF00] transition-colors">
                <fa-icon :icon="['fas', 'truck-fast']" class="text-[#00FF00] text-xs" />
              </div>
              <span class="text-[9px] text-[#A1A1AA] uppercase tracking-wider font-bold leading-tight">Despacho<br>Inmediato</span>
            </div>
            <div class="bg-[#161616] border border-[#262626] rounded-2xl p-3 flex flex-col items-center justify-center text-center gap-2 group hover:border-[#00FF00]/50 transition-colors">
              <div class="w-8 h-8 rounded-full bg-[#0D0D0D] flex items-center justify-center border border-[#262626] group-hover:border-[#00FF00] transition-colors">
                <fa-icon :icon="['fas', 'credit-card']" class="text-[#00FF00] text-xs" />
              </div>
              <span class="text-[9px] text-[#A1A1AA] uppercase tracking-wider font-bold leading-tight">Pago<br>Seguro</span>
            </div>
          </div>
        </div>
      </div>

      <!-- 4. SECCIÓN INFERIOR: Acordeones + Recomendados -->
      <div v-if="hasDescriptionContent || hasSpecsContent || recommendedProducts.length > 0" class="mt-16 lg:mt-24 space-y-10">

        <!-- Acordeones dinámicos (ancho completo) -->
        <div v-if="hasDescriptionContent || hasSpecsContent" class="space-y-4">
          <!-- Descripción & Beneficios -->
          <div v-if="hasDescriptionContent" class="bg-[#161616] border border-[#262626] rounded-3xl overflow-hidden">
            <button @click="expandedAccordion = expandedAccordion === 'desc' ? '' : 'desc'" class="w-full px-6 py-5 flex items-center justify-between hover:bg-white/5 transition-colors duration-300">
              <span class="font-oswald text-lg font-bold uppercase tracking-[0.08em] text-white">Descripción & Beneficios</span>
              <fa-icon :icon="['fas', expandedAccordion === 'desc' ? 'minus' : 'plus']" class="text-[#00FF00] text-sm" />
            </button>
            <div v-show="expandedAccordion === 'desc'" class="px-6 pb-6 pt-2 border-t border-[#262626]/50">
              <p v-if="product.description" class="text-[#A1A1AA] text-sm leading-[1.8]">{{ product.description }}</p>
              <ul v-if="product.benefits && product.benefits.length > 0" class="mt-5 space-y-3">
                <li v-for="(benefit, bIdx) in product.benefits" :key="bIdx" class="flex items-start gap-3">
                  <span class="text-[#00FF00] text-xs mt-0.5 shrink-0">✓</span>
                  <span class="text-[#A1A1AA] text-sm leading-relaxed">{{ benefit }}</span>
                </li>
              </ul>
            </div>
          </div>

          <!-- Modo de Uso / Especificaciones -->
          <div v-if="hasSpecsContent" class="bg-[#161616] border border-[#262626] rounded-3xl overflow-hidden">
            <button @click="expandedAccordion = expandedAccordion === 'specs' ? '' : 'specs'" class="w-full px-6 py-5 flex items-center justify-between hover:bg-white/5 transition-colors duration-300">
              <span class="font-oswald text-lg font-bold uppercase tracking-[0.08em] text-white">Modo de Uso / Especificaciones</span>
              <fa-icon :icon="['fas', expandedAccordion === 'specs' ? 'minus' : 'plus']" class="text-[#00FF00] text-sm" />
            </button>
            <div v-show="expandedAccordion === 'specs'" class="px-6 pb-6 pt-2 border-t border-[#262626]/50">
              <p v-if="product.usage" class="text-[#A1A1AA] text-sm leading-[1.8]">{{ product.usage }}</p>
              <p v-if="product.specs" class="text-[#A1A1AA] text-sm leading-[1.8]" :class="product.usage ? 'mt-4 pt-4 border-t border-[#262626]/30' : ''">{{ product.specs }}</p>
            </div>
          </div>
        </div>

        <!-- Recomendados (grid horizontal bonito) -->
        <div v-if="recommendedProducts.length > 0">
          <div class="flex items-center gap-4 mb-6">
            <div class="h-px flex-1 bg-[#262626]"></div>
            <h3 class="font-oswald text-xl font-bold uppercase tracking-[0.1em] text-white shrink-0">Completa tu Kit</h3>
            <div class="h-px flex-1 bg-[#262626]"></div>
          </div>
          <div class="grid grid-cols-2 sm:grid-cols-3 gap-4">
            <NuxtLink v-for="item in recommendedProducts" :key="item.id"
              :to="{ name: 'tienda-producto-slug', params: { slug: generateProductSlug(item.id, item.name) } }"
              class="group bg-[#161616] border border-[#262626] hover:border-[#00FF00]/40 rounded-2xl overflow-hidden transition-all duration-300 hover:-translate-y-1 hover:shadow-[0_8px_30px_rgba(0,255,0,0.08)]">
              <!-- Imagen -->
              <div class="aspect-square bg-white overflow-hidden">
                <img :src="optimizeImage(item.images?.[0] || item.image, 400)" :alt="item.name"
                  class="w-full h-full object-contain p-3 transition-transform duration-500 group-hover:scale-105" loading="lazy" />
              </div>
              <!-- Info -->
              <div class="p-4">
                <p class="text-[10px] text-[#00FF00] font-bold uppercase tracking-widest mb-1 truncate">{{ item.brand }}</p>
                <h4 class="text-xs font-bold text-white group-hover:text-[#00FF00] transition-colors line-clamp-2 leading-snug mb-2">{{ item.name }}</h4>
                <div class="flex items-center justify-between">
                  <span class="text-sm font-black text-[#00FF00]">{{ formatPrice(item.price) }}</span>
                  <button @click.prevent="addToCart(item, 1); showConfirm = true; setTimeout(() => { showConfirm = false }, 2500)"
                    class="w-7 h-7 rounded-full bg-[#0D0D0D] border border-[#262626] hover:border-[#00FF00] hover:bg-[#00FF00]/10 flex items-center justify-center transition-all">
                    <fa-icon :icon="['fas', 'plus']" class="text-[#A1A1AA] group-hover:text-[#00FF00] text-[10px]" />
                  </button>
                </div>
              </div>
            </NuxtLink>
          </div>
        </div>

      </div>
    </div>

    <!-- Producto no encontrado -->
    <div v-else class="flex flex-col items-center justify-center min-h-[70vh] gap-4">
      <fa-icon :icon="['fas', 'unlink']" class="text-5xl text-[#262626]" />
      <p class="text-[#A1A1AA] font-bold tracking-widest uppercase text-xs">{{ t('tienda.emptyTitle') }}</p>
      <NuxtLink to="/tienda" class="text-[#00FF00] hover:text-white border border-[#00FF00] hover:bg-[#00FF00]/10 rounded-xl px-6 py-2 transition-all text-xs font-black uppercase tracking-widest mt-2">← Volver al Catálogo</NuxtLink>
    </div>

    <!-- 5. COMPONENTE MÓVIL EXCLUSIVO (Sticky Bottom Bar) -->
    <Transition name="slide-up">
      <div v-if="product && !isLoading" class="lg:hidden fixed bottom-0 left-0 w-full z-50 bg-[#161616]/95 backdrop-blur-xl border-t border-[#262626] p-4 pb-safe flex items-center justify-between gap-4 shadow-[0_-10px_30px_rgba(0,0,0,0.8)]">
        <div class="flex items-center gap-3 flex-1 min-w-0">
          <div class="w-10 h-10 rounded-lg overflow-hidden bg-[#0D0D0D] border border-[#262626] shrink-0 hidden sm:block">
            <img :src="optimizeImage(product.images?.[0] || product.image, 100)" class="w-full h-full object-cover" />
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-[10px] text-[#A1A1AA] font-bold uppercase tracking-widest truncate">{{ product.name }}</p>
            <p class="text-sm font-black text-[#00FF00]">{{ formatPrice(product.price) }}</p>
          </div>
        </div>
        <button @click="handleBuyNow" :disabled="product.stock <= 0" class="shrink-0 bg-[#00FF00] text-black font-black uppercase tracking-widest text-[10px] sm:text-xs px-5 py-3 rounded-xl shadow-[0_0_15px_rgba(0,255,0,0.3)] disabled:bg-zinc-800 disabled:text-zinc-600 disabled:shadow-none transition-all">
          {{ product.stock > 0 ? 'Comprar' : 'Agotado' }}
        </button>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import type { Product } from '~/composables/useCatalog'
import { useLanguage } from '~/composables/useLanguage'
import { useDepartment } from '~/composables/useDepartment'
import { formatPrice } from '~/utils/format'
import { optimizeImage } from '~/utils/image'

const route = useRoute()
const router = useRouter()
const { addToCart, getProductQty, isStockFull } = useCart()
const { products, categories, isLoading, fetchCatalog } = useCatalog()
const { t, lang } = useLanguage()
const { setDepartment } = useDepartment()

// Helper simple para slug (puedes reemplazar con el tuyo de utils si existe)
const generateProductSlug = (id: number | string, name: string) => {
  return `${id}-${name.toLowerCase().replace(/[^a-z0-9]+/g, '-').replace(/(^-|-$)+/g, '')}`
}

const getIdFromSlug = (s: string) => {
  const match = s.match(/^(\d+)-/)
  return match ? parseInt(match[1], 10) : parseInt(s, 10)
}

const slug = computed(() => route.params.slug as string)
const productId = computed(() => getIdFromSlug(slug.value))

const product = computed(() => products.value.find(p => p.id === productId.value) as Product | undefined)

watchEffect(() => {
  if (product.value && categories.value.length > 0) {
    const cat = categories.value.find(c => c.id === product.value!.category)
    if (cat && cat.department) {
      setDepartment(cat.department)
    }
  }
})

// SEO dinámico
watchEffect(() => {
  if (product.value) {
    const title = `${product.value.name} | PersonalBarber`
    const description = `${product.value.name} de ${product.value.brand}. ${String(product.value.description).substring(0, 150)}...`
    const image = product.value.images?.[0] || product.value.image || '/og-image.webp'

    useSeoMeta({
      title,
      ogTitle: title,
      description,
      ogDescription: description,
      ogImage: image,
    })

    useHead({
      script: [{
        key: 'product-ld',
        type: 'application/ld+json',
        children: JSON.stringify({
          '@context': 'https://schema.org',
          '@type': 'Product',
          name: product.value.name,
          description: product.value.description,
          brand: { '@type': 'Brand', name: product.value.brand || 'PersonalBarber' },
          image: product.value.images || [product.value.image],
          url: `https://personalbarber.vip/tienda/producto/${slug.value}`,
          offers: {
            '@type': 'Offer',
            priceCurrency: 'COP',
            price: String(product.value.price).replace(/\D/g, ''),
            availability: product.value.stock > 0
              ? 'https://schema.org/InStock'
              : 'https://schema.org/OutOfStock',
            seller: { '@type': 'Organization', name: 'PersonalBarber' },
          }
        })
      }]
    })
  }
})

// UI State
const qty = ref(1)
const showConfirm = ref(false)
const carouselRef = ref<HTMLElement | null>(null)
const activeIdx = ref(0)
const selectedVariant = ref(0)
const expandedAccordion = ref('desc')

// Dynamic accordion visibility — only render if backend has content
const hasDescriptionContent = computed(() => {
  if (!product.value) return false
  return !!(product.value.description?.trim() || (product.value.benefits && product.value.benefits.length > 0))
})

const hasSpecsContent = computed(() => {
  if (!product.value) return false
  return !!(product.value.usage?.trim() || product.value.specs?.trim())
})

function handleScroll(e: Event) {
  const container = e.target as HTMLElement
  activeIdx.value = Math.round(container.scrollLeft / container.offsetWidth)
}

function scrollToImage(idx: number) {
  if (!carouselRef.value) return
  carouselRef.value.scrollTo({ left: idx * carouselRef.value.offsetWidth, behavior: 'smooth' })
  activeIdx.value = idx
}

watch(() => route.params.slug, () => {
  qty.value = 1
  activeIdx.value = 0
  selectedVariant.value = 0
  expandedAccordion.value = 'desc'
  if (carouselRef.value) carouselRef.value.scrollLeft = 0
  window.scrollTo({ top: 0, behavior: 'smooth' })
})

const recommendedProducts = computed(() => {
  if (!product.value || products.value.length === 0) return []
  const sameCategory = products.value
    .filter(p => p.category === product.value!.category && p.id !== product.value!.id)
    .sort(() => 0.5 - Math.random()).slice(0, 3)
  const otherCategories = products.value
    .filter(p => p.category !== product.value!.category)
    .sort(() => 0.5 - Math.random()).slice(0, 3)
  return [...sameCategory, ...otherCategories].slice(0, 3) // Mostrar máximo 3 en el cross-sell
})

function getCategoryLabel(catId: string) {
  const cat = categories.value.find(c => c.id === catId)
  return cat ? cat.label : catId
}

function handleAddToCart() {
  if (!product.value) return
  const result = addToCart(product.value as Parameters<typeof addToCart>[0], qty.value)
  if (result.success) {
    showConfirm.value = true
    qty.value = 1
    setTimeout(() => { showConfirm.value = false }, 3000)
  }
}

function handleBuyNow() {
  if (!product.value) return
  addToCart(product.value as Parameters<typeof addToCart>[0], qty.value)
  router.push('/checkout')
}

if (import.meta.server) {
  await fetchCatalog()
} else {
  fetchCatalog()
}
</script>

<style scoped>
.fade-enter-active, .fade-leave-active { transition: opacity 0.4s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }

.slide-up-enter-active, .slide-up-leave-active { transition: transform 0.4s cubic-bezier(0.16, 1, 0.3, 1), opacity 0.4s ease; }
.slide-up-enter-from, .slide-up-leave-to { transform: translateY(100%); opacity: 0; }

.hide-scrollbar::-webkit-scrollbar { display: none; }
.hide-scrollbar { -ms-overflow-style: none; scrollbar-width: none; }

.pb-safe { padding-bottom: env(safe-area-inset-bottom, 1rem); }
</style>
