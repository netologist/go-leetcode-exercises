# Insert Interval (Yeni Aralık Ekleme ve Birleştirme)

## LeetCode Bilgisi
- **Numara:** 57
- **Zorluk:** Medium
- **FinTech Şirketleri:** Stripe, Bloomberg, Citadel, Robinhood, Two Sigma

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Önceden sıralanmış zaman veya kademe listelerine yeni kural/aralık ekleme:
1. **Komisyon ve Ücret Kademeleri (Fee Brackets):** Mevcut işlem hacmi indirim kademelerine yeni bir promosyon aralığı ekleyip çakışan kademeleri tekilleştirme.
2. **Takas ve Açık Artırma Pencereleri (Settlement Windows):** Borsa takas takvimine yeni bir uzlaşma aralığı ekleme.
3. **Müşteri Limit ve Kredi Aralıkları:** Müşteriye tanımlanan dinamik kredi kullanım aralığını mevcut kotalarla birleştirme.

---

## Çözüm Yaklaşımı
Girdi dizisi zaten başlangıç zamanına göre sıralı ve çakışmasız olduğundan ek bir sıralama ($O(N \log N)$) yapmaya gerek yoktur. **3 Aşamalı Doğrusal Tarama ($O(N)$)** uygulanır:

1. **Aşama 1 (Öncesi):** `newInterval` başlamadan önce tamamen biten tüm aralıklar (`interval[1] < newInterval[0]`) doğrudan sonuca eklenir.
2. **Aşama 2 (Çakışanlar ve Birleştirme):** `newInterval` bittikten önce başlayan aralıklar (`interval[0] <= newInterval[1]`) ile `newInterval` tek bir aralıkta birleştirilir:
   - `newInterval[0] = min(newInterval[0], interval[0])`
   - `newInterval[1] = max(newInterval[1], interval[1])`
   Birleşen `newInterval` sonuca eklenir.
3. **Aşama 3 (Sonrası):** `newInterval` sonrasındaki kalan tüm aralıklar sonuca eklenir.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(N)$ - Dizi yalnızca tek bir döngüde bir kez taranır. Sıralama gerekmez.
- **Alan Karmaşıklığı (Space Complexity):** $O(N)$ - Sonuç aralıklarını tutan dilim (slice).

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- Boş aralık listesi `[]` durumuna ekleme.
- `newInterval`'ın tüm aralıklardan önce veya sonra olması.
- `newInterval`'ın mevcut tüm aralıkları yutması (`engulf`).
