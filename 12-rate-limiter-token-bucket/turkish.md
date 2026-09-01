# Token Bucket Rate Limiter (Jeton Kovası Hız Sınırlayıcı)

## Konu Bilgisi
- **Kategori:** FinTech Sistem Algoritmaları & Live Coding
- **Zorluk:** Medium
- **FinTech Şirketleri:** Stripe, Revolut, Robinhood, Brex, Adyen, Plaid, Coinbase

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Ödeme ve finans API'larında sistem stabilitesi ve DDoS/istismar koruması için en çok sorulan canlı kodlama sorusudur:
1. **Stripe API Hız Sınırlama (Rate Limiting):** Kullanıcı veya API anahtarı başına saniyede $N$ işlem (TPS) limitini koruma.
2. **Ani Yük Patlamaları (Burst Tolerance):** Normalde ortalama saniyede 10 istek kabul edilirken, anlık olarak 50 isteklik kısa patlamalara (burst) izin verilip sistemin tıkanmasını önleme.
3. **Kart Dolandırıcılığı Koruması (Card Testing Prevention):** Çalıntı kart numaralarını hızlıca deneyen botları engellemek için anlık istek kotası uygulama.

---

## Çözüm Yaklaşımı
Arka planda sürekli çalışan bir `time.Ticker` goroutine'i açmak yerine **Tembel Yenileme (Lazy Refill)** yaklaşımı kullanılır. Bu yaklaşım kaynak tüketimini sıfıra indirir:

1. **Alanlar:**
   - `capacity`: Kovadaki maksimum jeton kapasitesi (izin verilen anlık burst miktarı).
   - `tokens`: Kovada şu an bulunan kullanılabilir jeton miktarı (`float64`).
   - `refillRate`: Saniyede eklenen jeton sayısı (`float64`).
   - `lastRefillAt`: En son jeton yenilemesinin hesaplandığı zaman damgası (`time.Time`).
   - `mu sync.Mutex`: Eşzamanlı goroutine erişimlerinde race condition önlemek için kilit.

2. **İstek Geldiğinde (`AllowN`):**
   - Kilit alınır (`tb.mu.Lock()`).
   - `elapsed = time.Since(tb.lastRefillAt).Seconds()` hesaplanır.
   - Eklenen jetonlar: `tokens += elapsed * refillRate`. Eğer `tokens > capacity` ise `capacity` seviyesine çekilir.
   - `lastRefillAt = time.Now()` güncellenir.
   - Eğer `tokens >= n` ise, `tokens -= n` yapılarak `true` dönülür; yetersizse `false` dönülür.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(1)$ - İstek başına yalnızca basit matematiksel işlemler ve kilit yönetimi.
- **Alan Karmaşıklığı (Space Complexity):** $O(1)$ - Her hız sınırlayıcı nesnesi için sabit bellek.

---

## Go Eşzamanlılık (Concurrency) İpuçları
- Go'da `sync.Mutex` kullanımı ile thread-safety sağlanır.
- Lazy Refill sayesinde gereksiz goroutine sızıntısı (goroutine leak) veya timer yükü oluşmaz.
