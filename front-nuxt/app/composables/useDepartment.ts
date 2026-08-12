import { useState } from '#app'
import { computed } from 'vue'

export type Department = 'all' | 'men' | 'merch' | 'women'

export function useDepartment() {
  const activeDepartment = useState<Department>('active-department', () => 'all')

  const accentColor = computed(() => {
    switch (activeDepartment.value) {
      case 'all': return '#e2d9c8'   // warm neutral — ni verde ni rosa
      case 'men': return '#39FF14'
      case 'merch': return '#22d3ee'
      case 'women': return '#ec4899'
      default: return '#e2d9c8'
    }
  })

  const accentGlow = computed(() => {
    switch (activeDepartment.value) {
      case 'all': return 'rgba(226, 217, 200, 0.35)' // warm pearl glow
      case 'men': return 'rgba(57, 255, 20, 0.4)'
      case 'merch': return 'rgba(34, 211, 238, 0.4)'
      case 'women': return 'rgba(236, 72, 153, 0.4)'
      default: return 'rgba(226, 217, 200, 0.35)'
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
