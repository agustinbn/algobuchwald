---
description: Transcribe una clase (PPTX o PDF) a apuntes en Markdown, en español, con el estilo de /notes
argument-hint: <ruta-al-pptx-o-pdf> [MM-DD-AAAA]
---

Tu tarea es transcribir el material de una clase (un archivo `.pptx` o `.pdf`) en un apunte de clase en **Markdown**, siguiendo exactamente el estilo, tono e idioma (**español**) de los apuntes ya existentes en `notes/`.

Argumentos recibidos: `$ARGUMENTS`

- El primer argumento es la ruta al archivo `.pptx` o `.pdf` a transcribir (puede ser relativa al repo).
- Si se pasa un segundo argumento con formato `MM-DD-AAAA`, usalo como nombre del archivo de salida. Si no se pasa, inferí la fecha de la clase a partir del contenido del material (portada, fecha mencionada, nombre de archivo) y, si no hay ninguna pista, preguntale al usuario la fecha antes de crear el archivo.

## Paso 1: Extraer el contenido

- Si es `.pptx`: usá el skill de `pptx` para leer el contenido de las diapositivas (texto, código, tablas, notas del orador si las hay).
- Si es `.pdf`: usá el skill de `pdf` para extraer el texto (y OCR si el PDF es escaneado/de imágenes).
- Leé **todo** el material antes de escribir el resumen; no transcribas diapositiva por diapositiva, sino que organizá el contenido por temas como hacen los apuntes existentes.

## Paso 2: Revisar el estilo existente

Antes de escribir, mirá 2 o 3 archivos de `notes/*.md` (por ejemplo `notes/08-27-2026.md` que tiene bastante código, y `notes/09-09-2026.md` que tiene tablas) para calibrar el tono, la densidad y el formato exacto. Puntos clave del estilo a replicar:

- **Idioma**: todo en español (rioplatense, como habla un estudiante/docente de la UBA).
- **Encabezado**: `# Título de la clase` seguido de una línea `---`.
- **Secciones**: `##` para cada tema principal, `###` para subtemas.
- **Listas con guiones** (`-`) para conceptos, definiciones y observaciones. Usar **negrita** para resaltar términos clave/TDAs/nombres de algoritmos, y *cursiva* para énfasis puntual.
- **Bloques de código en Markdown** con el lenguaje correspondiente (` ```go `, ` ```python `, etc.) cada vez que el material muestre código, pseudocódigo o ejemplos de sintaxis. Si el material tiene pseudocódigo, escribilo como bloque de código igual (sin fence de lenguaje si no aplica a ningún lenguaje real).
- **Tablas** en Markdown cuando el material compara alternativas, tiempos, complejidades, etc.
- Si el material menciona ecuaciones o complejidades, usar el mismo estilo que los apuntes existentes (por ejemplo `O(n log n)`, `Θ(n)`, negrita para fórmulas importantes).
- Si la clase tiene avisos administrativos, fechas de entrega, checklists o recomendaciones de estudio, agregalos al final en una sección `## Avisos de la clase`, separada por `---` del resto, tal como en los apuntes existentes.
- No transcribas literalmente cada bullet de las diapositivas: sintetizá igual que un estudiante tomando apuntes, pero sin perder ningún concepto, ejemplo de código o dato importante (fórmulas, complejidades, nombres de algoritmos, ejercicios mencionados).

## Paso 3: Escribir el apunte

- Creá el archivo en `notes/MM-DD-AAAA.md` (mismo formato de nombre que los archivos existentes).
- Si ya existe un archivo con esa fecha, avisale al usuario y preguntá si hay que sobreescribirlo, fusionar el contenido o usar otra fecha, en vez de pisarlo directamente.
- Escribí el apunte completo respetando el estilo descripto arriba.

## Paso 4: Confirmar

Al terminar, mostrale al usuario la ruta del archivo creado y un resumen breve (2-3 líneas) de los temas cubiertos. No hace falta commitear ni pushear salvo que el usuario lo pida explícitamente.
