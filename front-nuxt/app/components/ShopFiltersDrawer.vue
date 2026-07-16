<template>
  <Transition name="drawer">
    <div v-if="isOpen" class="fixed inset-0 z-[80] flex justify-end">
      <!-- Backdrop -->
      <div 
        @click="close" 
        class="fixed inset-0 bg-black/80 backdrop-blur-md transition-opacity"
      ></div>

      <!-- Drawer Content -->
      <div class="relative w-full max-w-md bg-zinc-950 border-l border-zinc-800 h-full flex flex-col z-10 shadow-2xl overflow-hidden">
        <!-- Header -->
        <div class="p-6 border-b border-zinc-800 flex items-center justify-between bg-zinc-900/50">
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 rounded-xl dept-bg flex items-center justify-center text-black font-black text-xs">
              ⚡
            </div>
            <div>
              <h3 class="text-base font-black uppercase tracking-tight text-white">Filtros Avanzados</h3>
              <p class="text-[10px] text-zinc-400 font-bold tracking-widest uppercase">Explora por marca y categoría</p>
            </div>
          </div>
          <button @click="close" class="p-2 text-zinc-400 hover:text-white transition-colors">
            <fa-icon :icon="['fas', 'times']" class="text-lg" />
          </button>
        </div>

        <!-- Body Scrollable -->
        <div class="flex-1 overflow-y-auto p-6 space-y-8 no-scrollbar">
          <!-- 1. Búsqueda Rápida -->
          <div class="space-y-2">
            <label class="text-[10px] font-black uppercase tracking-[0.2em] text-zinc-400">Palabra Clave</label>
            <div class="relative group">
              <fa-icon :icon="['fas', 'search']" class="absolute left-4 top-1/2 -translate-y-1/2 text-zinc-500 group-focus-within:text-white transition-colors" />
              <input 
                v-model="localSearch" 
                type="text" 
                placeholder="Máquina, cera, rasuradora..." 
                class="w-full pl-11 pr-4 py-3 bg-zinc-900 border border-zinc-800 rounded-2xl text-sm text-white placeholder:text-zinc-600 focus:outline-none focus:border-white transition-all"
              />
              <button v-if="localSearch" @click="localSearch = ''" class="absolute right-4 top-1/2 -translate-y-1/2 text-zinc-500 hover:text-white">
                <fa-icon :icon="['fas', 'times-circle']" />
              </button>
            </div>
          </div>

          <!-- 2. Ordenar por -->
          <div class="space-y-2">
            <label class="text-[10px] font-black uppercase tracking-[0.2em] text-zinc-400">Ordenar Por</label>
            <div class="grid grid-cols-3 gap-2">
              <button 
                v-for="opt in sortOptions" 
                :key="opt.id"
                @click="localSort = opt.id"
                class="py-2.5 px-3 rounded-xl text-[11px] font-black uppercase tracking-tight border transition-all duration-300 flex flex-col items-center justify-center gap-1 text-center"
                :class="localSort === opt.id ? 'dept-bg text-black border-transparent shadow-lg' : 'bg-zinc-900/60 border-zinc-800 text-zinc-400 hover:border-zinc-700 hover:text-white'"
              >
                <span>{{ opt.label }}</span>
              </button>
            </div>
          </div>

          <!-- 3. Filtrar por Marcas -->
          <div v-if="brands.length > 0" class="space-y-3">
            <div class="flex items-center justify-between">
              <label class="text-[10px] font-black uppercase tracking-[0.2em] text-zinc-400">Marcas Disponibles</label>
              <button 
                v-if="localBrands.length > 0" 
                @click="localBrands = []" 
                class="text-[10px] font-bold text-red-400 hover:underline"
              >
                Limpiar marcas
              </button>
            </div>
            <div class="flex flex-wrap gap-2">
              <button 
                v-for="brand in brands" 
                :key="brand"
                @click="toggleBrand(brand)"
                class="px-3.5 py-1.5 rounded-full text-xs font-bold transition-all border flex items-center gap-1.5"
                :class="localBrands.includes(brand) ? 'dept-bg text-black border-transparent font-black shadow-lg scale-105' : 'bg-zinc-900 border-zinc-800 text-zinc-400 hover:border-zinc-700 hover:text-white'"
              >
                <span>{{ brand }}</span>
                <fa-icon v-if="localBrands.includes(brand)" :icon="['fas', 'check']" class="text-[9px]" />
              </button>
            </div>
          </div>

          <!-- 4. Categoría Activa -->
          <div class="space-y-3">
            <label class="text-[10px] font-black uppercase tracking-[0.2em] text-zinc-400">Categoría</label>
            <div class="grid grid-cols-2 gap-2">
              <button 
                v-for="f in categories" 
                :key="f.id"
                @click="localCategory = f.id"
                class="px-3 py-2.5 rounded-xl text-xs font-bold text-left border transition-all truncate flex items-center justify-between"
                :class="localCategory === f.id ? 'dept-bg text-black border-transparent font-black' : 'bg-zinc-900 border-zinc-800 text-zinc-400 hover:border-zinc-700 hover:text-white'"
              >
                <span class="truncate">{{ f.label }}</span>
                <fa-icon v-if="localCategory === f.id" :icon="['fas', 'check']" class="text-[10px] shrink-0 ml-1" />
              </button>
            </div>
          </div>
        </div>

        <!-- Footer / Apply -->
        <div class="p-6 border-t border-zinc-800 bg-zinc-900/80 flex items-center gap-3">
          <button 
            @click="resetAll" 
            class="flex-1 py-3 px-4 bg-zinc-800 hover:bg-zinc-700 text-zinc-300 rounded-2xl text-xs font-black uppercase tracking-wider transition-all"
          >
            Restablecer
          </button>
          <button 
            @click="applyFilters" 
            class="flex-[2] py-3.5 px-6 dept-bg text-black rounded-2xl text-xs font-black uppercase tracking-wider transition-all shadow-lg hover:scale-[1.02] active:scale-95 flex items-center justify-center gap-2"
          >
            <span>Ver Productos</span>
            <fa-icon :icon="['fas', 'arrow-right']" class="text-[10px]" />
          </button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'

const props = defineProps<{
  isOpen: boolean
  brands: string[]
  categories: { id: string; label: string }[]
  searchQuery: string
  selectedBrands: string[]
  selectedCategory: string
  sortBy: string
}>()

const emit = defineEmits<{
  (e: 'close'): void
  (e: 'update:filters', payload: { searchQuery: string; selectedBrands: string[]; selectedCategory: string; sortBy: string }): void
}>()

const localSearch = ref('')
const localBrands = ref<string[]>([])
const localCategory = ref('all')
const localSort = ref('default')

const sortOptions = [
  { id: 'default', label: '🚀 Recomendados' },
  { id: 'price-asc', label: '📉 Precio: Bajo' },
  { id: 'price-desc', label: '📈 Precio: Alto' },
]

watch(() => props.isOpen, (newVal) => {
  if (newVal) {
    localSearch.value = props.searchQuery || ''
    localBrands.value = [...(props.selectedBrands || [])]
    localCategory.value = props.selectedCategory || 'all'
    localSort.value = props.sortBy || 'default'
  }
})

function toggleBrand(brand: string) {
  const index = localBrands.value.indexOf(brand)
  if (index === -1) {
    localBrands.value.push(brand)
  } else {
    localBrands.value.splice(index, 1)
  }
}

function resetAll() {
  localSearch.value = ''
  localBrands.value = []
  localCategory.value = 'all'
  localSort.value = 'default'
  applyFilters()
}

function applyFilters() {
  emit('update:filters', {
    searchQuery: localSearch.value,
    selectedBrands: [...localBrands.value],
    selectedCategory: localCategory.value,
    sortBy: localSort.value
  })
  emit('close')
}

function close() {
  emit('close')
}
</script>

<style scoped>
.drawer-enter-active,
.drawer-leave-active {
  transition: opacity 0.3s ease;
}
.drawer-enter-active > div:last-child,
.drawer-leave-active > div:last-child {
  transition: transform 0.3s cubic-bezier(0.16, 1, 0.3, 1);
}
.drawer-enter-from,
.drawer-leave-to {
  opacity: 0;
}
.drawer-enter-from > div:last-child,
.drawer-leave-to > div:last-child {
  transform: translateX(100%);
}
.no-scrollbar::-webkit-scrollbar {
  display: none;
}
.no-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
</style>
