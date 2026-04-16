// server/api/verify-pin.post.ts
// Endpoint server-side para validar el PIN del admin
// El PIN se compara en el servidor donde process.env está disponible

export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  const { pin } = body

  if (!pin || typeof pin !== 'string') {
    throw createError({ statusCode: 400, statusMessage: 'PIN requerido' })
  }

  // Leer el PIN usando useRuntimeConfig para que Nitro extraiga correctamente el valor de .env/.env.local
  const config = useRuntimeConfig()
  const correctPin = String(config.adminPin || config.public.adminPin || '')

  if (!correctPin) {
    // Modo desarrollo sin PIN configurado — loguear error
    console.warn('[Admin] WARN: adminPin no está configurado en runtimeConfig')
    throw createError({ statusCode: 503, statusMessage: 'PIN no configurado en el servidor' })
  }

  const isValid = pin.trim() === correctPin.trim()

  return {
    ok: isValid,
    message: isValid ? 'Autenticado' : 'PIN incorrecto',
  }
})
