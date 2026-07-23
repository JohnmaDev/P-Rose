// composables/useCatalog.ts
// Catálogo SSR — intenta el endpoint unificado /api/get_catalog primero.
// Si no está disponible (entorno local pre-deploy), hace fallback a los
// endpoints individuales para mantener compatibilidad durante la transición.

export interface Product {
  id: number
  name: string
  brand: string
  price: string | number
  comparePrice?: number
  description: string
  category: string
  images?: string[]
  image?: string
  stock: number
  slug?: string
  benefits?: string[]
  usage?: string
  specs?: string
  variants?: string[]
  is_active?: boolean
  [key: string]: unknown
}

export interface Category {
  id: string
  label: string
  department: string
  style?: string
  accent?: string
  cover?: string
  icon?: string
  comingSoon?: boolean
  [key: string]: unknown
}

export function useCatalog() {
  const products   = useState<Product[]>('catalog-products',   () => [])
  const categories = useState<Category[]>('catalog-categories', () => [])
  const isLoaded   = useState<boolean>('catalog-isLoaded',     () => false)
  const isLoading  = useState<boolean>('catalog-isLoading',    () => false)
  const error      = useState<string | null>('catalog-error',  () => null)

  /**
   * Intenta primero el endpoint unificado /api/get_catalog (1 conexión MongoDB,
   * productos + categorías en paralelo vía goroutines en Go).
   * Si devuelve 404 (entorno local antes del deploy), hace fallback a los dos
   * endpoints individuales con Promise.all para mantener compatibilidad.
   */
  const fetchCatalog = async (force = false) => {
    if (isLoaded.value && !force) return { success: true }
    if (isLoading.value)          return { success: false }

    isLoading.value = true
    error.value     = null

    try {
      // ── Intento 1: endpoint unificado (producción) ──
      const res = await $fetch<{
        ok: boolean
        products: Product[]
        categories: Category[]
      }>('/api/get_catalog').catch(() => null)

      if (res?.ok) {
        products.value   = (res.products || []).filter(p => p.is_active !== false)
        categories.value = res.categories
        isLoaded.value   = true
        return { success: true }
      }

      // ── Fallback: endpoints individuales (local / pre-deploy) ──
      const [resProd, resCat] = await Promise.all([
        $fetch<{ ok: boolean; products: Product[] }>('/api/get_products'),
        $fetch<{ ok: boolean; categories: Category[] }>('/api/get_categories'),
      ])

      if (resProd.ok && resCat.ok) {
        products.value   = (resProd.products || []).filter(p => p.is_active !== false)
        categories.value = resCat.categories
        isLoaded.value   = true
        return { success: true }
      }

      throw new Error('Los endpoints devolvieron ok=false')

    } catch (err: unknown) {
      const msg = err instanceof Error ? err.message : 'Falló la conexión o error de red'
      if (import.meta.client) {
        error.value = msg
      }
      console.error('Error cargando catálogo:', err)
      return { success: false, error: msg }
    } finally {
      isLoading.value = false
    }
  }

  const invalidateCatalog = () => {
    isLoaded.value = false
    return fetchCatalog(true)
  }

  return {
    products,
    categories,
    isLoading,
    isLoaded,
    error,
    fetchCatalog,
    invalidateCatalog,
  }
}
