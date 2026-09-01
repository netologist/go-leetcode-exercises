# Number of Islands (Ada Sayısı / Dolandırıcılık Şebekesi Tespiti)

## LeetCode Bilgisi
- **Numara:** 200
- **Zorluk:** Medium
- **FinTech Şirketleri:** Bloomberg, Stripe, Citadel, Amazon, Robinhood, Revolut

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Finansal grafik (graph) ve ağ analizlerinde temel bileşen (connected components) bulma problemidir:
1. **Dolandırıcılık Halkası / Şebekesi Tespiti (Fraud Ring Detection):** Birbirine bağlı şüpheli banka hesapları, IP adresleri ve cihaz kimlikleri arasındaki izole dolandırıcılık kümelerini tespit etme.
2. **Finansal Likidite Havuzlarının Kümeler Halinde Modellenmesi:** Birbirine doğrudan veya dolaylı fon aktarımı yapabilen izole borsa/kripto likidite alt ağlarını belirleme.
3. **Müşteri Risk Yayılımı (Contagion Analysis):** Bir kurumun iflası durumunda doğrudan veya dolaylı etkilenen finansal kuruluş kümesini çıkarma.

---

## Çözüm Yaklaşımı
Matris üzerindeki bağlantılı '1' bloklarını bulmak için **Derinlik Öncelikli Arama (DFS)** veya **Genişlik Öncelikli Arama (BFS)** kullanılır:
- Matrisin tüm hücreleri `(r, c)` taranır.
- Eğer `grid[r][c] == '1'` (karşılaşılan yeni bir ada/küme) bulunursa:
  - Ada sayacı artırılır (`count++`).
  - DFS çağrısı başlatılarak bu adaya 4 yönden (yukarı, aşağı, sağ, sol) komşu olan tüm kara parçaları taranır.
  - Ziyaret edilen hücreler yerinde `'0'` (su) yapılarak (`in-place visited`) tekrar işlenmeleri engellenir.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(M \times N)$ - Matris içindeki her hücre en fazla iki kez ziyaret edilir ($M$ satır, $N$ sütun).
- **Alan Karmaşıklığı (Space Complexity):** $O(M \times N)$ - En kötü durumda (tüm matrisin kara olması) DFS çağrı yığını (call stack) derinliği.

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- Boş matris `[]` veya `[[]]`.
- Tamamı su `'0'` olan matris.
- Çapraz hücreler komşu sayılmaz (yalnızca 4 ana yön).
