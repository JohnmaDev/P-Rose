export default defineNuxtRouteMiddleware((to) => {
  // Solo se ejecuta si estamos procesando la petición en producción y no es localhost / 127.0.0.1
  const url = useRequestURL()
  const targetHost = 'personalbarber.co'

  if (
    url.hostname !== targetHost &&
    url.hostname !== 'localhost' &&
    url.hostname !== '127.0.0.1' &&
    !url.hostname.includes('deploy-preview')
  ) {
    return navigateTo(`https://${targetHost}${to.fullPath}`, {
      external: true,
      redirectCode: 301
    })
  }
})
