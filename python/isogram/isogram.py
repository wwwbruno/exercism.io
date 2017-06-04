import re

def is_isogram(word):
    regex = re.compile('[^a-zA-Z]')
    clean_word = regex.sub('', word).lower()

    return len(set(clean_word)) == len(clean_word)
