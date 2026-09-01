# Reorganize String (Karakterleri Yeniden Düzenleme / Görev Sıralayıcı)

## LeetCode Bilgisi
- **Numara:** 767
- **Zorluk:** Medium
- **FinTech Şirketleri:** Stripe, Citadel, Robinhood, Bloomberg

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Mülakatlarda çoğunlukla **"İşlem / Görev Zamanlayıcı (Task / Payment Throttle Dispatcher)"** senaryosu olarak sorulur:
1. **Aynı Kart / Üye İşyeri Hız Limiti (Card Velocity Throttle):** Aynı kredi kartından veya aynı satıcıdan gelen ardışık isteklerin arasına farklı işlemler yerleştirerek ödeme sağlayıcılarındaki ani rate limit blokajlarını önleme.
2. **Borsa Sembol Dağıtımı (Symbol Round-Robin Matching):** Tek bir hisse senedi sembolünün işlem kuyruğunu tekeline almasını engelleyerek adil ve sıralı eşleme yapma.
3. **API İstek Zamanlama:** Banka entegrasyonlarında aynı hedefe yönelik istekleri arka arkaya göndermemek için istekleri aralıklı dağıtma.

---

## Çözüm Yaklaşımı
1. **Frekans Sayımı:** Her karakterin (veya işlem türünün) toplam kaç kez geçtiği sayılır.
2. **Tavan Kontrolü (Pigeonhole / Güvercin Yuvası İlkesi):** Herhangi bir karakterin frekansı `(N + 1) / 2` değerinden büyükse, yan yana gelmeden dizilmeleri imkansızdır; hemen `""` döndürülür.
3. **Max-Heap (Öncelik Kuyruğu):**
   - En yüksek frekansa sahip iki farklı eleman heap'ten çekilir (`first`, `second`).
   - Sırayla sonuca eklenir ve sayaçları birer azaltılır.
   - Sayaçları hala 0'dan büyükse tekrar heap'e itilirler.
   - Bu açgözlü (greedy) yaklaşım, en sık geçen elemanların önce tüketilmesini ve asla yan yana gelmemesini garanti eder.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):** $O(N \log A)$ - $N$ karakter sayısı, $A$ farklı karakter sayısı (alfabe $A \le 26$). $A$ sabit kabul edildiğinde pratik olarak $O(N)$ sürer.
- **Alan Karmaşıklığı (Space Complexity):** $O(A)$ - Heap ve frekans haritası alfabe boyutu ile sınırlıdır ($O(1)$ ekstra alan).

---

## Go Mülakat Detayı: `container/heap`
Go mülakatlarında standart kütüphanedeki `container/heap` arayüzünü (`Len`, `Less`, `Swap`, `Push`, `Pop`) doğru implemente etmek en sık aranan becerilerden biridir.
