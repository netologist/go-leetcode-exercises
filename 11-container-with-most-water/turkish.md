# Container With Most Water (En Çok Su Tutan Konteyner)

## LeetCode Bilgisi
- **Numara:** 11
- **Zorluk:** Medium
- **FinTech Şirketleri:** Robinhood, Bloomberg, Citadel, Goldman Sachs, Two Sigma

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Bu problem, finansal optimizasyon ve likidite modellemelerinde temel bir analojidir:
1. **Likidite Havuzu Derinliği (Liquidity Spread Window):** İki fiyat seviyesi arasındaki mesafeyi (spread / width) ve desteklenen minimum hacmi (depth / height) maksimize eden optimal ticaret koridorunu belirleme.
2. **Kâr Marjı ve Hacim Çarpımı:** Fiyat aralığı genişliği ile işlem kapasitesinin çarpımı olarak kâr alanını maksimize etme.
3. **Risk/Getiri Bandı Optimizasyonu:** Portföy sınırları arasında en geniş ve en güvenli risk marjı penceresini seçme.

---

## Çözüm Yaklaşımı
Kaba kuvvet $O(N^2)$ her çifti dener.

Optimal **İki İşaretçi (Two Pointers - Açgözlü / Greedy)** yaklaşımı:
- Bir işaretçi en başta (`left = 0`), diğeri en sonda (`right = len - 1`) konumlandırılır.
- Alan formülü: `Area = min(height[left], height[right]) * (right - left)`.
- Alan hesaplanıp `maxWater` güncellenir.
- **Kritik Mantık:** Alanı daraltırken genişlik (`width`) her adımda 1 azalır. Alanı büyütebilmenin tek yolu yüksekliği artırmaktır. Alan, **kısa olan çubuk tarafından sınırlandığı için**, kısa olan çubuğun işaretçisi içeriye doğru kaydırılır (`hLeft < hRight ? left++ : right--`). Uzun olanı hareket ettirmek alanı asla artıramaz.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(N)$ - İki işaretçi birbirine doğru ilerler ve tüm elemanlar sadece bir kez taranır.
- **Alan Karmaşıklığı (Space Complexity):** $O(1)$ - Yalnızca işaretçiler ve alan sayaçları için sabit bellek.

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- Dizi boyutunun minimum 2 olması.
- Monotonik artan veya azalan yükseklikler.
- İki ucun çok yüksek, ortadaki tüm elemanların alçak olması.
