# Merge Intervals (Aralıkları Birleştirme)

## LeetCode Bilgisi
- **Numara:** 56
- **Zorluk:** Medium
- **FinTech Şirketleri:** Stripe, Bloomberg, Robinhood, Citadel, Two Sigma, Brex

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Aralık (Interval) problemleri, zaman ve finansal planlama motorlarının kalbidir:
1. **Emir Defteri ve Fiyat Kademeleri:** Aynı veya çakışan fiyat aralıklarındaki alım-satım tekliflerinin derinlik defterinde birleştirilmesi.
2. **Faiz ve Kredi Hesaplama Dönemleri:** Çakışan borçlanma dönemlerinde birden fazla faiz oranı çakışmasını birleştirip tek bir sürekli süre hesaplama.
3. **Borsa İşlem ve Takas Saatleri:** Farklı borsaların açık olduğu seans saatlerinin birleşik görünümünü çıkarma.
4. **Faturalandırma ve Abonelik Dönemleri (Stripe):** Kullanıcının çakışan abonelik periyotlarının tekilleştirilmesi.

---

## Çözüm Yaklaşımı
1. **Başlangıç Zamanına Göre Sıralama ($O(N \log N)$):** Aralıklar `interval[0]` değerine göre küçükten büyüğe sıralanır.
2. **Tek Geçişte Birleştirme ($O(N)$):**
   - İlk aralık sonuç listesine eklenir.
   - Sonraki her `current` aralık için, sonuçtaki son aralığın bitişi (`last[1]`) ile `current[0]` karşılaştırılır.
   - Eğer `current[0] <= last[1]` ise çakışma vardır: `last[1] = max(last[1], current[1])` yapılarak genişletilir.
   - Çakışma yoksa, `current` bağımsız yeni bir aralık olarak sonuç listesine eklenir.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(N \log N)$ - Go `sort.Slice` sıralaması nedeniyle. Birleştirme geçişi $O(N)$ sürer.
- **Alan Karmaşıklığı (Space Complexity):** $O(N)$ - Birleştirilmiş aralıkları saklamak için ayrılan sonuç dilimi (slice).

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- Uç uca değen aralıklar: `[1, 4]` ve `[4, 5]` -> `[1, 5]` olarak birleşmelidir.
- İç içe geçmiş aralıklar: `[1, 10]` ve `[2, 6]` -> `[1, 10]` olmalıdır (`max` kontrolü kritik).
- Karışık (sıralanmamış) girdi.
