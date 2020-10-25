import simple_replace as sr
import file_methods as fm

def main():

    #сначала запускаешь скрипт cipher_Hill.py - будет создан cipher.txt с шифром
    #запускаешь этот скрипт. в функции create_cipher 'en' - зашифровать, 'de' - расшифровать

    cipher = sr.create_cipher('cipher.txt', 'en')
    fm.write('encrypt.txt', sr.crypt('text.txt', cipher))


if __name__ == '__main__':
    main()