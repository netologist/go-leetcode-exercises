# Kth Largest Element in an Array (Dizideki En Büyük K. Eleman)

## LeetCode Bilgisi
- **Numara:** 215
- **Zorluk:** Medium
- **FinTech Şirketleri:** Citadel, Two Sigma, Bloomberg, Robinhood, Goldman Sachs

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Kantitatif finans, borsa emir eşleme ve risk analizinde çok kritik bir algoritmadır:
1. **Fiyat Persentili ve Medyanı (Price Percentile / Quantile):** Tüm veri kümesini pahalı bir şekilde sıralamadan ($O(N \log N)$) anlık emirlerin 95. veya 99. persentil (p99) fiyat seviyesini $O(N)$ ortalama sürede bulma.
2. **Karanlık Havuz (Dark Pool) Eşleme Fiyatı:** Belirli bir hacim kotasını dolduran eşik fiyatı (cutoff execution price) tespit etme.
3. **Piyasa Derinliği Sıralaması:** En yüksek teklif veya talepler arasından $K$'ıncı sıradaki kritik seviyeyi çekme.

---

## Çözüm Yaklaşımı
Tüm diziyi sıralamak $O(N \log N)$ sürer. Min-Heap kullanmak $O(N \log K)$ sürer.
Optimal çözüm **Rastgele Pivotlu Quickselect (Hoare Seçim Algoritması)** ile ortalama **$O(N)$** zamanda çalışır:

1. Dizideki en büyük $K$. eleman, sıralı bir dizide `targetIdx = len(nums) - k` indeksinde yer alır.
2. Dizi içinden rastgele bir pivot seçilir ve bölmeleme (partition) yapılır:
   - Pivottan küçük/eşit elemanlar sola, büyükler sağa yerleştirilir.
   - Pivotun nihai indeksi `pivotIdx` elde edilir.
3. **Karar Aşaması:**
   - Eğer `pivotIdx == targetIdx`: Aranan eleman bulunmuştur, hemen döndürülür!
   - Eğer `pivotIdx < targetIdx`: Aranan eleman sağ yarıdadır; sol sınır `left = pivotIdx + 1` yapılır.
   - Eğer `pivotIdx > targetIdx`: Aranan eleman sol yarıdadır; sağ sınır `right = pivotIdx - 1` yapılır.
4. QuickSort'un aksine iki yarıya birden dallanılmaz (yalnızca tek bir yarıya inilir). Dolayısıyla karmaşıklık: $N + N/2 + N/4 + \dots = 2N = O(N)$ olur.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** Ortalama $O(N)$ - Rastgele pivot seçimi sayesinde en kötü durum senaryosu ($O(N^2)$) pratikte bertaraf edilir.
- **Alan Karmaşıklığı (Space Complexity):** $O(1)$ - İteratif bölmeleme (in-place partition) yapıldığından ekstra bellek tüketilmez.

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- Tek elemanlı dizi ($K = 1$).
- Tekrar eden aynı sayılardan oluşan dizi.
- Negatif sayılar.
