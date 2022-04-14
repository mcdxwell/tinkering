

def palindrome(s):
    for i in range(0, int(len(s)/2) - 1):
        j = int(len(s)-i-1)

        if s[i] != s[j]:
            print(s[i], s[j])
            return False

    return True
