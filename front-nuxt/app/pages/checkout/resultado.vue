<template>
  <div class="bg-barber-black min-h-screen text-white flex flex-col items-center justify-center px-4">
    <!-- Header -->
    <header class="w-full border-b border-white/8 bg-barber-black fixed top-0 z-10">
      <div class="max-w-5xl mx-auto px-6 py-4 flex items-center justify-center">
        <NuxtLink to="/">
          <img src="/favicon.svg" alt="PersonalBarber" class="h-12 w-auto opacity-90 hover:opacity-100 transition-opacity" width="48" height="48" decoding="async" />
        </NuxtLink>
      </div>
    </header>

    <!-- Contenido principal -->
    <div class="max-w-md w-full text-center pt-24 pb-12">

      <!-- ⏳ Procesando -->
      <div v-if="orderStatus === 'loading' || orderStatus === 'PENDING' || orderStatus === 'AWAITING_PAYMENT'" class="space-y-6">
        <div class="w-24 h-24 rounded-full bg-neon-green/10 flex items-center justify-center mx-auto animate-pulse">
          <fa-icon :icon="['fas', 'spinner']" class="text-4xl text-neon-green fa-spin" />
        </div>
        <h1 class="text-2xl font-black text-white tracking-tight">Procesando tu pago...</h1>
        <p class="text-gray-400 text-sm leading-relaxed max-w-xs mx-auto">
          Estamos verificando tu transacción con Wompi. Esto puede tomar unos segundos.
        </p>
        <div class="flex items-center justify-center gap-2 text-xs text-gray-600">
          <div class="w-2 h-2 rounded-full bg-neon-green animate-ping"></div>
          Consultando estado...
        </div>
      </div>

      <!-- ✅ Aprobado -->
      <div v-else-if="orderStatus === 'APPROVED'" class="space-y-6">
        <div class="w-24 h-24 rounded-full bg-neon-green/15 flex items-center justify-center mx-auto result-icon">
          <fa-icon :icon="['fas', 'check-circle']" class="text-5xl text-neon-green" />
        </div>
        <h1 class="text-2xl font-black text-white tracking-tight">¡Pago exitoso!</h1>
        <p class="text-gray-400 text-sm leading-relaxed max-w-xs mx-auto">
          Tu pedido <span class="text-neon-green font-bold">{{ orderId }}</span> ha sido confirmado.
          Recibirás una notificación por WhatsApp con los detalles del envío.
        </p>
        <div class="bg-white/5 border border-white/10 rounded-2xl p-5 text-left space-y-3 mt-4">
          <div class="flex justify-between text-sm">
            <span class="text-gray-500">Orden</span>
            <span class="text-white font-bold">{{ orderId }}</span>
          </div>
          <div v-if="orderData?.total_format" class="flex justify-between text-sm">
            <span class="text-gray-500">Total pagado</span>
            <span class="text-neon-green font-bold">${{ orderData.total_format }} COP</span>
          </div>
          <div v-if="orderData?.shippingMethod" class="flex justify-between text-sm">
            <span class="text-gray-500">Envío</span>
            <span class="text-white text-xs">{{ shippingLabel }}</span>
          </div>
        </div>
        <div class="flex flex-col gap-3 mt-6">
          <NuxtLink to="/" class="w-full py-4 bg-neon-green hover:bg-neon-green-dark text-black font-black rounded-xl transition-all duration-300 uppercase text-sm text-center">
            Volver a la tienda
          </NuxtLink>
          <a :href="whatsappLink" target="_blank" class="w-full py-4 bg-[#25D366] hover:bg-[#20bd5a] text-black font-black rounded-xl transition-all duration-300 uppercase text-sm text-center flex items-center justify-center gap-2">
            <fa-icon :icon="['fab', 'whatsapp']" class="text-lg" />
            Contactar por WhatsApp
          </a>
        </div>
      </div>

      <!-- ❌ Rechazado / Error -->
      <div v-else-if="orderStatus === 'DECLINED' || orderStatus === 'ERROR' || orderStatus === 'VOIDED'" class="space-y-6">
        <div class="w-24 h-24 rounded-full bg-red-500/15 flex items-center justify-center mx-auto result-icon">
          <fa-icon :icon="['fas', 'times-circle']" class="text-5xl text-red-400" />
        </div>
        <h1 class="text-2xl font-black text-white tracking-tight">Pago no exitoso</h1>
        <p class="text-gray-400 text-sm leading-relaxed max-w-xs mx-auto">
          {{ declinedMessage }}
        </p>
        <div class="flex flex-col gap-3 mt-6">
          <NuxtLink to="/checkout" class="w-full py-4 bg-white/10 hover:bg-white/15 text-white font-black rounded-xl transition-all duration-300 uppercase text-sm text-center border border-white/10">
            Reintentar pago
          </NuxtLink>
          <a :href="whatsappLink" target="_blank" class="w-full py-4 bg-[#25D366] hover:bg-[#20bd5a] text-black font-black rounded-xl transition-all duration-300 uppercase text-sm text-center flex items-center justify-center gap-2">
            <fa-icon :icon="['fab', 'whatsapp']" class="text-lg" />
            Contactar soporte
          </a>
        </div>
      </div>

      <!-- ⏱️ Timeout -->
      <div v-else-if="orderStatus === 'timeout'" class="space-y-6">
        <div class="w-24 h-24 rounded-full bg-yellow-500/15 flex items-center justify-center mx-auto result-icon">
          <fa-icon :icon="['fas', 'clock']" class="text-5xl text-yellow-400" />
        </div>
        <h1 class="text-2xl font-black text-white tracking-tight">Pago en proceso</h1>
        <p class="text-gray-400 text-sm leading-relaxed max-w-xs mx-auto">
          Tu pago está siendo procesado y puede tardar unos minutos.
          Te notificaremos por WhatsApp cuando se confirme.
        </p>
        <p class="text-gray-600 text-xs mt-2">
          Orden: <span class="text-white font-bold">{{ orderId }}</span>
        </p>
        <div class="flex flex-col gap-3 mt-6">
          <NuxtLink to="/" class="w-full py-4 bg-white/10 hover:bg-white/15 text-white font-black rounded-xl transition-all duration-300 uppercase text-sm text-center border border-white/10">
            Volver a la tienda
          </NuxtLink>
          <a :href="whatsappLink" target="_blank" class="w-full py-4 bg-[#25D366] hover:bg-[#20bd5a] text-black font-black rounded-xl transition-all duration-300 uppercase text-sm text-center flex items-center justify-center gap-2">
            <fa-icon :icon="['fab', 'whatsapp']" class="text-lg" />
            Consultar estado
          </a>
        </div>
      </div>

      <!-- 🔒 Sin ID -->
      <div v-else-if="orderStatus === 'no_id'" class="space-y-6">
        <div class="w-24 h-24 rounded-full bg-white/5 flex items-center justify-center mx-auto">
          <fa-icon :icon="['fas', 'search']" class="text-4xl text-gray-600" />
        </div>
        <h1 class="text-xl font-black text-white tracking-tight">No encontramos tu orden</h1>
        <p class="text-gray-400 text-sm max-w-xs mx-auto">
          No se proporcionó un ID de orden válido. Si realizaste un pago, contacta soporte.
        </p>
        <NuxtLink to="/" class="inline-block mt-4 px-8 py-3 bg-white/10 hover:bg-white/15 text-white font-bold rounded-xl transition-all text-sm border border-white/10">
          Ir a la tienda
        </NuxtLink>
      </div>

      <!-- Footer seguridad -->
      <p class="text-center text-gray-700 text-[10px] mt-10">
        <fa-icon :icon="['fas', 'shield-alt']" class="mr-1" />
        Transacción procesada de forma segura por Wompi · SSL encriptado
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
useSeoMeta({ title: 'Resultado del Pago | PersonalBarber' })

const route = useRoute()

const orderId = ref('')
const orderStatus = ref<string>('loading')
const orderData = ref<any>(null)
const pollCount = ref(0)
const maxPolls = 20 // 20 × 3s = 60 segundos máximo de espera
let pollInterval: ReturnType<typeof setInterval> | null = null

// Mensaje descriptivo según el estado de rechazo
const declinedMessage = computed(() => {
  switch (orderStatus.value) {
    case 'DECLINED':
      return 'Tu entidad financiera no aprobó la transacción. Verifica tus datos o intenta con otro método de pago.'
    case 'VOIDED':
      return 'La transacción fue cancelada. Si no la cancelaste tú, contacta a tu banco.'
    case 'ERROR':
      return 'Ocurrió un error técnico durante el procesamiento. Por favor intenta nuevamente o contacta soporte.'
    default:
      return 'No se pudo completar el pago. Intenta de nuevo o contacta soporte.'
  }
})

// Etiqueta de envío legible
const shippingLabel = computed(() => {
  const methods: Record<string, string> = {
    express_valle: 'Valle de Aburrá · 24-48h',
    express_alrededores: 'Alrededores · 24-48h',
    express_nacional: 'Nacional · 2-5 días',
  }
  return methods[orderData.value?.shippingMethod] || orderData.value?.shippingMethod || ''
})

// Link de WhatsApp
const whatsappLink = computed(() => {
  const msg = orderId.value
    ? `Hola, tengo una consulta sobre mi pedido ${orderId.value}`
    : 'Hola, necesito ayuda con mi pago'
  return `https://api.whatsapp.com/send?phone=573337518070&text=${encodeURIComponent(msg)}`
})

// Consultar estado de la orden
async function checkOrderStatus() {
  if (!orderId.value) return

  try {
    const data = await $fetch<{ ok: boolean; order: any }>(`/api/order_status?id=${orderId.value}`)

    if (data.ok && data.order) {
      orderData.value = data.order
      const status = data.order.status

      // Si el estado es final, dejar de hacer polling
      if (['APPROVED', 'DECLINED', 'VOIDED', 'ERROR'].includes(status)) {
        orderStatus.value = status
        stopPolling()

        // Si fue aprobado, limpiar el carrito
        if (status === 'APPROVED') {
          try {
            const { clearCart } = useCart()
            clearCart()
          } catch { /* cart may not be initialized */ }
        }
        return
      }

      // Si sigue pendiente, actualizar el estado visual
      orderStatus.value = status
    }
  } catch (err) {
    console.error('Error consultando estado:', err)
  }

  // Incrementar contador de polls
  pollCount.value++
  if (pollCount.value >= maxPolls) {
    orderStatus.value = 'timeout'
    stopPolling()
  }
}

function stopPolling() {
  if (pollInterval) {
    clearInterval(pollInterval)
    pollInterval = null
  }
}

// Inicializar
onMounted(() => {
  // Leer el ID de la orden de los query params
  const id = route.query.id as string
  if (!id) {
    orderStatus.value = 'no_id'
    return
  }

  orderId.value = id
  orderStatus.value = 'loading'

  // Primera consulta inmediata
  checkOrderStatus()

  // Polling cada 3 segundos
  pollInterval = setInterval(checkOrderStatus, 3000)
})

onBeforeUnmount(() => {
  stopPolling()
})
</script>

<style scoped>
/* Animación de entrada para los iconos de resultado */
.result-icon {
  animation: resultBounce 0.6s cubic-bezier(0.34, 1.56, 0.64, 1);
}
@keyframes resultBounce {
  from {
    transform: scale(0.3);
    opacity: 0;
  }
  to {
    transform: scale(1);
    opacity: 1;
  }
}
</style>
