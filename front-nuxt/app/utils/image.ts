/**
 * Optimiza imágenes de Cloudinary dinámicamente.
 * Inyecta f_auto, q_auto y w_{width} con c_fill (tamaño exacto) para
 * garantizar que el navegador descargue solo los píxeles que necesita.
 *
 * @param url   - URL de Cloudinary u otra fuente
 * @param width - Ancho exacto en píxeles (0 = sin resize)
 *   Valores recomendados:
 *     80   → thumbnails en carrito mini / sticky bar
 *     150  → miniaturas de producto (galería)
 *     200  → tarjetas en móvil (2 columnas ~160px)
 *     300  → carrusel de relacionados
 *     400  → tarjetas en desktop
 *     800  → detalle de producto (imagen principal)
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
    // c_fill: genera exactamente w x h — elimina bytes innecesarios
    // vs c_limit que solo evita upscaling pero no fuerza el resize
    params.push(`w_${width}`, 'c_fill')
  }

  return `${base}/upload/${params.join(',')}/${path}`
}

export function optimizeSrcSet(url: string | undefined, widths: number[] = [200, 400]): string {
  if (!url || !url.includes('cloudinary.com')) return ''
  return widths.map(w => `${optimizeImage(url, w)} ${w}w`).join(', ')
}
