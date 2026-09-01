# Top K Frequent Elements (En Sık Geçen K Eleman)

## LeetCode Bilgisi
- **Numara:** 347
- **Zorluk:** Medium
- **FinTech Şirketleri:** Bloomberg, Citadel, Stripe, Robinhood, Two Sigma, Revolut

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Gerçek zamanlı analitik ve dolandırıcılık tespiti motorlarının vazgeçilmezidir:
1. **En Yüksek Hacimli Varlıklar (Top Traded Tickers):** Borsa veya kripto platformunda son $N$ dakikada en çok işlem gören ilk $K$ hisse/kripto parayı listeleme.
2. **Dolandırıcılık / Risk İzleme:** En çok başarısız ödeme veya ters ibraz (chargeback) alan ilk $K$ üye işyeri veya IP adresini tespit etme.
3. **Müşteri Harcama Kategorileri:** Kullanıcının en çok alışveriş yaptığı $K$ harcama kategorisini çıkarma.

---

## Çözüm Yaklaşımı
İki temel çözüm mevcuttur:

1. **Min-Heap Yaklaşımı ($O(N \log K)$):**
   - Frekans haritası oluşturulur ($O(N)$).
   - Boyutu $K$ ile sınırlı bir Min-Heap tutulur. Her yeni eleman eklendiğinde boyut $K$'yı aşarsa en küçük frekanslı eleman atılır.
   - Zaman: $O(N \log K)$, Alan: $O(N + K)$.

2. **Kova Sıralaması (Bucket Sort) Yaklaşımı ($O(N)$ - Optimal):**
   - Her elemanın frekansı `map[int]int` ile hesaplanır.
   - Maksimum frekans en fazla dizinin uzunluğu ($N$) olabileceğinden, boyutu $N+1$ olan bir kova dizisi `buckets = [][]int` tanımlanır.
   - `buckets[freq]`, tam olarak `freq` kez geçen sayıların listesini tutar.
   - Kova dizisi en yüksek frekanstan (sondan) başa doğru taranarak $K$ adet eleman toplanır.
   - Zaman: $O(N)$ doğrusal, Alan: $O(N)$.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(N)$ - Bucket Sort ile hiçbir karşılaştırmalı sıralama yapmadan doğrusal sürede tamamlanır.
- **Alan Karmaşıklığı (Space Complexity):** $O(N)$ - Frekans haritası ve kova dizisi.

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- $K = 1$ veya $K = N$ (tüm elemanların benzersiz olması).
- Negatif sayılar ve sıfırların frekans sayımına dahil olması.
- Birden fazla elemanın aynı frekansa sahip olması.
