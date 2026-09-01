# Best Time to Buy and Sell Stock (Hisse Alım-Satımı İçin En İyi Zaman)

## LeetCode Bilgisi
- **Numara:** 121
- **Zorluk:** Easy
- **FinTech Şirketleri:** Citadel, Bloomberg, Robinhood, Two Sigma, Goldman Sachs, Jane Street

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Alım-satım (trading) ve portföy yönetiminin en temel optimizasyon problemidir:
1. **Maksimum Kazanç Analizi:** Geçmiş fiyat hareketlerinde tek bir işlem döngüsünde (1 alım, 1 satım) elde edilebilecek teorik tavan kârı hesaplama.
2. **Akış Verisi (Streaming Data):** Fiyatlar anlık olarak gelirken ($O(1)$ ekstra bellek ile) minimum alış fiyatını güncelleyip anlık maksimum kârı hesaplamak.
3. **Maksimum Düşüş (Drawdown) Mantığı:** Bu algoritmanın tersi, bir varlığın zirveden dibe maksimum değer kaybını (Maximum Drawdown) hesaplamakta kullanılır.

---

## Çözüm Yaklaşımı
Kaba kuvvet yaklaşımı her gün alıp sonraki günlerde satmayı dener ($O(N^2)$).

Optimal **Tek Geçişli (One-Pass / Greedy)** yaklaşım:
- Bir `minPrice` değişkeni tutulur (başlangıçta sonsuz).
- Bir `maxProfit` değişkeni tutulur (başlangıçta 0).
- Dizi gezilirken:
  - Eğer mevcut fiyat `minPrice`'tan düşükse, `minPrice` güncellenir.
  - Değilse, `mevcutFiyat - minPrice` hesaplanır ve eğer `maxProfit`'ten büyükse `maxProfit` güncellenir.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(N)$ - Fiyat dizisi baştan sona yalnızca bir kez taranır.
- **Alan Karmaşıklığı (Space Complexity):** $O(1)$ - Yalnızca iki integer sayaç tutulur, fazladan bellek ayrılmaz.

---

## Edge Case'ler (Uç Durumlar)
- Fiyatların sürekli düşmesi (Kâr 0 olmalıdır, zarar edilmez çünkü işlem yapılmaz).
- Dizi boyutunun 0 veya 1 olması (İşlem yapılamaz, 0 döner).
- Tüm fiyatların aynı olması (Kâr 0).
