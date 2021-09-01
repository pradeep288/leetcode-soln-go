class Solution:
    def addStrings(self, num1: str, num2: str) -> str:
        m, n = len(num1) - 1, len(num2) - 1
        res = []
        i, j = m, n
        carry = 0
        while i >= 0 or j >= 0:
            total = carry
            if i >= 0:
                total = total + ord(num1[i]) - 48
                i -= 1
            if j >= 0:
                total = total + ord(num2[j]) - 48
                j -= 1

            carry = total // 10
            val = total % 10

            res.insert(0, str(val))

        if carry > 0:
            res.insert(0, str(carry))

        return ''.join(res)
