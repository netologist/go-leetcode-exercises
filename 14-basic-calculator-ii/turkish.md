# Basic Calculator II (Temel Hesap Makinesi II / İşlem Önceliği Motoru)

## LeetCode Bilgisi
- **Numara:** 227
- **Zorluk:** Medium
- **FinTech Şirketleri:** Bloomberg, Stripe, Robinhood, Citadel, Goldman Sachs

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Finansal kural motorlarında ve dinamik ücret hesaplayıcılarında sıklıkla karşımıza çıkar:
1. **Dinamik Komisyon ve Fiyatlandırma Formülleri:** Müşteri işlem hacmine veya borsa spreadine bağlı matematiksel ifadeleri harici bir `eval` kullanmadan güvenli ve deterministik olarak çalıştırma.
2. **Kredi / Faiz Formülü Yorumlayıcıları:** `anapara * faiz_orani + masraf / gun_sayisi` benzeri parametrik formülleri ayrıştırma.
3. **Güvenlik (Sandbox Expression Evaluation):** Kullanıcı tarafından girilen finansal formülleri güvenlik açığı oluşturmadan (kod enjeksiyonuna izin vermeden) AST / yığın mantığıyla çalıştırma.

---

## Çözüm Yaklaşımı
İşlem önceliğini (`*`, `/` önce, `+`, `-` sonra) yönetmek için **Yığın (Stack)** kullanılır:
- Metin karakter karakter taranır. Boşluklar atlanır.
- Sayılar okunurken `currentNum = currentNum * 10 + digit` şeklinde çok basamaklı sayılar inşa edilir.
- Bir operatörle veya dizinin sonuyla karşılaşıldığında, bir **önceki operatör (`lastOp`)** devreye sokulur:
  - `+`: `currentNum` doğrudan yığına eklenir (`stack = append(stack, currentNum)`).
  - `-`: `-currentNum` negatif olarak yığına eklenir.
  - `*`: Yığının tepesindeki sayı çekilir, `top * currentNum` çarpılır ve yığına geri konur.
  - `/`: Yığının tepesindeki sayı çekilir, `top / currentNum` bölünür ve yığına geri konur.
- Yeni operatör `lastOp = ch` yapılır ve `currentNum = 0` sıfırlanır.
- Döngü bittiğinde yığındaki tüm elemanlar toplanarak sonuç bulunur.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(N)$ - İfade dizisi yalnızca bir kez taranır.
- **Alan Karmaşıklığı (Space Complexity):** $O(N)$ - Ara sonuçları tutan yığın.

---

## Dikkat Edilmesi Gereken Noktalar (Edge Cases)
- Çok basamaklı sayılar (ör: `100 * 2`).
- Boşluk karakterleri (` 3/2 `).
- Tamsayı bölmesinde sıfıra doğru yuvarlama (Go'da tamsayı bölmesi `/` zaten sıfıra doğru keser).
- Dizinin son elemanına ulaşıldığında son sayının işleme alınması.
