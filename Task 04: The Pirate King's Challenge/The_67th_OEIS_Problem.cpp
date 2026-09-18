#include <bits/stdc++.h>
using namespace std;

int main() {
    // build a list of primes we can use
    int limit = 200000;
    vector<bool> isComposite(limit + 1, false);
    vector<long long> primes;

    for (int i = 2; i <= limit; i++) {
        if (isComposite[i] == false) {
            primes.push_back(i);
            for (int j = i + i; j <= limit; j = j + i) {
                isComposite[j] = true;
            }
        }
    }

    int t;
    cin >> t;
    while (t > 0) {
        t--;
        int n;
        cin >> n;
        vector<long long> a(n + 1);
        a[1] = primes[0];
        for (int i = 2; i <= n - 1; i++) {
            a[i] = primes[i - 2] * primes[i - 1];
        }
        if (n >= 2) {
            a[n] = primes[n - 2];
        }
        for (int i = 1; i <= n; i++) {
            cout << a[i];
            if (i != n) {
                cout << " ";
            }
        }
        cout << "\n";
    }
    return 0;
}
