# Coin Change (Bozuk Para / Nakit Para Çıkış Optimizasyonu)

## LeetCode Bilgisi
- **Numara:** 322
- **Zorluk:** Medium
- **FinTech Şirketleri:** Citadel, Bloomberg, Robinhood, Goldman Sachs, Two Sigma

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Bankacılık, döviz bürosu ve ödeme donanımlarının (ATM / POS) en klasik dinamik programlama problemidir:
1. **ATM Nakit Para Verme Optimizasyonu:** Müşterinin talep ettiği çekim tutarını ($Amount$), ATM'de bulunan mevcut banknot/madeni para kupürleriyle ($Coins$) **en az sayıda banknot** vererek karşılama.
2. **Kripto / Token UTXO Konsolidasyonu:** Belirli bir transfer tutarını karşılamak için cüzdandaki en az sayıda harcanmamış işlem çıktısını (UTXO) seçme.
3. **İşlem Ücreti Kırılımları:** Sabit ücret paketleri arasından hedeflenen komisyonu tam olarak en az adımda sağlama.

---

## Çözüm Yaklaşımı
Açgözlü (Greedy) yaklaşım her para birimi sistemi için optimal sonucu **vermez** (örneğin `coins = [1, 3, 4, 5]` ve `amount = 7` için Greedy `5+1+1` (3 adet) seçer; fakat DP ile optimal `4+3` (2 adet) bulunur).

Bu nedenle **Taban-Tavan Dinamik Programlama (Bottom-Up DP - Sınırsız Sırt Çantası / Unbounded Knapsack)** kullanılır:
- `dp[i]`: $i$ tutarını oluşturmak için gereken minimum madeni para sayısı.
- Başlangıçta `dp[0] = 0` ve diğer tüm `dp[1...amount] = Infinity` (erişilemez) atanır.
- $1$'den $amount$'a kadar her $i$ değeri ve her `coin` için:
  - Eğer `i - coin >= 0` ise: `dp[i] = min(dp[i], dp[i - coin] + 1)`.
- Sonuç `dp[amount]` sonsuzda kaldıysa bu tutara ulaşılamaz (`-1`), aksi halde `dp[amount]` döner.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(\text{Amount} \times C)$ - $C$ para birimi türü sayısı. İki katmanlı döngü.
- **Alan Karmaşıklığı (Space Complexity):** $O(\text{Amount})$ - DP tablosu için $amount + 1$ boyutunda dizi.

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- `amount = 0` (0 döner).
- Madeni paralarla tam tutarın oluşturulamaması (`-1` döner).
- Greedy yaklaşımın çuvalladığı para birimi kombinasyonları.
