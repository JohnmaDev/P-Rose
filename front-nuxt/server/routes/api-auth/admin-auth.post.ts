// server/routes/api-auth/admin-auth.post.ts
// Endpoint server-side para validar el PIN del admin.
// El PIN vive ÚNICAMENTE en runtimeConfig (raíz, lado servidor).
// NUNCA se expone al navegador.

export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  const { pin } = body

  if (!pin || typeof pin !== 'string') {
    throw createError({ statusCode: 400, statusMessage: 'PIN requerido' })
  }

  // Leer el PIN SOLO desde la raíz privada de runtimeConfig (servidor)
  const config = useRuntimeConfig()
  const correctPin = String(config.adminPin || '').trim()

  if (!correctPin) {
    console.error('[Admin] CRITICAL: adminPin no está configurado en runtimeConfig del servidor')
    throw createError({ statusCode: 503, statusMessage: 'PIN no configurado en el servidor' })
  }

  const isValid = pin.trim() === correctPin

  return {
    ok: isValid,
    message: isValid ? 'Autenticado' : 'PIN incorrecto',
  }
})

