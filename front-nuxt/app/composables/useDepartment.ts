import { useState } from '#app'
import { computed } from 'vue'

export type Department = 'all' | 'men' | 'merch' | 'women'

export function useDepartment() {
  const activeDepartment = useState<Department>('active-department', () => 'all')

  const accentColor = computed(() => {
    switch (activeDepartment.value) {
      case 'all':    return '#a78bfa'    // violeta neutro — fusión verde (men) + rosa (women)
      case 'men':    return '#39FF14'    // verde néon — identidad Barbería
      case 'merch':  return '#22d3ee'    // cyan — identidad Ropa & Merch
      case 'women':  return '#FF1493'    // rosa neón eléctrico — identidad Beauty
      default:       return '#a78bfa'
    }
  })

  const accentGlow = computed(() => {
    switch (activeDepartment.value) {
      case 'all':    return 'rgba(167, 139, 250, 0.35)'
      case 'men':    return 'rgba(57, 255, 20, 0.4)'
      case 'merch':  return 'rgba(34, 211, 238, 0.4)'
      case 'women':  return 'rgba(255, 20, 147, 0.45)'
      default:       return 'rgba(167, 139, 250, 0.35)'
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
