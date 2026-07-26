// composables/useLanguage.ts
// Sistema i18n propio — sin módulos externos
// Estado singleton reactivo que persiste en localStorage

import { es } from '~/locales/es'
import { en } from '~/locales/en'

type Lang = 'es' | 'en'
type DeepPartial<T> = { [P in keyof T]?: T[P] extends object ? DeepPartial<T[P]> : T[P] }

const translations: Record<Lang, typeof es> = { es, en }

// Singleton — compartido entre todos los componentes
const lang = ref<Lang>('es')
let initialized = false

export function useLanguage() {
  // Inicializar desde localStorage / browser language (solo cliente, una vez)
  const initLang = () => {
    if (initialized || !import.meta.client) return
    initialized = true
    const stored = localStorage.getItem('pb-lang') as Lang | null
    // Evitar descalce de hidratación en SSR: solo actualizar si el usuario guardó 'en' explícitamente
    if (stored === 'en') {
      nextTick(() => {
        lang.value = 'en'
      })
    }
  }

  // Función de traducción con soporte a dot notation: t('hero.line1')
  const t = (key: string): string => {
    const keys = key.split('.')
    let obj: unknown = translations[lang.value]
    for (const k of keys) {
      if (obj && typeof obj === 'object') {
        obj = (obj as Record<string, unknown>)[k]
      } else {
        return key // fallback al key si no existe
      }
    }
    return typeof obj === 'string' ? obj : key
  }

  const setLang = (l: Lang) => {
    lang.value = l
    if (import.meta.client) localStorage.setItem('pb-lang', l)
  }

  const toggleLang = () => setLang(lang.value === 'es' ? 'en' : 'es')

  return { lang: readonly(lang), t, setLang, toggleLang, initLang }
}
