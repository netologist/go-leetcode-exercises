# Group Anagrams (Anagramları Gruplama / Üye İşyeri İsmi Normalizasyonu)

## LeetCode Bilgisi
- **Numara:** 49
- **Zorluk:** Medium
- **FinTech Şirketleri:** Stripe, Plaid, Bloomberg, Robinhood, Revolut

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Ödeme işlem zenginleştirme (enrichment) ve üye işyeri (merchant) temizleme algoritmalarında kullanılır:
1. **İşlem Açıklaması ve Üye İşyeri Temizleme (Merchant Name Normalization):** Ham banka ekstrelerindeki yazım varyasyonlarını veya harf permütasyonlarını gruplayarak tek bir ana marka altında birleştirme.
2. **Kripto / Hash İmzası Eşleme:** Farklı sıralarla gelen işlem parametrelerinin veya token sembollerinin aynı anahtar imzaya sahip olup olmadığını doğrulama.
3. **Dolandırıcılık Tespiti:** Benzer harf kombinasyonları ile oluşturulan sahte alan adlarını (typosquatting) tespit etme.

---

## Çözüm Yaklaşımı
İki temel yaklaşım vardır:
1. **Sıralama (Sorting) Yaklaşımı:** Her kelimenin harfleri sıralanır (`"eat" -> "aet"`). Zaman: $O(N \times K \log K)$.
2. **Karakter Frekansı İmzası (Optimal $O(N \times K)$):**
   - Her kelime için 26 boyutlu sabit bir dizi `[26]byte` oluşturulur.
   - Kelimedeki her harf için `count[char - 'a']++` yapılır.
   - **Go Dili Avantajı:** Go'da sabit boyutlu diziler (`[26]byte`) karşılaştırılabilirdir (`comparable`) ve doğrudan `map[[26]byte][]string` içinde anahtar (key) olarak kullanılabilir! Bu sayede string dönüşümü veya bellek tahsisi (heap allocation) olmadan $O(1)$ harita indekslemesi yapılır.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(N \times K)$ - $N$ kelime sayısı, $K$ maksimum kelime uzunluğu.
- **Alan Karmaşıklığı (Space Complexity):** $O(N \times K)$ - Hash map ve gruplanmış sonuç dilimi.

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- Boş string `[""]`.
- Tek harfli kelimeler.
- Tamamı aynı kelimelerden oluşan girdi.
