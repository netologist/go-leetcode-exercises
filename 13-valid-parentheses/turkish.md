# Valid Parentheses (Geçerli Parantez Kontrolü)

## LeetCode Bilgisi
- **Numara:** 20
- **Zorluk:** Easy
- **FinTech Şirketleri:** Stripe, Bloomberg, Robinhood, Plaid, Two Sigma

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Finansal veri formatları ve formül ayrıştırmanın temel yapıtaşıdır:
1. **Finansal Protokol ve Mesaj Doğrulama (ISO 8583 / FIX / SWIFT):** Finansal mesaj paketlerinin, JSON veya XML payload'larının ve iç içe geçmiş blokların sözdizimsel (syntax) doğruluğunu kontrol etme.
2. **Kural ve Akıllı Sözleşme Motorları:** Kredi değerlendirme veya otomatik alım-satım kurallarında `(risk_score > 700 AND (income > 50k OR collateral >= 100k))` gibi iç içe mantıksal ifadelerin doğrulanması.
3. **Finansal Formül Ayrıştırma (Financial Formula Parser):** Muhasebe formüllerindeki parantez hiyerarşisini kontrol etme.

---

## Çözüm Yaklaşımı
**Yığın (Stack)** veri yapısı kullanılır:
- Açılan her parantez (`(`, `{`, `[`) yığına itilir (`push`).
- Kapanan her parantez (`)`, `}`, `]`) için:
  - Yığın boşsa veya yığının tepesindeki eleman ilgili açılış parantezi ile eşleşmiyorsa sözdizimi geçersizdir (`false`).
  - Eşleşiyorsa yığının tepesindeki eleman çıkarılır (`pop`).
- Dizi sonuna gelindiğinde yığın tamamen boşsa parantezler dengelidir (`true`).

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(N)$ - Metin tek bir döngüde taranır; her karakter için yığın işlemleri $O(1)$'dir.
- **Alan Karmaşıklığı (Space Complexity):** $O(N)$ - En kötü durumda (örneğin `(((((`) tüm karakterler yığına itilir.

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- Karakter uzunluğunun tek sayı olması (`len % 2 != 0` anında `false`).
- Doğrudan kapanış paranteziyle başlama (ör: `]`).
- Açılış yapılıp kapanmadan dizinin bitmesi (ör: `({[`).
