<template>
  <div class="bg-barber-black min-h-screen text-white flex flex-col">

    <!-- Header -->
    <header class="w-full border-b border-white/8 bg-barber-black">
      <div class="max-w-5xl mx-auto px-6 py-5 flex flex-col items-center gap-3">
        <NuxtLink to="/">
          <img src="/favicon.svg" alt="PersonalBarber" class="h-9 w-auto opacity-90 hover:opacity-100 transition-opacity" />
        </NuxtLink>
        <!-- Breadcrumb -->
        <nav class="flex items-center gap-2 text-xs font-semibold tracking-wide">
          <button @click="goToStep(1)" :class="step >= 1 ? 'text-neon-green' : 'text-white/30'" class="hover:underline">Información</button>
          <span class="text-white/20">›</span>
          <button @click="goToStep(2)" :class="step >= 2 ? 'text-neon-green' : 'text-white/30'" :disabled="step < 2" class="disabled:cursor-not-allowed hover:underline">Envío</button>
          <span class="text-white/20">›</span>
          <span :class="step === 3 ? 'text-white font-bold' : 'text-white/30'">Pago</span>
        </nav>
      </div>
    </header>

    <main class="flex-1">
      <ClientOnly>
        <!-- Carrito vacío -->
        <div v-if="cartItems.length === 0" class="flex flex-col items-center justify-center min-h-[60vh] gap-4">
          <fa-icon :icon="['fas', 'shopping-bag']" class="text-5xl text-white/10" />
          <p class="text-gray-500">Tu carrito está vacío</p>
          <NuxtLink to="/tienda" class="text-neon-green hover:underline text-sm font-bold">← Explorar productos</NuxtLink>
        </div>

        <div v-else class="max-w-5xl mx-auto px-4 py-10">
          <div class="grid grid-cols-1 lg:grid-cols-5 gap-8 items-start">

            <!-- Columna principal -->
            <div class="lg:col-span-3 space-y-6">

              <!-- ═══ PASO 1: INFORMACIÓN ═══ -->
              <div v-if="step === 1">
                <div class="bg-white/5 rounded-2xl p-6 border border-white/10">
                  <h2 class="text-lg font-bold text-white mb-5 flex items-center gap-2">
                    <span class="w-6 h-6 bg-neon-green text-black text-xs font-black rounded-full flex items-center justify-center">1</span>
                    Información de contacto
                  </h2>
                  <div class="space-y-4">
                    <div class="grid grid-cols-2 gap-4">
                      <div>
                        <label class="label-xs">Nombre *</label>
                        <input v-model="form.firstName" @blur="touched.firstName=true" type="text" placeholder="Juan" class="input-field" :class="{'border-red-500/50': touched.firstName && !form.firstName.trim()}" />
                        <p v-if="touched.firstName && !form.firstName.trim()" class="err">Campo obligatorio</p>
                      </div>
                      <div>
                        <label class="label-xs">Apellido *</label>
                        <input v-model="form.lastName" @blur="touched.lastName=true" type="text" placeholder="García" class="input-field" :class="{'border-red-500/50': touched.lastName && !form.lastName.trim()}" />
                        <p v-if="touched.lastName && !form.lastName.trim()" class="err">Campo obligatorio</p>
                      </div>
                    </div>
                    <div>
                      <label class="label-xs">Email *</label>
                      <input v-model="form.email" @blur="touched.email=true" type="email" placeholder="juan@email.com" class="input-field" :class="{'border-red-500/50': touched.email && !isEmailValid}" />
                      <p v-if="touched.email && !isEmailValid" class="err">Ingresa un correo válido</p>
                    </div>
                    <div>
                      <label class="label-xs">Teléfono / WhatsApp *</label>
                      <input v-model="form.phone" @input="form.phone=form.phone.replace(/[^0-9+]/g,'')" @blur="touched.phone=true" type="tel" placeholder="+57 300 123 4567" class="input-field" :class="{'border-red-500/50': touched.phone && !isPhoneValid}" />
                      <p v-if="touched.phone && !isPhoneValid" class="err">Número colombiano inválido</p>
                    </div>
                    <div class="border-t border-white/10 pt-4">
                      <h3 class="text-sm font-bold text-white mb-3">Dirección de envío</h3>
                      <div class="space-y-3">
                        <div>
                          <label class="label-xs">Ciudad *</label>
                          <input v-model="form.city" @blur="touched.city=true" type="text" placeholder="Medellín, Bogotá..." class="input-field" :class="{'border-red-500/50': touched.city && !form.city.trim()}" />
                          <p v-if="touched.city && !form.city.trim()" class="err">La ciudad es obligatoria</p>
                        </div>
                        <div>
                          <label class="label-xs">Dirección completa *</label>
                          <input v-model="form.address" @blur="touched.address=true" type="text" placeholder="Calle 50 # 30-10 Apto 201" class="input-field" :class="{'border-red-500/50': touched.address && !form.address.trim()}" />
                          <p v-if="touched.address && !form.address.trim()" class="err">La dirección es obligatoria</p>
                        </div>
                        <div>
                          <label class="label-xs">Notas adicionales <span class="text-gray-600 font-normal">(opcional)</span></label>
                          <input v-model="form.notes" type="text" placeholder="Torre A, conjunto cerrado, timbre 302..." class="input-field" />
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="mt-6 flex items-center justify-between">
                    <NuxtLink to="/tienda" class="text-gray-400 hover:text-white text-sm flex items-center gap-1 transition-colors">
                      <fa-icon :icon="['fas', 'arrow-left']" class="text-xs" /> Volver a la tienda
                    </NuxtLink>
                    <button @click="nextStep" :disabled="!step1Valid"
                      class="px-8 py-3 font-black text-sm rounded-xl transition-all duration-300 flex items-center gap-2"
                      :class="step1Valid ? 'bg-neon-green hover:bg-neon-green-dark text-black shadow-[0_0_20px_rgba(57,255,20,0.3)]' : 'bg-white/10 text-gray-600 cursor-not-allowed'">
                      Continuar con el envío <fa-icon :icon="['fas', 'arrow-right']" class="text-xs" />
                    </button>
                  </div>
                </div>
              </div>

              <!-- ═══ PASO 2: ENVÍO ═══ -->
              <div v-if="step === 2">
                <!-- Resumen de contacto (modo lectura) -->
                <div class="bg-white/5 rounded-2xl p-5 border border-white/10 mb-4">
                  <div class="flex items-center justify-between text-sm">
                    <div class="flex flex-col gap-1">
                      <div class="flex gap-4">
                        <span class="text-gray-500 w-16">Contacto</span>
                        <span class="text-white">{{ form.email }}</span>
                      </div>
                      <div class="flex gap-4">
                        <span class="text-gray-500 w-16">Enviar a</span>
                        <span class="text-white">{{ form.firstName }} {{ form.lastName }}, {{ form.city }}</span>
                      </div>
                    </div>
                    <button @click="step = 1" class="text-neon-green text-xs hover:underline font-semibold">Cambiar</button>
                  </div>
                </div>

                <!-- Método de envío -->
                <div class="bg-white/5 rounded-2xl p-6 border border-white/10">
                  <div class="flex items-center justify-between mb-5">
                    <h2 class="text-lg font-bold text-white flex items-center gap-2">
                      <span class="w-6 h-6 bg-neon-green text-black text-xs font-black rounded-full flex items-center justify-center">2</span>
                      Método de envío
                    </h2>
                    <!-- Chip ciudad detectada -->
                    <span class="text-[10px] font-bold px-3 py-1 rounded-full flex items-center gap-1.5"
                      :class="isMetroCity ? 'bg-neon-green/15 text-neon-green border border-neon-green/30' : 'bg-white/8 text-gray-400 border border-white/15'">
                      <span class="w-1.5 h-1.5 rounded-full" :class="isMetroCity ? 'bg-neon-green' : 'bg-gray-500'"></span>
                      {{ isMetroCity ? 'Área metro Medellín' : 'Envío nacional' }}
                    </span>
                  </div>
                  <div class="space-y-3">
                    <label v-for="s in shippingMethods" :key="s.id"
                      class="flex items-center gap-4 p-4 rounded-xl border cursor-pointer transition-all duration-300"
                      :class="selectedShipping === s.id ? 'border-neon-green bg-neon-green/10' : 'border-white/10 hover:border-white/20'">
                      <input type="radio" v-model="selectedShipping" :value="s.id" class="hidden" />
                      <div class="w-5 h-5 rounded-full border-2 flex items-center justify-center flex-shrink-0"
                        :class="selectedShipping === s.id ? 'border-neon-green' : 'border-gray-600'">
                        <div v-if="selectedShipping === s.id" class="w-2.5 h-2.5 rounded-full bg-neon-green"></div>
                      </div>
                      <div class="flex-1">
                        <div class="flex items-center gap-2 mb-0.5">
                          <span class="text-xl">{{ s.emoji }}</span>
                          <p class="text-white font-bold text-sm">{{ s.label }}</p>
                          <span v-if="s.badge" class="text-[9px] bg-neon-green/20 text-neon-green border border-neon-green/30 px-2 py-0.5 rounded-full font-black">{{ s.badge }}</span>
                        </div>
                        <p class="text-gray-500 text-xs">{{ s.desc }}</p>
                      </div>
                      <span class="text-white font-bold text-sm flex-shrink-0">{{ s.price }}</span>
                    </label>
                  </div>

                  <!-- Banner info envio nacional -->
                  <div v-if="!isMetroCity && selectedShipping === 'nacional'" class="mt-4 p-4 bg-yellow-500/10 border border-yellow-500/20 rounded-xl">
                    <p class="text-yellow-300 text-xs leading-relaxed">
                      <fa-icon :icon="['fas', 'truck']" class="mr-1" />
                      <strong>Sin costo adicional sorpresa:</strong> Confirmado tu pedido, te enviamos el costo exacto de envío a <strong>{{ form.city }}</strong> por WhatsApp antes de cobrar.
                      Trabajamos con <span class="text-white">Envia, Coordinadora, Interrapidísimo y Servientrega</span> según lo que mejor le sirva a tu ciudad.
                    </p>
                  </div>

                  <div class="mt-6 flex items-center justify-between">
                    <button @click="step = 1" class="text-gray-400 hover:text-white text-sm flex items-center gap-1 transition-colors">
                      <fa-icon :icon="['fas', 'arrow-left']" class="text-xs" /> Volver
                    </button>
                    <button @click="step = 3" :disabled="!selectedShipping"
                      class="px-8 py-3 font-black text-sm rounded-xl transition-all duration-300 flex items-center gap-2"
                      :class="selectedShipping ? 'bg-neon-green hover:bg-neon-green-dark text-black shadow-[0_0_20px_rgba(57,255,20,0.3)]' : 'bg-white/10 text-gray-600 cursor-not-allowed'">
                      Continuar al pago <fa-icon :icon="['fas', 'arrow-right']" class="text-xs" />
                    </button>
                  </div>
                </div>
              </div>

              <!-- ═══ PASO 3: PAGO ═══ -->
              <div v-if="step === 3">
                <!-- Resúmenes de pasos anteriores -->
                <div class="bg-white/5 rounded-2xl p-5 border border-white/10 mb-4 divide-y divide-white/8">
                  <div class="flex items-center justify-between text-sm pb-3">
                    <div class="flex gap-4">
                      <span class="text-gray-500 w-16">Contacto</span>
                      <span class="text-white">{{ form.email }}</span>
                    </div>
                    <button @click="step = 1" class="text-neon-green text-xs hover:underline font-semibold">Cambiar</button>
                  </div>
                  <div class="flex items-center justify-between text-sm py-3">
                    <div class="flex gap-4">
                      <span class="text-gray-500 w-16">Enviar a</span>
                      <span class="text-white">{{ form.firstName }} {{ form.lastName }}, {{ form.address }}, {{ form.city }}</span>
                    </div>
                    <button @click="step = 1" class="text-neon-green text-xs hover:underline font-semibold">Cambiar</button>
                  </div>
                  <div class="flex items-center justify-between text-sm pt-3">
                    <div class="flex gap-4">
                      <span class="text-gray-500 w-16">Envío</span>
                      <span class="text-white">{{ currentShipping?.label }} · {{ currentShipping?.price }}</span>
                    </div>
                    <button @click="step = 2" class="text-neon-green text-xs hover:underline font-semibold">Cambiar</button>
                  </div>
                </div>

                <!-- Métodos de pago -->
                <div class="bg-white/5 rounded-2xl p-6 border border-white/10">
                  <h2 class="text-lg font-bold text-white mb-5 flex items-center gap-2">
                    <span class="w-6 h-6 bg-neon-green text-black text-xs font-black rounded-full flex items-center justify-center">3</span>
                    Método de pago
                  </h2>
                  <div class="space-y-3">
                    <label v-for="method in paymentMethods" :key="method.id"
                      class="flex items-center gap-4 p-4 rounded-xl border cursor-pointer transition-all duration-300"
                      :class="selectedPayment === method.id ? 'border-neon-green bg-neon-green/10' : 'border-white/10 hover:border-white/20'">
                      <input type="radio" v-model="selectedPayment" :value="method.id" class="hidden" />
                      <div class="w-5 h-5 rounded-full border-2 flex items-center justify-center flex-shrink-0"
                        :class="selectedPayment === method.id ? 'border-neon-green' : 'border-gray-600'">
                        <div v-if="selectedPayment === method.id" class="w-2.5 h-2.5 rounded-full bg-neon-green"></div>
                      </div>
                      <div class="flex items-center gap-3 flex-1">
                        <div class="text-2xl">{{ method.emoji }}</div>
                        <div>
                          <p class="text-white font-bold text-sm">{{ method.label }}</p>
                          <p class="text-gray-500 text-xs">{{ method.desc }}</p>
                        </div>
                      </div>
                      <span v-if="method.badge" class="text-[10px] bg-neon-green/20 text-neon-green border border-neon-green/30 px-2 py-0.5 rounded-full font-bold">{{ method.badge }}</span>
                    </label>
                  </div>
                  <div class="mt-4 p-4 bg-blue-500/10 border border-blue-500/20 rounded-xl">
                    <p class="text-blue-300 text-xs leading-relaxed">
                      <fa-icon :icon="['fas', 'info-circle']" class="mr-1" />
                      <strong>Modo Demo:</strong> Los pagos online se activarán con Wompi muy pronto. Por ahora selecciona <strong>WhatsApp</strong> para coordinar.
                    </p>
                  </div>

                  <div class="mt-6 flex items-center justify-between">
                    <button @click="step = 2" class="text-gray-400 hover:text-white text-sm flex items-center gap-1 transition-colors">
                      <fa-icon :icon="['fas', 'arrow-left']" class="text-xs" /> Volver
                    </button>
                    <button @click="handleCheckout" :disabled="isProcessing"
                      class="px-8 py-3 bg-neon-green hover:bg-neon-green-dark text-black font-black text-sm rounded-xl transition-all duration-300 shadow-[0_0_20px_rgba(57,255,20,0.3)] flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed">
                      <fa-icon v-if="isProcessing" :icon="['fas', 'spinner']" class="fa-spin" />
                      <fa-icon v-else :icon="['fas', 'lock']" class="text-xs" />
                      {{ isProcessing ? 'Procesando...' : `Pagar ${cartTotalFormatted}` }}
                    </button>
                  </div>
                </div>
              </div>
            </div>

            <!-- Resumen del pedido (sticky) -->
            <div class="lg:col-span-2 sticky top-6">
              <div class="bg-white/5 rounded-2xl p-5 border border-white/10">
                <h2 class="text-sm font-bold text-white mb-4 tracking-wide uppercase">Resumen del pedido</h2>
                <div class="space-y-3 mb-5">
                  <div v-for="item in cartItems" :key="item.id" class="flex items-center gap-3">
                    <div class="w-10 h-10 rounded-lg overflow-hidden bg-white/5 flex-shrink-0">
                      <img :src="optimizeImage(item.images?.[0] || item.image, 100)" :alt="item.name" class="w-full h-full object-cover" />
                    </div>
                    <div class="flex-1 min-w-0">
                      <p class="text-white text-xs font-semibold leading-tight truncate">{{ item.name }}</p>
                      <p class="text-gray-500 text-xs">× {{ item.qty }}</p>
                    </div>
                    <span class="text-neon-green text-xs font-bold flex-shrink-0">{{ formatPrice(parsePrice(item.price) * item.qty) }}</span>
                  </div>
                </div>
                <div class="border-t border-white/10 pt-4 space-y-2">
                  <div class="flex justify-between text-sm"><span class="text-gray-400">Subtotal</span><span class="text-white font-semibold">{{ cartTotalFormatted }}</span></div>
                  <div class="flex justify-between text-sm">
                    <span class="text-gray-400">Envío</span>
                    <span class="text-neon-green font-semibold">{{ step >= 2 && currentShipping ? currentShipping.price : 'Calculando...' }}</span>
                  </div>
                  <div class="flex justify-between text-base font-black border-t border-white/10 pt-3 mt-2">
                    <span class="text-white">TOTAL</span>
                    <span class="text-neon-green">{{ cartTotalFormatted }}</span>
                  </div>
                </div>
                <p class="text-center text-gray-600 text-[10px] mt-4">
                  <fa-icon :icon="['fas', 'shield-alt']" class="mr-1" />Pago seguro · SSL encriptado
                </p>
                <div class="flex items-center justify-center gap-2 mt-3 flex-wrap">
                  <span v-for="logo in ['Nequi', 'PSE', 'Visa', 'Mastercard']" :key="logo" class="text-[10px] bg-white/5 border border-white/10 px-2 py-1 rounded-md text-gray-500 font-bold">{{ logo }}</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </ClientOnly>
    </main>

    <!-- Footer de políticas -->
    <footer class="w-full border-t border-white/8 mt-auto">
      <div class="max-w-5xl mx-auto px-6 py-5 flex flex-wrap items-center justify-center gap-x-6 gap-y-2">
        <NuxtLink to="/tienda" class="checkout-footer-link">Tienda</NuxtLink>
        <NuxtLink to="/agendar" class="checkout-footer-link">Agendar Cita</NuxtLink>
        <NuxtLink to="/politicas/envios" class="checkout-footer-link">Política de Envíos</NuxtLink>
        <NuxtLink to="/politicas/reembolsos" class="checkout-footer-link">Reembolsos</NuxtLink>
        <NuxtLink to="/politicas/privacidad" class="checkout-footer-link">Privacidad</NuxtLink>
        <NuxtLink to="/politicas/terminos" class="checkout-footer-link">Términos</NuxtLink>
        <a href="https://api.whatsapp.com/send?phone=573045840264" target="_blank" rel="noopener" class="checkout-footer-link">Contacto</a>
      </div>
    </footer>

    <!-- Modal próximamente -->
    <Transition name="fade">
      <div v-if="showSoonAlert" class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm">
        <div class="bg-barber-black border border-white/10 rounded-3xl p-8 max-w-sm w-full text-center shadow-2xl">
          <div class="w-20 h-20 bg-neon-green/10 rounded-full flex items-center justify-center mx-auto mb-6 text-neon-green">
            <fa-icon :icon="['fas', 'tools']" class="text-3xl" />
          </div>
          <h3 class="text-xl font-black text-white mb-2 uppercase tracking-tight">¡Próximamente!</h3>
          <p class="text-gray-400 text-sm leading-relaxed mb-8">
            Estamos integrando Wompi. Por ahora selecciona <span class="text-neon-green font-bold">"WhatsApp"</span> para coordinar directamente.
          </p>
          <button @click="showSoonAlert = false" class="w-full py-4 bg-neon-green hover:bg-neon-green-dark text-black font-black rounded-xl transition-all duration-300 uppercase text-sm">
            Entendido
          </button>
        </div>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
useSeoMeta({ title: 'Finalizar Compra | PersonalBarber' })

const router = useRouter()
const { cartItems, cartTotalFormatted, formatPrice, parsePrice, clearCart } = useCart()

const step = ref(1)
const form = reactive({ firstName: '', lastName: '', email: '', phone: '', city: '', address: '', notes: '' })
const touched = reactive({ firstName: false, lastName: false, email: false, phone: false, city: false, address: false })
const selectedShipping = ref('')
const selectedPayment = ref('wompi')
const showSoonAlert = ref(false)
const isProcessing = ref(false)

// Ciudades del área metropolitana de Medellín
const MEDELLIN_METRO = [
  'medellín', 'medellin', 'bello', 'itagüí', 'itagui', 'envigado', 'sabaneta',
  'la estrella', 'caldas', 'copacabana', 'girardota', 'barbosa', 'rionegro',
  'guarne', 'el retiro', 'la ceja', 'marinilla', 'el santuario',
]

const isMetroCity = computed(() => {
  const city = form.city.trim().toLowerCase()
  return MEDELLIN_METRO.some(m => city.includes(m))
})

const shippingMethods = computed(() => {
  if (isMetroCity.value) {
    return [
      {
        id: 'express',
        emoji: '⚡',
        label: 'PersonalBarber Express',
        desc: 'Entrega en 24–48 horas · Medellín y área metropolitana',
        price: '$8.000 COP',
        badge: 'Más rápido',
      },
      {
        id: 'pickup',
        emoji: '🏠',
        label: 'Recogida coordinada',
        desc: 'Acuerda el punto de entrega directo por WhatsApp · Gratis',
        price: 'Gratis',
      },
    ]
  }
  return [
    {
      id: 'nacional',
      emoji: '🚚',
      label: 'PersonalBarber Envíos Nacionales',
      desc: 'Buscamos la transportadora más conveniente para tu destino (Envia, Coordinadora, Interrapidísimo, Servientrega, etc.)',
      price: 'A cotizar',
      badge: 'Te contactamos',
      info: 'Una vez confirmes tu pedido, te enviamos por WhatsApp el costo exacto de envío a tu ciudad antes de procesar el pago.',
    },
    {
      id: 'pickup',
      emoji: '🏠',
      label: 'Recogida coordinada',
      desc: 'Acuerda el punto de entrega directo por WhatsApp · Gratis',
      price: 'Gratis',
    },
  ]
})

// Resetear método de envío cuando cambia la ciudad
watch(isMetroCity, () => { selectedShipping.value = '' })

const paymentMethods = [
  { id: 'wompi', emoji: '💳', label: 'Wompi', desc: 'Nequi, PSE, tarjetas débito/crédito', badge: 'Recomendado' },
  { id: 'nequi', emoji: '💜', label: 'Nequi', desc: 'Pago directo desde tu app Nequi' },
  { id: 'pse', emoji: '🏦', label: 'PSE', desc: 'Débito bancario en línea' },
  { id: 'whatsapp', emoji: '💬', label: 'Coordinar por WhatsApp', desc: 'Contacta al barber para acordar el pago' },
]

const currentShipping = computed(() => shippingMethods.value.find(s => s.id === selectedShipping.value))

const isEmailValid = computed(() => /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(form.email))
const isPhoneValid = computed(() => /^(57)?3\d{9}$/.test(form.phone.replace(/[\s\-+]/g, '')))
const step1Valid = computed(() =>
  form.firstName.trim() && form.lastName.trim() &&
  isEmailValid.value && isPhoneValid.value &&
  form.city.trim() && form.address.trim()
)

function nextStep() {
  if (step.value === 1) {
    Object.keys(touched).forEach(k => (touched as Record<string, boolean>)[k] = true)
    if (!step1Valid.value) return
    // Pre-seleccionar primer método de envío disponible
    if (!selectedShipping.value && shippingMethods.value.length > 0) {
      selectedShipping.value = shippingMethods.value[0].id
    }
  }
  step.value++
}

function goToStep(n: number) {
  if (n <= step.value) step.value = n
}

async function handleCheckout() {
  if (selectedPayment.value === 'whatsapp') {
    isProcessing.value = true
    try {
      const payload = {
        customer: { firstName: form.firstName, lastName: form.lastName, email: form.email, phone: form.phone, city: form.city, address: form.address },
        items: cartItems.map(i => ({ id: i.id, qty: i.qty })),
        paymentMethod: 'whatsapp',
      }
      const data = await $fetch<{ ok: boolean; order: { id: string; total_format: string } }>('/api/create_order', {
        method: 'POST', headers: { 'Content-Type': 'application/json' }, body: payload,
      })
      if (!data.ok) throw new Error('Error')
      const phone = '573045840264'
      const itemsList = cartItems.map(i => `• ${i.name} x${i.qty}`).join('\n')
      const shipping = currentShipping.value?.label || 'PersonalBarber Envíos'
      const msg = `¡Hola Andrés! Acabo de hacer un pedido:\n\n*ID:* ${data.order.id}\n\n${itemsList}\n\n*Envío:* ${shipping}\n*Total:* $${data.order.total_format} COP\n\nNombre: ${form.firstName} ${form.lastName}\nCiudad: ${form.city}\nDirección: ${form.address}`
      window.open(`https://api.whatsapp.com/send?phone=${phone}&text=${encodeURIComponent(msg)}`, '_blank')
      clearCart()
      router.push('/tienda')
    } catch (e) {
      console.error(e)
      alert('Error procesando tu orden. Intenta de nuevo.')
    } finally {
      isProcessing.value = false
    }
  } else {
    showSoonAlert.value = true
  }
}
</script>

<style scoped>
.input-field {
  width: 100%; background: rgba(255,255,255,0.05); border: 1px solid rgba(255,255,255,0.1);
  border-radius: 0.75rem; padding: 0.75rem 1rem; color: white; font-size: 0.875rem;
  outline: none; transition: all 0.2s;
}
.input-field::placeholder { color: #4b5563; }
.input-field:focus { border-color: rgba(57,255,20,0.5); background: rgba(255,255,255,0.08); }
.label-xs { display: block; font-size: 0.75rem; color: #9ca3af; font-weight: 600; margin-bottom: 0.25rem; }
.err { font-size: 0.625rem; color: #f87171; margin-top: 0.25rem; }
.checkout-footer-link { font-size: 0.7rem; font-weight: 600; color: #6b7280; text-decoration: none; letter-spacing: 0.05em; transition: color 0.2s; }
.checkout-footer-link:hover { color: #fff; }
</style>
