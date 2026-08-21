import { useState } from '#app'
import { computed } from 'vue'

export type Department = 'all' | 'men' | 'merch' | 'women'

export function useDepartment() {
  const activeDepartment = useState<Department>('active-department', () => 'all')

  const accentColor = computed(() => {
    switch (activeDepartment.value) {
      case 'all':    return '#39FF14'    // verde néon original para catálogo general
      case 'men':    return '#39FF14'    // verde néon — identidad Barbería
      case 'merch':  return '#22d3ee'    // cyan — identidad Ropa & Merch
      case 'women':  return '#FF1493'    // rosa neón eléctrico — identidad Beauty
      default:       return '#39FF14'
    }
  })

  const accentGlow = computed(() => {
    switch (activeDepartment.value) {
      case 'all':    return 'rgba(57, 255, 20, 0.4)'
      case 'men':    return 'rgba(57, 255, 20, 0.4)'
      case 'merch':  return 'rgba(34, 211, 238, 0.4)'
      case 'women':  return 'rgba(255, 20, 147, 0.45)'
      default:       return 'rgba(57, 255, 20, 0.4)'
    }
  })

  function setDepartment(dept: Department) {
    activeDepartment.value = dept
  }

  return {
    activeDepartment,
    accentColor,
    accentGlow,
    setDepartment
  }
}
