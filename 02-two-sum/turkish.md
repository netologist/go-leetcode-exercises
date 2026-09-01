# Two Sum (İki Sayının Toplamı)

## LeetCode Bilgisi
- **Numara:** 1
- **Zorluk:** Easy
- **FinTech Şirketleri:** Stripe, Bloomberg, Plaid, Robinhood, Adyen

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Two Sum, algoritmik mülakatların temeli kabul edilir; ancak FinTech bağlamında çok pratik kullanım senaryoları vardır:
1. **Çift Girişli Muhasebe (Double-Entry Ledger Balancing):** Borç ve alacak kayıtlarının birbirini sıfırlaması veya belirli bir mutabakat tutarını sağlaması.
2. **İşlem Eşleme (Transaction Reconciliation):** Banka ekstresindeki bir ana ödeme tutarını oluşturan iki parçalı alt ödemeyi bulma.
3. **FX / Arbitraj Çiftleri:** İki varlığın fiyat farkının veya toplamının hedeflenen marjı sağladığı ikilileri yakalama.

---

## Çözüm Yaklaşımı
Kaba kuvvet (brute force) $O(N^2)$ iç içe iki döngü yerine **Tek Geçişli Hash Map (One-Pass Hash Table)** yaklaşımı kullanılır:
- Dizi taranırken her $num$ için `complement = target - num` hesaplanır.
- Eğer `complement` daha önce hash map'e kaydedildiyse, çözüm anında `[]int{seen[complement], i}` olarak döndürülür.
- Kaydedilmediyse, mevcut eleman `seen[num] = i` şeklinde haritaya eklenir.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(N)$ - Dizi tek bir döngüde gezilir; hash map okuma ve yazma ortalama $O(1)$ sürede gerçekleşir.
- **Alan Karmaşıklığı (Space Complexity):** $O(N)$ - En kötü durumda $N-1$ eleman hash map'e kaydedilir.

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- Aynı indeksteki eleman iki kez kullanılamaz.
- Negatif sayılar ve sıfır hedef değeri.
- Dizi içinde tekrar eden sayılar (aynı değere sahip iki farklı indeks).
