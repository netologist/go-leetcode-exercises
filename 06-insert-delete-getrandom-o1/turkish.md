# Insert Delete GetRandom O(1)

## LeetCode Bilgisi
- **Numara:** 380
- **Zorluk:** Medium
- **FinTech Şirketleri:** Robinhood, Citadel, Two Sigma, Bloomberg, Stripe

---

## FinTech Mülakatlarındaki Yeri ve Önemi
Bu veri yapısı, likidite havuzları ve rastgele örnekleme gerektiren finansal motorlarda sıkça kullanılır:
1. **Piyasa Yapıcı / Rastgele Likidite Sağlama:** Aktif emir havuzundan (order pool) eşit olasılıkla rastgele bir emir veya token seçip piyasaya sunma.
2. **Dolandırıcılık / Denetim Örneklemesi (Audit Sampling):** Anlık akan milyonlarca finansal işlem arasından $O(1)$ sürede rastgele işlem çekip derinlemesine denetim (fraud inspection) kuyruğuna yönlendirme.
3. **Kupon ve Ödül Dağıtımı:** Kullanıcı havuzuna dinamik ekleme/çıkarma yaparken her çekilişte $O(1)$ sürede rastgele bir hak sahibini belirleme.

---

## Çözüm Yaklaşımı
Yalnızca bir `map` kullanılırsa $O(1)$ ekleme ve silme yapılabilir, fakat rastgele bir anahtarı eşit olasılıkla seçmek $O(N)$ sürer.
Yalnızca bir `slice` kullanılırsa $O(1)$ `GetRandom` (indeks üzerinden) yapılabilir, ancak ortadaki bir elemanı silmek $O(N)$ kaydırma gerektirir.

**İkisinin Birleşimi:**
- `nums []int`: Değerleri tutar (`GetRandom` için $O(1)$ indeksleme).
- `pos map[int]int`: Hangi değerin `nums` içindeki hangi indekste olduğunu tutar (`val -> index`).

### $O(1)$ Silme Numarası: Swap and Pop (Yer Değiştir ve At)
Bir elemanı silmek istediğimizde:
1. Silinecek elemanın indeksini `pos` haritasından buluruz.
2. `nums` içindeki **son elemanı**, silinecek elemanın yerine yazarız (`nums[idx] = lastVal`).
3. `pos` içindeki `lastVal` indeksini güncelleriz (`pos[lastVal] = idx`).
4. `nums` diliminin son elemanını atarız (`nums = nums[:lastIdx]`) ve `pos` haritasından silinen değeri sileriz.
Bu sayede hiçbir eleman kaydırması yapılmaz ve silme işlemi $O(1)$ zamanda biter.

---

## Zaman ve Alan Karmaşıklığı
- **Zaman Karmaşıklığı (Time Complexity):**
  - `Insert`: Ortalama $O(1)$
  - `Remove`: Ortalama $O(1)$
  - `GetRandom`: $O(1)$
- **Alan Karmaşıklığı (Space Complexity):** $O(N)$ - $N$ eleman için slice ve hash map.
