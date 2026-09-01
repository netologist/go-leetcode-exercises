# Subarray Sum Equals K (Toplamı K Olan Alt Dizi Sayısı)

## LeetCode Bilgisi
- **Numara:** 560
- **Zorluk:** Medium
- **FinTech Şirketleri:** Stripe, Robinhood, Citadel, Bloomberg, Two Sigma

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Finansal defter (ledger) ve anomali tespiti sistemlerinde kritik bir algoritmadır:
1. **İşlem Defteri Sıfırlama ve Mutabakat (Ledger Reconciliation):** Sürekli akan para transferleri (pozitif ve negatif sayılar) içinde toplamı tam olarak $K$ (örneğin 0) eden kesintisiz işlem bloklarını tespit etme.
2. **Kara Para Aklama ve Dolandırıcılık Tespiti (Structuring / Smurfing):** Belirli bir eşik değere ($K$) ulaşan ardışık şüpheli transfer kalıplarını bulma.
3. **P&L (Kâr/Zarar) Hedef Segment Analizi:** Belirli bir zaman penceresinde net kârı $K$ seviyesine getiren işlem serilerini hesaplama.

---

## Çözüm Yaklaşımı
Kaba kuvvet yaklaşımı her alt diziyi toplar ($O(N^2)$).
Negatif sayılar da olabileceğinden kayan pencere (Sliding Window / İki İşaretçi) burada **çalışmaz**.

**Ön Ek Toplamı ve Hash Map (Prefix Sum + Hash Map) Yaklaşımı:**
- Matematiksel olarak: `Sum(i, j) = PrefixSum[j] - PrefixSum[i-1]`
- Eğer `Sum(i, j) == k` ise: `PrefixSum[j] - k == PrefixSum[i-1]` olmalıdır.
- Bu nedenle, dizi taranırken kümülatif toplam (`currentSum`) hesaplanır.
- Hash map'te `currentSum - k` değerinin daha önce kaç kez görüldüğü sorgulanır (`prefixCount[currentSum - k]`) ve toplam sonuca eklenir.
- Ardından `prefixCount[currentSum]++` yapılarak mevcut toplam haritaya işlenir.
- **Başlangıç Kuralı:** `prefixCount[0] = 1` yapılmalıdır; çünkü dizinin en başından başlayan ve tam olarak `k`'ya eşit olan alt diziler için `currentSum - k == 0` olur.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(N)$ - Dizi tek bir döngüde gezilir; hash map okuma/yazma $O(1)$ sürer.
- **Alan Karmaşıklığı (Space Complexity):** $O(N)$ - Ön ek toplamlarının frekanslarını tutan hash map.

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- Negatif sayılar ve sıfırlar: Toplamın artıp azalması veya aynı kalması durumunda hash map frekansları doğru işletilmelidir.
- `k` değerinin sıfır veya negatif olması.
- `prefixCount[0] = 1` başlangıç değerinin unutulmaması.
