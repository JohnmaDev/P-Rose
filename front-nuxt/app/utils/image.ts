/**
 * Optimiza imágenes de Cloudinary dinámicamente.
 * Inyecta f_auto, q_auto y opcionalmente w_{width} con c_limit.
 *
 * @param url   - URL de Cloudinary u otra fuente
 * @param width - Ancho máximo en píxeles (0 = sin resize)
 *   Valores recomendados:
 *     100  → thumbnails en carrito/checkout
 *     400  → tarjetas de tienda
 *     800  → detalle de producto
 *     1200 → hero/headers
 */
export function optimizeImage(url: string | undefined, width = 0): string {
  if (!url) return '/hero_barber.webp'

  // Si no es URL de Cloudinary, retornar tal cual
  if (!url.includes('cloudinary.com')) {
    return url
  }

  const parts = url.split('/upload/')
  if (parts.length !== 2) return url

  const [base, rest] = parts

  // Detectar si el primer segmento de rest contiene transformaciones existentes
  let path = rest
  const firstSegment = rest.split('/')[0]

  if (firstSegment.includes(',') || firstSegment.includes('_') || (!firstSegment.startsWith('v') && !firstSegment.includes('.'))) {
    path = rest.substring(firstSegment.length + 1)
  }

  const params = ['f_auto', 'q_auto']
  if (width > 0) {
    params.push(`w_${width}`, 'c_limit')
  }

  return `${base}/upload/${params.join(',')}/${path}`
}
