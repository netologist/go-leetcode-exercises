# Accounts Merge (Hesapları Birleştirme / KYC Kimlik Çözümleme)

## LeetCode Bilgisi
- **Numara:** 721
- **Zorluk:** Medium
- **FinTech Şirketleri:** Stripe, Plaid, Robinhood, Brex, Coinbase, Revolut

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Kimlik çözümleme (Identity Resolution), Müşterini Tanı (KYC) ve dolandırıcılık önleme sistemlerinin en kritik problemidir:
1. **KYC / Tekil Müşteri Kimliği (Single Customer View):** Bir kullanıcının farklı e-posta, telefon veya banka hesapları ile açtığı hesapların ortak e-postalar üzerinden tek bir müşteri profili altında birleştirilmesi.
2. **Çoklu Hesap ve Bonus İstismarı Tespiti (Sybil / Promo Abuse):** Aynı kişinin farklı isimler veya hesaplarla kaydolup ortak bir e-posta üzerinden birbirine bağlandığını tespit etme.
3. **Plaid Benzeri Hesap Bağlama:** Kullanıcının farklı finansal kurumlardaki hesaplarını birleştirilmiş net varlık (Net Worth) görünümünde toplama.

---

## Çözüm Yaklaşımı
Bu problem bir **Ayrık Kümeler / Bağlantılı Bileşenler (Disjoint Set Union - Union Find)** problemidir:

1. **Union-Find Veri Yapısı:**
   - `parent []int`: Her hesabın üst/kök temsilcisini tutar (Yol Sıkıştırması / Path Compression ile $O(\alpha(N))$ hızında).
   - `rank []int`: Ağaç derinliğini dengeli tutmak için (Union by Rank).
2. **E-posta İndeksleme ve Birleştirme:**
   - Her e-posta taranır (`emailToAccountIdx[email]`).
   - Eğer bir e-posta daha önce başka bir hesap indeksinde görüldüyse, mevcut hesap ile eski hesap `uf.Union(i, oldIdx)` yapılarak aynı kümeye bağlanır.
3. **Kök Hesap Altında E-postaları Toplama:**
   - Her e-posta ait olduğu kök hesaba (`uf.Find(idx)`) eklenir.
4. **Çıktı Formatı:**
   - Her kök hesap için müşteri adı alınır, e-postalar alfabetik sıralanır (`sort.Strings`) ve sonuç listesine eklenir.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(N \cdot K \log(N \cdot K) + N \cdot K \cdot \alpha(N))$ - $N$ hesap sayısı, $K$ hesap başına maksimum e-posta sayısı. E-postaların sıralanması baskın faktördür.
- **Alan Karmaşıklığı (Space Complexity):** $O(N \cdot K)$ - E-posta haritası ve Union-Find dizileri.

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- İki farklı kişinin isimlerinin aynı olması (`"John"`), fakat e-postalarının farklı olması (birleşmemelidir).
- Zincirleme bağlar: A hesabı B ile, B hesabı C ile ortak e-postaya sahipse, A, B ve C tek bir hesapta toplanmalıdır.
