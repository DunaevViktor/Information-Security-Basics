from string import ascii_lowercase
from random import choice
import file_methods as fm


def create_cipher_file(path):
    alphabet = list(ascii_lowercase)
    amount_letters = len(alphabet)
    cipher = list()
    for i in range(amount_letters):
        letter = choice(alphabet)
        cipher.append(chr(i + 97) + '-' + letter)
        alphabet.remove(letter)
    fm.write(path, cipher)


def create_cipher(path, option):
    contents = fm.read(path)
    cipher = dict()
    if option == 'en':
        for str in contents:
            sub_str = str.split('-')
            cipher[sub_str[0]] = sub_str[1]
    elif option == 'de':
        for str in contents:
            sub_str = str.split('-')
            cipher[sub_str[1]] = sub_str[0]
    return cipher


def crypt(path, cipher):
    contents = fm.read(path)
    new_contents = list()
    new_str = list()
    for str in contents:
        for i in str:
            new_str.append(cipher[i])
        new_contents.append(''.join(new_str))
        new_str = []
    return new_contents


if __name__ == '__main__':
    create_cipher_file('cipher.txt')
    print('Cipher has been created.')
