# Go Image Hash

## Fundamentos teóricos de *image hashing* (hashes perceptuales)

---

### 1. ¿Qué es un hash perceptual?

Un **hash perceptual** (o *perceptual hash, p-hash*) es un resumen binario corto que describe *cómo se ve* una imagen, no su contenido exacto de bytes.

* A diferencia de un hash criptográfico (SHA-256), busca que **imágenes “visualmente” parecidas produzcan hashes cercanos** (poca distancia de Hamming) y que imágenes muy distintas queden lejos.
* El objetivo típico es **detección de duplicados, búsqueda inversa y filtrado de contenido** con tolerancia a ediciones menores (redimensionar, recortar, comprimir, cambiar brillo).

---

### 3. Pipeline genérico

1. **Normalización**

   * Redimensionar a tamaño fijo (ej. 8×8 px o 32×32 px).
   * Convertir a escala de grises (luminancia Y) para evitar sesgo de color.

2. **Transformación** (dependiendo del algoritmo)

   * Estadística simple (media o diferencias).
   * Análisis de frecuencias (DCT) o multirresolución (Wavelet).

3. **Cuantización**

   * Comparar cada coeficiente con un umbral (media/mediana).
   * Bit = 1 si coeficiente ≥ umbral, 0 en caso contrario.

4. **Empaquetado**

   * Agrupar bits en `uint64`, `uint128` o `[]byte`.
   * Proveer métodos `Distance(other)` que devuelvan la distancia de Hamming.

---

### 4. Algoritmos clásicos

| Algoritmo                   | Idea central                                             | Pasos (resumen)                                                                       | Pros / Contras                                                              |
| --------------------------- | -------------------------------------------------------- | ------------------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| **Average Hash (aHash)**    | Comparar cada píxel con la media global.                 | 1) Escalar 8×8 → 64 px. 2) Calcular media. 3) Bit=1 si píxel ≥ media.                 | Rapidísimo; sensible a contraste extremo y ruido.                           |
| **Difference Hash (dHash)** | Capturar gradientes horizontales o verticales.           | 1) Escalar 9×8. 2) Bit=1 si `p[x] > p[x+1]`.                                          | Mejor con variaciones de luz que aHash; sigue ignorando info de frecuencia. |
| **Perceptual Hash (pHash)** | Conservar solo las *bajas frecuencias* de la DCT.        | 1) Escalar 32×32. 2) DCT 2D. 3) Tomar sub-matriz 8×8 (excepto DC). 4) Umbral=mediana. | Robusto a compresión y blur; algo más costoso (O(N log N)).                 |
| **Wavelet Hash (wHash)**    | Similar a pHash pero usando transformada Wavelet (Haar). | 1) Escalar 32×32. 2) 2-Nivel Haar. 3) Mantener sub-bandas LL. 4) Umbral=mediana.      | Mejora rotaciones leves; algo más compleja de implementar.                  |

---

### 5. Distancia de Hamming y umbral

La **distancia de Hamming** cuenta bits distintos entre dos hashes del mismo tamaño.

* Para hashes de 64 bits, un umbral de 5–10 suele balancear *recall* y *precision* en duplicados comunes.
* La elección ideal depende del conjunto de datos; conviene trazar una curva ROC sobre imágenes de prueba.

---

### 6. Métricas y validación

* **Tasa de colisiones** en imágenes distintas.
* **Robustez** frente a transformaciones sintéticas (ruido, compresión).
* **Tiempo medio** de cómputo y de comparación.
* **Distribución de bits 0/1** (se busca \~50 % para evitar sesgo).
