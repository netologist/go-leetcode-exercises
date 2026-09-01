# Meeting Rooms II (Minimum Sunucu / Kaynak Kapasitesi)

## LeetCode Bilgisi
- **Numara:** 253
- **Zorluk:** Medium
- **FinTech Şirketleri:** Bloomberg, Stripe, Citadel, Two Sigma, Robinhood, Brex

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Bu problem mülakatlarda çoğunlukla **"Aynı Anda Kaç Sunucu / Worker / Ödeme Kanalı Gerekli?"** şeklinde sorulur:
1. **Eşzamanlı Ödeme İşleme Kapasitesi (Payment Gateway Peak Concurrency):** Günün belirli saatlerinde başlayan ve biten ödeme işlemlerini işlemek için anlık kaç izole worker thread/process gerektiği.
2. **Borsa Emir Eşleme Motoru (Matching Engine Partitioning):** Çakışan açık artırma pencerelerini karşılamak için gereken asgari işlem kuyruğu sayısı.
3. **Finansal API Kota ve Sunucu Maliyet Optimizasyonu:** Sunucuları ne zaman ayağa kaldırıp ne zaman kapatacağımızı (auto-scaling) planlamak için zirve eşzamanlılığı hesaplama.

---

## Çözüm Yaklaşımı
İki temel çözüm vardır: **Min-Heap** veya **İki İşaretçi (Two Pointers - Başlangıç ve Bitiş Ayrıştırması)**. İki işaretçi yaklaşımı daha az bellek ayırır ve Go dilinde çok daha hızlı çalışır:

1. Başlangıç zamanları (`starts`) ve bitiş zamanları (`ends`) ayrı dizilere çıkarılır ve her ikisi de bağımsız olarak küçükten büyüğe sıralanır.
2. `startIdx` ve `endIdx` işaretçileri ile taranır:
   - `starts[startIdx] < ends[endIdx]` ise: En erken biten işlem dahi henüz tamamlanmadan yeni bir işlem başlamıştır. Yeni bir oda / sunucu tahsis edilir (`rooms++`).
   - Aksi takdirde: En erken biten işlem tamamlanmıştır, dolayısıyla o oda boşa çıkar ve yeni işlem için tekrar kullanılır (`endIdx++`).

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(N \log N)$ - Başlangıç ve bitiş dizilerinin sıralanması maliyeti. Tarama döngüsü $O(N)$ sürer.
- **Alan Karmaşıklığı (Space Complexity):** $O(N)$ - Başlangıç ve bitiş zamanlarını tutan iki dilim (slice).

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- Bir toplantının bitiş saati ile diğerinin başlangıç saati eşit olduğunda (örneğin `[1, 5]` ve `[5, 10]`): Aynı oda yeniden kullanılabilir, yeni oda gerekmez. (`<` kontrolü yapılır, `<=` değil).
