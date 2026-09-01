# Design Hit Counter (Gerçek Zamanlı İstek Sayacı / TPS Monitörü)

## LeetCode Bilgisi
- **Numara:** 362
- **Zorluk:** Medium
- **FinTech Şirketleri:** Stripe, Robinhood, Bloomberg, Citadel, Datadog

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Ödeme ve alım-satım ağ geçitlerinin anlık metrik ve telemetri altyapısında kullanılır:
1. **Anlık TPS (Transactions Per Second) Hesaplama:** Son 5 dakika veya son 1 dakika içinde gerçekleşen başarılı işlem sayısını anlık raporlama.
2. **Hata ve İstisna İzleme (Error Rate Metric):** Son 300 saniyedeki 5xx veya başarısız ödeme isteklerinin sayısını tutarak otomatik devre kesiciyi (Circuit Breaker) tetikleme.
3. **Piyasa Fiyat Güncelleme Frekansı:** WebSocket fiyat yayını yapan sunucunun saniyelik mesaj hacmini ölçme.

---

## Çözüm Yaklaşımı
Kuyruk (Queue / Slice) yaklaşımı her gelen istek için yeni bir eleman ekler. Milyonlarca istek geldiğinde bellek tükenir ($O(N)$ bellek) ve eski elemanları temizlemek maliyetli olur.

**Dairesel Kova Dizisi (Circular Bucket Buffer) Yaklaşımı ($O(1)$ Alan & $O(1)$ Zaman):**
- 300 saniyelik sabit boyutlu iki dizi oluşturulur:
  - `times [300]int`: İlgili kovanın temsil ettiği zaman damgası.
  - `hits [300]int`: İlgili saniyede kaydedilen vuruş (hit) sayısı.
- **`Hit(timestamp)`:**
  - Kova indeksi: `idx = timestamp % 300`.
  - Eğer `times[idx] != timestamp` ise: Kova eski bir 5 dakikalık periyoda aittir; `times[idx] = timestamp` ve `hits[idx] = 1` yapılarak sıfırlanır.
  - Eğer `times[idx] == timestamp` ise: Aynı saniye içinde gelen ek bir istektir; `hits[idx]++` yapılır.
- **`GetHits(timestamp)`:**
  - Sabit 300 kova taranır. `timestamp - times[i] < 300` olan aktif kovaların `hits[i]` değerleri toplanır.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):**
  - `Hit`: $O(1)$
  - `GetHits`: $O(1)$ - Her zaman tam olarak 300 kova döner (sabit maliyet).
- **Alan Karmaşıklığı (Space Complexity):** $O(1)$ - Saniyede 1 milyon istek gelse dahi bellek yalnızca iki adet 300 elemanlı dizi ile sınırlıdır!

---

## Go Eşzamanlılık (Concurrency)
- Eşzamanlı HTTP isteklerini karşılamak için `sync.Mutex` eklenmiştir.
