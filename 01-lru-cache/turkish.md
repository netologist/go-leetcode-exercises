# LRU Cache (En Az Son Kullanılanı Tahliye Eden Önbellek)

## LeetCode Bilgisi
- **Numara:** 146
- **Zorluk:** Medium
- **FinTech Şirketleri:** Stripe, Robinhood, Citadel, Bloomberg, Two Sigma, Brex, Revolut

---

## FinTech Mülakatlarındaki Yeri ve Önemi
FinTech sistemlerinde (özellikle yüksek frekanslı alım-satım / HFT, ödeme ağ geçitleri ve borsa emir defterleri) veri erişim gecikmesi (latency) kritik öneme sahiptir:
1. **Piyasa Fiyatı ve Kotasyon Önbellekleme:** Anlık hisse senedi veya kripto fiyatlarını veritabanına gitmeden mikrosaniyeler mertebesinde $O(1)$ sürede sunmak.
2. **Idempotency Key (Tekrarlanabilirlik) Saklama:** Stripe benzeri ödeme sistemlerinde aynı ödemenin iki kez çekilmemesi için son $N$ işlemin anahtarlarını bellekte tutmak.
3. **Müşteri Oturum ve Bakiye Bilgisi:** Sık işlem yapan kullanıcıların hesap bakiyesini hızlıca sorgulamak.

---

## Çözüm Yaklaşımı ve Mimari
$O(1)$ `Get` ve $O(1)$ `Put` gereksinimi için iki veri yapısının birleşimi kullanılır:
1. **Hash Map (`map[int]*Node`):** Anahtar üzerinden ilgili düğüme $O(1)$ zamanda doğrudan erişim sağlar.
2. **Çift Yönlü Bağlı Liste (Doubly Linked List):** Elemanların kullanım sırasını tutar. En son kullanılanlar (MRU) listenin başına (`head`), en eski kalanlar (LRU) listenin sonuna (`tail`) yerleştirilir.

### Neden Dummy Head ve Dummy Tail?
Listenin başına ve sonuna sahte (sentinel) düğümler koymak, boş liste veya tek elemanlı liste kontrollerini ortadan kaldırır. Düğüm ekleme/çıkarma operasyonlarındaki `nil` pointer hatalarını (edge case) tamamen engeller.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):**
  - `Get(key)`: $O(1)$ - Hash map erişimi + bağlı listede düğüm kaydırma.
  - `Put(key, value)`: $O(1)$ - Hash map güncellemesi/eklemesi + tahliye işlemi.
- **Alan Karmaşıklığı (Space Complexity):**
  - $O(\text{Capacity})$ - Kapasite kadar `Node` ve hash map girdisi.

---

## Go Mülakat İpuçları ve Concurrency
Mülakatçılar genellikle şunu sorar: *"Bu yapıyı thread-safe (eşzamanlı kullanıma uygun) nasıl yaparsın?"*
- `sync.RWMutex` kullanarak: `Get` için `c.mu.Lock()` (çünkü `moveToHead` işlemi listeyi mutasyona uğratır; salt okuma değildir!), `Put` için `c.mu.Lock()` eklenmelidir.
