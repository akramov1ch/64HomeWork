# Uy vazifasi: Go yordamida Redis ma'lumotlar tuzilmalari bilan ishlash

## Maqsad
Ushbu vazifaning maqsadi `Redis` ma'lumotlar tuzilmalari (stringlar, hashlar, setlar, listlar) va pub/sub xabar almashish modelidan foydalanishni Go tilida real dunyo ilovasida amalga oshirishni mashq qilishdir. 

## Vazifalar
1. **Strings**

    **Maqsad:** Oddiy kalit-qiymat juftlarini saqlash va olish uchun `Redis` stringlaridan foydalanish.

    **Vazifa:**
    1. `Redis`-da oddiy kalit-qiymat juftligini saqlang.
    2. Kalitning qiymatini oling.
    3. `Redis`-da qiymatni oshiradigan hisoblagichni amalga oshiring.


2. **Hashes**

    **Maqsad:** Strukturaviy ma'lumotlarni saqlash va olish uchun `Redis` xeshlaridan foydalanish.

    **Vazifa:**
    1. Foydalanuvchi ma'lumotlarini (masalan, foydalanuvchi identifikatori, nomi, elektron pochta manzili) `Redis` xeshida saqlang.
    2. Xeshning bir nechta maydonlarini atomik ravishda yangilash funktsiyasini amalga oshiring.
    3. Xeshning barcha maydonlari va qiymatlarini olish va ularni ko'rsatish uchun funktsiyani amalga oshiring.

3. **Lists**

    **Maqsad:** Elementlar roʻyxatini saqlash va olish uchun `Redis` lists dan foydalanish.

    **Vazifa:**
    1. List ga elementlar qo'shing va teskari tartibda elementlarni oling.
    2. Shart asosida list dan elementlarni olib tashlash uchun funktsiyani amalga oshiring.
    3. Redis list yordamida `pub/sub` modelini amalga oshiring.


## Qo'shimcha talabl
**Har bir vazifa uchun bajarilgan komandalarni va natijalarni ko'rsatadigan skrinshot qo'shing**

