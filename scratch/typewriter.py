
import sys
import time

message = """
🚀 Refactorización "tranquila" completada con éxito.
✅ Constantes THEMES y DEPARTMENTS extraídas fuera del setup.
✅ Lógica de syncFilter simplificada y más legible.
✅ Commit realizado en la rama main.
❌ El push requiere credenciales manuales en este entorno.

¡Proyecto limpio y listo para la acción! 🥂
"""

for char in message:
    sys.stdout.write(char)
    sys.stdout.flush()
    time.sleep(0.02)
print()
