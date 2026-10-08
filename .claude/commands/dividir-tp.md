---
description: Divide un TP entre Agustin y Muriel con la mínima dependencia entre tareas y genera el mensaje de WhatsApp
argument-hint: <link-al-enunciado-del-tp>
---

Tu tarea es leer el enunciado de un TP y armar una **división de tareas entre Agustin y Muriel** (los dos únicos integrantes) que minimice la dependencia entre ambos, y devolver un **mensaje listo para pegar en WhatsApp**.

Argumento recibido: `$ARGUMENTS` (link al enunciado del TP).

## Paso 1: Leer el enunciado

- Traé el contenido del link (WebFetch; si es un PDF o Google Doc/Drive, usá la herramienta o skill que corresponda; si es un PDF local o descargado, el skill `pdf`).
- Si no podés acceder al link, avisale al usuario y pedile el texto o el archivo. No inventes el enunciado.
- Leelo completo: funciones/TDAs a implementar, firmas, interfaces dadas, pruebas provistas, informe/análisis pedido, fechas de entrega.

## Paso 2: Revisar el repo

Mirá rápido `tp0/`, `tp1/` y `tdas/` para ver la estructura típica (paquetes, archivos `.go`, tests, informe) y nombrar las tareas con archivos reales cuando sea posible.

## Paso 3: Dividir con mínima dependencia

Criterios, en orden de prioridad:

1. **Cortar por interfaz, no por flujo**: si B usa algo de A, definir primero la firma/contrato (interfaz del TDA, firma de la función) en el mensaje, así cada uno trabaja contra ese contrato y no espera al otro. Que B nunca tenga que esperar el código de A.
2. **Un módulo/archivo/TDA completo por persona** (implementación + sus propios tests), así no editan los mismos archivos y se evitan conflictos de git.
3. **Tareas independientes primero**: lo que no depende de nada (TDA aislado, parser, utilidades, informe, análisis de complejidad) se reparte sin vínculos.
4. **Balancear la carga** (esfuerzo parecido entre los dos).
5. Lo que sí o sí requiere a los dos (integración, `main`, revisión cruzada, armado final de la entrega) va al final y se indica claramente quién lo hace.

Si una dependencia es inevitable, dejala explícita: qué necesita, de quién y qué contrato (firma) usar mientras tanto.

## Paso 4: Responder

Respondé **solo** con el mensaje de WhatsApp, en un bloque de código para copiar fácil. Reglas:

- Español rioplatense, tono casual entre compañeros.
- Formato WhatsApp: `*negrita*` (un asterisco), `_cursiva_`, guiones para listas. Nada de `#`, tablas ni `**`.
- Estructura:
  - Una línea con el nombre del TP y la fecha de entrega (si figura).
  - `*Agustin:*` con sus tareas.
  - `*Muriel:*` con sus tareas.
  - `*Contratos/firmas acordadas:*` solo si hay dependencias (firmas exactas que ambos deben respetar).
  - `*Juntos:*` integración / revisión / entrega, si corresponde.
- Corto: cada tarea en una línea, sin explicaciones largas.

Después del bloque, agregá máximo 1–2 líneas fuera del mensaje solo si hay algo que el usuario deba revisar (supuestos, ambigüedades del enunciado).
