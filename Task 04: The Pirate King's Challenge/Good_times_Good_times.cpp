#include <bits/stdc++.h>
using namespace std;

int main() {
    int t;
    cin >> t;

    while (t > 0) {
        t--;

        long long x;
        cin >> x;

        // count the digits of x.
        int d = 0;
        long long temp = x;
        while (temp > 0) {
            d++;
            temp = temp / 10;
        }

        // build y = 10^d + 1.
        long long y = 1;
        for (int i = 0; i < d; i++) {
            y = y * 10;
        }
        y = y + 1;

        cout << y << "\n";
    }

    return 0;
}
